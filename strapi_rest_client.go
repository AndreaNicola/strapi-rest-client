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
	BaseUrl string
}

func New() StrapiRestClient {
	return NewWithUlr(os.Getenv("STRAPI_BASE_URL"))
}

func NewWithUlr(baseUrl string) (src StrapiRestClient) {

	if baseUrl == "" {
		panic("STRAPI BASE URL IS MANDATORY")
	}

	src.BaseUrl = baseUrl
	return

}

type callback func(map[string]interface{})

func call(method, url string, callback callback) error {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var bodyMap map[string]interface{}
	err = json.Unmarshal(body, &bodyMap)

	if err != nil {
		return err
	}

	callback(bodyMap)

	return nil

}

func (src StrapiRestClient) GetUser(id int) (*StrapiUser, error) {
	var result StrapiUser
	err := call("GET", src.BaseUrl+"/users/"+strconv.Itoa(id), result.New)
	return &result, err
}

func (src StrapiRestClient) GetProduct(id int) (*StrapiProduct, error) {
	var result StrapiProduct
	err := call("GET", src.BaseUrl+"/products/"+strconv.Itoa(id), result.New)
	return &result, err
}
