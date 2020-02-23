package models

type Worker struct {
	ID              string       `json:"id"`
	WorkerType      *int         `json:"worker_type" bson:"worker_type"`
	Description     *string      `json:"description"`
	Location        Location     `json:"location"`
	Public          *bool        `json:"public"`
	Experience      []Experience `json:"experience"`
	ProfileIDPublic string       `json:"profile_id_public" bson:"profile_id_public"`
	Names           string       `json:"names"`
	Email           string       `json:"email"`
	Birthdate       string       `json:"birthdate"`
	Phone           string       `json:"phone"`
	Img             string       `json:"img"`
}

type Experience struct {
	Description *string `json:"description"`
	Phone       *string `json:"phone"`
}
