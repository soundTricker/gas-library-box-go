package gaslibrarybox

import (
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	"appengine/user"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Member struct {
	MemberKey    string    `json:"memberKey" datastore:"-" endpoints:"req"`
	Nickname     string    `json:"nickname" datastore:"nickname"`
	UserIconUrl  string    `json:"userIconUrl" datastore:"userIconUrl,unindex"`
	Url          string    `json:"url" datastore:"url,unindex"`
	RegisteredAt time.Time `json:"registeredAt" datastore:"registeredAt"`
	ModifiedAt   time.Time `json:"modifiedAt" datastore:"modifiedAt"`
}

const memberKind = "Member"
const memberMemcacheKey = "Member_%s"

func GetMember(c appengine.Context, u *user.User, m *Member) error {
	return GetMemberByEmail(c, u.Email, m)
}

func GetMemberByEmail(c appengine.Context, email string, m *Member) error {
	item, err := memcache.Get(c, fmt.Sprintf(memberMemcacheKey, email))
	switch err {
	case nil:
		json.Unmarshal(item.Value, m)
		return nil
	case memcache.ErrCacheMiss:
	default:
		return err
	}
	k := datastore.NewKey(c, memberKind, email, 0, nil)
	if err := datastore.Get(c, k, m); err != nil {
		return err
	}
	m.MemberKey = k.Encode()

	putMember2Cache(c, email, m)

	return nil
}

func putMember2Cache(c appengine.Context, email string, m *Member) {
	PutEntity2Memcache(c, fmt.Sprintf(memberMemcacheKey, email), m)
}

func PutMember(c appengine.Context, u *user.User, m *Member) error {

	if err := GetMember(c, u, &Member{}); err != datastore.ErrNoSuchEntity {
		return errors.New("Duplicate")
	}

	if m.Nickname == "" || m.Url == "" {
		return errors.New("Bad Reuest")
	}

	k := datastore.NewKey(c, memberKind, u.Email, 0, nil)

	m.ModifiedAt = time.Now()
	m.RegisteredAt = time.Now()
	m.MemberKey = k.Encode()
	_, err := datastore.Put(c, k, m)

	if err != nil {
		return err
	}

	putMember2Cache(c, u.Email, m)

	return nil
}
