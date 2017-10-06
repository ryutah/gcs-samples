package main

import (
	"gcs"
	"net/http"
)

func init() {
	http.HandleFunc("/notify", gcs.Notify)
}
