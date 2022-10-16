package system

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func StartServer() {
	port := flag.Int("p", 9000, "upload port")
	path := flag.String("f", "./test", "upload path")
	flag.Parse()
	fileInfo, err := os.Stat(*path)
	if err != nil {
		os.Mkdir(*path, os.ModePerm)
	}
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer listen.Close()
	fmt.Printf("server start at port:%d\n", *port)
	fmt.Printf("接收端启动成功，等待发送端发送文件！上传路径为%s\n", fileInfo.Name())
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		reader := bufio.NewReader(conn)
		for {
			fileName, err := reader.ReadSlice('\n')
			if err != nil {
				if err != io.EOF {
					log.Println(err)
				} else {
					break
				}
			}

			fileContent, err := reader.ReadSlice('\n')

			if err != nil {
				if err != io.EOF {
					log.Println(err)
				} else {
					break
				}
			}
			name := fmt.Sprintf("%s/%s", *path, strings.TrimSpace(string(fileName)))
			fmt.Printf(name + "\n")
			file, err := os.Create(name)
			file.Write(fileContent)
		}
	}
}
