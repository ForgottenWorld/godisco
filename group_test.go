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

	for _, group := range gl.Groups {
		n := "TestGroup-" + group.Name
		t.Run(n, func(t *testing.T) {
			gi, err := GetGroup(c, group.Name)
			if err != nil {
				t.Errorf("unexpected failure retrieving %s group: %v", group.Name, err)
				return
			}

			t.Logf("Group %s retrieved", gi.Group.Name)

			if gi.Group.VisibleMembers {
				gm, err := GetGroupMembers(c, group.Name, 0)
				if err != nil {
					t.Errorf("unexpected failure retrieving %s group members: %v", group.Name, err)
					return
				}

				t.Logf("Group %s has %d members", group.Name, len(gm.Members))
			}
		})
	}
}
