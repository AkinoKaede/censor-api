package sensitivewords

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/studentmain/censor"
)

type Resp struct {
	StatusCode     int      `json:"-"`
	SensitiveWords []string `json:"sensitiveWords,omitempty"`
	Message        string   `json:"message,omitempty"`
}

func (r *Resp) Marshal() []byte {
	res, _ := json.Marshal(r)

	return res
}

func Handler(w http.ResponseWriter, r *http.Request) {
	resp := &Resp{
		StatusCode: http.StatusOK,
	}

	if r.Method == "GET" {
		if q, _ := url.ParseQuery(r.URL.RawQuery); q.Get("text") != "" {
			resp.SensitiveWords = censor.FindSensitiveWords(q.Get("text"))
		} else {
			resp.StatusCode = http.StatusNotFound
			resp.Message = "Not found sensitive words from the text"
		}
	} else {
		resp.StatusCode = http.StatusBadRequest
		resp.Message = "Bad request"
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(resp.StatusCode)
	w.Write(resp.Marshal())
}
