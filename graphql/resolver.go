package graphql

import (
	"chambeape/db"
	"chambeape/models"
	"context"
	"errors"
	"math/rand"
	"strings"
	"time"
	"unicode"

	"github.com/globalsign/mgo"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"gopkg.in/mgo.v2/bson"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

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

func (r *Resolver) Job() JobResolver {
	return &jobResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Profile() ProfileResolver {
	return &profileResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type jobResolver struct{ *Resolver }

func (r *jobResolver) Location(ctx context.Context, obj *models.Job) (*Location, error) {
	var location Location
	location.Route = obj.Location.Route
	location.Locality = obj.Location.Locality
	return &location, nil
}
func (r *jobResolver) Owner(ctx context.Context, obj *models.Job) (*Owner, error) {
	var owner Owner
	owner.Phone = obj.Owner.Phone
	return &owner, nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateProfile(ctx context.Context, input NewProfile) (*models.Profile, error) {
	var user models.Profile
	err := r.profiles.Insert(bson.M{
		"id_public":  input.IDPublic,
		"names":      input.Names,
		"email":      input.Email,
		"img":        input.Img,
		"updated_at": time.Now().Local()})
	if err != nil {
		return &models.Profile{}, err
	}

	err = r.profiles.Find(bson.M{"id_ublic": input.IDPublic}).One(&user)
	if err != nil {
		return &models.Profile{}, err
	}

	return &user, nil
}
func (r *mutationResolver) UpdateProfile(ctx context.Context, idPublic string, input UpdateProfile) (*models.Profile, error) {
	var fields = bson.M{}
	var profile models.Profile
	update := false
	if input.Phone != "" {
		count, err := r.profiles.Find(bson.M{"phone": input.Phone, "id_public": bson.M{"$ne": idPublic}}).Count()
		if err != nil {
			return &models.Profile{}, err
		} else if count > 0 {
			return &models.Profile{}, errors.New("Número de celular ya esta registrado en otro perfil")
		}
		fields["phone"] = input.Phone
		update = true
	}
	if input.Birthdate != nil && *input.Birthdate != "" {
		fields["birthdate"] = *input.Birthdate
		update = true
	}
	if input.Names != nil {
		fields["names"] = input.Names
		update = true
	}
	if input.Img != nil {
		fields["img"] = input.Img
		update = true
	}
	if input.Email != nil {
		fields["email"] = input.Email
		update = true
	}
	if input.Location != nil {
		fields["location"] = *input.Location
		update = true
	}
	if input.WorkerType != nil {
		fields["worker_type"] = *input.WorkerType
		update = true
	}
	if input.WorkerDescription != nil {
		fields["worker_description"] = *input.WorkerDescription
		update = true
	}
	if input.WorkerPublic != nil {
		fields["worker_public"] = *input.WorkerPublic
		update = true
	}
	if input.WorkerExperience != nil {
		fields["worker_experience"] = input.WorkerExperience
		update = true
	}
	update = true
	if !update {
		return &models.Profile{}, errors.New("No hay campos para actualizar")
	}
	fields["updated_at"] = time.Now().Local()
	r.profiles.Update(bson.M{"id_public": idPublic}, bson.M{"$set": fields})

	err := r.profiles.Find(bson.M{"id_public": idPublic}).One(&profile)
	if err != nil {
		return &models.Profile{}, err
	}
	return &profile, nil
}
func (r *mutationResolver) CreateJob(ctx context.Context, input NewJob) (*models.Job, error) {
	var job models.Job
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	if input.Location.ToSearch != nil {
		upperLocality, _, _ := transform.String(t, *input.Location.ToSearch)
		*input.Location.ToSearch = strings.ToLower(upperLocality)
	}
	tim := time.Now().Local()
	var owner models.Owner
	owner.Phone = &input.ContactPhone
	err := r.jobs.Insert(bson.M{"title": strings.ToLower(input.Title),
		"publication_date":     tim.Format("2006-01-02"),
		"job_type":             input.JobType,
		"job_type_description": strings.ToLower(input.JobTypeDescription),
		"id_public":            input.IDPublic,
		"owner":                owner,
		"calls":                0,
		"state":                true,
		"location":             input.Location,
		"updated_at":           time.Now().Local()})
	err = r.jobs.Find(bson.M{"id_public": input.IDPublic}).One(&job)
	if err != nil {
		return &models.Job{}, errors.New("No se puedo crear Trabajo")
	}

	return &job, nil

}
func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: sinespacios marks
}
func (r *mutationResolver) UpdateJob(ctx context.Context, idPublic string, input UpdateJob) (*models.Job, error) {
	var fields = bson.M{}
	var job models.Job

	update := false
	if input.Calls != nil {
		fields["calls"] = *input.Calls
		update = true
	}
	if input.State != nil {
		fields["state"] = *input.State
		update = true
		fields["updated_at"] = time.Now().Local()
	}
	if input.Reports != nil {
		fields["reports"] = *input.Reports
		update = true
	}
	if !update {
		return &models.Job{}, errors.New("No hay campos para actualizar")
	}
	err := r.jobs.Update(bson.M{"id_public": idPublic}, bson.M{"$set": fields})
	if err != nil {
		return &models.Job{}, err
	}

	err = r.jobs.Find(bson.M{"id_public": idPublic}).One(&job)
	if err != nil {
		return &models.Job{}, err
	}
	return &job, nil
}

type profileResolver struct{ *Resolver }

func (r *profileResolver) Location(ctx context.Context, obj *models.Profile) (*Location, error) {
	var location Location
	location.Route = obj.Location.Route
	location.Locality = obj.Location.Locality
	return &location, nil
}
func (r *profileResolver) WorkerExperience(ctx context.Context, obj *models.Profile) ([]Experience, error) {
	var experiences []Experience
	for i := 0; i < len(obj.WorkerExperience); i++ {
		var favorite Experience
		favorite.Description = obj.WorkerExperience[i].Description
		favorite.Phone = obj.WorkerExperience[i].Phone
		experiences = append(experiences, favorite)
	}
	return experiences, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Profile(ctx context.Context, idPublic *string, phone *string, workerPublic *bool) (*models.Profile, error) {
	var profile models.Profile
	var fields = bson.M{}
	if phone != nil {
		fields["phone"] = phone
	}
	if workerPublic != nil {
		fields["worker_public"] = workerPublic
	}
	if idPublic != nil {
		fields["id_public"] = idPublic
	}
	r.profiles.Find(fields).One(&profile)
	return &profile, nil
}
func (r *queryResolver) Profiles(ctx context.Context, limit int, search *string, workerType *int, random *bool, workerPublic *bool) ([]*models.Profile, error) {
	//Actualziar perfiles
	type OldWorker struct {
		WorkerType  int                 `json:"worker_type" bson:"worker_type"`
		Description string              `json:"description" bson:"description"`
		Location    models.Location     `json:"location" bson:"location"`
		Public      bool                `json:"public" bson:"public"`
		Experience  []models.Experience `json:"experience" bson:"experience"`
	}
	type oldProfiles struct {
		IDPublic  string `json:"id_public" bson:"id_public"`
		Names     string `json:"names" bson:"names"`
		Img       string `json:"img" bson:"worker_type"`
		Phone     string
		Email     string
		Birthdate string
		Worker    OldWorker `json:"worker" bson:"worker"`
	}
	var allProfiles []oldProfiles
	r.profiles.Find(bson.M{}).All(&allProfiles)
	for i := 0; i < len(allProfiles); i++ {
		var fields = bson.M{}

		fields["worker_type"] = &allProfiles[i].Worker.WorkerType
		fields["worker_description"] = &allProfiles[i].Worker.Description
		fields["worker_public"] = &allProfiles[i].Worker.Public
		fields["worker_experience"] = allProfiles[i].Worker.Experience
		fields["location"] = allProfiles[i].Worker.Location
		r.profiles.Update(bson.M{"id_public": allProfiles[i].IDPublic}, bson.M{"$set": fields})
	}
	var profiles []*models.Profile
	var fields = bson.M{}

	if search != nil {
		fields["names"] = bson.M{"$regex": *search, "$options": "i"}
	}
	if workerType != nil {
		fields["worker_type"] = workerType
	}
	if workerPublic != nil {
		fields["worker_public"] = workerPublic
	}
	r.profiles.Find(fields).Limit(limit).Sort("-updated_at").All(&profiles)
	if random != nil {
		ShuffleProfiles(profiles)
	}
	return profiles, nil
}
func (r *queryResolver) Job(ctx context.Context, idPublic string) (*models.Job, error) {
	var job models.Job
	if err := r.jobs.Find(bson.M{"id_public": idPublic}).One(&job); err != nil {
		return &models.Job{}, err
	}
	job.ID = bson.ObjectId(job.ID).Hex()

	return &job, nil

}
func (r *queryResolver) Jobs(ctx context.Context, profileIDPublic *string, state *bool, search *string, limit int, jobType *int, random *bool) ([]*models.Job, error) {
	var jobs []*models.Job
	var fields = bson.M{}
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
	if jobType != nil {
		fields["job_type"] = jobType
	}
	r.jobs.Find(fields).Limit(limit).Sort("-updated_at").All(&jobs)
	if random != nil {
		Shuffle(jobs)
	}
	return jobs, nil

}
func Shuffle(slc []*models.Job) {
	for i := 1; i < len(slc); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slc[r], slc[i] = slc[i], slc[r]
		}
	}
}
func ShuffleProfiles(slc []*models.Profile) {
	for i := 1; i < len(slc); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slc[r], slc[i] = slc[i], slc[r]
		}
	}
}
