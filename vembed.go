package main

import (
	"bytes"
	"compress/gzip"
	"embed"
	_ "embed"
	"fmt"
	"main/tools"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
)

// Cache du fichier Single Page index.html zippé
var fic_cache string
var bytes_cache []byte
var mu sync.RWMutex

// Embarque les static Vue dans le binaire à la compilation
//
//go:embed vues/*.vue index.html
var vues embed.FS

// 1er Construction : Avant de lancer les goroutines httpServe
func InitSpa() {

	// La Single page est désormais zippée dans le cache en mémoire
	packEmbeds("index.html")
}

func readFile(fic string) ([]byte, error) {

	if tools.Cfg.WebUi.Cache {
		// lit le fichier embarqué dans le binaire à la compilation
		return vues.ReadFile(fic)
	}
	// lit le disque
	return os.ReadFile(fic)
}

// Fusionne les fichiers html, avec server side include <embed src="xxx">
func packEmbeds(fic string) {

	// En exclusion mutuelle, en dev, refresh et goroutines
	mu.Lock()
	defer mu.Unlock()

	// Fichier racine à lire
	content, err := readFile(fic)
	if err != nil {
		fmt.Println("Fichier index.html non trouvé :", err)
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
		included, err := readFile(result[1])
		if err != nil {
			// Fichier non trouvé
			fmt.Println("Fichier v-emded", result[1], "non trouvé : ", err)
			html = strings.ReplaceAll(html, result[0], "<v-emded src='"+result[1]+"' title='not found'>")
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
}

// Compresse en gzip la chaine
func compress(html string) []byte {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	gz.Write([]byte(html))
	gz.Close()
	return buf.Bytes()
}

// Traite le point d'entrée en html ( "/" ou "/index.html")
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
	// En Dev : reconstruit la single page avec les .vue à chaque refresh !
	if !tools.Cfg.WebUi.Cache {
		packEmbeds("index.html")
	}

	//  index.html and / : the single vuejs page
	if len(bytes_cache) == 0 {
		http.NotFound(w, r)
		return
	}

	// Resultat final
	w.Header().Add("Content-Encoding", "gzip")
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	w.Write(bytes_cache)
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
