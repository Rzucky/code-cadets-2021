package pokemon

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
)

const pokemonURL = "https://pokeapi.co/api/v2/pokemon/"

func linearBackoff(retry int) time.Duration {
	return time.Duration(retry) * time.Second
}

type pokemonName struct {
	Name string `json:"name"`
}

type pokemonId struct {
	Id int `json:"id"`
}

type locationAreaName struct {
	Name string `json:"name"`
}

// locationAreaName is used to get the name of the area for each location_area
type pokemonLocalEncounters struct {
	LocalEncounterArea locationAreaName `json:"location_area"`
}

// combined data for a specific pokemon
type pokemonData struct {
	Name      string   `json:"name"`
	Locations []string `json:"location_areas"`
}

// returns []byte from a given URL
func gettingJSONBodyContent(url string) ([]byte, error) {
	httpClient := pester.New()
	httpClient.Backoff = linearBackoff

	httpResponse, err := httpClient.Get(url)
	if err != nil {
		return nil, errors.WithMessage(err, "HTTP get towards pokemon API")
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "reading body of pokemon API response")
	}

	return bodyContent, nil
}

func getPokemonName(id int) (string, error) {
	pokemonIdUrl := pokemonURL + strconv.Itoa(id)

	bodyContent, err := gettingJSONBodyContent(pokemonIdUrl)
	if err != nil {
		return "", err
	}

	var pokemonName pokemonName
	err = json.Unmarshal(bodyContent, &pokemonName)
	if err != nil {
		return "", errors.WithMessage(err, "unmarshalling the JSON body content")
	}

	return pokemonName.Name, nil
}

func getPokemonId(name string) (int, error) {

	pokemonNameUrl := pokemonURL + name

	bodyContent, err := gettingJSONBodyContent(pokemonNameUrl)
	if err != nil {
		return 0, err
	}

	var pokemonId pokemonId
	err = json.Unmarshal(bodyContent, &pokemonId)
	if err != nil {
		return 0, errors.WithMessage(err, "unmarshalling the JSON body content")
	}

	return pokemonId.Id, nil
}

func getPokemonEncounters(id int) ([]pokemonLocalEncounters, error) {
	pokemonEncountersUrl := pokemonURL + strconv.Itoa(id) + "/encounters"

	bodyContent, err := gettingJSONBodyContent(pokemonEncountersUrl)
	if err != nil {
		return nil, err
	}

	var pokemonFoundLocations []pokemonLocalEncounters
	err = json.Unmarshal(bodyContent, &pokemonFoundLocations)
	if err != nil {
		return nil, errors.WithMessage(err, "unmarshalling the JSON body content")
	}

	return pokemonFoundLocations, nil
}

func GetPokemonInfo(id int, name string) (pokemonData, error) {

	if id == 0 && len(name) == 0 {
		return pokemonData{}, errors.New("pokemon name or id not specified")
	}

	if id < 0 {
		return pokemonData{}, errors.New("pokemon id is negative, should be positive")
	}

	// if someone entered the name with an uppercase letter
	if len(name) != 0 {
		name = strings.ToLower(name)
	}

	// getting name over id
	if id != 0 && len(name) == 0 {
		foundName, err := getPokemonName(id)
		if err != nil {
			return pokemonData{}, err
		}
		name = foundName
	}

	// getting id over pokemons id
	if len(name) != 0 && id == 0 {
		foundId, err := getPokemonId(name)
		if err != nil {
			return pokemonData{}, err
		}
		id = foundId
	}

	// getting encounters over id
	pokemonFoundLocations, err := getPokemonEncounters(id)
	if err != nil {
		return pokemonData{}, err
	}

	var NameAndLocationsCombined pokemonData
	// the name was either given from the start or gotten via id
	NameAndLocationsCombined.Name = name

	// adding all locations into one string splice
	for _, location := range pokemonFoundLocations {
		NameAndLocationsCombined.Locations = append(NameAndLocationsCombined.Locations, location.LocalEncounterArea.Name)
	}

	return NameAndLocationsCombined, nil
}
