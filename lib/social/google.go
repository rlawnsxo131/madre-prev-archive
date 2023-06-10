package social

import (
	"encoding/json"
	"io"
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

type googlePeopleApiResponse struct {
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

type googlePeopleProfile struct {
	SocialID    string
	Email       string
	PhotoUrl    *string
	DisplayName *string
}

type googlePeopleAPI struct {
	accessToken string
}

func NewGooglePeopleAPI(actk string) *googlePeopleAPI {
	return &googlePeopleAPI{
		accessToken: actk,
	}
}

func (g *googlePeopleAPI) Do() (*googlePeopleProfile, error) {
	req, err := g.createRequest(g.accessToken)
	if err != nil {
		return nil, err
	}

	res, err := g.excuteRequest(req)
	if err != nil {
		return nil, err
	}

	gapiRes, err := g.mapToGooglePeopleApiResponse(res)
	if err != nil {
		return nil, err
	}

	return g.mapToGooglePeopleProfile(gapiRes), nil
}

func (g *googlePeopleAPI) createRequest(accessToken string) (*http.Request, error) {
	url := "https://people.googleapis.com/v1/people/me?personFields=names,emailAddresses,photos"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "social googlePeopleAPI createRequest")
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	return req, nil
}

func (g *googlePeopleAPI) excuteRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "social googlePeopleAPI excuteRequest")
	}
	return res, nil
}

func (g *googlePeopleAPI) mapToGooglePeopleApiResponse(res *http.Response) (*googlePeopleApiResponse, error) {
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "social googlePeopleAPI mapTogooglePeopleApiResponse res.Body read")
	}

	var gapiRes googlePeopleApiResponse
	if err := json.Unmarshal(body, &gapiRes); err != nil {
		return nil, errors.Wrap(err, "social googlePeopleAPI mapTogooglePeopleApiResponse Unmarshal")
	}
	return &gapiRes, nil
}

func (g *googlePeopleAPI) mapToGooglePeopleProfile(gapiRes *googlePeopleApiResponse) *googlePeopleProfile {
	var socialId string
	var email string
	var photoUrl *string
	var displayName *string

	// uniq key
	socialId = strings.ReplaceAll(
		gapiRes.ResourceName,
		"people/",
		"",
	)

	if len(gapiRes.EmailAddresses) > 0 {
		value := gapiRes.EmailAddresses[0].Value
		if value != "" {
			email = value
		}
	}

	if len(gapiRes.Photos) > 0 {
		url := gapiRes.Photos[0].Url
		if len(url) != 0 {
			photoUrl = &url
		}
	}

	if len(gapiRes.Names) > 0 {
		name := strings.Split(gapiRes.Names[0].DisplayName, " ")[0][:50]
		if name != "" {
			displayName = &name
		}
	}

	return &googlePeopleProfile{
		SocialID:    socialId,
		Email:       email,
		PhotoUrl:    photoUrl,
		DisplayName: displayName,
	}
}
