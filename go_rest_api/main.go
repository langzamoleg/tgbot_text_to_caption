package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/fogleman/gg"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func getImage(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()
	params := r.URL.Query().Get("text")
	const S = 1024
	im, err := gg.LoadImage("template.png")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("font.ttf", 96); err != nil {
		panic(err)
	}

	dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringAnchored(params, S/2, S/2, 0.5, 0.5)
	dc.Clip()
	dc.SavePNG(id.String() + ".png")

	fileBytes, err := ioutil.ReadFile(id.String() + ".png")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	os.Remove(id.String() + ".png")
}

func main() {

	router := mux.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	router.HandleFunc("/api", getImage).Methods("GET")

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
