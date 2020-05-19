package poste

import (
	"github.com/go-resty/resty/v2"
)

type PosteAPI struct {
	origin   string
	username string
	password string
}

func New(origin string, username string, password string) PosteAPI {
	api := PosteAPI{origin, username, password}
	return api
}

func (api PosteAPI) get(url string, query map[string]string) (*resty.Response, error) {
	client := resty.New()
	return client.R().SetBasicAuth(api.username, api.password).SetQueryParams(query).Get(api.origin + url)
}

func (api PosteAPI) post(url string, body map[string]interface{}) (*resty.Response, error) {
	client := resty.New()
	return client.R().SetBasicAuth(api.username, api.password).SetBody(body).Post(api.origin + url)
}

func (api PosteAPI) delete(url string) (*resty.Response, error) {
	client := resty.New()
	return client.R().SetBasicAuth(api.username, api.password).Delete(api.origin + url)
}
