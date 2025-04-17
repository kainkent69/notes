package main

import (
	"bufio"
	"cmp"
	"encoding/json"
	"log"
	"os"
)

type Animal struct {
	Type  string `json:"type"`
	Color string `json:"color"`
	Name  string `json:"name"`
}

func readAndGetFile() []byte {
	fs, err := os.Open("./animals.json")
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 0)

	scanner := bufio.NewScanner(fs)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Panicf("Scanner Errored: %v\n", err)
		}
		buf = append(buf, []byte(scanner.Text())...)
	}
	return buf

}

func BasicUsage() {
	slc := make([]Animal, 0)
	file := readAndGetFile()
	// convert the data to non json data
	log.Println("Converting Data Into NonJson")
	err := json.Unmarshal(file, &slc)
	if err != nil {
		log.Panicf("Unmarshal Failed %v\n", err)
	}
	log.Printf("\n\nThe UnMarshaled=>\n\n %#+v\n\n", slc)
	// convert the unmarshaled data into marshal again
	log.Println("Marshaling Unmarshaled Data")
	data, err := json.Marshal(slc)
	if err != nil {
		log.Printf("Marshaling failed: %v", err)
		log.Fatal(0)
	}
	// printing the marshlaed text
	log.Printf("The Marshaled=>\n %v\n", string(data))
}

func MuateData() {
	file := readAndGetFile()
	animals := make([]Animal, 0)
	json.Unmarshal(file, &animals)
	for i, animal := range animals {
		if cmp.Compare("Coby", animal.Name) == 0 {
			animals[i].Name = "Mr. Coby"
		} else if cmp.Compare(animal.Color, "orange") == 0 {
			animals[i].Color = "somewhat-orange"
		}
	}

	log.Printf("The New List=>: \n\n%+v\n", animals)
}
func main() {
	log.Println("Testing making Json Data")
	log.Print("\n\n=================Basic Usage=========================\n\n")
	BasicUsage()
	log.Print("\n\n==================Mutate Data========================\n\n")
	MuateData()

}
