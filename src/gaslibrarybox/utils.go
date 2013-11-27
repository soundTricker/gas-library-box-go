package gaslibrarybox

import (
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	"appengine/user"
	"errors"
	"github.com/soundTricker/go-endpoints/endpoints"

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
	} else {
		c.Errorf(err.Error())
	}
}

type QueryMarker struct {
	datastore.Cursor
}

func (qm *QueryMarker) MarshalJSON() ([]byte, error) {
	return []byte(`"` + qm.String() + `"`), nil
}

func (qm *QueryMarker) UnmarshalJSON(buf []byte) error {
	if len(buf) < 2 || buf[0] != '"' || buf[len(buf)-1] != '"' {
		return errors.New("QueryMarker: bad cursor value")
	}
	cursor, err := datastore.DecodeCursor(string(buf[1 : len(buf)-1]))
	if err != nil {
		return err
	}
	*qm = QueryMarker{cursor}
	return nil
}
