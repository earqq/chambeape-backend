package graphql

import (
	"chambeape/db"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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

	count, err := r.profiles.Find(bson.M{"id_public": input.IDPublic}).Count()
	if err != nil {
		return &Profile{}, err
	} else if count > 0 {
		return &Profile{}, errors.New("user with that id public already exists")
	}
	err = r.profiles.Insert(bson.M{"email": input.Email,
		"birthdate":       input.Birthdate,
		"names":           input.Names,
		"profile_type":    input.ProfileType,
		"id_public":       input.IDPublic,
		"phone":           input.Phone,
		"available_weeks": 4,
		"updated_at":      time.Now().Local(),
		"img":             input.Img})
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

	if input.Names != nil && *input.Names != "" {
		fields["names"] = *input.Names
		update = true
	}
	if input.Phone != nil && *input.Phone != "" {
		fields["phone"] = *input.Phone
		update = true
	}
	if &input.IDPublic != nil && input.IDPublic != "" {
		fields["id_public"] = input.IDPublic
		update = true
	}
	if input.Img != nil && *input.Img != "" {
		fields["img"] = *input.Img
		update = true
	}
	if input.Email != nil && *input.Email != "" {
		fields["email"] = *input.Email
		update = true
	}
	if input.Birthdate != nil && *input.Birthdate != "" {
		fields["birthdate"] = *input.Birthdate
		update = true
	}
	if input.ProfileType != nil {
		fields["profile_type"] = *input.ProfileType
		update = true
	}
	if input.AvailableWeeks != nil {
		fields["available_weeks"] = *input.AvailableWeeks
		update = true
	}
	if input.Worker != nil {
		fields["worker"] = *input.Worker
		update = true
	}
	if !update {
		return &Profile{}, errors.New("no fields present for updating data")
	}
	fields["updated_at"] = time.Now().Local()
	err := r.profiles.Update(bson.M{"id_public": input.IDPublic}, bson.M{"$set": fields})
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
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	if input.Location.ToSearch != nil {
		upperLocality, _, _ := transform.String(t, *input.Location.ToSearch)
		*input.Location.ToSearch = strings.ToLower(upperLocality)
	}
	err = r.jobs.Insert(bson.M{"title": strings.ToLower(input.Title),
		"end_date":             input.EndDate,
		"job_type":             input.JobType,
		"job_type_description": strings.ToLower(input.JobTypeDescription),
		"id_public":            input.IDPublic,
		"owner":                input.Owner,
		"visits":               input.Visits,
		"validate":             input.Validate,
		"calls":                input.Calls,
		"state":                input.State,
		"location":             input.Location,
		"updated_at":           time.Now().Local()})
	err = r.jobs.Find(bson.M{"id_public": input.IDPublic}).One(&job)
	if err != nil {
		return &Job{}, err
	}

	return &job, nil
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
func (r *mutationResolver) UpdateJob(ctx context.Context, input UpdateJob) (*Job, error) {
	var fields = bson.M{}
	var job Job

	update := false
	newUpdatedAt := false
	if input.Title != nil && *input.Title != "" {
		fields["title"] = strings.ToLower(*input.Title)
		update = true
		newUpdatedAt = true
	}

	if input.Calls != nil {
		fields["calls"] = *input.Calls
		update = true
	}
	if input.Visits != nil {
		fields["visits"] = *input.Visits
		update = true
	}
	if &input.IDPublic != nil && input.IDPublic != "" {
		fields["id_public"] = input.IDPublic
		update = true
	}
	if input.JobType != nil {
		fields["job_type"] = *input.JobType
		update = true
		newUpdatedAt = true
	}
	if input.JobTypeDescription != nil && *input.JobTypeDescription != "" {
		fields["job_type_description"] = strings.ToLower(*input.JobTypeDescription)
		update = true
		newUpdatedAt = true
	}
	if input.EndDate != nil && *input.EndDate != "" {
		fields["end_date"] = *input.EndDate
		update = true
		newUpdatedAt = true
	}
	if input.Location != nil {
		t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
		if input.Location.ToSearch != nil && *input.Location.ToSearch != "" {
			upperToSearch, _, _ := transform.String(t, *input.Location.ToSearch)
			*input.Location.ToSearch = strings.ToLower(upperToSearch)
		}
		fields["location"] = *input.Location
		update = true
		newUpdatedAt = true
	}
	if input.State != nil {
		fields["state"] = *input.State
		update = true
		newUpdatedAt = true
	}

	if input.Owner != nil {
		fields["owner"] = *input.Owner
		update = true
		newUpdatedAt = true
	}

	if !update {
		return &Job{}, errors.New("no fields present for updating data")
	}
	if newUpdatedAt {
		fields["updated_at"] = time.Now().Local()
	}
	err := r.jobs.Update(bson.M{"id_public": input.IDPublic}, bson.M{"$set": fields})
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
func (r *queryResolver) Profiles(ctx context.Context, limit int, profile_type *int, search *string, worker_type *int, random *bool) ([]*Profile, error) {
	var profiles []*Profile
	var fields = bson.M{}
	if profile_type != nil {
		fields["profile_type"] = profile_type
	}

	if search != nil {
		fields["$or"] = []bson.M{
			bson.M{"worker.location.to_search": bson.M{"$regex": strings.ToLower(*search)}},
			bson.M{"names": bson.M{"$regex": strings.ToLower(*search)}}}
	}
	if worker_type != nil {
		fields["worker.worker_type"] = worker_type
	}
	r.profiles.Find(fields).Limit(limit).Sort("-updated_at").All(&profiles)
	if random != nil {
		ShuffleProfile(profiles)
	}
	return profiles, nil
}

func (r *queryResolver) Job(ctx context.Context, id_public string) (*Job, error) {
	var job Job
	fmt.Println("probando ando2 ")
	if err := r.jobs.Find(bson.M{"id_public": id_public}).One(&job); err != nil {
		return &Job{}, err
	}
	job.ID = bson.ObjectId(job.ID).Hex()

	return &job, nil
}
func (r *queryResolver) Jobs(ctx context.Context, profileIDPublic *string, endDate *string, state *bool, search *string, limit int, job_type *int, random *bool) ([]*Job, error) {
	var jobs []*Job
	var fields = bson.M{}
	if endDate != nil {
		fields["end_date"] = bson.M{"$lt": endDate}
	}
	if state != nil {
		arr := []*bool{state}
		fields["state"] = bson.M{"$in": arr}
	}
	if search != nil {
		fields["$or"] = []bson.M{
			bson.M{"location.to_search": bson.M{"$regex": strings.ToLower(*search)}},
			bson.M{"job_type_description": bson.M{"$regex": strings.ToLower(*search)}},
			bson.M{"title": bson.M{"$regex": strings.ToLower(*search)}}}
	}
	if profileIDPublic != nil {
		fields["owner.id_public"] = profileIDPublic
	}
	if job_type != nil {
		fields["job_type"] = job_type
	}
	r.jobs.Find(fields).Limit(limit).Sort("-updated_at").All(&jobs)
	if random != nil {
		Shuffle(jobs)
	}
	return jobs, nil
}
func Shuffle(slc []*Job) {
	for i := 1; i < len(slc); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slc[r], slc[i] = slc[i], slc[r]
		}
	}
}
func ShuffleProfile(slc []*Profile) {
	for i := 1; i < len(slc); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slc[r], slc[i] = slc[i], slc[r]
		}
	}
}
