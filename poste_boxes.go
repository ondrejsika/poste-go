package poste

import (
	"fmt"
)

func (api PosteAPI) CreateBox(email string, passwordPlaintext string) error {
	rawResp, err := api.RawCreateBox(email, passwordPlaintext)
	if err != nil {
		return err
	}
	if rawResp.StatusCode() == 201 {
		return nil
	}
	return fmt.Errorf("Err")
}

func (api PosteAPI) DeleteBox(email string) error {
	rawResp, err := api.RawDeleteBox(email)
	if err != nil {
		return err
	}
	if rawResp.StatusCode() == 204 {
		return nil
	}
	return fmt.Errorf("Err")
}
