package spotify

import (
	"fmt"
	"net/http"

	"github.com/ayberktandogan/melody/config"
)

func openWebServerForCallback(state string, res chan<- string, err chan<- error) {
	http.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		if state != r.URL.Query().Get("state") {
			err <- fmt.Errorf("Error")
			return
		}

		res <- r.URL.Query().Get("code")
	})

	http.ListenAndServe(":"+config.Config.Port, nil)
}
