package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/odoo-backend/internal/app/config/env"
	"github.com/odoo-backend/internal/app/project"
	"github.com/odoo-backend/internal/app/responses"
	"github.com/skilld-labs/go-odoo/api"
)

func getAllEmployee() []*project.Employee {
	godoohost := env.Get("GODOO_HOST")
	c, err := api.NewClient(godoohost, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = c.Login("roketout", "admin@roketout.com", "admin")
	if err != nil {
		fmt.Println(err.Error())
	}
	s := api.NewHrEmployeeService(c)
	so, err := s.GetAll()
	if err != nil {
		fmt.Println(err.Error())
	}

	employees := make([]*project.Employee, 0)

	for index := 0; index < len(*so); index++ {

		var elem project.Employee

		elem.Nama = (*so)[index].DisplayName
		elem.Department = (*so)[index].DepartmentId.Name
		elem.Email = (*so)[index].WorkEmail

		employees = append(employees, &elem)
	}

	return employees
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {

	res := responses.NewResponse("you actual fucking pepeg")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func GetAllEmployee(w http.ResponseWriter, r *http.Request) {

	res := getAllEmployee()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func GetJawaban(w http.ResponseWriter, r *http.Request) {

	var qnadata *project.PertanyaanJawaban

	err := json.NewDecoder(r.Body).Decode(&qnadata)
	if err != nil {
		printError(err, w)
		return
	}

	jawaban, err := srv.FetchShowPertanyaanJawaban(qnadata.Pertanyaan)
	if err != nil {
		printError(err, w)
		return
	}

	res := responses.NewResponse(jawaban)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func StoreQNA(w http.ResponseWriter, r *http.Request) {

	var qnadata *project.PertanyaanJawaban

	err := json.NewDecoder(r.Body).Decode(&qnadata)
	if err != nil {
		printError(err, w)
		return
	}

	jawaban, err := srv.FetchStorePertanyaanJawaban(qnadata)
	if err != nil {
		printError(err, w)
		return
	}

	res := responses.NewResponse(jawaban)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func StoreKeluhan(w http.ResponseWriter, r *http.Request) {

	// ID              interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	// Pengguna        Pengguna    `json:"pengguna,omitempty" bson:"pengguna,omitempty"`
	// Department      string      `json:"department,omitempty" bson:"department,omitempty"`
	// PenanggungJawab Employee    `json:"penanggungjawab,omitempty" bson:"penanggungjawab,omitempty"`
	// Isi             string      `json:"isi,omitempty" bson:"isi,omitempty"`
	// Jawaban         string      `json:"jawaban,omitempty" bson:"jawaban,omitempty"`

	var keluhan *project.Keluhan

	err := json.NewDecoder(r.Body).Decode(&keluhan)
	if err != nil {
		printError(err, w)
		return
	}

	_, err = srv.FetchStorePengguna(&keluhan.Pengguna)
	if err != nil {
		printError(err, w)
		return
	}

	keluhan.Jawaban = "Berikut adalah contoh jawaban dari keluhan"

	galen := getAllEmployee()

	for _, data := range galen {
		if data.Department == keluhan.Department {
			keluhan.PenanggungJawab = *data
		}
	}

	keluhan, err = srv.FetchStoreKeluhan(keluhan)
	if err != nil {
		printError(err, w)
		return
	}

	res := responses.NewResponse(keluhan)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func StorePengguna(w http.ResponseWriter, r *http.Request) {

	var pengguna *project.Pengguna

	err := json.NewDecoder(r.Body).Decode(&pengguna)
	if err != nil {
		printError(err, w)
		return
	}

	penggunares, err := srv.FetchStorePengguna(pengguna)
	if err != nil {
		printError(err, w)
		return
	}

	res := responses.NewResponse(penggunares)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func GetPengguna(w http.ResponseWriter, r *http.Request) {

	pengguna, err := srv.ShowAllPengguna()
	if err != nil {
		printError(err, w)
		return
	}

	res := responses.NewResponse(pengguna)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
