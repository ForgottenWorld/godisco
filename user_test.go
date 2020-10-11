package godisco

import (
	"testing"
)

func TestUsers(t *testing.T) {
	c, err := NewClient(forum, "", "")
	if err != nil {
		t.Errorf("client creation unexpected failure: %v", err)
		return
	}

	ur, err := GetUser(c, "codinghorror")
	if err != nil {
		t.Errorf("unexpected failure retrieving user: %v", err)
		return
	}

	t.Logf("User %s retrieved", ur.User.Username)
}
