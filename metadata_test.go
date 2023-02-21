package nostrevent

import (
	"testing"
)

func TestNewMetadata(t *testing.T) {
	c := MetadataContent{
		Name: "john",
	}
	NewMetadata(c)
}
