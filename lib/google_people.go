package lib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	googlePeopleUrl = "https://people.googleapis.com/v1/people/me?personFields=names,emailAddresses,photos"
)

// {
// 	"resourceName": "",
// 	"etag": "",
// 	"names": [
// 	  {
// 		"metadata": {
// 		  "primary": true,
// 		  "source": {
// 			"type": "PROFILE",
// 			"id": ""
// 		  }
// 		},
// 		"displayName": "",
// 		"familyName": "",
// 		"givenName": "",
// 		"displayNameLastFirst": "",
// 		"unstructuredName": ""
// 	  }
// 	],
// 	"photos": [
// 	  {
// 		"metadata": {
// 		  "primary": true,
// 		  "source": {
// 			"type": "PROFILE",
// 			"id": ""
// 		  }
// 		},
// 		"url": "",
// 		"default": true
// 	  }
// 	],
// 	"emailAddresses": [
// 	  {
// 		"metadata": {
// 		  "primary": true,
// 		  "verified": true,
// 		  "source": {
// 			"type": "ACCOUNT",
// 			"id": ""
// 		  },
// 		  "sourcePrimary": true
// 		},
// 		"value": "email@gmail.com"
// 	  }
// 	]
// }
type RawGoogleProfile struct {
	ResourceName string `json:"resourceName"`
	Etag         string `json:"etag"`
	Names        []struct {
		Metadata struct {
			Primary bool `json:"primary"`
			Source  struct {
				Type string `json:"type"`
				Id   string `json:"id"`
			} `json:"source"`
		} `json:"metadata"`
		DisplayName          string `json:"displayname"`
		FamilyName           string `json:"familyName"`
		GivenName            string `json:"givenName"`
		DisplayNameLastFirst string `json:"displayNameLastFirst"`
		UnstructuredName     string `json:"unstructuredName"`
	}
	Photos []struct {
		MetaData struct {
			Primary bool `json:"primary"`
			Source  struct {
				Type string `json:"type"`
				Id   string `json:"id"`
			} `json:"struct"`
		} `json:"metadata"`
		Url     string `json:"url"`
		Default bool   `json:"default"`
	}
	EmailAddresses []struct {
		Metadata struct {
			Primary  bool `json:"primary"`
			Verified bool `json:"verified"`
			Source   struct {
				Type string `json:"type"`
				Id   string `json:"id"`
			} `json:"struct"`
			SourcePrimary bool `json:"sourcePrimary"`
		} `json:"metadata"`
		Value string `json:"value"`
	} `json:"emailAddresses"`
}

type GoogleProfile struct {
	SocialId    string
	Email       string
	PhotoUrl    string
	DisplayName string
}

type googlePeopleApi struct {
	accessToken string
}

type GooglePeopleApi interface {
	GetGoogleProfile() *GoogleProfile
	createRequest() *http.Request
	excuteRequest(req *http.Request) *http.Response
	convertToRawGoogleProfile(res *http.Response) *RawGoogleProfile
	convertToGoogleProfile(rawProfile *RawGoogleProfile) *GoogleProfile
}

func NewGooglePeopleApi(accessToken string) GooglePeopleApi {
	api := &googlePeopleApi{
		accessToken: accessToken,
	}
	return api
}

func (g *googlePeopleApi) GetGoogleProfile() *GoogleProfile {
	req := g.createRequest()
	res := g.excuteRequest(req)
	defer res.Body.Close()
	rawProfile := g.convertToRawGoogleProfile(res)
	profile := g.convertToGoogleProfile(rawProfile)

	return profile
}

func (g *googlePeopleApi) createRequest() *http.Request {
	req, err := http.NewRequest("GET", googlePeopleUrl, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+g.accessToken)
	return req
}

func (g *googlePeopleApi) excuteRequest(req *http.Request) *http.Response {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return res
}

func (g *googlePeopleApi) convertToRawGoogleProfile(res *http.Response) *RawGoogleProfile {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var rawGoogleProfile RawGoogleProfile
	if err := json.Unmarshal(body, &rawGoogleProfile); err != nil {
		panic(err)
	}
	return &rawGoogleProfile
}

func (g *googlePeopleApi) convertToGoogleProfile(rawProfile *RawGoogleProfile) *GoogleProfile {
	// socialId: data.resourceName?.replace('people/', '') ?? '',
	// email: data.emailAddresses?.[0].value ?? '',
	// photoUrl: data.photos?.[0].url ?? null,
	// displayName: data.names?.[0].displayName?.split(' (')[0] ?? '',

	var socialId string
	var email string
	var photoUrl string
	var displayName string

	if rawProfile.ResourceName != "" {
		replaceResourceName := strings.ReplaceAll(rawProfile.ResourceName, "people/", "")
		if replaceResourceName != "" {
			socialId = replaceResourceName
		}
	}

	if len(rawProfile.EmailAddresses) > 0 {
		value := rawProfile.EmailAddresses[0].Value
		if value != "" {
			email = value
		}
	}

	if len(rawProfile.Photos) > 0 {
		url := rawProfile.Photos[0].Url
		if url != "" {
			photoUrl = url
		}
	}

	if len(rawProfile.Names) > 0 {
		name := strings.Split(rawProfile.Names[0].DisplayName, " ")[0]
		if name != "" {
			displayName = name
		}
	}

	return &GoogleProfile{
		SocialId:    socialId,
		Email:       email,
		PhotoUrl:    photoUrl,
		DisplayName: displayName,
	}
}
