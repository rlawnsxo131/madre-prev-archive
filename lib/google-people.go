package lib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// Google people api Response
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
	} `json:"names"`
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
	} `json:"photos"`
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

type GooglePeopleApi interface {
	GetGoogleProfile() (*GoogleProfile, error)
	createRequest() (*http.Request, error)
	excuteRequest(req *http.Request) (*http.Response, error)
	convertToRawGoogleProfile(res *http.Response) (*RawGoogleProfile, error)
	convertToGoogleProfile(rawProfile *RawGoogleProfile) *GoogleProfile
}

type googlePeopleApi struct {
	accessToken string
}

func NewGooglePeopleApi(accessToken string) GooglePeopleApi {
	api := &googlePeopleApi{
		accessToken: accessToken,
	}
	return api
}

func (g *googlePeopleApi) GetGoogleProfile() (*GoogleProfile, error) {
	req, err := g.createRequest()
	if err != nil {
		err = errors.Wrap(err, "GetGoogleProfile: createRequest error")
		return nil, err
	}

	res, err := g.excuteRequest(req)
	if err != nil {
		err = errors.Wrap(err, "GetGoogleProfile: excuteRequest error")
		return nil, err
	}

	rawProfile, err := g.convertToRawGoogleProfile(res)
	if err != nil {
		err = errors.Wrap(err, "GetGoogleProfile: convertToRawGoogleProfile error")
		return nil, err
	}

	profile := g.convertToGoogleProfile(rawProfile)

	return profile, nil
}

func (g *googlePeopleApi) createRequest() (*http.Request, error) {
	url := "https://people.googleapis.com/v1/people/me?personFields=names,emailAddresses,photos"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+g.accessToken)
	return req, nil
}

func (g *googlePeopleApi) excuteRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (g *googlePeopleApi) convertToRawGoogleProfile(res *http.Response) (*RawGoogleProfile, error) {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var rawProfile RawGoogleProfile
	if err := json.Unmarshal(body, &rawProfile); err != nil {
		return nil, err
	}

	return &rawProfile, nil
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
