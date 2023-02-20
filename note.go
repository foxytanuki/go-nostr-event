package nostrevent

import "github.com/nbd-wtf/go-nostr"

func NewNote(content string) CustomEvent {
	return NewEvent(nostr.KindTextNote, nostr.Tags{}, content)
}
