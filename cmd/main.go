package main

import (
	"fmt"
	djan_go "github.com/Jonathanpatta/djan-go"
	"net/http"
)

type Product struct {
	Id    string `json:"id,omitempty"`
	Code  string `json:"code,omitempty"`
	Price uint   `json:"price,omitempty"`
}

type Person struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func main() {

	c, err := djan_go.NewDefaultConfig()
	if err != nil {
		fmt.Println(err)
	}

	djan_go.RegisterHttpModel(Product{}, &djan_go.DataModelConfig{
		EndPointName: "product",
		GlobalConfig: c,
	})

	djan_go.RegisterHttpModel(Person{}, &djan_go.DataModelConfig{
		EndPointName: "person",
		GlobalConfig: c,
	})

	http.Handle("/", c.Router)

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
