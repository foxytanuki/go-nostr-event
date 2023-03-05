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
	evt := NewMetadata(content)
	fmt.Println(evt.Content)
}
