package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Behavioural Pattern-> Observer Design Pattern

// An interface with method must be implement
type Observer interface {
	Update(stockName string, price float64)
}

// Subject interface defindes methods to manage observers.

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// Stock is the concrete subject that maintains stock price and notifies observers.
type Stock struct {
	observers []Observer
	name      string
	price     float64
}

// RegisterObserver adds a new observer to the stock
func (s *Stock) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

// RemoveObserver removes an observer from the stock
func (s *Stock) RemoveObserver(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers inform all registered observers about a price change
func (s *Stock) NotifyObservers() {
	for _, ob := range s.observers {
		ob.Update(s.name, s.price)
	}
}

// SetPrice sets a new price for the stock and notifies observers
func (s *Stock) SetPrice(price float64) {
	s.price = price
	fmt.Printf("Stock %s new price: %.2f\n", s.name, s.price)
	s.NotifyObservers()
}

// StockClient is a concrete observer that listens to stock price updates
type StockClient struct {
	id string
}

// Update is called when the stcok price changes
func (c *StockClient) Update(stockName string, price float64) {
	fmt.Printf("Client %s received update: %s new price is %.2f\n", c.id, stockName, price)
}

func main() {
	//Create a stock subject
	appleStock := &Stock{name: "Apple"}

	// Create stock clients (observers).
	client1 := &StockClient{id: "Client1"}
	client2 := &StockClient{id: "Client2"}
	client3 := &StockClient{id: "Client3"}

	// Register clients to the stock
	appleStock.RegisterObserver(client1)
	appleStock.RegisterObserver(client2)
	appleStock.RegisterObserver(client3)

	//Simulate stock price changes and notify observers
	go func() {
		i := 0
		rand.New(rand.NewSource(time.Now().UnixNano()))

		for {
			newPrice := 100 + rand.Float64()*10 // random price between 100 and 110
			appleStock.SetPrice(newPrice)
			time.Sleep(2 * time.Second)
			// removing one observer dynamically after few iteration
			if i == 5 {
				appleStock.RemoveObserver(client3)
			}
			i++
		}
	}()

	//Keep the main function running
	// select{}

	//break when you press any key
	fmt.Scanln()
}
