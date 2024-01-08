// client.go

/*
Il file client.go rappresenta il client che effettua chiamate RPC al Load Balancer.
Carica la configurazione dal file di configurazione, mostra l'indirizzo del Load Balancer a cui connettersi e inizia la connessione.

Se non vengono inseriti argomenti sufficienti, il programma termina con un messaggio di errore.
Successivamente, converte gli argomenti in numeri interi, crea un'istanza di Args con i valori desiderati e di Result.
Infine, effettua una chiamata alla procedura remota in modo sincrono, stampando il risultato ottenuto dal Load Balancer.
*/

package main

import (
	configuration "ProgettoSDCC"
	"ProgettoSDCC/serviceLB"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
)

func main() {
	// Caricamento della configurazione dal file di configurazione
	config, err := configuration.LoadConfig()
	if err != nil {
		fmt.Println("Errore durante il caricamento della configurazione nel client:", err)
		return
	}

	// Mostra l'indirizzo del Load Balancer a cui connettersi
	fmt.Println("Indirizzo del Load Balancer a cui connettersi:", config.LoadBalancer)

	// Connessione al Load Balancer tramite TCP
	conn, err := rpc.Dial("tcp", config.LoadBalancer)
	if err != nil {
		fmt.Println("Errore durante la connessione al Load Balancer:", err)
		return
	}
	defer func(conn *rpc.Client) {
		// Chiusura della connessione
		err := conn.Close()
		if err != nil {
			// Gestisce eventuali errori durante la chiusura della connessione
		}
	}(conn)

	// Verifica se sono presenti argomenti sufficienti
	if len(os.Args) < 3 {
		fmt.Printf("Argomenti non inseriti\n")
		os.Exit(1)
	}

	// Converti gli argomenti in numeri interi
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	// Crea un'istanza di Args con i valori desiderati e di Result
	args := serviceLB.Args{A: n1, B: n2}
	var result serviceLB.Result

	// Effettua una chiamata alla procedura remota in modo sincrono
	log.Printf("Effettuata chiamata sincrona\n")
	err = conn.Call("Service.Sum", args, &result)
	if err != nil {
		log.Fatal("Errore in ForwardRequest:", err)
	}

	// Stampa il risultato ottenuto dal Load Balancer
	fmt.Printf("Risultato: %d + %d = %d\n", args.A, args.B, result)
	fmt.Println("Risultato ottenuto:", result)
}
