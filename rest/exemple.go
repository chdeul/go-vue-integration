package rest

import (
	"main/tools"
	"net/http"
)

// REST call : retourne des élément de configuration
func apiExemple(w http.ResponseWriter, r *http.Request) {

	// Json with Capitalize name fields marshalled
	var body struct{ Data string }

	if r.Method == http.MethodPost {

		// Post data -> json
		tools.Decode(r, &body)
		if body.Data != "" {
			tools.Warn("Reçu POST with data : %s", body.Data)
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	ret := struct{ Data string }{"Hello world!"}

	tools.Log("Reçu GET -> return : %s", ret.Data)

	// Return as json
	tools.Encode(w, &ret)
}
