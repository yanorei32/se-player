package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	appDir := filepath.Dir(os.Args[0])

	listen := ":3000"

	if len(os.Args) == 2 {
		listen = os.Args[1]
	}

	log.Println("Starting HTTP server...")

	srv := http.Server{Addr: listen}

	http.HandleFunc("/api/ls", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "This endpoint requires a GET request.\n");
			return
		}

		files, err := ioutil.ReadDir(filepath.Join(appDir, "contents/se"))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to read se dir: %v", err)
			return
		}

		filenames := make([]string, 0)

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			if file.Name()[0] == '.' {
				continue
			}

			filenames = append(filenames, file.Name())
		}

		jsonBytes, err := json.Marshal(filenames)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "JSON marshal error: %v", err)
			return
		}

		w.Write(jsonBytes)
	})

	http.Handle(
		"/se",
		http.StripPrefix(
			"/se/",
			http.FileServer(http.Dir(filepath.Join(appDir, "se"))),
		),
	)

	http.Handle(
		"/",
		http.FileServer(http.Dir(filepath.Join(appDir, "contents"))),
	)

	hub := newHub()
	go hub.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

