package rest

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
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
	payload, _ = ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	// if err != nil {
	// 	e = errors.New(util.CannotReadPayloadMsg)
	// 	util.Logger.ERROR.Log(fmt.Sprintf(util.CannotReadPayload, err))
	// 	return
	// }
	// if len(payload) == 0 {
	// 	e = errors.New(util.PayloadMissing)
	// 	util.Logger.ERROR.Log(util.PayloadMissing)
	// 	return
	// }
	return
}
