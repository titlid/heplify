package save

import (
	"fmt"
	"net"
	"os"
)

func SendRTP(data []byte) {
	udpAddr, _ := net.ResolveUDPAddr("udp4", "192.168.2.78:1234")

	//连接udpAddr，返回 udpConn
	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println("udp dial ok ")
	var i = 0
	// 发送数据
	//fmt.Println(t)
	len, err := udpConn.Write(data)
	i++
	if err != nil {
		return
	}
	fmt.Println("client write len:", len)
}
