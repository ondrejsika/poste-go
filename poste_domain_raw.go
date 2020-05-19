package poste

import (
	"github.com/go-resty/resty/v2"
)

func (api PosteAPI) RawListDomains() (*resty.Response, error) {
	return api.get("/v1/domains", map[string]string{})
}

func (api PosteAPI) RawCreateDomain(name string) (*resty.Response, error) {
	return api.post("/v1/domains", map[string]interface{}{"name": name})
}

func (api PosteAPI) RawDeleteDomain(name string) (*resty.Response, error) {
	return api.delete("/v1/domains/" + name)
}
