package trip

import (
	"testing"
)

func TestErrorReturnedWhenNotLoggedIn(t *testing.T) {
	_, err := GetTripsByUser(nil)
	if err.Error() != "User Not Logged In" {
		t.Errorf("Expected User Not Logged In Error, got - %v", err)
	}
}
