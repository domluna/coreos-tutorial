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
        for _, droplet := range droplets {
                fmt.Printf("%v %d %v %s\n", droplet.Name, droplet.ID, droplet.Status, droplet.IPV4())
                // err := client.DeleteDroplet(droplet.ID)
                // if err != nil {
                //         fmt.Println(err)
                // }
        }

}
