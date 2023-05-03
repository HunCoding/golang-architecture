package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
	"sync"
)

type singletonDatabase struct {
	capital []string
}

func (sd *singletonDatabase) GetCapitals() []string {
	return sd.capital
}

var once sync.Once
var instance *singletonDatabase

func NewSingletonDatabase() *singletonDatabase {
	godotenv.Load(".env")
	capitals := strings.Split(os.Getenv("CAPITALS"), ",")
	once.Do(func() {
		db := singletonDatabase{capitals}
		fmt.Println("Iniciou a variavel")
		instance = &db
	})
	fmt.Println(&instance)
	return instance
}

func main() {
	fmt.Println(NewSingletonDatabase(), NewSingletonDatabase())
}
