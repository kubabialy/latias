package name

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func GenerateName() string {
    return fetchRandomPokemonName()
}

type Pokemon struct {
    Name string     `json:"name"`
}

const NUMBER_OF_POKES = 913

func getRandomId() int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(NUMBER_OF_POKES)
}

func fetchRandomPokemonName() string {
    url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v/", getRandomId())
    res, err := http.Get(url)

    if err != nil {
      log.Fatalln("Could not fetch the data")
    }

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
      log.Fatalln("Could not read from response")
    }

    pkm := Pokemon{}
    if json.Unmarshal(body, &pkm) != nil {
        log.Fatalln("Could not parse json input")
    }

    return pkm.Name
}
