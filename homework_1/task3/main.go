package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/pkg/errors"

	"code-cadets-2021/homework1/task3/pokemon"
)

// Finding all local encounters of a specific pokemon.
// Can be done either with id directly or via name from which we get the id
func main() {

	var id int
	var name string

	// we can enter either pokemon name or it's id
	flag.IntVar(&id, "id", 0, "Value for the pokemon name")
	flag.StringVar(&name, "name", "", "Value for the pokemon id")

	flag.Parse()
	NameAndLocationsCombined, err := pokemon.GetPokemonInfo(id, name)
	if err != nil {
		log.Fatal(err)
	}

	// printOutput is []byte
	printOutput, err := json.Marshal(NameAndLocationsCombined)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "marshalling the pokemon data into JSON"),
		)
	}
	// printing []byte into a readable JSON
	_, err = os.Stdout.Write(printOutput)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "writing bytes into JSON"),
		)
	}

}
