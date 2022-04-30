package social

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/utils"
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

type googlePeopleProfileOriginal struct {
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
	SocialId    string
	Email       string
	PhotoUrl    string
	DisplayName string
}

func GetGoogleProfile(accessToken string) (*googlePeopleProfile, error) {
	req, err := createRequest(accessToken)
	if err != nil {
		err = errors.Wrap(
			err,
			"GetPeopleProfile: createRequest error",
		)
		return nil, err
	}

	res, err := excuteRequest(req)
	if err != nil {
		err = errors.Wrap(
			err,
			"GetPeopleProfile: excuteRequest error",
		)
		return nil, err
	}

	originProfile, err := mapToGooglePeopleProfileOriginal(res)
	if err != nil {
		err = errors.Wrap(
			err,
			"GetPeopleProfile: convertToRawPeopleProfile error",
		)
		return nil, err
	}

	googlePeopleProfile := mapToGooglePeopleProfile(originProfile)

	return googlePeopleProfile, nil
}

func createRequest(accessToken string) (*http.Request, error) {
	url := "https://people.googleapis.com/v1/people/me?personFields=names,emailAddresses,photos"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	return req, nil
}

func excuteRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func mapToGooglePeopleProfileOriginal(res *http.Response) (*googlePeopleProfileOriginal, error) {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var rawProfile googlePeopleProfileOriginal
	if err := json.Unmarshal(body, &rawProfile); err != nil {
		return nil, err
	}

	return &rawProfile, nil
}

func mapToGooglePeopleProfile(originProfile *googlePeopleProfileOriginal) *googlePeopleProfile {
	var socialId string
	var email string
	var photoUrl string
	var displayName string

	// must not be null
	if originProfile.ResourceName != "" {
		replaceResourceName := strings.ReplaceAll(originProfile.ResourceName, "people/", "")
		if replaceResourceName != "" {
			socialId = replaceResourceName
		} else {
			socialId = utils.GenerateUUIDString()
		}
	} else {
		socialId = utils.GenerateUUIDString()
	}

	if len(originProfile.EmailAddresses) > 0 {
		value := originProfile.EmailAddresses[0].Value
		if value != "" {
			email = value
		}
	}

	if len(originProfile.Photos) > 0 {
		url := originProfile.Photos[0].Url
		if len(url) != 0 {
			photoUrl = url
		}
	}

	if len(originProfile.Names) > 0 {
		name := strings.Split(originProfile.Names[0].DisplayName, " ")[0]
		if name != "" {
			displayName = name
		}
	}

	return &googlePeopleProfile{
		SocialId:    socialId,
		Email:       email,
		PhotoUrl:    photoUrl,
		DisplayName: displayName,
	}
}
