package main

import (
	"fmt"
	"net"
)

//udp server demo
func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
	})
	if err != nil {
		fmt.Printf("listener failed,err:%v\n", err)
		return
	}
	defer listener.Close()
	for {
		var buf [1024]byte
		n, addr, err := listener.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("read from udp failed,err:%v\n", err)
			return
		}
		fmt.Println("接收到的数据：", string(buf[:n]))
		_, err = listener.WriteToUDP(buf[:n], addr)
		if err != nil {
			fmt.Printf("write to %v failed.err:%v\n", addr, err)
			return
		}

	}
}
