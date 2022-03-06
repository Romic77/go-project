package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

//编码
func Encode(message string) ([]byte, error) {
	//读取消息长度，转换成int32类型（占4个字节）
	var length = int(len(message))
	var pkg = new(bytes.Buffer)
	//写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	//写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

//解码
func Decode(reader *bufio.Reader) (string, error) {
	//读取消息的长度
	//读取前四个字节
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	//buffer返回缓冲中现有可读的字符串
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	//读取真正的消息数据
	pack := make([]byte, int(length+4))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
