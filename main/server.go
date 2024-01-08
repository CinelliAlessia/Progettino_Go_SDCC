// server.go

/*
Il seguente programma implementa un server RPC che fornisce un servizio di calcolo aritmetico.
Il server si registra per ascoltare le chiamate RPC sulla porta specificata come argomento del programma.
Viene creato un'istanza della struttura che implementa l'interfaccia Arith per gestire le richieste.

Il server ascolta su una porta specificata come argomento e accetta connessioni RPC in arrivo su tale porta.
Inoltre, il servizio Arith viene registrato presso il server RPC.

Nota: Questo codice deve essere esteso con la logica specifica del servizio Arith implementato dalla struttura.
*/

package main

import (
	"ProgettoSDCC/service"
	"log"
	"net"
	"net/rpc"
	"os"
)

func main() {
	// Stampa l'indirizzo e la porta su cui il server RPC è in ascolto
	log.Printf("RPC main in ascolto sulla porta %s", os.Args[1])

	// Crea un'istanza della struttura che implementa l'interfaccia Arith
	arith := new(service.Arith)

	// Crea un nuovo server RPC
	server := rpc.NewServer()

	// Registra il servizio Arith presso il server RPC
	err := server.RegisterName("Arithmetic", arith)
	if err != nil {
		log.Fatal("Formato del servizio Arith non corretto: ", err)
	}

	// Crea un listener per accettare connessioni sulla porta specificata
	lis, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Fatal("Errore durante l'avvio del server RPC:", err)
	} else {
		// Consente al server RPC di accettare connessioni in arrivo sul listener
		// e di gestire le richieste per ogni connessione in arrivo.
		log.Printf("RPC main in ascolto sulla porta %s", lis.Addr())
		server.Accept(lis)
	}
}
