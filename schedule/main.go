package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type Transaction struct {
	Amount float32
	Date   time.Time
	Title  string
}

func main() {
	counter, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid counter value")
		return
	}

	fmt.Println("Counter:", counter)

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case _ = <-ticker.C:
			i := 0
			for i < counter {
				fmt.Println("Transaction", i)
				transaction := Transaction{
					Amount: float32(gofakeit.Product().Price),
					Date:   gofakeit.PastDate(),
					Title:  gofakeit.Product().Name,
				}
				fmt.Println("transaction:", transaction)
				i++
			}

		}
	}
}
