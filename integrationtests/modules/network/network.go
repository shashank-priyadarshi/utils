package network

import (
	"github.com/shashank-priyadarshi/utilities/network"
	"github.com/shashank-priyadarshi/utilities/network/models"
)

func Test() {
	netw, _ := network.New(&models.Config{})
	client, _ := netw.Start()
	group, _ := client.Group()
	_ = group.Handler()
	_ = client.Handler()
}
