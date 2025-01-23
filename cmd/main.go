package main

import (
	"fmt"
	djan_go "github.com/Jonathanpatta/djan-go"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
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

//type Product struct {
//	Id         string
//	Code       string
//	Price      uint
//	Creator    Person
//	VerifiedBy []Person
//}
//
//type Person struct {
//	Id    string
//	Name  string
//	Email string
//}

func main() {

	err := godotenv.Load("test.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	c, err := djan_go.NewDefaultConfig()
	if err != nil {
		fmt.Println(err)
	}

	pgurl := os.Getenv("PGURL")
	db, err := gorm.Open(postgres.Open(pgurl), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	c.GormDb = db

	djan_go.RegisterHttpModel(Product{}, &djan_go.DataModelConfig{
		EndPointName: "product",
		GlobalConfig: c,
	})

	djan_go.RegisterHttpModel(Person{}, &djan_go.DataModelConfig{
		EndPointName: "person",
		GlobalConfig: c,
	})

	//c.Router.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
	//	b, err := os.ReadFile("ObjectPage.html") // just pass the file name
	//	if err != nil {
	//		fmt.Print(err)
	//	}
	//
	//	str := string(b) // convert content to a 'string'
	//
	//	_, err = fmt.Fprint(w, str)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//	}
	//})

	fmt.Println(djan_go.GetObjectSchemaJson(&Product{}))

	http.Handle("/", c.Router)

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
