package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	nostrevent "github.com/foxytanuki/go-nostr-event"
	"github.com/nbd-wtf/go-nostr"
)

func main() {
	sk := nostr.GeneratePrivateKey()
	cev := nostrevent.NewNote("hi")
	if err := cev.SignPk(sk); err != nil {
		log.Fatal(err)
	}

	cev.SignPk(sk)

	// print the event
	b, err := json.MarshalIndent(cev, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	// publish the event
	url := "wss://relay.damus.io"
	relay, _ := nostr.RelayConnect(context.Background(), url)
	pk, _ := nostr.GetPublicKey(sk)
	fmt.Println("published to ", url, relay.Publish(context.Background(), cev.Event), "from", pk)

	// Check the result on web clients, e.g.) https://snort.social
	// or CLI tools like the below with https://github.com/blakejakopovic/nostreq and https://github.com/blakejakopovic/nostcat
	// $ nostreq --authors <PubKey> | nostcat wss://relay.damus.io
}
