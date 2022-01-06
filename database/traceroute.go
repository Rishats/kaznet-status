package database

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/pixelbender/go-traceroute/traceroute"
)

func trace(iip string) []net.IP {

	replyIps := []net.IP{}

	t := &traceroute.Tracer{
		Config: traceroute.Config{
			Delay:    50 * time.Millisecond,
			Timeout:  time.Second,
			MaxHops:  30,
			Count:    3,
			Networks: []string{"ip4:icmp", "ip4:ip"},
		},
	}
	defer t.Close()
	err := t.Trace(context.Background(), net.ParseIP(iip), func(reply *traceroute.Reply) {
		log.Printf("%d. %v %v", reply.Hops, reply.IP, reply.RTT)
		replyIps = append(replyIps, reply.IP)
	})
	if err != nil {
		log.Fatal(err)
	}

	return replyIps
}
