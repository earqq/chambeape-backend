// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

type AddLocation struct {
	Latitude   *string `json:"latitude"`
	Longitude  *string `json:"longitude"`
	PostalCode *string `json:"postal_code"`
	Route      *string `json:"route"`
	Locality   *string `json:"locality"`
	AreaLevel1 *string `json:"area_level_1"`
	AreaLevel2 *string `json:"area_level_2"`
	Country    *string `json:"country"`
	ToSearch   *string `json:"to_search"`
}

type AddShares struct {
	Facebook *string `json:"facebook"`
	Whatsapp *string `json:"whatsapp"`
}

type AddWorker struct {
	WorkerType  *int         `json:"worker_type"`
	Description *string      `json:"description"`
	Location    *AddLocation `json:"location"`
	EndDate     *string      `json:"end_date"`
	Public      *bool        `json:"public"`
	Shares      *AddShares   `json:"shares"`
}

type Job struct {
	ID                 string   `json:"id"`
	Title              string   `json:"title"`
	IDPublic           string   `json:"id_public"`
	EndDate            string   `json:"end_date"`
	JobType            int      `json:"job_type"`
	JobTypeDescription string   `json:"job_type_description"`
	Visits             int      `json:"visits"`
	Calls              int      `json:"calls"`
	Validate           bool     `json:"validate"`
	State              bool     `json:"state"`
	Location           Location `json:"location"`
	Owner              JobOwner `json:"owner"`
}

type JobOwner struct {
	IDPublic string `json:"id_public"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Img      string `json:"img"`
}

type Location struct {
	Latitude   *string `json:"latitude"`
	Longitude  *string `json:"longitude"`
	PostalCode *string `json:"postal_code"`
	Route      *string `json:"route"`
	Locality   *string `json:"locality"`
	AreaLevel1 *string `json:"area_level_1"`
	AreaLevel2 *string `json:"area_level_2"`
	Country    *string `json:"country"`
	ToSearch   *string `json:"to_search"`
}

type NewJob struct {
	Title              string      `json:"title"`
	EndDate            string      `json:"end_date"`
	IDPublic           string      `json:"id_public"`
	State              bool        `json:"state"`
	Validate           bool        `json:"validate"`
	Visits             int         `json:"visits"`
	Calls              int         `json:"calls"`
	JobType            int         `json:"job_type"`
	JobTypeDescription string      `json:"job_type_description"`
	Location           AddLocation `json:"location"`
	Owner              NewJobOwner `json:"owner"`
}

type NewJobOwner struct {
	IDPublic string `json:"id_public"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Img      string `json:"img"`
}

type NewProfile struct {
	Email       *string `json:"email"`
	Names       string  `json:"names"`
	IDPublic    string  `json:"id_public"`
	Birthdate   *string `json:"birthdate"`
	Phone       *string `json:"phone"`
	ProfileType int     `json:"profile_type"`
	Img         *string `json:"img"`
}

type NewVideo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type Profile struct {
	ID             string `json:"id"`
	IDPublic       string `json:"id_public"`
	ProfileType    int    `json:"profile_type"`
	Names          string `json:"names"`
	Email          string `json:"email"`
	AvailableWeeks int    `json:"available_weeks"`
	Birthdate      string `json:"birthdate"`
	Phone          string `json:"phone"`
	Img            string `json:"img"`
	Worker         Worker `json:"worker"`
}

type Shares struct {
	Facebook *string `json:"facebook"`
	Whatsapp *string `json:"whatsapp"`
}

type UpdateJob struct {
	Title              *string      `json:"title"`
	IDPublic           string       `json:"id_public"`
	State              *bool        `json:"state"`
	EndDate            *string      `json:"end_date"`
	Validate           bool         `json:"validate"`
	JobType            *int         `json:"job_type"`
	JobTypeDescription *string      `json:"job_type_description"`
	Visits             *int         `json:"visits"`
	Calls              *int         `json:"calls"`
	Location           *AddLocation `json:"location"`
	Owner              *NewJobOwner `json:"owner"`
}

type UpdateProfile struct {
	IDPublic       string     `json:"id_public"`
	Names          *string    `json:"names"`
	Img            *string    `json:"img"`
	Email          *string    `json:"email"`
	AvailableWeeks *int       `json:"available_weeks"`
	Birthdate      *string    `json:"birthdate"`
	Phone          *string    `json:"phone"`
	ProfileType    *int       `json:"profile_type"`
	Worker         *AddWorker `json:"worker"`
}

type Video struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type Worker struct {
	WorkerType  *int      `json:"worker_type"`
	Description *string   `json:"description"`
	Location    *Location `json:"location"`
	EndDate     *string   `json:"end_date"`
	Public      *bool     `json:"public"`
	Shares      *Shares   `json:"shares"`
}
