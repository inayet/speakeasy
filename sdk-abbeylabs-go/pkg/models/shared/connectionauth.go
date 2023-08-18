// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type ConnectionAuthType string

const (
	ConnectionAuthTypeOauth2 ConnectionAuthType = "Oauth2"
)

type ConnectionAuth struct {
	Oauth2Flow *Oauth2Flow

	Type ConnectionAuthType
}

func CreateConnectionAuthOauth2(oauth2 Oauth2Flow) ConnectionAuth {
	typ := ConnectionAuthTypeOauth2
	typStr := ConnectionAuthTypeEnum(typ)
	oauth2.ConnectionAuthTypeEnum = &typStr

	return ConnectionAuth{
		Oauth2Flow: &oauth2,
		Type:       typ,
	}
}

func (u *ConnectionAuth) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	type discriminator struct {
		ConnectionAuthTypeEnum string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.ConnectionAuthTypeEnum {
	case "Oauth2":
		d = json.NewDecoder(bytes.NewReader(data))
		oauth2Flow := new(Oauth2Flow)
		if err := d.Decode(&oauth2Flow); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.Oauth2Flow = oauth2Flow
		u.Type = ConnectionAuthTypeOauth2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ConnectionAuth) MarshalJSON() ([]byte, error) {
	if u.Oauth2Flow != nil {
		return json.Marshal(u.Oauth2Flow)
	}

	return nil, nil
}
