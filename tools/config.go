package tools

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"go.yaml.in/yaml/v4"
)

// Structure de données de config
type Config struct {
	Logging string
	LogTime bool

	WebUi struct {
		ContextPath string
		Cache       bool
	}
	Server struct {
		Protocol string
		ListenIp string
		Port     string
		CertFile string
		KeyFile  string
		UsedUrl  string
		BackUrl  string
	}
}

// Config globale chargée
var Cfg Config

// Recuperer du json
func Decode(r *http.Request, obj interface{}) {
	json.NewDecoder(r.Body).Decode(&obj)
}

// Envoyer du json
func Encode(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func LoadConfig(dir string) {

	// Lecture fichier
	err := LoadAnyConfig(dir+"/config.yml", &Cfg)
	if err != nil {
		panic(err)
	}

	// Normalisation des context path configuré  "/"" => ""   "/toto" => "/toto"
	if Cfg.WebUi.ContextPath == "/" {
		Cfg.WebUi.ContextPath = ""
	}
	// Pas de date et heure sur les log et coloré (dev)
	if !Cfg.LogTime {
		LogWithColor(true)
		log.SetFlags(0)
	}
	// Positionne le niveau de log
	SetLogLevel(Cfg.Logging)
}

// Charge une configuration yaml
func LoadAnyConfig(filepath string, ptr any) error {

	data, err := os.ReadFile(filepath)
	if err != nil {
		Error("Fichier config %s non trouvé", filepath)
		return err
	}
	if err := yaml.Unmarshal(data, ptr); err != nil {
		Error("Fichier config erroné : %s", err)
		return err
	}
	return err
}
