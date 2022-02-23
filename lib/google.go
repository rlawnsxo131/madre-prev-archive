package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	googlePeopleURL = "https://people.googleapis.com/v1/people/me?personFields=names,emailAddresses,photos"
)

func GetGoogleProfile(accessToken string) (p interface{}, err error) {
	req, err := http.NewRequest("GET", googlePeopleURL, nil)
	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	profile := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&profile); err != nil {
		panic(err)
	}

	// socialId: data.resourceName?.replace('people/', '') ?? '',
	// email: data.emailAddresses?.[0].value ?? '',
	// photo: data.photos?.[0].url ?? null,
	// displayName: data.names?.[0].displayName?.split(' (')[0] ?? '',

	resourceName := profile["resourceName"]
	emailAddresses := profile["emailAddresses"]
	photos := profile["photos"]
	displayName := profile["names"]

	log.Println(resourceName)
	log.Println(emailAddresses)
	log.Println(photos)
	log.Println(displayName)

	return "", err
}
