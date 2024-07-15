package main

import (
	"fmt"
	djan_go "github.com/Jonathanpatta/djan-go"
	"net/http"
)

type Product struct {
	//gorm.Model
	Id    string `json:"id,omitempty"`
	Code  string `json:"code,omitempty"`
	Price uint   `json:"price,omitempty"`
}

func main() {

	//var data MyS[Product]
	//
	//fmt.Println(data.testy())

	//db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}

	// Migrate the schema
	//db.AutoMigrate(&Product{})

	//pmodel := djan_go.RegisterDataModel(Product{}, db)

	//data := Product{Id: "5", Code: "D42", Price: 100}
	//Create
	//_, err = pmodel.Post(data)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//t, err := pmodel.Get("5")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(t)
	//if err != nil {
	//	fmt.Println(err)
	//}

	// Read
	//var product Product
	//db.First(&product, 1)                 // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(Product{}, "id = ?", "5")

	//_, err = pmodel.Delete("5")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//router := mux.NewRouter()
	//router.StrictSlash(false)

	c, err := djan_go.NewDefaultConfig()
	if err != nil {
		fmt.Println(err)
	}

	djan_go.RegisterHttpModel(Product{}, &djan_go.DataModelConfig{
		Auth:         false,
		EndPointName: "product",
		GlobalConfig: c,
	})

	//router.HandleFunc("/product/list", httpmodel.List).Methods("GET", "OPTIONS")
	//router.HandleFunc("/product", httpmodel.Post).Methods("POST", "OPTIONS")
	//router.HandleFunc("/product", httpmodel.Put).Methods("PUT", "OPTIONS")
	//router.HandleFunc("/product/{id}", httpmodel.Get).Methods("GET", "OPTIONS")
	//router.HandleFunc("/product/{id}", httpmodel.Delete).Methods("DELETE", "OPTIONS")
	http.Handle("/", c.Router)

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
