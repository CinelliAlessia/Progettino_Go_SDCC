// myService.go

/*
Il file arith.go contiene la definizione del servizio Arith, utilizzato per gestire le chiamate RPC.
L'interfaccia esposta include un tipo di argomenti e un risultato, oltre a un metodo Sum che somma due numeri interi.

Args: Struttura che rappresenta gli argomenti del metodo Sum, con due campi A e B di tipo intero.
Arith: Struttura vuota che rappresenta il servizio RPC, implementando l'interfaccia per gestire le richieste.
Result: Tipo che rappresenta il risultato di una chiamata RPC, un intero.

Sum: Metodo dell'interfaccia Arith che esegue la somma dei due numeri passati come argomenti e restituisce il risultato.
*/

package service

import "log"

// Args rappresenta gli argomenti del metodo Sum
type Args struct {
	A, B int
}

// Arith è una struttura vuota che rappresenta il servizio RPC
type Arith struct{}

// Result rappresenta il risultato di una chiamata RPC, un intero
type Result int

// Sum è un metodo dell'interfaccia Arith che somma due numeri interi
func (a *Arith) Sum(args Args, result *Result) error {
	*result = Result(args.A + args.B)
	log.Printf("Richiesta servita: %d + %d = %d", args.A, args.B, *result)
	return nil
}
