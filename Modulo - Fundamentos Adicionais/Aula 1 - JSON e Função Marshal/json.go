package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
	Dono  string `json:"dono"`
}

func main() {
	c := cachorro{"Huddy", "Poodle", 12, "Natan"}
	fmt.Println(c)

	cachorroJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal("Erro ao converter o cachorro")
	}

	fmt.Println(cachorroJson)

	fmt.Println(bytes.NewBuffer(cachorroJson))

	c2 := map[string]string{
		"nome":  "Teddy",
		"raca":  "Dálmata",
		"idade": "5",
	}

	cachorroJson2, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal("Erro ao converter o cachorro")
	}

	fmt.Println(bytes.NewBuffer(cachorroJson2))

}
