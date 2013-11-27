package gaslibrarybox

import (
	"appengine/datastore"
	"encoding/json"
	"errors"
	"github.com/soundTricker/go-endpoints/endpoints"
	"net/http"
)

type MemberService struct {
}

type MemberGetReq struct {
	MemberKey string `json:"memberKey" endpoints:"req,required, desc= The member key"`
}

func (m *MemberService) Get(r *http.Request, req *MemberGetReq, resp *Member) error {
	c := endpoints.NewContext(r)

	if req.MemberKey != "me" {
		k, err := datastore.DecodeKey(req.MemberKey)
		if err != nil {
			return err
		}

		if err := GetMemberByEmail(c, k.StringID(), resp); err != nil {
			return err
		}
		if resp.MemberKey == "" {
			return errors.New("Not Found User")
		}
	}

	u, err := GetCurrentUser(c)
	if err != nil {
		return err
	}

	if err := GetMember(c, u, resp); err != nil {
		return err
	}

	return nil
}

func (m *MemberService) Insert(r *http.Request, req *Member, resp *Member) error {
	c := endpoints.NewContext(r)
	u, err := GetCurrentUser(c)

	if err != nil {
		return err
	}

	if u == nil {
		return errors.New("Unauthorized")
	}

	if err := PutMember(c, u, req); err != nil {
		return err
	}

	b, err := json.Marshal(req)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, resp); err != nil {
		return err
	}

	return nil
}
