package models

type Profile struct {
	ID          string `json:"id"`
	IDPublic    string `json:"id_public" bson:"id_public"`
	ProfileType int    `json:"profile_type" bson:"profile_type"`
}
