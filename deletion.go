package nostrevent

import (
	"fmt"

	"github.com/nbd-wtf/go-nostr"
)

func NewDeletion(ids []string) WrappedEvent {
	var tags nostr.Tags
	for _, id := range ids {
		t := nostr.Tag{"e", id}
		tags = append(tags, t)
	}
	fmt.Println(tags)
	return NewEvent(nostr.KindDeletion, tags, "")
}
