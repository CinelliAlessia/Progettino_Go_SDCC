// Package main implementa un semplice load balancer per gestire connessioni RPC
// tra un client e diversi server. Utilizza il pacchetto net/rpc e consente
// di gestire connessioni in modo concorrente.

package main

import (
	configuration "ProgettoSDCC"
	"ProgettoSDCC/serviceLB"
	"log"
	"net"
	"net/rpc"
)

func main() {
	// Carica la configurazione dal file di configurazione
	config, err := configuration.LoadConfig()
	if err != nil {
		log.Println("Errore durante il caricamento della configurazione in LB:", err)
		return
	}

	// Inizializza gli indirizzi dei server da distribuire
	servers := make([]string, config.NumberOfServer)
	for i := 0; i < config.NumberOfServer; i++ {
		servers[i] = config.Servers[i]
	}

	// Crea una struttura del servizioLB per lavorare con le funzioni esposte
	serviceLoadB := &serviceLB.ServiceLB{}

	// Crea un nuovo server RPC
	server := rpc.NewServer()
	err = server.RegisterName("Service", serviceLoadB)
	if err != nil {
		log.Fatal("Formato del servizio Sum non è corretto: ", err)
	}

	// Crea un listener per accettare connessioni su un indirizzo specificato
	listener, err := net.Listen("tcp", config.LoadBalancer)
	if err != nil {
		log.Fatal("Errore durante la creazione del listener:", err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Println("Errore durante la chiusura del listener:", err)
		}
	}(listener)
	log.Println("Load Balancer in ascolto su", config.LoadBalancer)

	// Loop infinito per gestire connessioni da diversi clienti
	currServer := 0
	for {

		serverAddr := servers[currServer]
		currServer = (currServer + 1) % config.NumberOfServer

		// Accetta una connessione e gestisci le richieste RPC
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Errore durante l'accettazione della connessione:", err)
			continue
		}
		// Imposta sull'istanza di ServiceLB l'indirizzo del server corrente
		go func() {
			serviceLoadB.SetServerAddress(serverAddr)

			// Servi la connessione RPC
			server.ServeConn(conn)

			// Chiudi la connessione quando non è più necessaria
			err := conn.Close()
			if err != nil {
				return
			}
		}()

		/* Versione sincrona bloccante ma che usa realmente rpc per la accept
		serviceLoadB.SetServerAddress(serverAddr)
		log.Printf("Impostato indirizzo %s\n", serverAddr)

		// Accetta una connessione e gestisci le richieste RPC
		server.Accept(listener)
		*/
	}
}
