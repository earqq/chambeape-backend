// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import "time"

type AddLocation struct {
	Latitude   *string `json:"latitude"`
	Longitude  *string `json:"longitude"`
	PostalCode *string `json:"postal_code" bson:"postal_code"`
	Route      *string `json:"route"`
	Locality   *string `json:"locality"`
	AreaLevel1 *string `json:"area_level_1" bson:"area_level_1"`
	AreaLevel2 *string `json:"area_level_2" bson:"area_level_2"`
	Country    *string `json:"country"`
}

type Job struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Tasks     []Task    `json:"tasks"`
	IDPublic  string    `json:"id_public" bson:"id_public"`
	EndDate   string    `json:"end_date" bson:"end_date"`
	JobType   int       `json:"job_type" bson:"job_type"`
	Price     float64   `json:"price"`
	State     bool      `json:"state"`
	Location  Location  `json:"location"`
	Owner     JobOwner  `json:"owner"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at"`
}

type JobOwner struct {
	IDPublic string `json:"id_public" bson:"id_public"`
	Phone    string `json:"phone"`
	Img      string `json:"img"`
}

type Location struct {
	Latitude   *string `json:"latitude"`
	Longitude  *string `json:"longitude"`
	PostalCode *string `json:"postal_code" bson:"id_public"`
	Route      *string `json:"route"`
	Locality   *string `json:"locality"`
	AreaLevel1 *string `json:"area_level_1" bson:"area_level_1"`
	AreaLevel2 *string `json:"area_level_2" bson:"area_level_2"`
	Country    *string `json:"country"`
}

type NewJob struct {
	Title    string      `json:"title"`
	Tasks    []NewTask   `json:"tasks"`
	EndDate  string      `json:"end_date" bson:"end_date"`
	IDPublic string      `json:"id_public" bson:"id_public"`
	State    bool        `json:"state"`
	JobType  int         `json:"job_type" bson:"job_type"`
	Price    float64     `json:"price"`
	Location AddLocation `json:"location"`
	Owner    NewJobOwner `json:"owner"`
}

type NewJobOwner struct {
	IDPublic string `json:"id_public" bson:"id_public"`
	Phone    string `json:"phone"`
	Img      string `json:"img"`
}

type NewProfile struct {
	Email       string  `json:"email"`
	Names       string  `json:"names"`
	IDPublic    string  `json:"id_public" bson:"id_public"`
	Birthdate   *string `json:"birthdate"`
	Phone       *string `json:"phone"`
	ProfileType int     `json:"profile_type" bson:"profile_type"`
	Img         *string `json:"img"`
}

type NewTask struct {
	Description *string `json:"description"`
}

type Profile struct {
	ID          string    `json:"id"`
	IDPublic    string    `json:"id_public" bson:"id_public"`
	ProfileType int       `json:"profile_type" bson:"profile_type"`
	Names       string    `json:"names"`
	Email       string    `json:"email"`
	Birthdate   string    `json:"birthdate"`
	Phone       string    `json:"phone"`
	Img         string    `json:"img"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updated_at"`
}

type Task struct {
	Description string `json:"description"`
}

type UpdateJob struct {
	Title    string      `json:"title"`
	IDPublic string      `json:"id_public" bson:"id_public"`
	Tasks    []NewTask   `json:"tasks"`
	State    bool        `json:"state"`
	EndDate  string      `json:"end_date" bson:"end_date"`
	JobType  int         `json:"job_type" bson:"job_type"`
	Price    float64     `json:"price"`
	Location AddLocation `json:"location"`
	Owner    NewJobOwner `json:"owner"`
}

type UpdateProfile struct {
	IDPublic    string `json:"id_public" bson:"id_public"`
	Names       string `json:"names"`
	Img         string `json:"img"`
	Email       string `json:"email"`
	Birthdate   string `json:"birthdate"`
	Phone       string `json:"phone"`
	ProfileType int    `json:"profile_type" bson:"profile_type"`
}
