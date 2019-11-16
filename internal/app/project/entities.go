package project

import (
	"time"
)

const (
	UserUrl = "/api/v1/users"
)

/*
|--------------------------------------------------------------------------
| Other Entities
|--------------------------------------------------------------------------
|
*/
type Timestamp struct {
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	CreatedBy AuthUser  `json:"created_by,omitempty" bson:"created_by,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	UpdatedBy AuthUser  `json:"updated_by,omitempty" bson:"updated_by,omitempty"`
	DeletedAt time.Time `json:"-,omitempty" bson:"deleted_at,omitempty"`
}

type PertanyaanJawaban struct {
	ID         interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Pengguna   Pengguna    `json:"pengguna,omitempty" bson:"pengguna,omitempty"`
	Pertanyaan string      `json:"pertanyaan,omitempty" bson:"pertanyaan,omitempty"`
	Jawaban    string      `json:"jawaban,omitempty" bson:"jawaban,omitempty"`
}

type Keluhan struct {
	ID              interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Pengguna        Pengguna    `json:"pengguna,omitempty" bson:"pengguna,omitempty"`
	Department      string      `json:"department,omitempty" bson:"department,omitempty"`
	PenanggungJawab Employee    `json:"penanggungjawab,omitempty" bson:"penanggungjawab,omitempty"`
	Isi             string      `json:"isi,omitempty" bson:"isi,omitempty"`
	Jawaban         string      `json:"jawaban,omitempty" bson:"jawaban,omitempty"`
}

type Pengguna struct {
	ID    interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Email string      `json:"email,omitempty" bson:"email,omitempty"`
}

type Employee struct {
	Nama       string `json:"nama,omitempty" bson:"nama,omitempty"`
	Email      string `json:"email,omitempty" bson:"email,omitempty"`
	Department string `json:"department,omitempty" bson:"department,omitempty"`
}

/*
|--------------------------------------------------------------------------
| In Source Entities
|--------------------------------------------------------------------------
|
*/
type Auth struct {
	User  AuthUser `json: "auth_user, omitempty" bson: "auth_user,omitempty"`
	Token string   `json:"token"`
}

type AuthUser struct {
	ID       string `json:"id,omitempty" bson:"id,omitempty"`
	FullName string `json:"full_name,omitempty" bson:"full_name,omitempty"`
}

type User struct {
	ID       interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	FullName string      `json:"full_name,omitempty" bson:"full_name,omitempty"`

	Timestamp
}
