package hateSpeech

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render(w, "index.html", r)
}

func Detect(w http.ResponseWriter, r *http.Request) {

	render(w, "detect.html", r)

}

// Loop through the sentence, if buhari(for example) is found, check if the current word is the hashmap
// containing the hate words
// if true, return true
// otherwise, return false

func isSpeechHateWord(speech string) {

}
