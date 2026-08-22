package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const pathTempFile = "~/.baguette/config.json"

// Load configs from json file
func Load() {
	log.Fatal("To implement")
}

// Select a config
func Pick() {
	log.Fatal("To implement")
}

// Create config
func Create() {
	log.Fatal("To implement")
}

// Delete config
func Delete() {
	log.Fatal("To implement")
}

//Todo: Apprend à écrire dans un json
// puis apprend à écrire par dessus
// charge tout puis manipule puis écris
// puis limit à 4 config max

type LaConf struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Target     string `json:"target"`
	Client     string `json:"client"`
	SocketPAth string `json:"socketPath"`
	Selected   bool   `json:"selected"`
}

func OpenFile(path string) (*os.File, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func CloseFile(file *os.File) error {
	err := file.Close()

	if err != nil {
		return err
	}

	return nil
}

func ReadJson(path string) error {

	value, err := ReadAtAll(path)
	if err != nil {
		return err
	}

	fmt.Println(value)
	return nil
}

func ReadAtAll(path string) (*[]LaConf, error) {

	var config []LaConf
	file, err := OpenFile(path)
	if err != nil {
		return nil, err
	}

	buf := bufio.NewReader(file)
	endOfJson := ']'

	jsonBlob, err := buf.ReadBytes(byte(endOfJson))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBlob, &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
