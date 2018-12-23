package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var imageFileExtentionsList = []string{
	".png",
	".apng",
	".jpg",
	".jpeg",
	".gif",
	".webp",
}

func main() {
	t, err := template.New("index").Parse(index)
	if err != nil {
		log.Fatal(err)
		return
	}

	imagePaths := make([]string, 0, 1)
	if len(os.Args) > 1 {
		files, err := ioutil.ReadDir(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fileName := file.Name()
			for _, ext := range imageFileExtentionsList {
				if ext == filepath.Ext(fileName) {
					imagePaths = append(imagePaths, fileName)
					break
				}
			}
		}
	}
	if len(imagePaths) == 0 {
		imagePaths = append(imagePaths, "/default.png")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			if err := t.ExecuteTemplate(w, "index", imagePaths); err != nil {
				log.Fatal(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
		http.FileServer(http.Dir(".")).ServeHTTP(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
