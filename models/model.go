package models

import "time"

type DogBreed struct {
	ID               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weightLowLbs"`
	WeightHighLbs    int    `json:"weightHighLbs"`
	AverageWeight    int    `json:"averageWeight"`
	Lifespan         int    `json:"lifespan"`
	Details          string `json:"details"`
	AlternativeNames string `json:"alternativeNames"`
	GeographicOrigin string `json:"geographicOrigin"`
}

type CatBreed struct {
	ID               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weightLowLbs"`
	WeightHighLbs    int    `json:"weightHighLbs"`
	AverageWeight    int    `json:"averageWeight"`
	Lifespan         int    `json:"lifespan"`
	Details          string `json:"details"`
	AlternativeNames string `json:"alternativeNames"`
	GeographicOrigin string `json:"geographicOrigin"`
}

type Dog struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	BreedID          int       `json:"breedID"`
	BreederID        int       `json:"breederID"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"dateOfBirth"`
	SpayedOrNeutered int       `json:"spayedOrNeutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            string    `json:"breed"`
	Breeder          string    `json:"breeder"`
}

type Cat struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	BreedID          int       `json:"breedID"`
	BreederID        int       `json:"breederID"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"dateOfBirth"`
	SpayedOrNeutered int       `json:"spayedOrNeutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            string    `json:"breed"`
	Breeder          string    `json:"breeder"`
}

type Breeder struct {
	ID          int         `json:"id"`
	BreederName string      `json:"breederName"`
	Address     string      `json:"address"`
	Phone       string      `json:"phone"`
	Email       string      `json:"email"`
	Active      int         `json:"active"`
	DogBreeds   []*DogBreed `json:"dogBreeds"`
	CatBreeds   []*CatBreed `json:"catBreeds"`
}

type Pet struct {
	Species     string `json:"species"`
	Breed       string `json:"breed"`
	MinWeight   int    `json:"minWeight"`
	MaxWeight   int    `json:"maxWeight"`
	Description string `json:"description"`
	LifeSpan    int    `json:"lifeSpan"`
}
