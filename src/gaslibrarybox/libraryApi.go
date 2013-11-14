package gaslibrarybox

import (
	"appengine/datastore"
	"errors"
	"fmt"
	"gaslibrarybox/models"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

type LibraryList struct {
	Items []*models.Library `json:items`
}

type LibraryService struct {
}

type LibraryGetReq struct {
	LibraryKey string `json:"libraryKey" endpoints:"req,desc=library's project key"`
}

// defined with "/library/{libraryKey}" path template
func (l *LibraryService) Get(r *http.Request, req *LibraryGetReq, resp *models.Library) error {
	c := endpoints.NewContext(r)

	err := models.GetLibrary(c, req.LibraryKey, resp)

	switch err {
	case datastore.ErrNoSuchEntity:
		return fmt.Errorf("Not Found Not found key:%s library", req.LibraryKey)
	case nil:
	default:
		return err
	}
	return nil
}

func (ls *LibraryService) Insert(r *http.Request, req *models.Library, resp *models.Library) error {
	c := endpoints.NewContext(r)

	err := models.PutLibrary(c, req)

	switch err {
	case models.DuplicateEntity:
		return errors.New("Duplicate")
	case err:
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
