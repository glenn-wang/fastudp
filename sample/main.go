package main

import (
	"fmt"
	"log"
	"net"

	"github.com/glenn-wang/fastudp"
)

func myHandler(frame []byte, addr *net.UDPAddr) {
	log.Println("myHandler remote addr len: ", len(addr.IP))

	log.Println("myHandler recv packet")
}

func main() {
	fmt.Println("Hi")

	s, err := fastudp.NewUDPServer("udp4", "0.0.0.0:4321", true, 4, 1024, myHandler)

	// s, err := fastudp.NewUDPServer("udp6", "[fe80::604:4ff:fe16:1352]:4321", true, 4, 1024, eh)
	if err != nil {
		panic(err)
	}

	// var poller Poller
	// poller.fd =

	defer s.Shutdown()

	select {}

	// eh.Statd.Tick()

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
