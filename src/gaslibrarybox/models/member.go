package models

import (
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	"appengine/user"
	"encoding/json"
	"fmt"
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

const memberKind = "Member"
const memberMemcacheKey = "Member_%s"

func GetMember(c appengine.Context, u *user.User, m *Member) error {
	item, err := memcache.Get(c, fmt.Sprintf(memberMemcacheKey, u.Email))
	switch err {
	case nil:
		json.Unmarshal(item.Value, m)
		return nil
	case memcache.ErrCacheMiss:
	default:
		return err
	}
	k := datastore.NewKey(c, memberKind, u.Email, 0, nil)
	if err := datastore.Get(c, k, m); err != nil {
		return err
	}
	m.MemberKey = k.StringID()

	putMember2Cache(c, m)

	return nil
}

func putMember2Cache(c appengine.Context, m *Member) {
	PutEntity2Memcache(c, fmt.Sprintf(memberMemcacheKey, m.MemberKey), m)
}

func PutMember(c appengine.Context, u *user.User, m *Member) error {

}
