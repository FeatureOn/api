package middleware

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/comm/rest/dto"
	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// FromJSON deserializes the object from JSON string
// in an io.Reader to the given interface
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}

func readPayload(r *http.Request) (payload []byte, e error) {
	payload, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		e = errors.New(viper.GetString("CannotReadPayloadMsg"))
		log.Error().Err(err).Msg(viper.GetString("CannotReadPayloadMsg"))
		return
	}
	if len(payload) == 0 {
		e = errors.New(viper.GetString("PayloadMissingMsg"))
		log.Error().Err(err).Msg(viper.GetString("PayloadMissingMsg"))
		return
	}
	return
}

// ExtractUserPayload extracts user data from the request body
// Returns UserRequest model if found, error otherwise
func ExtractUserPayload(r *http.Request) (user *dto.UserRequest, e error) {
	payload, e := readPayload(r)
	if e != nil {
		return
	}
	err := json.Unmarshal(payload, &user)
	if err != nil {
		e = errors.New(viper.GetString("CannotParsePayloadMsg"))
		log.Error().Err(err).Msg(viper.GetString("CannotParsePayloadMsg"))
		return
	}
	return
}

// ExtractLoginPayload extracts login data from the request body
// Returns LoginRequest model if found, error otherwise
func ExtractLoginPayload(r *http.Request) (login *dto.LoginRequest, e error) {
	payload, e := readPayload(r)
	if e != nil {
		return
	}
	err := json.Unmarshal(payload, &login)
	if err != nil {
		e = errors.New(viper.GetString("CannotParsePayloadMsg"))
		log.Error().Err(err).Msg(viper.GetString("CannotParsePayloadMsg"))
		return
	}
	return
}