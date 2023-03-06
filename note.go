package nostrevent

import "github.com/nbd-wtf/go-nostr"

func NewNote(content string) WrappedEvent {
	return NewEvent(nostr.KindTextNote, nostr.Tags{}, content)
}
