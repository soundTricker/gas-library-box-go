package gaslibrarybox

import (
	"appengine/user"
	"time"
)

type Member struct {
	User         *user.User `json:"-" datastore:"user,unindex"`
	MemberKey    string     `json:"memberKey" datastore:"-" endpoints:"req"`
	Nickname     string     `json:"nickname" datastore:"nickname"`
	UserIconUrl  string     `json:"userIconUrl" datastore:"userIconUrl,unindex"`
	Url          string     `json:"url" datastore:"url,unindex"`
	RegisteredAt time.Time  `json:"registeredAt" datastore:"registeredAt"`
	ModifiedAt   time.Time  `json:"modifiedAt" datastore:"modifiedAt"`
}
