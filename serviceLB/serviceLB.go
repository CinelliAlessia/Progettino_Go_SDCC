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
	"fmt"
	"log"
	"net/rpc"
)

// Args rappresenta gli argomenti delle procedure remote
type Args struct {
	A, B int
}

// Result rappresenta il risultato delle procedure remote
type Result int

// ServiceLB gestisce gli indirizzi dei server per realizzare la politica RR
type ServiceLB struct {
	ServerAddr    string
	CurrentServer int
	Servers       []string
}

// forwardRequest effettua una chiamata RPC al server per eseguire la procedura richiesta dal client
func (s *ServiceLB) forwardRequest(args Args, result *Result, operation string) error {

	s.ServerAddr = s.Servers[s.CurrentServer]
	s.CurrentServer = (s.CurrentServer + 1) % len(s.Servers)

	fmt.Printf("serviceLB -> impostato sul serviceLB: %s \n", s.ServerAddr)

	// Connessione al server tramite RPC
	serverConn, err := rpc.Dial("tcp", s.ServerAddr)
	if err != nil {
		fmt.Println("serviceLB -> Errore durante la connessione al server in ServiceLB:", err)
		return nil
	}
	defer func(serverConn *rpc.Client) {
		// Chiusura della connessione al server
		err := serverConn.Close()
		if err != nil {
			fmt.Println("serviceLB -> Errore durante la chiusura della connessione al server in ServiceLB:", err)
		}
	}(serverConn)

	// Chiamata RPC alla procedura remota del server
	if operation == "Sum" {
		err := serverConn.Call("Arithmetic.Sum", args, result)
		if err != nil {
			return err
		}
	} else if operation == "Factorial" {
		err := serverConn.Call("Arithmetic.Factorial", args, result)
		if err != nil {
			return err
		}
	} else if operation == "IsPrime" {
		err := serverConn.Call("Arithmetic.IsPrime", args, result)
		if err != nil {
			return err
		}
	} else if operation == "Power" {
		err := serverConn.Call("Arithmetic.Power", args, result)
		if err != nil {
			return err
		}
	} else if operation == "MCD" {
		err := serverConn.Call("Arithmetic.MCD", args, result)
		if err != nil {
			return err
		}
	} else {
		log.Fatal("serviceLB -> Operazione non supportata: ", operation)
	}

	return nil
}

// Sum è un metodo utilizzato solo dal client per effettuare la chiamata RPC tramite forwardRequest
func (s *ServiceLB) Sum(args Args, result *Result) error {
	operation := "Sum"
	err := s.forwardRequest(args, result, operation)
	if err != nil {
		return err
	}
	return nil
}

// Factorial è un metodo utilizzato solo dal client per effettuare la chiamata RPC tramite forwardRequest
func (s *ServiceLB) Factorial(args Args, result *Result) error {
	operation := "Factorial"
	err := s.forwardRequest(args, result, operation)
	if err != nil {
		return err
	}
	return nil
}

// IsPrime è un metodo utilizzato solo dal client per effettuare la chiamata RPC tramite forwardRequest
func (s *ServiceLB) IsPrime(args Args, result *Result) error {
	operation := "IsPrime"
	err := s.forwardRequest(args, result, operation)
	if err != nil {
		return err
	}
	return nil
}

// Power è un metodo utilizzato solo dal client per effettuare la chiamata RPC tramite forwardRequest
func (s *ServiceLB) Power(args Args, result *Result) error {
	operation := "Power"
	err := s.forwardRequest(args, result, operation)
	if err != nil {
		return err
	}
	return nil
}

// MCD è un metodo utilizzato solo dal client per effettuare la chiamata RPC tramite forwardRequest
func (s *ServiceLB) MCD(args Args, result *Result) error {
	operation := "MCD"
	err := s.forwardRequest(args, result, operation)
	if err != nil {
		return err
	}
	return nil
}
