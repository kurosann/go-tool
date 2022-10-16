package system

import (
	"fmt"
	"io"
	"net"
)

func RunForward() {
	listen, err := net.Listen("tcp", ":2333")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		conn, err := listen.Accept()
		dial, err := net.Dial("tcp", ":23333")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		go io.Copy(conn, dial)
		go io.Copy(dial, conn)
	}
}
