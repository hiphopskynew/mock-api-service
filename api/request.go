package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	vPath, vQuery, vHeader := r.RequestURI, r.URL.Query(), r.Header
	var vBody interface{} = nil
	if bytes, e := ioutil.ReadAll(r.Body); e == nil {
		doc := map[string]interface{}{}
		json.Unmarshal(bytes, &doc)
		if len(doc) > 0 {
			vBody = doc
		} else {
			vBody = string(bytes)
		}
	}
	fmt.Println("--------------------------------")
	fmt.Println("PATH =>", vPath)
	fmt.Println("QUERY =>", vQuery)
	fmt.Println("HEADER =>", vHeader)
	fmt.Println("BODY =>", vBody)
	fmt.Println("HTTP METHOD =>", r.Method)
	fmt.Println("--------------------------------")

	result := map[string]interface{}{
		"Request-Endpoint":    vPath,
		"Request-Query-Param": vQuery,
		"Request-Header":      vHeader,
		"Request-Body":        vBody,
		"Request-Method":      r.Method,
	}
	bytes, _ := json.Marshal(result)
	fmt.Println("=>", string(bytes))
	fmt.Println("")

	sanitizeHeader := func(h map[string][]string) {
		for k, v := range h {
			if strings.ToLower(k) == "content-length" {
				w.Header().Del(k)
				continue
			}
			w.Header()[k] = []string{v[0]}
		}
	}

	w.Header().Add("content-type", "application/json")
	for k, v := range vHeader {
		w.Header().Add(k, fmt.Sprintf("%s", v[0]))
	}

	sanitizeHeader(w.Header())
	encoder := json.NewEncoder(w)
	encoder.Encode(result)
}
