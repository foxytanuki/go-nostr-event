package nostrevent

import (
	"testing"

	"github.com/nbd-wtf/go-nostr"
)

func TestSignPk(t *testing.T) {
	evt := NewEvent(1, nostr.Tags{}, "test")

	sk := "bcba0165fde1788455ec1788204ecc40f503bdde9bd4bdb254c047b9595e82e8"
	expected := "db804e02dc87159a7ce1db9ec47c3232b339a015cdfacb8b2e3c4f7530c58ea4"

	evt.SignPk(sk)

	if evt.PubKey != expected {
		t.Errorf("invalid PubKey")
	}
	ok, err := evt.CheckSignature()
	if err != nil {
		t.Errorf(err.Error())
	}
	if !ok {
		t.Errorf("invalid Signature")
	}
}
