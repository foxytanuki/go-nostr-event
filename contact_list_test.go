package nostrevent

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/nbd-wtf/go-nostr"
)

func TestFormatRelays(t *testing.T) {
	r := []Relay{
		{
			Url:   "wss://relay.damus.io",
			Read:  true,
			Write: true,
		},
		{
			Url:   "wss://nos.lol",
			Read:  true,
			Write: false,
		},
	}

	result := formatRelays(r)
	// {\"wss://relay.damus.io\":{\"read\":true,\"write\":true},\"wss://nos.lol\":{\"read\":true,\"write\":false}}
	expected := "{\\\"wss://relay.damus.io\\\":{\\\"read\\\":true,\\\"write\\\":true},\\\"wss://nos.lol\\\":{\\\"read\\\":true,\\\"write\\\":false}}"

	if strings.Compare(result, expected) != 0 {
		fmt.Println(result)
		fmt.Println(expected)
		t.Errorf("fail: %s", result)
	}
}

func TestFormatContacts(t *testing.T) {
	expected := nostr.Tags{
		{"p", "c6dc2b963a3125b06dc4007fa21075405f53bbcafd3d1ae98d77ba2e434f6947"},
		{"p", "32e1827635450ebb3c5a7d12c1f8e7b2b514439ac10a67eef3d9fd9c5c68e245"},
	}
	c := formatContacts([]string{"c6dc2b963a3125b06dc4007fa21075405f53bbcafd3d1ae98d77ba2e434f6947", "32e1827635450ebb3c5a7d12c1f8e7b2b514439ac10a67eef3d9fd9c5c68e245"})
	if !reflect.DeepEqual(c, expected) {
		t.Errorf("expected: %v, got: %v", expected, c)
	}
}

// TODO: impl
func TestNewContactList(t *testing.T) {
	c := []string{
		"c6dc2b963a3125b06dc4007fa21075405f53bbcafd3d1ae98d77ba2e434f6947",
		"32e1827635450ebb3c5a7d12c1f8e7b2b514439ac10a67eef3d9fd9c5c68e245",
	}

	r := []Relay{
		{
			Url:   "wss://relay.damus.io",
			Read:  true,
			Write: true,
		},
		{
			Url:   "wss://nos.lol",
			Read:  true,
			Write: false,
		},
	}

	wevt := NewContactList(c, r)
	b, _ := json.Marshal(wevt)
	fmt.Println(string(b))
}
