package googleoauth2

import (
	"golang.org/x/net/context"
	"testing"
)

func TestGetToken(t *testing.T) {
	_, err := GetToken(context.Background(), "123")
	if err != nil {
		t.Error(err)
	}
}
