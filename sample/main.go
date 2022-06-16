package main

import (
	"log"
	"net"

	"github.com/glenn-wang/fastudp"
)

var fastUS *fastudp.Server = nil

func myHandler(frame []byte, addr *net.UDPAddr) {
	if len(addr.IP) == 4 {
		log.Println("recv packet; remote addr:", addr.String())

		// fixme: 接收路径 端口处理有误
		fastUS.WriteTo(frame, addr)
	}
}

func main() {
	log.Println("Hi")

	var err error

	fastUS, err = fastudp.NewUDPServer("udp4", "0.0.0.0:4321", true, 4, 1024, myHandler)

	// s, err := fastudp.NewUDPServer("udp6", "[fe80::604:4ff:fe16:1352]:4321", true, 4, 1024, eh)
	if err != nil {
		panic(err)
	}

	// var poller Poller
	// poller.fd =

	defer fastUS.Shutdown()

	select {}
}

// type eh3 struct {
// 	// Statd
// }

// func (it *eh3) OnReaded(frame []byte, addr *net.UDPAddr) {
// 	// it.Statd.Incr(len(frame))

// 	log.Println("remote addr len: ", len(addr.IP))

// 	log.Println("recv packet")
// }

//  eh := &eh3{}
// s, err := fastudp.NewUDPServer("udp4", "0.0.0.0:4321", true, 4, 1024, eh)

// eh.Statd.Tick()
