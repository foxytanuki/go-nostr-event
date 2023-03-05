package nostrevent

import (
	"time"

	"github.com/nbd-wtf/go-nostr"
)

type WrappedEvent struct {
	nostr.Event
}

func (wevt *WrappedEvent) SignPk(sk string) error {
	pub, err := nostr.GetPublicKey(sk)
	if err != nil {
		return err
	}
	wevt.PubKey = pub
	if err := wevt.Sign(sk); err != nil {
		return err
	}

	return nil
}

func NewEvent(kind int, tags nostr.Tags, content string) WrappedEvent {
	evt := nostr.Event{
		Kind:      kind,
		CreatedAt: time.Now(),
		Tags:      tags,
		Content:   content,
	}
	wevt := WrappedEvent{
		Event: evt,
	}

	return wevt
}
