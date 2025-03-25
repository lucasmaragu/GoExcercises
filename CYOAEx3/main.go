package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

var storyArc map[string]StoryArc

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))
	data, err := os.ReadFile("gopher.json")
	if err != nil {
		fmt.Println("Error leyendo el archivo:", err)
		return
	}

	err = json.Unmarshal(data, &storyArc)
	if err != nil {
		fmt.Println("Error al deserializar JSON:", err)
		return
	}

	fmt.Println("Historias cargadas:", len(storyArc))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		arcKey := r.URL.Path[1:] // Obtener la clave desde la URL
		if arcKey == "" {
			arcKey = "intro" // Historia inicial por defecto
		}

		arc, exists := storyArc[arcKey]
		if !exists {
			http.NotFound(w, r)
			return
		}

		tmpl.Execute(w, arc)
	})
	http.ListenAndServe(":80", nil)
}
