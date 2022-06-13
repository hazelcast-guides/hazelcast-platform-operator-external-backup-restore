package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"

	"github.com/hazelcast/hazelcast-go-client"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You should pass an argument to run: fill or size")
	} else if !(os.Args[1] == "fill" || os.Args[1] == "size") {
		fmt.Println("Wrong argument, you should pass: fill or size")
	} else {
		config := hazelcast.Config{}
		cc := &config.Cluster
		cc.Network.SetAddresses("34.70.165.31:5701")
		ctx := context.TODO()
		client, err := hazelcast.StartNewClientWithConfig(ctx, config)
		if err != nil {
			panic(err)
		}
		fmt.Println("Successful connection!")
		m, err := client.GetMap(ctx, "persistent-map")
		if err != nil {
			panic(err)
		}
		if os.Args[1] == "fill" {
			fmt.Println("Starting to fill the map with random entries.")

			for {
				num := rand.Intn(100_000)
				key := fmt.Sprintf("key-%d", num)
				value := fmt.Sprintf("value-%d", num)
				if _, err = m.Put(ctx, key, value); err != nil {
					fmt.Println("ERR:", err.Error())
				} else {
					if mapSize, err := m.Size(ctx); err != nil {
						fmt.Println("ERR:", err.Error())
					} else {
						fmt.Println("Current map size:", mapSize)
					}
				}
			}
		} else {
			if mapSize, err := m.Size(ctx); err != nil {
				fmt.Println("ERR:", err.Error())
			} else {
				fmt.Println("Current map size:", mapSize)
			}
		}
	}

}
