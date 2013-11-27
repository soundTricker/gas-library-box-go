package gaslibrarybox

import (
	"appengine/datastore"
	"fmt"
	"github.com/soundTricker/go-endpoints/endpoints"
	"net/http"
	"strings"
)

type LibraryList struct {
	Items []*Library   `json:items`
	Next  *QueryMarker `json:"next,omitempty"`
}

type LibraryService struct {
}

type LibraryGetReq struct {
	LibraryKey string `json:"libraryKey" endpoints:"req,required,desc=library's project key"`
}

type LibraryListReq struct {
	Limit int          `json:"limit,string" endpoints:"d=100,max=200"`
	Page  *QueryMarker `json:"cursor"`
}

func (l *LibraryService) List(r *http.Request, req *LibraryListReq, resp *LibraryList) {
	c := endpoints.NewContext(r)

	limit := 100

	if req.Limit > 0 && req.Limit <= 200 {
		limit = req.Limit
	}

	resp.Items = make([]*Library, 0, limit)

	q := datastore.NewQuery("Library").Limit(limit)

	if req.Page != nil {
		q = q.Start(req.Page.Cursor)
	}

	var iter *datastore.Iterator
	for iter := q.Run(c); ; {
		var l Library
		key, err := iter.Next(&l)
		if err == datastore.Done {
			break
		}
		if err != nil {
			return err
		}
		l.LibraryKey = key.StringID()
		resp.Items = append(resp.Items, &l)
	}

	cur, err := iter.Cursor()
	if err != nil {
		return err
	}

	resp.Next = &QueryMarker{cur}
	return nil

}

// defined with "/library/{libraryKey}" path template
func (l *LibraryService) Get(r *http.Request, req *LibraryGetReq, resp *Library) error {
	c := endpoints.NewContext(r)

	if req.LibraryKey == "" {
		return endpoints.NewBadRequestError("Library key is required")
	}

	err := GetLibrary(c, req.LibraryKey, resp)

	switch err {
	case datastore.ErrNoSuchEntity:
		return endpoints.NewNotFoundError(fmt.Sprintf("Not found key:%s library", req.LibraryKey))
	case nil:
	default:
		return err
	}
	return nil
}

func (ls *LibraryService) Insert(r *http.Request, req *Library, resp *Library) error {

	c := endpoints.NewContext(r)

	if req == nil {
		return endpoints.NewBadRequestError("Need JSON Body")
	}

	if req.LibraryKey == "" || req.Label == "" || req.SourceUrl == "" {
		return endpoints.NewBadRequestError("Need libraryKey, label, sourceUrl")
	}

	if !strings.HasPrefix(req.SourceUrl, "https://") && !strings.HasPrefix(req.SourceUrl, "http://") {
		return endpoints.NewBadRequestError("The source url require https:// or http://")
	}

	u, err := GetCurrentUser(c)
	if err != nil || u == nil {
		return endpoints.NewUnauthorizedError("Given login2")
	}

	m := &Member{}

	if err := GetMember(c, u, m); err != nil {
		if err != datastore.ErrNoSuchEntity {
			return err
		} else {
			return endpoints.NewUnauthorizedError("Given register to site")
		}
	}

	err = PutLibrary(c, m, req)

	switch err {
	case DuplicateEntity:
		return endpoints.NewConflictError(fmt.Sprintf("Key %s have been exist.", req.LibraryKey))
	case nil:
	default:
		return err
	}

	resp.LibraryKey = req.LibraryKey
	resp.Label = req.Label
	resp.Desc = req.Desc
	resp.LongDesc = req.LongDesc
	resp.SourceUrl = req.SourceUrl
	resp.RegisteredAt = req.RegisteredAt
	resp.ModifiedAt = req.ModifiedAt
	resp.AuthorName = req.AuthorName
	resp.AuthorUrl = req.AuthorUrl
	resp.AuthorIconUrl = req.AuthorIconUrl
	resp.AuthorKey = req.AuthorKey

	return nil

}
