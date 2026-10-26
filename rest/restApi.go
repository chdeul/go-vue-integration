package rest

import (
	"context"
	"main/tools"
	"net/http"
	"regexp"
	"strings"
)

type myReqCtx struct {
}

// Données contexte de requête
type agremReqCtx struct {
	path []string // Url et param
}

// Une route d'entrée d'API
type route struct {
	method  string           // GET POST PUT DELETE ...
	regex   *regexp.Regexp   // Url path with matching elements
	handler http.HandlerFunc // http func handler
	// profil string         // Auth profil to manage access on this route
}

// Tableau des routes / handler methods
var routes = []route{

	// REST  /api/exemple/...
	Route("GET", "exemple", apiExemple),
	Route("POST", "exemple", apiExemple),

	// TODO
	// List of API Paths and theirs controllers
	// REST  /api/xxxxx/...
	// Route("GET", "xxxxx", apiXxxxx),
}

// Declare une route
func Route(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

// Extrait un element (matching) de la route
func getField(r *http.Request, index int) string {
	fields := r.Context().Value(myReqCtx{}).(agremReqCtx).path
	return fields[index]
}

// Aiguille sur la bonne route
func ServeRestApi(w http.ResponseWriter, r *http.Request) {

	// Catch rest errors
	defer func() {
		if r := recover(); r != nil {
			tools.Error("Exception dans  ServeRestApi %v", r)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	// Pour message erreur final
	var allow []string

	// Cherche parmi les routes
	for _, route := range routes {

		// Est-ce que cela match ?
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {

			// Trouvé, mais pas bon verbe
			if r.Method != route.method {
				// Pour message erreur final
				allow = append(allow, route.method)
				continue
			}
			// Contexte de requete : avec url
			reqCtx := agremReqCtx{path: matches[1:]}

			// Auth  et Profil d'accès  gérable ici
			//    Jeton d'auth présent Dans Header ou Cookie
			//        Pas de jeton , Pas Authentifie
			//           http.Error(w, "Accès non authentifié", http.StatusUnauthorized)
			//           return

			// Authentifie mais Pas les bons droits
			//    http.Error(w, "Vous n'avez pas les droits", http.StatusForbidden)
			//    return

			// aiguille vers la methode, avec le contexte de requete
			ctx := context.WithValue(r.Context(), myReqCtx{}, reqCtx)
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		// Trouvé mauvais verbe
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "Verbe non autorisé", http.StatusMethodNotAllowed)
		return
	}
	// Url Non trouvée
	http.NotFound(w, r)
}
