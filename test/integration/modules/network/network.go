package network

import (
	"go.ssnk.in/utils/network"
	"go.ssnk.in/utils/network/types"
)

func Test() {
	netw, _ := network.New(&types.Config{})
	client, _ := netw.Start()
	group, _ := client.Group()
	_ = group.Handler()
	_ = client.Handler()
}
