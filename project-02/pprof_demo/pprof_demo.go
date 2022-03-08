package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

//.\pprof_demo.exe -cpu=true 等待30秒 会生成一个cpu.pprof
func main() {
	var isCpuPprof bool //是否开启CpuProfile标志位
	var isMemPprof bool //是否开启内存profile的标志位

	flag.BoolVar(&isCpuPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCpuPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}

		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}

	for i := 0; i < 6; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)

	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem  pprof failed, err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}
}

func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan ,value:%v\n", v)
		default:

		}
	}
}
