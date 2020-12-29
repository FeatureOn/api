package middleware

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/FeatureOn/api/adapters/comm/rest/dto"
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

// ExtractAddUserPayload extracts user data from the request body
// Returns UserRequest model if found, error otherwise
func ExtractAddUserPayload(r *http.Request) (user *dto.AddUserRequest, e error) {
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

// ExtractAddEnvironmentPayload extracts AddEnvironmentRequest data from the request body
// Returns AddEnvironmentRequest model if found, error otherwise
func ExtractAddEnvironmentPayload(r *http.Request) (env *dto.AddEnvironmentRequest, e error) {
	payload, e := readPayload(r)
	if e != nil {
		return
	}
	err := json.Unmarshal(payload, &env)
	if err != nil {
		e = errors.New(viper.GetString("CannotParsePayloadMsg"))
		log.Error().Err(err).Msg(viper.GetString("CannotParsePayloadMsg"))
		return
	}
	return
}

// ExtractAddProductPayload extracts AddProductRequest data from the request body
// Returns AddProductRequest model if found, error otherwise
func ExtractAddProductPayload(r *http.Request) (prod *dto.AddProductRequest, e error) {
	payload, e := readPayload(r)
	if e != nil {
		return
	}
	err := json.Unmarshal(payload, &prod)
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

// ExtractAddFeaturePayload extracts AddFeature data from the request body
// Returns AddFeatureRequest model if found, error otherwise
func ExtractAddFeaturePayload(r *http.Request) (env *dto.AddFeatureRequest, e error) {
	payload, e := readPayload(r)
	if e != nil {
		return
	}
	err := json.Unmarshal(payload, &env)
	if err != nil {
		e = errors.New(viper.GetString("CannotParsePayloadMsg"))
		log.Error().Err(err).Msg(viper.GetString("CannotParsePayloadMsg"))
		return
	}
	return
}
