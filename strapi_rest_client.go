package strapi_rest_client

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type StrapiRestClient struct {
	baseUrl string
}

func New() StrapiRestClient {
	return NewWithUlr(os.Getenv("STRAPI_BASE_URL"))
}

func NewWithUlr(baseUrl string) (src StrapiRestClient) {

	if baseUrl == "" {
		panic("STRAPI BASE URL IS MANDATORY")
	}

	src.baseUrl = baseUrl
	return

}

type callback func(map[string]interface{})

func call(method, url string, callback callback) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var bodyMap map[string]interface{}
	err = json.Unmarshal(body, &bodyMap)

	if err != nil {
		panic(err)
	}

	callback(bodyMap)

}


func (src StrapiRestClient) GetUser(id int) *StrapiUser {
	var result StrapiUser
	call("GET", src.baseUrl+"/users/"+strconv.Itoa(id), result.New)
	return &result
}

func (src StrapiRestClient) GetProduct(id int) *StrapiProduct {
	var result StrapiProduct
	call("GET", src.baseUrl+"/products/"+strconv.Itoa(id), result.New)
	return &result
}