package rest

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

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
