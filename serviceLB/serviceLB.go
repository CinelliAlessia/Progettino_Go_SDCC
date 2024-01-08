// serviceLB.go

/*
Il file serviceLB.go implementa la gestione delle chiamate RPC lato client.
La struttura ServiceLB contiene informazioni sull'indirizzo del server e fornisce funzionalità per effettuare chiamate RPC.

Args: Struttura che rappresenta gli argomenti delle procedure remote, con campi A e B di tipo intero.
Result: Tipo che rappresenta il risultato delle procedure remote, un intero.
ServiceLB: Struttura principale che gestisce la connessione al server e fornisce funzioni per effettuare chiamate RPC.

SetServerAddress: Metodo che imposta l'indirizzo del server secondo la politica Round Robin.
forwardRequest: Metodo che effettua una chiamata RPC al server per eseguire la procedura richiesta dal client.
Sum: Metodo utilizzato solo dal client per effettuare la chiamata RPC tramite forwardRequest.
*/

package serviceLB

import (
	"log"
	"net/rpc"
)

// Args rappresenta gli argomenti delle procedure remote
type Args struct {
	A, B int
}

// Result rappresenta il risultato delle procedure remote
type Result int

// ServiceLB gestisce la connessione al server e fornisce funzionalità per effettuare chiamate RPC
type ServiceLB struct {
	ServerAddr string
}

// SetServerAddress imposta l'indirizzo del server secondo la politica Round Robin
func (s *ServiceLB) SetServerAddress(addr string) {
	s.ServerAddr = addr
}

// forwardRequest effettua una chiamata RPC al server per eseguire la procedura richiesta dal client
func (s *ServiceLB) forwardRequest(args Args, result *Result) error {
	log.Println(s.ServerAddr)

	// Connessione al server tramite RPC
	serverConn, err := rpc.Dial("tcp", s.ServerAddr)
	if err != nil {
		log.Println("Errore durante la connessione al server in ServiceLB:", err)
		return nil
	}
	defer func(serverConn *rpc.Client) {
		// Chiusura della connessione al server
		err := serverConn.Close()
		if err != nil {
			log.Println("Errore durante la chiusura della connessione al server in ServiceLB:", err)
		}
	}(serverConn)

	// Chiamata RPC alla procedura remota del server
	err = serverConn.Call("Arithmetic.Sum", args, result)
	if err != nil {
		log.Fatal("Errore in Arithmetic.Sum: ", err)
	}

	return nil
}

// Sum è un metodo utilizzato solo dal client per effettuare la chiamata RPC tramite forwardRequest
func (s *ServiceLB) Sum(args Args, result *Result) error {
	err := s.forwardRequest(args, result)
	if err != nil {
		return err
	}
	return nil
}

/*
In questo file ci sono tre funzioni:
	SetServerAddress: imposta l'indirizzo del server secondo la politica Round Robin.
	forwardRequest: effettua la chiamata RPC al server che svolge la procedura richiesta dal client.
	Sum: Una funzione utilizzata solo dal client che effettua la chiamata RPC tramite forwardRequest.
*/
