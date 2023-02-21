package nostrevent

import (
	"fmt"
	"testing"
)

func TestNewMetadata(t *testing.T) {
	c := MetadataContent{
		Name: "john",
	}
	ev := NewMetadata(c)
	fmt.Println(ev.Content)
}
