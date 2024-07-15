package djan_go

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type HttpModel[T any] struct {
	EndPointName string
	DataModel    *DataModel[T]
}

func RegisterHttpModel[T any](data T, c *DataModelConfig) *HttpModel[T] {
	httpmodel := &HttpModel[T]{
		DataModel:    RegisterDataModel(data, c),
		EndPointName: c.EndPointName,
	}

	AddHttpModelSubrouter(c.GlobalConfig.Router, c.EndPointName, httpmodel)

	return httpmodel
}

func (d *HttpModel[T]) Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	data, err := d.DataModel.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	outData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprint(w, string(outData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (d *HttpModel[T]) Post(w http.ResponseWriter, r *http.Request) {
	var data T
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err = d.DataModel.Post(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	outData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprint(w, string(outData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (d *HttpModel[T]) Put(w http.ResponseWriter, r *http.Request) {
	var data T
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err = d.DataModel.Put(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	outData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprint(w, string(outData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (d *HttpModel[T]) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	data, err := d.DataModel.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	outData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprint(w, string(outData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (d *HttpModel[T]) List(w http.ResponseWriter, r *http.Request) {
	data, err := d.DataModel.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	outData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprint(w, string(outData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddHttpModelSubrouter[T any](r *mux.Router, epn string, httpmodel *HttpModel[T]) {
	router := r.PathPrefix("/api/" + epn).Subrouter()

	router.HandleFunc("/list", httpmodel.List).Methods("GET", "OPTIONS")
	router.HandleFunc("", httpmodel.Post).Methods("POST", "OPTIONS")
	router.HandleFunc("", httpmodel.Put).Methods("PUT", "OPTIONS")
	router.HandleFunc("/{id}", httpmodel.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/{id}", httpmodel.Delete).Methods("DELETE", "OPTIONS")
}
