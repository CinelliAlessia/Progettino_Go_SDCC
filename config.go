// config.go

/*
Il file config.go gestisce la lettura del file di configurazione config.json.
Contiene una struttura Config che rappresenta gli indirizzi del Load Balancer e dei Servers, e il numero totale di Servers.
La funzione LoadConfig apre il file di configurazione, lo decodifica in una struttura Config e restituisce un'istanza di Config.
*/

package configuration

import (
	"encoding/json"
	"log"
	"os"
)

// Costante che rappresenta il percorso del file di configurazione
const filename = "..\\configuration\\config.json"

// Config rappresenta gli indirizzi del Load Balancer e dei Servers, e il numero totale di Servers
type Config struct {
	LoadBalancer   string   `json:"addressLoadBalancer"`
	Servers        []string `json:"addressServers"`
	NumberOfServer int      `json:"numberOfServers"`
}

// LoadConfig legge il file di configurazione e restituisce un'istanza di Config
func LoadConfig() (*Config, error) {
	// Apertura del file di configurazione
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Errore durante l'apertura del file di configurazione in config.go: ", err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			// Gestione degli errori durante la chiusura del file
		}
	}(file)

	// Decodifica del file JSON nella struttura Config
	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Errore durante il parsing del file di configurazione:", err)
		return nil, err
	}

	// Restituzione della configurazione
	return &config, nil
}
