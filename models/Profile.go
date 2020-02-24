package models

type Profile struct {
	ID                string       `json:"id"`
	IDPublic          string       `json:"id_public" bson:"id_public"`
	Names             string       `json:"names"`
	Email             string       `json:"email"`
	Birthdate         string       `json:"birthdate"`
	Phone             string       `json:"phone"`
	Img               string       `json:"img"`
	Location          Location     `json:"location"`
	WorkerPublic      *bool        `json:"worker_public"  bson:"worker_public"`
	WorkerDescription *string      `json:"worker_description"  bson:"worker_description"`
	WorkerType        *int         `json:"worker_type" bson:"worker_type"`
	WorkerExperience  []Experience `json:"worker_experience"  bson:"worker_experience"`
}
type Experience struct {
	Description *string `json:"description"`
	Phone       *string `json:"phone"`
}
