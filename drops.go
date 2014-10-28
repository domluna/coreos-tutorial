package main

import (
	"fmt"

	"github.com/domluna/dogo"
)

func main() {
	client, err := dogo.NewClient("")
	if err != nil {
		fmt.Println(err)
		return
	}
	droplets, err := client.ListDroplets()
	if err != nil {
		fmt.Println(err)
		return
	}

	opts := &dogo.CreateDropletOpts{
		Region:            "nyc3",
		Size:              "512mb",
		Image:             5914637,
		// Your key ID here
		Keys:              []string{"136188"},
		PrivateNetworking: true,
	}

	droplet, err := client.CreateDroplet(opts)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(droplet)

	for _, droplet := range droplets {
		fmt.Printf("%v %d %v %s\n", droplet.Name, droplet.ID, droplet.Status, droplet.IPV4Addr())
		//         err := client.DeleteDroplet(droplet.ID)
		//         if err != nil {
		//                 fmt.Println(err)
		//         }
	}

}
