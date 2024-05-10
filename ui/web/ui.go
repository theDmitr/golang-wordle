package web

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/logic"
	"main/ui/web/page"
	"net/http"
)

type WordleUi struct {
	logic logic.WordleGame
}

func NewWordleUi(logic *logic.WordleGame) *WordleUi {
	ui := new(WordleUi)
	ui.logic = *logic
	return ui
}

func WriteFile(w *http.ResponseWriter, filename string) {
	file := page.NewFile(filename)
	fmt.Fprint(*w, string(*file.Body))
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	WriteFile(&w, "resources/index.html")
}

func StylesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	WriteFile(&w, "resources/styles.css")
}

func ScriptHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	WriteFile(&w, "resources/script.js")
}

func (ui *WordleUi) RegenerateHandler(w http.ResponseWriter, r *http.Request) {
	ui.logic.Regenerate()
}

func (ui *WordleUi) TryHandler(w http.ResponseWriter, r *http.Request) {
	buff, _ := io.ReadAll(r.Body)
	attempt := string(buff)
	resultMap, _ := ui.logic.Attempt(&attempt)
	jsonString, _ := json.Marshal(resultMap)
	fmt.Fprint(w, string(jsonString))
}

func (ui *WordleUi) Start() {
	http.HandleFunc("/game", GameHandler)
	http.HandleFunc("/styles.css", StylesHandler)
	http.HandleFunc("/script.js", ScriptHandler)
	http.HandleFunc("/regenerate", ui.RegenerateHandler)
	http.HandleFunc("/try", ui.TryHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
