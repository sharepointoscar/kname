package themes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Theme is a struct
type Theme struct {
	Name        string
	ClusterName string
}

// NewTheme This func is uppercased. Packages that import `theme` can instantiate a
// new theme using this method. This allows us more control around car creation.
func NewTheme(t string) Theme {

	_theme := Theme{
		Name: t,
	}

	if strings.Compare(_theme.Name, "yoga") == 0 {

		fmt.Println("using yoga theme object")
		fmt.Println("generating clustername based on yoga pose name...")

		_theme.ClusterName = getYogaPoseName()

	} else if strings.Compare(_theme.Name, "cocktails") == 0 {

		fmt.Println("using cocktails theme object")
		fmt.Println("generating clustername based on cocktail name...")

		_theme.ClusterName = getCocktailName()

	} else if strings.Compare(_theme.Name, "national-parks") == 0 {

		fmt.Println("using national-parks theme object")
		fmt.Println("generating clustername based on National Park name...")

		_theme.ClusterName = getNationalParkName()

	} else {
		fmt.Println("theme not supported. Want to contribute and create it? Fork the repo!")
	}

	return _theme
}

/* getCocktailName returns a random cocktail name
from our json data source
*/
func getCocktailName() string {

	jsonFile, err := os.Open("data/cocktails.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	type Cocktail struct {
		Name  string `json:"Name"`
		Glass string `json:"Glass"`
	}

	type Cocktails struct {
		Cocktails []Cocktail `json:"items"`
	}

	var _cocktails Cocktails

	json.Unmarshal(byteValue, &_cocktails)

	rand.Seed(time.Now().UnixNano())
	var randomCocktail = _cocktails.Cocktails[rand.Intn(len(_cocktails.Cocktails))]

	_cocktailName := strings.ToLower(strings.Replace(randomCocktail.Name, " ", "", -1))

	return _cocktailName
}

/* getYogaPoseName returns a random yoga pose name
from our json data source
*/
func getYogaPoseName() string {

	jsonFileYoga, err := os.Open("data/yoga.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFileYoga.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFileYoga)

	type YogaPose struct {
		Sanksrit string `json:"sanskritName"`
		English  string `json:"englishName"`
	}

	type YogaPoses struct {
		YogaPoses []YogaPose `json:"items"`
	}

	var _poses YogaPoses

	json.Unmarshal(byteValue, &_poses)

	rand.Seed(time.Now().UnixNano())
	var randomPose = _poses.YogaPoses[rand.Intn(len(_poses.YogaPoses))]

	_poseName := strings.ToLower(strings.Replace(randomPose.Sanksrit, " ", "", -1))

	return _poseName
}

/* getNationalParkName returns a random National Park name
from our json data source
*/
func getNationalParkName() string {

	jsonFileParks, err := os.Open("data/national_parks.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFileParks.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFileParks)

	type Park struct {
		Name  string `json:"Name"`
		State string `json:"State"`
	}

	type Parks struct {
		Parks []Park `json:"items"`
	}

	var _parks Parks

	json.Unmarshal(byteValue, &_parks)

	rand.Seed(time.Now().UnixNano())
	var _randomPark = _parks.Parks[rand.Intn(len(_parks.Parks))]

	_parkName := strings.ToLower(strings.Replace(_randomPark.Name, " ", "", -1))

	return _parkName
}
