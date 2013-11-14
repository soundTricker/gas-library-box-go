package gaslibrarybox

import (
	"github.com/crhym3/go-endpoints/endpoints"
)

func init() {
	libraryService := &LibraryService{}
	api, err := endpoints.RegisterService(libraryService, "libraries", "v1", "The Google Apps Script Library Box API", true)
	if err != nil {
		panic(err.Error())
	}

	info := api.MethodByName("Get").Info()
	info.Path, info.HttpMethod, info.Name = "{libraryKey}", "GET", "get"

	info = api.MethodByName("Insert").Info()
	info.Path, info.HttpMethod, info.Name = "insert", "POST", "insert"

	endpoints.HandleHttp()
}
