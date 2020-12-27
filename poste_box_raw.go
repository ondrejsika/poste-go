package poste

import (
	"github.com/go-resty/resty/v2"
)

func (api PosteAPI) RawCreateBox(email string, passwordPlaintext string) (*resty.Response, error) {
	return api.post("/v1/boxes", map[string]interface{}{"email": email, "passwordPlaintext": passwordPlaintext})
}

func (api PosteAPI) RawGetBoxSieve(email string) (*resty.Response, error) {
	return api.get("/v1/boxes/"+email+"/sieve", map[string]string{})
}

func (api PosteAPI) RawDeleteBox(email string) (*resty.Response, error) {
	return api.delete("/v1/boxes/" + email)
}
