package fastudp

import "net"

// type EventHandler interface {
// 	OnReaded([]byte, *net.UDPAddr)
// }

type Handler func([]byte, *net.UDPAddr)
