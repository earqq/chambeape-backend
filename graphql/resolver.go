package graphql

import (
	"context"
	"errors"
	"fmt"
	"tuchamba/db"

	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	profiles *mgo.Collection
	jobs     *mgo.Collection
}

func New() Config {
	return Config{
		Resolvers: &Resolver{
			profiles: db.GetCollection("profiles"),
			jobs:     db.GetCollection("jobs"),
		},
	}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateProfile(ctx context.Context, input NewProfile) (*Profile, error) {
	var user Profile

	count, err := r.profiles.Find(bson.M{"email": input.Email}).Count()
	if err != nil {
		return &Profile{}, err
	} else if count > 0 {
		return &Profile{}, errors.New("user with that email already exists")
	}
	err = r.profiles.Insert(bson.M{"email": input.Email, "birthdate": input.Birthdate, "names": input.Names, "profile_type": input.ProfileType, "id_public": input.IDPublic, "phone": input.Phone, "img": input.Img})
	if err != nil {
		return &Profile{}, err
	}

	err = r.profiles.Find(bson.M{"email": input.Email}).One(&user)
	if err != nil {
		return &Profile{}, err
	}

	return &user, nil
}
func (r *mutationResolver) UpdateProfile(ctx context.Context, input UpdateProfile) (*Profile, error) {
	var fields = bson.M{}
	var user Profile

	update := false

	if &input.Names != nil && input.Names != "" {
		fields["names"] = input.Names
		update = true
	}
	if &input.Phone != nil && input.Phone != "" {
		fields["phone"] = input.Phone
		update = true
	}
	if &input.IDPublic != nil && input.IDPublic != "" {
		fields["id_public"] = input.IDPublic
		update = true
	}
	if &input.Img != nil && input.Img != "" {
		fields["img"] = input.Img
		update = true
	}
	if &input.Email != nil && input.Email != "" {
		fields["email"] = input.Email
		update = true
	}
	if &input.Birthdate != nil && input.Birthdate != "" {
		fields["birthdate"] = input.Birthdate
		update = true
	}
	if &input.ProfileType != nil {
		fields["profile_type"] = input.ProfileType
		update = true
	}

	if !update {
		return &Profile{}, errors.New("no fields present for updating data")
	}

	err := r.profiles.Update(bson.M{"id_public": input.IDPublic}, fields)
	if err != nil {
		fmt.Print("errorr", input.IDPublic)
		return &Profile{}, err
	}

	err = r.profiles.Find(bson.M{"id_public": input.IDPublic}).One(&user)
	if err != nil {
		return &Profile{}, err
	}
	user.ID = bson.ObjectId(user.ID).Hex()
	return &user, nil
}

func (r *mutationResolver) CreateJob(ctx context.Context, input NewJob) (*Job, error) {
	var job Job
	count, err := r.jobs.Find(bson.M{"id_public": input.IDPublic}).Count()
	if err != nil {
		return &Job{}, err
	} else if count > 0 {
		return &Job{}, errors.New("user with that id public already exists")
	}
	err = r.jobs.Insert(bson.M{"title": input.Title,
		"end_date":  input.EndDate,
		"job_type":  input.JobType,
		"id_public": input.IDPublic,
		"owner":     input.Owner,
		"price":     input.Price,
		"location":  input.Location,
		"tasks":     input.Tasks})
	err = r.jobs.Find(bson.M{"id_public": input.IDPublic}).One(&job)
	if err != nil {
		return &Job{}, err
	}

	return &job, nil
}
func (r *mutationResolver) UpdateJob(ctx context.Context, input UpdateJob) (*Job, error) {
	var fields = bson.M{}
	var job Job

	update := false

	if &input.Title != nil && input.Title != "" {
		fields["title"] = input.Title
		update = true
	}
	if &input.Tasks != nil {
		fields["tasks"] = input.Tasks
		update = true
	}
	if &input.IDPublic != nil && input.IDPublic != "" {
		fields["id_public"] = input.IDPublic
		update = true
	}
	if &input.JobType != nil {
		fields["job_type"] = input.JobType
		update = true
	}
	if &input.EndDate != nil && input.EndDate != "" {
		fields["end_date"] = input.EndDate
		update = true
	}
	if &input.Location != nil {
		fields["location"] = input.Location
		update = true
	}
	if &input.Price != nil {
		fields["price"] = input.Price
		update = true
	}
	if &input.Owner != nil {
		fields["owner"] = input.Owner
		update = true
	}
	if !update {
		return &Job{}, errors.New("no fields present for updating data")
	}

	err := r.jobs.Update(bson.M{"id_public": input.IDPublic}, fields)
	if err != nil {
		fmt.Print("errorr", input.IDPublic)
		return &Job{}, err
	}

	err = r.jobs.Find(bson.M{"id_public": input.IDPublic}).One(&job)
	if err != nil {
		return &Job{}, err
	}
	job.ID = bson.ObjectId(job.ID).Hex()
	return &job, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Profile(ctx context.Context, public_id string) (*Profile, error) {
	var user Profile

	if err := r.profiles.Find(bson.M{"id_public": public_id}).One(&user); err != nil {
		return &Profile{}, err
	}
	user.ID = bson.ObjectId(user.ID).Hex()

	return &user, nil
}
func (r *queryResolver) Profiles(ctx context.Context) ([]*Profile, error) {
	var profiles []*Profile
	r.profiles.Find(bson.M{}).All(&profiles)
	fmt.Print(profiles)
	return profiles, nil
}

func (r *queryResolver) Job(ctx context.Context, id_public string) (*Job, error) {
	var job Job

	if err := r.profiles.Find(bson.M{"id_public": id_public}).One(&job); err != nil {
		return &Job{}, err
	}
	job.ID = bson.ObjectId(job.ID).Hex()

	return &job, nil
}
func (r *queryResolver) Jobs(ctx context.Context) ([]*Job, error) {
	var jobs []*Job
	r.jobs.Find(bson.M{}).All(&jobs)
	return jobs, nil
}
