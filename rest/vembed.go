package rest

import (
	"bytes"
	"compress/gzip"
	"main/tools"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// Debug : mettre à faux (défaut)
var cached bool = false

var (
	// Cache d'un fichier
	fic_cache   string
	bytes_cache []byte
)

// Fusionne les fichiers html, avec server side include <embed src="xxx">
func packEmbeds(w http.ResponseWriter, r *http.Request, fic string) {

	// Test si le Cache matche
	if !cached || fic != fic_cache || bytes_cache == nil {

		// Fichier racine à lire
		content, err := os.ReadFile(fic)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		// Contenu html initial lu
		html := string(content)

		// Cherche des balises
		for {
			// Les <embed src= >
			exp := regexp.MustCompile(`<v-embed[ \t]+src=["']([^'"]+)['"][ \t]*></v-embed>`)
			result := exp.FindStringSubmatch(html)
			if result == nil {
				// Terminé
				break
			}
			// Lit le fichier src indiqué
			included, err := os.ReadFile(result[1])
			if err != nil {
				// Fichier non trouvé
				html = strings.ReplaceAll(html, result[0], "<v-emded src='"+result[1]+"' not found>")
				continue
			}
			// Injecte le ContextPath de l'application dans app.vue ou autre v-embed
			v_embed := string(included)
			v_embed = strings.Replace(v_embed, "const ContextPath = ''", "const ContextPath = '"+tools.Cfg.WebUi.ContextPath+"'", 1)

			// Remplace la balise par le contenu inclus
			html = strings.ReplaceAll(html, result[0], v_embed)
		}
		// Cache the result  : TODO Zip content
		bytes_cache = compress(html)
		fic_cache = fic

	} // Test cache

	// Resultat final
	w.Header().Add("Content-Encoding", "gzip")
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	w.Write(bytes_cache)
}

// Compresse en gzip la chaine
func compress(html string) []byte {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	gz.Write([]byte(html))
	gz.Close()
	return buf.Bytes()
}

// Traite le point d'entrée en html
func ServeIndex(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "" && r.URL.Path != "index.html" {

		// Demande icone par le navigateur
		if r.URL.Path == "favicon.ico" {
			icoUrl := tools.Cfg.WebUi.ContextPath + "/resources/img/favicon.png"
			http.Redirect(w, r, icoUrl, http.StatusSeeOther)
			return
		}
		// Autre url à la racine : non géré
		if strings.HasPrefix(r.URL.Path, ".well-known/") {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// No found
		tools.Warn("Unknown resource: %s", r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Cache ou non la single page (traitement des v-embed)
	cached = tools.Cfg.WebUi.Cache

	//  index.html and / : the single vuejs page
	packEmbeds(w, r, "index.html")
}

// Traite le refresh ou favori browser sur route /vue/*
func ServeVue(w http.ResponseWriter, r *http.Request) {

	// Query param to return
	params := r.URL.RawQuery

	// Check loop
	if strings.HasPrefix(params, "route=/vue/") {
		// On boucle
		w.WriteHeader(http.StatusLoopDetected)
		return
	}
	// retour en redirection
	ret := tools.Cfg.WebUi.ContextPath + "/?route=/vue/" + r.URL.Path

	if params != "" {
		ret = ret + "&" + params
	}
	// Redirect to index.html?route=/vue/xxxxx
	http.Redirect(w, r, ret, http.StatusSeeOther)
}

// Traite le point d'entrée en html
func ServeNone(w http.ResponseWriter, r *http.Request) {

	// No found
	tools.Warn("Unknown context : %s", r.URL.Path)
	w.WriteHeader(http.StatusNotFound)
}
