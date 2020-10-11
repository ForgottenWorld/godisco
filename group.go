package godisco

import (
	"encoding/json"
	"fmt"
)

// Group information about a group
type Group struct {
	ID             int    `json:"id"`
	Automatic      bool   `json:"automatic"`
	Name           string `json:"name"`
	UserCount      int    `json:"user_count"`
	Primary        bool   `json:"primary_group"`
	Title          string `json:"title"`
	TrustLevel     int    `json:"grant_trust_level"`
	Mentionable    bool   `json:"mentionable"`
	VisibleMembers bool   `json:"can_see_members"`
}

// GroupResponse expected struct for a group
type GroupResponse struct {
	Group Group `json:"group"`
}

// GroupList expected struct for groups list
type GroupList struct {
	Groups []Group
}

// GroupMembersResponse defines list of members in a group
type GroupMembersResponse struct {
	Members []Member `json:"members"`
	Owners  []Member `json:"owners"`
	Meta    struct {
		Total  int `json:"total"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
}

// Member information about a member
type Member struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar_template"`
	Name     string `json:"name"`
	Title    string `json:"title"`
	Posted   string `json:"last_posted_at"`
	Seen     string `json:"last_seen_at"`
}

// GetGroup show details of a given group
func GetGroup(req Requester, groupName string) (*GroupResponse, error) {
	endpoint := fmt.Sprintf("/groups/%s.json", groupName)
	body, _, err := req.Get(endpoint)
	if err != nil {
		return nil, err
	}
	var gr GroupResponse
	err = json.Unmarshal(body, &gr)
	return &gr, err
}

// GetGroups list some groups
func GetGroups(req Requester, page int) (*GroupList, error) {
	endpoint := fmt.Sprintf("/groups.json?page=%d", page)
	body, _, err := req.Get(endpoint)
	if err != nil {
		return nil, err
	}
	var groups GroupList
	err = json.Unmarshal(body, &groups)
	return &groups, err
}

// GetGroupMembers list some members of a given group
func GetGroupMembers(req Requester, groupName string, page int) (*GroupMembersResponse, error) {
	endpoint := fmt.Sprintf("/groups/%s/members.json?page=%d", groupName, page)
	body, _, err := req.Get(endpoint)
	if err != nil {
		return nil, err
	}
	var gmr GroupMembersResponse
	err = json.Unmarshal(body, &gmr)
	return &gmr, err
}

//@TODO
// Implement add to group - Name: "DockerForMacWinBeta" - id: 45
// APi Call: POST - '%s/admin/groups/%s?api_key=%s&api_username=%s'
// alias_level = 0

type groupUpdate struct {
	GroupID string   `json:"group_id"`
	Users   []string `json:"users"`
}

// GroupInfo describes the group update received
type GroupInfo struct {
	Basic struct {
		ID           int    `json:"id"`
		Automatic    bool   `json:"automatic"`
		Name         string `json:"name"`
		UserCount    int    `json:"user_count"`
		AliasLevel   int    `json:"alias_level"`
		Visible      bool   `json:"visible"`
		Domains      string `json:"automatic_membership_email_domains"`
		Retroactive  bool   `json:"automatic_membership_retroactive"`
		Primary      bool   `json:"primary_group"`
		Title        string `json:"title"`
		Trust        string `json:"grant_trust_level"`
		Incoming     string `json:"incoming_email"`
		Notification int    `json:"notification_level"`
		Messages     bool   `json:"has_messages"`
		Mentionable  bool   `json:"mentionable"`
	} `json:"basic_group"`
}

func updateGroupMembers(req Requester, groupName string, groupID string, members []string) (groupInfo *GroupInfo, err error) {
	update := &groupUpdate{
		GroupID: groupID,
		Users:   members,
	}
	data, err := json.Marshal(update)
	endpoint := "/admin/groups/bulk"
	body, _, err := req.Post(endpoint, data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &groupInfo)
	return groupInfo, err
}
