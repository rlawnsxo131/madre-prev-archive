package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	googlePeopleURL = "https://people.googleapis.com/v1/people/me?personFields=names,emailAddresses,photos"
)

func GetGoogleProfile(accessToken string) *profile {
	req := createRequest(accessToken)
	res := excuteGoogleProfileRequest(req)
	mapData := convertToMap(res)
	googleProfileStruct := generateGoogleProfileStruct(mapData)

	return googleProfileStruct
}

func createRequest(accessToken string) *http.Request {
	req, err := http.NewRequest("GET", googlePeopleURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	return req
}

func excuteGoogleProfileRequest(req *http.Request) *http.Response {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	return res
}

func convertToMap(res *http.Response) map[string]interface{} {
	mapData := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&mapData); err != nil {
		panic(err)
	}
	return mapData
}

type profile struct {
	socialId    string
	email       string
	photo       string
	displayName string
}

func generateGoogleProfileStruct(mapData map[string]interface{}) *profile {
	resourceName := mapData["resourceName"]
	emailAddresses := mapData["emailAddresses"]
	photos := mapData["photos"]
	displayName := mapData["names"]

	// socialId: data.resourceName?.replace('people/', '') ?? '',
	// email: data.emailAddresses?.[0].value ?? '',
	// photo: data.photos?.[0].url ?? null,
	// displayName: data.names?.[0].displayName?.split(' (')[0] ?? '',

	log.Println(resourceName)
	log.Println(emailAddresses)
	log.Println(photos)
	log.Println(displayName)

	return &profile{}
}
