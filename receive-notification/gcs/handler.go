package gcs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func Notify(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "Receive notify")
	log.Infof(ctx, "Method : %v", r.Method)
	var (
		headers    = make([]string, 0, len(r.Header))
		parameters = make([]string, 0, len(r.Form))
	)

	for key, val := range r.Header {
		header := fmt.Sprintf("%s = %s", key, strings.Join(val, ", "))
		headers = append(headers, header)
	}
	for key, val := range r.Form {
		param := fmt.Sprintf("%s = %s", key, strings.Join(val, ", "))
		parameters = append(parameters, param)
	}
	if msg, err := ioutil.ReadAll(r.Body); err != nil {
		log.Errorf(ctx, "Error : %v", err)
	} else {
		log.Infof(ctx, "Message : %v", string(msg))
	}
	log.Infof(ctx, "Headers : %v", headers)
	log.Infof(ctx, "Params : %v", parameters)
}
