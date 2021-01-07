package pushover

import (
	"strings"
	"testing"
)

func TestPushover(t *testing.T) {
	var want string = "\"success\": 1"
	if got := sendMessage("Title", "Message", false, "", ""); strings.Contains(got, want) {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
