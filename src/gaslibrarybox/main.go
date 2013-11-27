package gaslibrarybox

import (
	"github.com/soundTricker/go-endpoints/endpoints"
)

func init() {
	libraryService := &LibraryService{}
	api, err := endpoints.RegisterService(libraryService, "libraries", "v1", "The Google Apps Script Library Box API", true)
	if err != nil {
		panic(err.Error())
	}

	info := api.MethodByName("Get").Info()
	info.Path, info.HttpMethod, info.Name = "{libraryKey}", "GET", "get"
	info.Scopes, info.Audiences, info.ClientIds = Scopes, Audiences, ClientIds

	info = api.MethodByName("Insert").Info()
	info.Path, info.HttpMethod, info.Name = "insert", "POST", "insert"
	info.Scopes, info.Audiences, info.ClientIds = Scopes, Audiences, ClientIds

	memberService := &MemberService{}

	api, err = endpoints.RegisterService(memberService, "members", "v1", "The Google Apps Script Library Box Member API", true)
	if err != nil {
		panic(err.Error())
	}

	info = api.MethodByName("Get").Info()
	info.Path, info.HttpMethod, info.Name = "{memberKey}", "GET", "get"
	info.Scopes, info.Audiences, info.ClientIds = Scopes, Audiences, ClientIds

	info = api.MethodByName("Insert").Info()
	info.Path, info.HttpMethod, info.Name = "insert", "POST", "insert"
	info.Scopes, info.Audiences, info.ClientIds = Scopes, Audiences, ClientIds

	endpoints.HandleHttp()
}
