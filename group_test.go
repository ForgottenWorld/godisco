package godisco

import (
	"testing"
)

func TestGroups(t *testing.T) {
	c, err := NewClient(forum, "", "")
	if err != nil {
		t.Errorf("client creation unexpected failure: %v", err)
		return
	}

	gl, err := GetGroups(c, 0)

	for _, v := range gl.Groups {
		n := "TestGroup-" + v.Name
		t.Run(n, func(t *testing.T) {
			gi, err := GetGroup(c, v.Name)
			if err != nil {
				t.Errorf("unexpected failure retrieving %s group: %v", v.Name, err)
				return
			}

			t.Logf("Group %s retrieved", gi.Group.Name)

			if gi.Group.VisibleMembers {
				gm, err := GetGroupMembers(c, v.Name, 0)
				if err != nil {
					t.Errorf("unexpected failure retrieving %s group members: %v", v.Name, err)
					return
				}

				t.Logf("Group %s has %d members", v.Name, len(gm.Members))
			}
		})
	}
}
