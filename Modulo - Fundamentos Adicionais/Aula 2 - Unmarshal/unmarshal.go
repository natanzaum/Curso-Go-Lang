package main

import (
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

	cachorroEmJSON := `{"nome":"Huddy","raca":"Poodle","idade":12,"dono":"Natan"}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal("Erro na conversão do cão -> ", erro)
	}

	fmt.Println(c)

	cachorroEmJSON2 := `{"nome":"Rex","raca":"Caramelo","idade":"2","dono":"Natan"}`

	c2 := map[string]string{}

	if erro := json.Unmarshal([]byte(cachorroEmJSON2), &c2); erro != nil {
		log.Fatal("Erro na conversão do cão -> ", erro)
	}

	fmt.Println(c2)

}
