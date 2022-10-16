package system

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func Run() {
	port := flag.Int("p", 9000, "upload port")
	path := flag.String("f", "./upload", "upload path")
	flag.Parse()
	_, err := os.Stat(*path)
	if err != nil {
		log.Println("文件不存在")
		return
	}
	dir, err := os.ReadDir(*path)
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return
	}
	filterDir := make([]os.DirEntry, 0)
	for _, v := range dir {
		if !v.IsDir() {
			if v.Name()[0] != '.' {
				filterDir = append(filterDir, v)
			}
		}
	}
	for i, v := range filterDir {
		file, _ := os.Open(*path + "/" + v.Name())
		nbuf := []byte(v.Name())
		nbuf = append(nbuf, '\n')
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		buf := make([]byte, 4096)
		n, _ := file.Read(buf)
		buf[n] = '\n'
		nbuf = append(nbuf, buf[:n+1]...)
		_, err = conn.Write(nbuf)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// 更新进度条
		percent := float64(i+1) / float64(len(filterDir))
		processor := strings.Repeat("=", int(50*percent))
		if percent == 1 {
			processor = strings.Repeat("=", 50*int(percent)) + ">"
		} else {
			processor += ">" + strings.Repeat(" ", 50-len(processor))
		}
		time.Sleep(1 * time.Second)
		fmt.Printf("\r%.0f%%\t[%s]", percent*100, processor)
	}
	fmt.Println()
}
