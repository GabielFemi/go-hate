package hateSpeech

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Detect(w http.ResponseWriter, r *http.Request) {
	render(w, "index.html", r)
	vars := mux.Vars(r)
	speechGottenFromUser := vars["speech"]
	isSpeechHateWord(speechGottenFromUser)
}

func isSpeechHateWord(speech string) {

}
