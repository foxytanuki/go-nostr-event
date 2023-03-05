package nostrevent

import (
	"fmt"
	"testing"
)

func TestNewMetadata(t *testing.T) {
	c := MetadataContent{
		Name: "john",
	}
	content, err := PrepareMetadataContent(c)
	if err != nil {
		t.Error(err)
	}
	ev := NewMetadata(content)
	fmt.Println(ev.Content)
}
