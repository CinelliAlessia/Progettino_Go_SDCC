package main

import (
	configuration "ProgettoSDCC"
	"ProgettoSDCC/serviceLB"
	"fmt"
	"log"
	"net/rpc"
	_ "os"
	"strconv"
	"strings"
)

func main() {
	// Carica la configurazione dal file di configurazione
	config, err := configuration.LoadConfig()
	if err != nil {
		fmt.Println("Errore durante il caricamento della configurazione nel client:", err)
		return
	}

	// Visualizza l'indirizzo del Load Balancer a cui ci si sta connettendo
	fmt.Println("Indirizzo del Load Balancer a cui connettersi:", config.LoadBalancer)

	// Connessione al Load Balancer tramite TCP
	conn, err := rpc.Dial("tcp", config.LoadBalancer)
	if err != nil {
		fmt.Println("Errore durante la connessione al Load Balancer:", err)
		return
	}
	defer func(conn *rpc.Client) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// Loop infinito per consentire all'utente di inserire richieste continuamente
	for {
		fmt.Print("Inserisci due numeri separati da spazio (o 'stop' per terminare): ")
		var input string
		fmt.Scanln(&input)

		// Verifica se l'utente ha inserito "stop" per terminare il loop
		if strings.ToLower(input) == "stop" {
			fmt.Println("Chiusura del client.")
			break
		}

		// Parsa l'input dell'utente per ottenere gli argomenti della chiamata RPC
		args := parseInput(input)
		var result serviceLB.Result

		// Effettua una chiamata alla procedura remota in modo sincrono
		log.Printf("Effettuata chiamata sincrona\n")
		err = conn.Call("Service.Sum", args, &result)
		if err != nil {
			log.Fatal("Errore in ForwardRequest:", err)
		}

		// Stampa il risultato ottenuto dalla chiamata RPC
		fmt.Printf("Risultato: %d + %d = %d\n", args.A, args.B, result)
		fmt.Println("Risultato ottenuto:", result)
	}
}

// Funzione per parsare l'input dell'utente e ottenere gli argomenti della chiamata RPC
func parseInput(input string) serviceLB.Args {
	// Divide l'input in parti basate su spazi
	parts := strings.Fields(input)
	// Verifica se ci sono esattamente due parti
	if len(parts) != 2 {
		fmt.Println("Input non valido. Inserisci due numeri separati da spazio.")
		return serviceLB.Args{}
	}

	// Converti le parti in numeri interi
	n1, err1 := strconv.Atoi(parts[0])
	n2, err2 := strconv.Atoi(parts[1])

	// Verifica se la conversione è avvenuta con successo
	if err1 != nil || err2 != nil {
		fmt.Println("Input non valido. Assicurati di inserire numeri validi.")
		return serviceLB.Args{}
	}

	// Restituisce gli argomenti della chiamata RPC
	return serviceLB.Args{A: n1, B: n2}
}
