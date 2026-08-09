package envLoader

import (
	"fmt"
	"log"
	"runtime"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
}

func PrintEnv() {
	//fmt.Printf("%s: %s \n", key, os.Getenv(key))
	fmt.Printf("%s: %s \n", "GOARCH", runtime.GOARCH)
	fmt.Printf("%s: %s \n", "GOOS", runtime.GOOS)
}
