package main

import (
	"log"
	"os/exec"
)

func main() {
	config, err := ProgettoSDCC.LoadConfig()
	if err != nil {
		log.Println("Errore durante il caricamento della configurazione:", err)
		return
	}

	// Comando da eseguire
	for i := 0; i < config.NumberMicroserver; i++ {

		command := "go"
		arguments := []string{"run", ".\\server.go", config.Microservers[i]}
		log.Printf("Comando: %s %s\n", command, arguments)

		// Esegui il comando
		exec.Command(command, arguments...)
		/*
			// Stampa l'output del comando
			log.Println("Output del comando:")
			log.Println(string(output))*/
	}
}
