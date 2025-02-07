package djan_go

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/qor/roles"
	"net/http"
)

type HttpDataModel[T any] struct {
	EndPointName    string
	DataModel       DataModel[T]
	DataModelConfig *DataModelConfig
	Auth            bool
	Debug           bool
	Permissions     *roles.Permission
	GlobalConfig    *Config
}

func RegisterDefaultHttpModel[T any](h *HttpDataModel[T]) (*HttpDataModel[T], error) {
	h.DataModelConfig = &DataModelConfig{
		GlobalConfig: h.GlobalConfig,
		Permission:   h.Permissions,
	}
	dm := NewGormDataModel[T](h.DataModelConfig)
	httpDataModel := &HttpDataModel[T]{
		DataModel:       dm,
		Auth:            h.Auth,
		DataModelConfig: h.DataModelConfig,
		EndPointName:    h.EndPointName,
		Permissions:     h.Permissions,
		Debug:           h.GlobalConfig.Debug || h.Debug,
	}

	if h.EndPointName == "" {
		panic("empty endpoint name")
	}

	AddHttpModelSubrouter(h.GlobalConfig.Router, h.EndPointName, httpDataModel)

	return httpDataModel, nil
}

func RegisterHttpCustomModel[T any](h *HttpDataModel[T]) (*HttpDataModel[T], error) {
	httpDataModel := &HttpDataModel[T]{
		DataModel:       h.DataModel,
		Auth:            h.Auth,
		DataModelConfig: h.DataModelConfig,
		EndPointName:    h.EndPointName,
		Permissions:     h.Permissions,
		Debug:           h.GlobalConfig.Debug || h.Debug,
	}

	if h.EndPointName == "" {
		panic("empty endpoint name")
	}

	AddHttpModelSubrouter(h.GlobalConfig.Router, h.EndPointName, httpDataModel)

	return httpDataModel, nil
}

func (d *HttpDataModel[T]) Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	data, err := d.DataModel.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	HttpOutput(w, data)
}

func (d *HttpDataModel[T]) Post(w http.ResponseWriter, r *http.Request) {
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
	HttpOutput(w, data)

}

func (d *HttpDataModel[T]) Put(w http.ResponseWriter, r *http.Request) {
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
	HttpOutput(w, data)

}

func (d *HttpDataModel[T]) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	data, err := d.DataModel.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	HttpOutput(w, data)

}

func (d *HttpDataModel[T]) List(w http.ResponseWriter, r *http.Request) {
	data, err := d.DataModel.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	HttpOutput(w, data)

}

func HttpOutput(w http.ResponseWriter, v interface{}) {
	outData, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprint(w, string(outData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddHttpModelSubrouter[T any](r *mux.Router, epn string, httpmodel *HttpDataModel[T]) {
	router := r.PathPrefix("/api/" + epn).Subrouter()
	//fmt.Println(r)
	//fmt.Println(router)
	if !httpmodel.Debug && httpmodel.Auth {
		router.Use(httpmodel.JWTMiddleware)
	}
	router.HandleFunc("/list", httpmodel.List).Methods("GET", "OPTIONS")
	router.HandleFunc("", httpmodel.Post).Methods("POST", "OPTIONS")
	router.HandleFunc("", httpmodel.Put).Methods("PUT", "OPTIONS")
	router.HandleFunc("/{id}", httpmodel.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/{id}", httpmodel.Delete).Methods("DELETE", "OPTIONS")
}
