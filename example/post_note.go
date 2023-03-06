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
	nev := generateNoteEvent(sk)
	mev := generateMetadataEvent(sk)

	// publish the event
	url := "wss://relay.damus.io"
	relay, _ := nostr.RelayConnect(context.Background(), url)
	for _, ev := range []nostr.Event{nev, mev} {
		fmt.Println("published to ", url, relay.Publish(context.Background(), ev))
	}

	// Check the result on web clients, e.g.) https://snort.social, https://iris.to/
	// or CLI tools like the below with https://github.com/blakejakopovic/nostreq and https://github.com/blakejakopovic/nostcat
	// $ nostreq --authors <PubKey> | nostcat wss://relay.damus.io
}

func generateNoteEvent(sk string) nostr.Event {
	wevt := nostrevent.NewNote("hi")
	if err := wevt.SignPk(sk); err != nil {
		log.Fatal(err)
	}

	wevt.SignPk(sk)

	// print the event
	b, err := json.MarshalIndent(wevt, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	return wevt.Event
}

func generateMetadataEvent(sk string) nostr.Event {

	c := nostrevent.MetadataContent{
		Name:        "example_post_note",
		DisplayName: "example",
	}
	wevt := nostrevent.NewMetadata(c)
	wevt.SignPk(sk)
	// print the event
	b, err := json.MarshalIndent(wevt, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	return wevt.Event
}
