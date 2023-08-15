package sanctum

import (
	"fmt"
	"testing"
)

func TestHawaldar(t *testing.T) {
	length := len(randomString())

	fmt.Printf("Length %d\n", length)
	if length != 20 {
		t.Errorf("Length is not 20")
	}
}
