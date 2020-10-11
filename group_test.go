package godisco_test

import (
	"testing"

	. "github.com/ForgottenWorld/godisco"
)

func TestGroups(t *testing.T) {
	c, err := NewClient(forum, "", "")
	if err != nil {
		t.Errorf("client creation unexpected failure: %v", err)
		return
	}

	gl, err := GetGroups(c, 0)
	if err != nil {
		t.Errorf("unexpected failure retrieving group list: %v", err)
		return
	}

	for _, g := range gl.Groups {
		n := "TestGroup-" + g.Name
		t.Run(n, func(t *testing.T) {
			gi, err := GetGroup(c, g.Name)
			if err != nil {
				t.Errorf("unexpected failure retrieving %s group: %v", g.Name, err)
				return
			}

			t.Logf("Group %s retrieved", gi.Group.Name)

			if gi.Group.VisibleMembers {
				gm, err := GetGroupMembers(c, g.Name, 0)
				if err != nil {
					t.Errorf("unexpected failure retrieving %s group members: %v", g.Name, err)
					return
				}

				t.Logf("Group %s has %d members", g.Name, len(gm.Members))
			}
		})
	}
}
