package test

import (
	"bufio"
	"chapter13/chatbycellnet/cellnet"
	"chapter13/chatbycellnet/cellnet/packet"
	"chapter13/chatbycellnet/chat/proto"
	"chapter13/chatbycellnet/cellnet/socket"

	"fmt"
	"os"
	"strings"
)

func Readconsole(callback func(string))  {
	for{
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err!=nil{
			break
		}

		text = strings.TrimSpace(text)

		callback(text)
	}
}


func onMessage(ses cellnet.Session, raw interface{})  {
	switch ev:=raw.(type) {
	case socket.ConnectedEvent:
		fmt.Println("Connected")
	case packet.MsgEvent:
		msg := ev.Msg.(*proto.ChatACK)
		fmt.Println(msg)
	case socket.SessionClosedEvent:
		fmt.Println("disconnected")

	}
}


func main()  {
	queue := cellnet.NewEventQueue()

	peer := socket.NewConnector(packet.NewMessageCallback(onMessage), queue)

	peer.Start("127.0.0.1:8801")

	peer.SetName("client")

	queue.StartLoop()

	Readconsole(func(str string) {
		ses := peer.(interface{Session() cellnet.Session}).Session()

		ses.Send(&proto.ChatREQ{
			Content:str,
		})
	})

}


