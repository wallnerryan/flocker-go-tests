package main

import (
	"fmt"
	flockerclient "github.com/ClusterHQ/flocker-go"
)

func main() {
	client, err := flockerclient.NewClient("ip-172-20-0-183", 4523, "172.20.0.182", "/etc/flocker/cluster.crt", "/etc/flocker/api.key", "/etc/flocker/api.crt")
	if err != nil {
		fmt.Printf("Could not create client")
	}

	id, err := client.GetDatasetID("testvol-poc")
	if err != nil {
        	fmt.Println(err)
        	fmt.Printf("Could not find dataset ID")
	}
	fmt.Printf(id+"\n")

	s, err := client.GetDatasetState(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(s.Primary+"\n")
	fmt.Printf(s.Path+"\n")
}

