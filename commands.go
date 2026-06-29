package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/mahdee-123/pokedexcli/internal/pokecache"
)


var cache pokecache.Cache = pokecache.NewCache(5*time.Second)
var locationApiBaseUrl string = "https://pokeapi.co/api/v2/location"

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type LocationResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type LocationAreaResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon PokemonReference `json:"pokemon"`
}

type PokemonReference struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}


type cliCommand struct {
	name        string
	description string
	callback    func(string) error
}

var commands = map[string]cliCommand{}

func init() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "show help",
			callback: commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the program",
			callback: commandExit,
		},
		
		"map" : {
			name : "map", 
			description: "",
			callback: showMap,
		},

		"explore" : {
			name : "explore", 
			description: "",
			callback: exploreArea,
		},
	}
}



func commandExit(_ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}



func showMap(_ string) error {


	// duita case 
	/// case 1 : cach exist
	var result LocationResponse

	/// cach checking
	// data,ok := locationsCache.Get(locationApiUrl)

	if data,ok := cache.Get(locationApiBaseUrl); ok {
		err := json.Unmarshal(data, &result)
		if err != nil {
			 return err
		}
		for _, location := range result.Results {
			fmt.Println(location.Name)
		} 	
		fmt.Println("data comes from cache....")
		return nil
	} 

	res, err := http.Get(locationApiBaseUrl)
	
	if err != nil {
		return err
	}
	
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	// add in cach
	if err != nil {
		return err
	}
	cache.Add(locationApiBaseUrl,body)
	if res.StatusCode > 299 {
		return fmt.Errorf("status code: %d", res.StatusCode)
	}


	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	for _ , location := range result.Results {
		fmt.Println(location.Name)
	}
	fmt.Println("data comes from server....")
		
	return nil
}

func commandHelp(_ string) error {
	fmt.Println("welcome to the pokedex")
	fmt.Println("usages : ")
	fmt.Println()

	for _, cmd := range commands  {
		fmt.Printf("%s : %s\n", cmd.name, cmd.description)
	}


	return nil

}


func exploreArea(location string) error {

	var url string

	var result LocationAreaResponse
	id, err := strconv.Atoi(location)
	
	if err != nil {
		url = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)
	} else {
		url = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d", id)
	}



	data , ok := cache.Get(url)

	if ok {
		err := json.Unmarshal(data, &result)
		if err != nil {
			return err
		}
		for _, encounter := range result.PokemonEncounters {
			fmt.Println(encounter.Pokemon.Name)
		}
		return nil
	}
	
	res, err := http.Get(url)

	if err != nil {
		return err 
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)


	if err != nil {
		return err
	}

	cache.Add(url, body)


	err = json.Unmarshal(body, &result)
	
	if err != nil {
		return err
	}

	
	if res.StatusCode > 299 {
    fmt.Printf("status code: %d", res.StatusCode)
	} 
	for _, encounter := range result.PokemonEncounters {
			fmt.Println(encounter.Pokemon.Name)
	}


	return nil
}

