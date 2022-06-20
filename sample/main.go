package main

import (
	"encoding/binary"
	"log"
	"net"

	fu "github.com/glenn-wang/fastudp"

	l "github.com/glenn-wang/fastudp/mylog"
)

var fastUS *fu.Server = nil

func onRead(frame []byte, addr *net.UDPAddr) {
	if len(addr.IP) == 4 {

		addr.Port = int(ntohs(uint16(addr.Port)))

		l.DEBUG("recv packet; remote addr: %s", addr.String())

		fastUS.WriteTo(frame, addr)
	}
}

func main() {
	log.Println("Hi")

	var err error

	fastUS, err = fu.NewUDPServer("udp4", "0.0.0.0:4321", true, 4, 1024, onRead)

	// s, err := fastudp.NewUDPServer("udp6", "[fe80::604:4ff:fe16:1352]:4321", true, 4, 1024, eh)
	if err != nil {
		panic(err)
	}

	// var poller Poller
	// poller.fd =

	defer fastUS.Shutdown()

	select {}
}

func ntohs(number uint16) uint16 {

	by := make([]byte, 2)
	binary.LittleEndian.PutUint16(by, uint16(number))

	return binary.BigEndian.Uint16(by)
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
