package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func main() {

	config := api.DefaultConfig()
	config.Address = "192.168.122.10:8500"
	config.Scheme = "http"
	config.Datacenter = "datacenter1"

	// Get a new client
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: "test/url", Value: []byte("https://www.example.com")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the exisitng Pair
	exisitngPair, _, err := kv.Get("tv1/url", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", exisitngPair.Key, exisitngPair.Value)

	// Lookup new pair
	pair, _, err := kv.Get("test/url", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)

	// List keys
	keys, _, err := kv.List("tv2", nil)
	if err != nil {
		panic(err)
	}

	for _, v := range keys {
		fmt.Printf("Key: %v\n", v.Key)
	}

}
