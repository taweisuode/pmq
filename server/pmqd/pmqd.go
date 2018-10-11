package pmqd

import (
	"fmt"
	"net"
)

type Topic struct {
	tcpListener  net.TCPListener
	httpListener net.TCPListener
}

func TcpListen() {
	topicObject := &Topic{}
	tcpListener, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP("127.0.0.1"), 9999, ""})
	topicObject.tcpListener = *tcpListener
	if err != nil {
		fmt.Println("tcp connect fail", err.Error())
		return
	}
	id := 0
	for {
		conn, err := topicObject.tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println("tcp accept fail", err.Error())
		}
		id++
		go HandleConn(conn, id)
	}
	fmt.Println("hello world")
}

func HandleConn(conn *net.TCPConn, id int) {
	fmt.Println("send your message")
	message := ""
	for {
		fmt.Scanf("%s", &message)
		switch message {
		case "create_topic":
			topic := ""
			channel := ""
			fmt.Printf("please input topic and message", &topic, &channel)
		}
		_, err := conn.Write(([]byte(message)))
		if err != nil {
			fmt.Println("write data error", err)
		}
		if message == "exit" {
			break
		}
	}
	fmt.Println("connect close")
	defer conn.Close()
}