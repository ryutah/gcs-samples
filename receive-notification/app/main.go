package main

import (
	"net/http"

	"github.com/ryutah/gcs-samples/receive-notification/gcs"
)

func init() {
	http.HandleFunc("/notify", gcs.Notify)
}
