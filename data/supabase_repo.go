package data

import (
	"TheLazyLemur/openfit/config"
	"fmt"
	"io/ioutil"
	"net/http"
)

func LoadAll() {
	key := config.GoDotEnvVariable("key")
	url := config.GoDotEnvVariable("url")

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("apikey", key)
	req.Header.Add("Authorization", "Bearer "+key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("request failed")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func LoadAllBodyPart(bodyPart string) {
	key := config.GoDotEnvVariable("key")
	url := config.GoDotEnvVariable("url")

	req, _ := http.NewRequest("GET", url+"?bodyPart=eq."+bodyPart, nil)

	req.Header.Add("apikey", key)
	req.Header.Add("Authorization", "Bearer "+key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("request failed")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func LoadById(id string) {
	key := config.GoDotEnvVariable("key")
	url := config.GoDotEnvVariable("url")

	req, err := http.NewRequest("GET", url+"?id=eq."+id, nil)
	if err != nil {
		panic("failed creating request")
	}

	req.Header.Add("apikey", key)
	req.Header.Add("Authorization", "Bearer "+key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("request failed")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("failed reading request body")
	}

	fmt.Println(string(body))
}
