package nostrevent

import (
	"fmt"
	"strings"

	"github.com/nbd-wtf/go-nostr"
)

type Relay struct {
	Url   string
	Read  bool
	Write bool
}

func formatRelays(relays []Relay) string {
	var ss []string
	for _, r := range relays {
		s := fmt.Sprintf("\\\"%s\\\":{\\\"read\\\":%t,\\\"write\\\":%t}", r.Url, r.Read, r.Write)
		ss = append(ss, s)
	}
	return fmt.Sprintf("{%s}", strings.Join(ss, ","))
}

func formatContacts(contacts []string) nostr.Tags {
	tt := nostr.Tags{}
	for _, contact := range contacts {
		t := nostr.Tag{"p", contact}

		tt = append(tt, t)
	}
	return tt
}

func NewContactList(contacts []string, relays []Relay) WrappedEvent {
	return NewEvent(nostr.KindContactList, formatContacts(contacts), formatRelays(relays))
}
