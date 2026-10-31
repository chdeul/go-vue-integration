package rest

import (
	"main/tools"
	"net/http"
	"strconv"
)

// REST call : retourne des élément de configuration
func apiExemple(w http.ResponseWriter, r *http.Request) {

	// Json with Capitalize name fields marshalled
	var body struct{ Data string }

	// POST /exemple called
	if r.Method == http.MethodPost {

		// Post data -> json
		tools.Decode(r, &body)
		if body.Data != "" {
			tools.Warn("Reçu POST with data : %s", body.Data)
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	// else GET Call eiyther /exemple  or    /exemple/(word)/print/(number)

	// Get my Context
	mot := ""
	ctx := GetMyCtx(r)
	if len(ctx.Path) == 2 { // Two matchs : word and number
		mot = ctx.Path[0]
		nb, _ := strconv.Atoi(ctx.Path[1])

		// Log it
		tools.Info("Path  mot: %s  nb: %d ", mot, nb)
	}
	// Get data injected before routing
	data := ctx.MyData
	tools.Info("Field data: %s", data)

	// Json return data to javascript
	ret := struct{ Data string }{"Hello world! " + mot}

	tools.Log("Reçu GET -> return : %s", ret.Data)

	// Return as json
	tools.Encode(w, &ret)
}
