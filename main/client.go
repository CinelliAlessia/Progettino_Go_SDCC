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
	"ProgettoSDCC/service"
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// Caricamento della configurazione dal file di configurazione
	config, err := configuration.LoadConfig()
	if err != nil {
		fmt.Println("Client -> Errore durante il caricamento della configurazione nel client:", err)
		return
	}

	// Mostra l'indirizzo del Load Balancer a cui connettersi
	fmt.Println("Client -> Indirizzo del Load Balancer a cui connettersi:", config.LoadBalancer)

	// Connessione al Load Balancer tramite TCP
	conn, err := rpc.Dial("tcp", config.LoadBalancer)
	if err != nil {
		fmt.Println("Client -> Errore durante la connessione al Load Balancer:", err)
		return
	}
	defer func(conn *rpc.Client) { // Chiusura della connessione
		err := conn.Close()
		if err != nil {
			fmt.Println("Client -> Errore durante la chiusura della connessione al Load Balancer:", err)
		}
	}(conn)

	for {
		// Stampa il menu interattivo
		fmt.Println("Scegli un'operazione:")
		fmt.Println("1. Somma")
		fmt.Println("2. Fattoriale")
		fmt.Println("3. Verifica se un numero è primo")
		fmt.Println("4. Potenza")
		fmt.Println("5. MCD")
		fmt.Println("6. Uscita")

		// Leggi l'input dell'utente per l'operazione
		fmt.Print("\nInserisci il numero dell'operazione desiderata: ")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Client -> Errore durante la lettura dell'input:", err)
			break
		}

		// Esegui l'operazione scelta
		switch choice {
		case 1, 2, 3, 4, 5:
			runOperation(conn, choice)
		case 6:
			fmt.Println("Uscita dal programma.")
			return
		default:
			fmt.Println("Scelta non valida. Riprova.")
		}
	}
}

func runOperation(conn *rpc.Client, choice int) {
	var args service.Args
	var result service.Result

	// Leggi i parametri per l'operazione
	for {
		fmt.Print("Inserisci il primo parametro: ")
		_, err := fmt.Scan(&args.A)
		if err == nil {
			break
		}
		fmt.Println("Client -> Errore durante la lettura del primo parametro:", err)
	}

	if choice != 2 && choice != 3 { // Fattoriale e numero primo non hanno bisogno di inserire il secondo parametro

		for {
			fmt.Print("Inserisci il secondo parametro: ")
			_, err := fmt.Scan(&args.B)
			if err == nil {
				break
			}
			fmt.Println("Client -> Errore durante la lettura del secondo parametro:", err)
		}
	}

	// Esegui l'operazione scelta
	switch choice {
	case 1:
		runSumOperation(conn, args, &result)
	case 2:
		runFactorialOperation(conn, args, &result)
	case 3:
		runIsPrimeOperation(conn, args, &result)
	case 4:
		runPowerOperation(conn, args, &result)
	case 5:
		runMCDOperation(conn, args, &result)
	}

	//fmt.Printf("Risultato dell'operazione: %v\n\n", result)
}

func runSumOperation(conn *rpc.Client, args service.Args, result *service.Result) {
	call := conn.Go("Service.Sum", args, result, nil)
	<-call.Done
	if call.Error != nil {
		log.Fatal("Errore durante la chiamata RPC:", call.Error)
	}
	fmt.Printf("Risultato dell'operazione: %d+%d = %d\n\n", args.A, args.B, *result)
}

func runFactorialOperation(conn *rpc.Client, args service.Args, result *service.Result) {
	call := conn.Go("Service.Factorial", args, result, nil)
	<-call.Done
	if call.Error != nil {
		log.Fatal("Errore durante la chiamata RPC:", call.Error)
	}
	fmt.Printf("Risultato dell'operazione: %d! = %d\n\n", args.A, *result)

}

func runIsPrimeOperation(conn *rpc.Client, args service.Args, result *service.Result) {
	call := conn.Go("Service.IsPrime", args, result, nil)
	<-call.Done
	if call.Error != nil {
		log.Fatal("Errore durante la chiamata RPC:", call.Error)
	}
	fmt.Printf("Risultato dell'operazione: %d è primo? %v\n\n", args.A, *result)

}

func runPowerOperation(conn *rpc.Client, args service.Args, result *service.Result) {
	call := conn.Go("Service.Power", args, result, nil)
	<-call.Done
	if call.Error != nil {
		log.Fatal("Errore durante la chiamata RPC:", call.Error)
	}

	fmt.Printf("Risultato dell'operazione: %d^%d = %v\n\n", args.A, args.B, *result)
}

func runMCDOperation(conn *rpc.Client, args service.Args, result *service.Result) {
	call := conn.Go("Service.MCD", args, result, nil)
	<-call.Done
	if call.Error != nil {
		log.Fatal("Errore durante la chiamata RPC:", call.Error)
	}
	fmt.Printf("Risultato dell'operazione: il MCD di %d è %v\n\n", args.A, *result)

}
