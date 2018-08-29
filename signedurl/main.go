package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
)

var (
	googleAccessID = os.Getenv("GOOGLE_ACCESS_ID")
	pricateKey     = strings.Replace(os.Getenv("PRIVATE_KEY"), "\\n", "\n", -1)
	bucket         = os.Getenv("BUCKET")
)

type fileInfo struct {
	ContentType string `json:"content_type"`
	FileName    string `json:"file_name"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/signedurl", signedURL)

	http.ListenAndServe(":8080", nil)
}

func signedURL(w http.ResponseWriter, r *http.Request) {
	info := new(fileInfo)
	if err := json.NewDecoder(r.Body).Decode(info); err != nil {
		log.Printf("failed to decode body: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	signedURL, err := storage.SignedURL(bucket, info.FileName, &storage.SignedURLOptions{
		GoogleAccessID: googleAccessID,
		PrivateKey:     []byte(pricateKey),
		ContentType:    info.ContentType,
		Expires:        time.Now().Add(15 * time.Minute),
		Method:         http.MethodPut,
	})
	// In AppEngine
	//  account, err := appengine.ServiceAccount(ctx)
	//  if err != nil {
	//  	// error handling...
	//  }
	//  signedURL, err := storage.SignedURL(bucket, info.FileName, &storage.SignedURLOptions{
	//  	GoogleAccessID: account,
	//  	ContentType:    info.ContentType,
	//  	Expires:        time.Now().Add(15 * time.Minute),
	//  	Method:         http.MethodPut,
	//  })
	if err != nil {
		log.Printf("failed to create signed url: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"upload_url": signedURL,
	})
}
