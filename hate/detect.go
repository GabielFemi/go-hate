package hateSpeech

import (
	"fmt"
	"net/http"
)

func Detect(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		_ = r.ParseForm()
		speechGottenFromUser := r.FormValue("speech")
		fmt.Println(speechGottenFromUser)
		isSpeechHateWord(speechGottenFromUser)
		http.Redirect(w, r, "/", 301)
	} else {
		render(w, "index.html", r)
	}

}

// Loop through the sentence, if buhari(for example) is found, check if the current word is the hashmap
// containing the hate words
// if true, return true
// otherwise, return false

func isSpeechHateWord(speech string) {

}
