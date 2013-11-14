package models

import (
	"appengine"
	"appengine/memcache"
	"appengine/user"
	"errors"
	"github.com/crhym3/go-endpoints/endpoints"

	"encoding/json"
)

const clientId = ""

var (
	Scopes    = []string{endpoints.EmailScope}
	ClientIds = []string{clientId, endpoints.ApiExplorerClientId}
	Audiences = []string{clientId}
)

func GetCurrentUser(c endpoints.Context) (*user.User, error) {
	u, err := endpoints.CurrentUser(c, Scopes, Audiences, ClientIds)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("Unauthorized: Please, sign in.")
	}
	c.Debugf("Current user: %#v", u)
	return u, nil
}

func PutEntity2Memcache(c appengine.Context, key string, e interface{}) {
	if b, err := json.Marshal(e); err == nil {
		memcache.Set(c, &memcache.Item{Key: key, Value: b})
	}
}
