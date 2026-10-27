package main

import (
	"context"
	"crypto/tls"
	"embed"
	_ "embed"
	"io/fs"
	"main/rest"
	"main/tools"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// static web server content embed in binary.
//
//go:embed resources/*
var resources embed.FS

func main() {

	// ---- Initialisations générales ----
	tools.Info("Début")

	// Initialisation  Config
	//time.LoadLocation("Europe/Paris")
	tools.LoadConfig(".")

	// Context path de la partie appli web  /
	path := tools.Cfg.WebUi.ContextPath

	// Http serve static resources ( js, css, img, fonts )
	var fsr http.Handler
	if tools.Cfg.WebUi.Cache {
		sfs, e := fs.Sub(resources, "resources")
		if e != nil {
			tools.Fatal("FS Server fail :  %v", e)
		}
		fsr = http.FileServer(http.FS(sfs))
	} else {
		fsr = http.FileServer(http.Dir("resources"))
	}
	http.Handle(path+"/resources/", http.StripPrefix(path+"/resources/", fsr))

	// Http serve static vues  ( debug ? )
	//fv := http.FileServer(http.Dir("vues/"))
	//http.Handle(path+"/vues/", http.StripPrefix(path+"/vues/", fv))

	// Http serve Rest API  /api/*
	http.Handle(path+"/api/", http.StripPrefix(path+"/api/", http.HandlerFunc(rest.ServeRestApi)))

	// Http server refresh and favoris of browser on /vue/* routes
	http.Handle(path+"/vue/", http.StripPrefix(path+"/vue/", http.HandlerFunc(ServeVue)))

	// Assemble la single page ( en prod 1 fois, en dev, une 1ere fois )
	InitSpa()

	// Point d'entree, Accueil ( single page ) et autres  /  => /index.html
	http.Handle(path+"/", http.StripPrefix(path+"/", http.HandlerFunc(ServeIndex)))
	if path != "" {
		// Wrong context path
		http.Handle("/", http.StripPrefix("/", http.HandlerFunc(ServeNone)))
	}

	// Serveur HTTP
	svr := &http.Server{
		TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12},
		Addr:      tools.Cfg.Server.ListenIp + ":" + tools.Cfg.Server.Port,
		Handler:   http.DefaultServeMux,
	}
	// Sur Ctrl C ou kill
	go detecterArret(svr)

	if tools.Cfg.Server.Protocol == "https" {
		// https : en production pour transfert mots de passe crypté
		err := svr.ListenAndServeTLS(tools.Cfg.Server.CertFile, tools.Cfg.Server.KeyFile)
		if err != nil {
			tools.Fatal("TLS Server fail to start :  %v", err)
		}
	} else {
		// Http pour test
		err := svr.ListenAndServe()
		if err != nil {
			tools.Fatal("Server fail to start :  %v", err)
		}
	}

	tools.Info("Fin ")
}

func detecterArret(svr *http.Server) {

	// Capture CTRL+C ou kill -SIGINT
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	tools.Info("Serveur web prêt : %s", tools.Cfg.Server.BackUrl)

	<-stop // On attend le signal

	tools.Info("Arrêt demandé...")

	// Timeout pour l'arrêt propre
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := svr.Shutdown(ctx); err != nil {
		tools.Fatal("Arrêt forcé: %v", err)
	}
}
