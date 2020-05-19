package poste

import (
	"github.com/go-resty/resty/v2"
)

func (api PosteAPI) RawListDomains() (*resty.Response, error) {
	return api.get("/v1/domains", map[string]string{})
}
