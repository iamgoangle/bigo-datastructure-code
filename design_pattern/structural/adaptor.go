package main

import "fmt"

// Problem
// Imagine that you’re creating a stock market monitoring app.
// The app downloads the stock data from multiple sources in XML format and then displays nice-looking charts and diagrams for the user.

// At some point, you decide to improve the app by integrating a smart 3rd-party analytics library.
// But there’s a catch: the analytics library only works with data in JSON format.

// Solution
// You can create an adapter.
// This is a special object that converts the interface of one object so that another object can understand it.
// https://refactoring.guru/design-patterns/adapter

func legacy() []string {
	return []string{"Apple", "Facebook"}
}

type Adaptor interface {
	Stocks() map[string]int
}

type adaptor struct {
	legacy []string
}

func newAdaptor(l []string) Adaptor {
	return &adaptor{
		legacy: l,
	}
}

func (a *adaptor) Stocks() map[string]int {
	result := make(map[string]int)
	for i, v := range a.legacy {
		result[v] = i * 100
	}

	return result
}

func main() {
	ad := newAdaptor(legacy())
	fmt.Println(ad.Stocks())
}
