package models

type Job struct {
	ID                 string   `json:"id"`
	Title              string   `json:"title"`
	IDPublic           string   `json:"id_public" bson:"id_public"`
	ContactPhone       string   `json:"contact_phone" bson:"contact_phone"`
	PublicationDate    string   `json:"publication_date" bson:"publication_date"`
	JobType            int      `json:"job_type" bson:"job_type"`
	JobTypeDescription string   `json:"job_type_description" bson:"job_type_description"`
	Visits             int      `json:"visits"`
	Calls              int      `json:"calls"`
	Validate           bool     `json:"validate"`
	State              bool     `json:"state"`
	Reports            int      `json:"reports"`
	Owner              Owner    `json:"owner"`
	Location           Location `json:"location"`
}
type Owner struct {
	Phone *string `json:"phone"`
}
type Location struct {
	Route    *string `json:"route"`
	Locality *string `json:"locality"`
	ToSearch *string `json:"to_search" bson:"to_search"`
}
