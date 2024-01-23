// Package main implementa un semplice load balancer per gestire connessioni RPC
// tra un client e diversi server. Utilizza il pacchetto net/rpc e consente
// di gestire connessioni in modo concorrente.

package main

import (
	configuration "ProgettoSDCC"
	"ProgettoSDCC/serviceLB"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func main() {
	config, err := configuration.LoadConfig()
	if err != nil {
		fmt.Println("LB -> Errore durante il caricamento della configurazione in LB:", err)
		return
	}

	serviceLoadB := &serviceLB.ServiceLB{}

	server := rpc.NewServer()
	err = server.RegisterName("Service", serviceLoadB)
	if err != nil {
		log.Fatal("LB -> Formato del servizio non è corretto: ", err)
	}

	listener, err := net.Listen("tcp", config.LoadBalancer)
	if err != nil {
		log.Fatal("LB -> Errore durante la creazione del listener:", err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("LB -> Errore durante la chiusura del listener:", err)
		}
	}(listener)
	fmt.Println("LB -> Load Balancer in ascolto su", config.LoadBalancer)

	serviceLoadB.CurrentServer = 0
	serviceLoadB.Servers = config.Servers
	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("LB -> Errore durante l'accettazione della connessione:", err)
			continue
		}

		// Gestisce ogni connessione in una goroutine, cosi da poter gestire più di una connessione
		go server.ServeConn(conn)
	}

}
