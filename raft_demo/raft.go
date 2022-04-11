package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/rpc"
	"sync"
	"time"
)

// 3个节点常量
const raftCount = 3

//
// Leader
// @Description leader 对象
//
type Leader struct {
	//任期
	Term int
	//LeaderId
	LeaderId int
}

//
// Raft
// @Description 1.实现3节点选举 2.实现分布式选举，接入rpc调用  3.
//
type Raft struct {
	//锁
	mu sync.RWMutex
	//节点编号
	me int
	//当前任期
	currentTerm int
	//为那个节点投票
	voteFor int
	//状态 0-follower 1-candidate 2-leader
	state int
	//发送最后一条数据的时间
	lastMessageTime int64
	//当前的leader
	currentLeader int
	//节点间发信息的通道
	message chan bool
	//选举的通道
	electCh chan bool
	//心跳信息通道
	heartBeat chan bool
	//返回心跳信息的通道
	heartBeatRe chan bool
	//超时时间
	timeout int
}

//0 还没上任
var leader = Leader{0, -1}

func main() {
	//过程，3个节点，最初都是follower
	//如果有candidate状态，进行投票和拉票，选举Leader

	for i := 0; i < raftCount; i++ {
		Init(i)
	}

	//加入服务端监听
	rpc.Register(new(Raft))
	rpc.HandleHTTP()
	//监听服务
	http.ListenAndServe(":8080", nil)

	for {

	}
}

//
// Init
// @Description 初始化节点
// @param me int
// @return *Raft
//
func Init(me int) *Raft {
	rf := &Raft{}
	rf.me = me
	rf.voteFor = -1
	rf.state = 0
	rf.timeout = 0
	rf.currentLeader = -1
	rf.setTerm(0)
	rf.message = make(chan bool)
	rf.electCh = make(chan bool)
	rf.heartBeat = make(chan bool)
	rf.heartBeatRe = make(chan bool)
	//设置随机种子
	rand.Seed(time.Now().UnixNano())

	//选举的协程
	go rf.election()

	//心跳检测的协程
	go rf.sendLeaderHeartBeat()
	return rf
}

//
// setTerm
// @Description 设置任期
// @receiver rf *Raft
// @param term int
//
func (rf *Raft) setTerm(term int) {
	rf.currentTerm = term
}

func (rf *Raft) election() {
	var result bool
	for {
		timeout := randRange(150, 300)
		rf.lastMessageTime = millisecond()
		select {
		//延迟等待1毫秒
		case <-time.After(time.Duration(timeout) * time.Millisecond):
			fmt.Println("当前节点的状态为:", rf.state)
		}
		//如果每个节点都是false 那么就需要选主
		result = false
		if !result {
			//选主逻辑
			result = rf.electionLeader(&leader)
		}
	}
}

//
// randRange
// @Description 随机值
// @param min int64
// @param max int64
// @return int64
//
func randRange(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

//
// millisecond
// @Description 获得当前时间，发送最后一条数据的时间
// @return int64
//
func millisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//
// election_leader
// @Description 实现选主逻辑
// @receiver rf *Raft
// @param leader *Leader
// @return bool
//
func (rf *Raft) electionLeader(leader *Leader) bool {
	//定义超时
	var timeout int64 = 100

	//投票数量
	var vote int
	//心跳
	var triggerHeartBeat bool
	//时间
	last := millisecond()
	//用户返回值
	success := false

	//当前节点变成candidate
	rf.mu.Lock()
	rf.becomeCandidate()
	rf.mu.Unlock()
	fmt.Println("start election leader")

	for {
		//遍历所有节点拉取选票
		for i := 0; i < raftCount; i++ {
			if i != rf.me {
				go func() {
					if leader.LeaderId < 0 {
						//设置投票
						rf.electCh <- true
					}
				}()
			}
		}
		//投票数量
		vote = 1
		//遍历节点
		for i := 0; i < raftCount; i++ {
			//计算投票数量
			select {
			case ok := <-rf.electCh:
				if ok {
					vote++
					//选票个数>节点个数/2 则成功
					success = vote > raftCount/2

					if success && !triggerHeartBeat {
						//选主成功，触发心跳
						triggerHeartBeat = true
						rf.mu.Lock()
						rf.becomeLeader()
						rf.mu.Unlock()

						//leader向其他节点发送心跳信号
						rf.heartBeat <- true
						fmt.Println(rf.me, "节点成为了leader")
						fmt.Println("leader开始发送心跳信号")
					}
				}
			}
		}
		if timeout+last < millisecond() || (vote > raftCount/2 || rf.currentLeader > -1) {
			break
		} else {
			select {
			case <-time.After(time.Duration(10) * time.Millisecond):

			}
		}
	}
	return success
}

//
// becomeCandidate
// @Description 修改自身状态成为candidate候选人
// @receiver rf *Raft
//
func (rf *Raft) becomeCandidate() {
	rf.state = 1
	rf.setTerm(rf.currentTerm + 1)
	rf.voteFor = rf.me
	rf.currentLeader = -1
}

//
// becomeLeader
// @Description 成为leader
// @receiver rf *Raft
//
func (rf *Raft) becomeLeader() {
	rf.state = 2
	rf.setTerm(rf.currentTerm + 1)
	rf.currentLeader = rf.me
}

//
// sendLeaderHeartBeat
// @Description leader 发送心跳给follower节点
// @receiver rf *Raft
//
func (rf *Raft) sendLeaderHeartBeat() {
	for {
		select {
		case <-rf.heartBeat:
			rf.sendAppendEntriesImpl()
		}
	}
}

//
// sendAppendEntriesImpl
// @Description 用于返回给leader的确认信号
// @receiver rf *Raft
//
func (rf *Raft) sendAppendEntriesImpl() {
	if rf.currentLeader == rf.me {
		//此时是leader
		var success_count int64 = 0
		for i := 0; i < raftCount; i++ {
			if i != rf.me {
				go func() {
					//rf.heartBeatRe <- true
					rpcClient, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
					if err != nil {
						log.Fatal(err)
					}
					var ok bool = false
					rpcClient.Call("Raft.Communication", Param{Msg: "hello"}, &ok)

					if ok {
						rf.heartBeatRe <- true
					}
				}()
			}
		}

		for i := 0; i < raftCount; i++ {
			select {
			case ok := <-rf.heartBeatRe:
				if ok {
					success_count++
					if success_count > raftCount/2 {
						fmt.Println("投票选举成功，心跳信号ok")
						log.Fatal("程序结束")
					}
				}
			}
		}
	}
}

type Param struct {
	Msg string
}

//
// Communication
// @Description rpc通信
// @receiver rf *Raft
// @param p Param
// @param a *bool
//
func (rf *Raft) Communication(p Param, a *bool) error {
	fmt.Println(p.Msg)
	*a = true
	return nil
}
