package nostrevent

import (
	"reflect"
	"testing"

	"github.com/nbd-wtf/go-nostr"
)

func TestNewDeletion(t *testing.T) {
	ids := []string{"d17414186fa1ea0d738b07bf96d12901a55e2c9ddc2008ebd43ab6a1eebdde18"}
	expected := nostr.Tags{
		{
			"e",
			"d17414186fa1ea0d738b07bf96d12901a55e2c9ddc2008ebd43ab6a1eebdde18",
		},
	}
	wevt := NewDeletion(ids)
	if !reflect.DeepEqual(wevt.Tags, expected) {
		t.Errorf("expected: %v, got: %v", expected, wevt.Tags)
	}
}
