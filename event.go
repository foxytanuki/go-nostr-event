package nostrevent

import (
	"time"

	"github.com/nbd-wtf/go-nostr"
)

type WrappedEvent struct {
	nostr.Event
}

func (wev *WrappedEvent) SignPk(sk string) error {
	pub, err := nostr.GetPublicKey(sk)
	if err != nil {
		return err
	}
	wev.PubKey = pub
	if err := wev.Sign(sk); err != nil {
		return err
	}

	return nil
}

func NewEvent(kind int, tags nostr.Tags, content string) WrappedEvent {
	ev := nostr.Event{
		Kind:      kind,
		CreatedAt: time.Now(),
		Tags:      tags,
		Content:   content,
	}
	wev := WrappedEvent{
		Event: ev,
	}

	return wev
}
