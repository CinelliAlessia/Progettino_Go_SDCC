// serviceS.go

/*
Il file serviceS.go contiene la definizione del servizio Arith, utilizzato per gestire le chiamate RPC.
L'interfaccia esposta include un tipo di argomenti e un risultato, oltre a vari metodi per eseguire operazioni matematiche.

Args: Struttura che rappresenta gli argomenti dei metodi, con due campi A e B di tipo intero.
Arith: Struttura vuota che rappresenta il servizio RPC, implementando l'interfaccia per gestire le richieste.
Result: Tipo che rappresenta il risultato di una chiamata RPC, un intero.
*/

package service

import (
	"fmt"
	"math"
)

// Args rappresenta gli argomenti dei metodi
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
	fmt.Printf("serviceS -> Richiesta servita: %d + %d = %d\n", args.A, args.B, *result)
	return nil
}

// Factorial è un nuovo metodo dell'interfaccia Arith che calcola il fattoriale di un numero
func (a *Arith) Factorial(args Args, result *Result) error {
	if args.A < 0 {
		fmt.Printf("serviceS -> Richiesta servita: Impossibile calcolare il fattoriale di un numero negativo")
		return nil
	} else if args.A > 25 {
		fmt.Printf("serviceS -> Richiesta servita: Impossibile calcolare il fattoriale di un numero maggiore di 25")
		return nil
	}

	fact := 1
	for i := 1; i <= args.A; i++ {
		fact *= i
	}
	*result = Result(fact)
	fmt.Printf("serviceS -> Richiesta servita: %d! = %d\n", args.A, *result)
	return nil
}

// Power è un metodo dell'interfaccia Power che calcola la potenza di un numero
func (a *Arith) Power(args Args, result *Result) error {
	res := 1
	for i := 0; i < args.B; i++ {
		res *= args.A
	}
	*result = Result(res)
	fmt.Printf("serviceS -> Richiesta servita: %d^%d = %d\n", args.B, args.A, *result)
	return nil
}

// MCD è un nuovo metodo dell'interfaccia Arith che calcola il Massimo Comun Divisore
func (a *Arith) MCD(args Args, result *Result) error {
	// Algoritmo di Euclide per il calcolo del MCD
	b, c := args.A, args.B
	for c != 0 {
		b, c = c, b%c
	}
	*result = Result(b)
	fmt.Printf("serviceS -> Richiesta servita: MCD(%d, %d) = %d\n", args.A, args.B, *result)
	return nil
}

// IsPrime è un nuovo metodo dell'interfaccia Arith che verifica se un numero è primo
func (a *Arith) IsPrime(args Args, result *Result) error {
	n := args.A

	// Verifica se il numero è minore o uguale a 1
	if n <= 1 {
		*result = Result(0) // 0 indica che il numero non è primo
	} else {
		// Verifica se il numero è divisibile per qualche numero tra 2 e la radice quadrata di n
		sqrtN := int(math.Sqrt(float64(n)))
		isPrime := true

		for i := 2; i <= sqrtN; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			*result = Result(1) // 1 indica che il numero è primo
		} else {
			*result = Result(0) // 0 indica che il numero non è primo
		}
	}

	fmt.Printf("serviceS -> Richiesta servita: IsPrime(%d) = %d\n", args.A, *result)
	return nil
}
