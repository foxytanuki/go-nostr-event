package nostrevent

import (
	"time"

	"github.com/nbd-wtf/go-nostr"
)

type CustomEvent struct {
	nostr.Event
}

func (cev *CustomEvent) SignPk(sk string) error {
	pub, err := nostr.GetPublicKey(sk)
	if err != nil {
		return err
	}
	cev.PubKey = pub
	if err := cev.Sign(sk); err != nil {
		return err
	}

	return nil
}

func NewEvent(kind int, tags nostr.Tags, content string) CustomEvent {
	ev := nostr.Event{
		Kind:      kind,
		CreatedAt: time.Now(),
		Tags:      tags,
		Content:   content,
	}
	cev := CustomEvent{
		Event: ev,
	}

	return cev
}
