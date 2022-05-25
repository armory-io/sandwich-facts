package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	fmt.Println("Sandwich Facts Started!")

	facts := []string{
		"Sandwiches!",
		"Sandwich is not spelled 'sandwhich'",
		"It is debated that hotdogs are sanwiches",
		"The sandwich is named after John Montagu (1718-92), the 4th Earl of Sandwich, who started a craze for eating beef between two slices of toast",
		"The first recorded use of the word 'Sandwich' (with a capital S), was in 1762 by Edward Gibbon writing of the Beef Steak Club in London",
		"The only novel by Jane Austen containing the word ‘sandwich’ is Mansfield Park",
		"Around 12 billion sandwiches are eaten every year in the UK",
		"The world’s largest sandwich currently weighs 5,440 pounds",
		"In 2008, an attempt in Iran to beat the record for the world's biggest sandwich failed when the impatient crowd ate it before it was measured",
		"The earliest reference to a ‘bacon sandwich’ listed in the Oxford English Dictionary was by George Orwell in 1931",
		"The record for creating the most expensive sandwich ever made is claimed by chef Tom Bridge whose Lancaster Cheese Sandwich which was sold on eBay in 2006 for £345",
		"The Sandwich Islands (now Hawaii) were also named after the 4th Earl of Sandwich in 1778",
		"Sandwich' is also a town in Kent, although the name has no direct connection with a sandwich",
	}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		fact := facts[rand.Intn(len(facts))]
		fmt.Println(fact)
		io.WriteString(w, fact)
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
