package nostrevent

import (
	"encoding/json"
	"strconv"

	"github.com/nbd-wtf/go-nostr"
)

type MetadataContent struct {
	Banner      string `json:"banner"`
	Website     string `json:"website"`
	Picture     string `json:"picture"`
	Lud16       string `json:"lud16"`
	DisplayName string `json:"display_name"`
	About       string `json:"about"`
	Name        string `json:"name"`
	NIP05       string `json:"nip05"`
	NIP05Valid  bool   `json:"nip05valid"`
}

func NewMetadata(content MetadataContent) CustomEvent {
	b, _ := json.Marshal(content)
	s := strconv.Quote(string(b))
	return NewEvent(nostr.KindSetMetadata, nostr.Tags{}, s)
}
