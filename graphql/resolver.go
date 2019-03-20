package graphql

import (
	"context"
	"errors"
	"tuchamba/db"

	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	profiles *mgo.Collection
}

func New() Config {
	return Config{
		Resolvers: &Resolver{
			profiles: db.GetCollection("profiles"),
		},
	}
}

func (r *Resolver) Mutation() MutationResolver {
	r.profiles = db.GetCollection("profiles")
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	r.profiles = db.GetCollection("profiles")
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

	err = r.profiles.Insert(bson.M{"email": input.Email, "names": input.Names, "token": input.Token, "phone": input.Phone, "img": input.Img})
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
		fields["first"] = *input.Names
		update = true
	}
	if input.Phone != nil && *input.Phone != "" {
		fields["last"] = *input.Phone
		update = true
	}
	if input.Token != nil && *input.Token != "" {
		fields["last"] = *input.Token
		update = true
	}
	if input.Img != nil && *input.Img != "" {
		fields["last"] = *input.Img
		update = true
	}
	if input.Email != nil && *input.Email != "" {
		fields["email"] = *input.Email
		update = true
	}

	if !update {
		return &Profile{}, errors.New("no fields present for updating data")
	}

	err := r.profiles.UpdateId(bson.ObjectIdHex(input.ID), fields)
	if err != nil {
		return &Profile{}, err
	}

	err = r.profiles.Find(bson.M{"_id": bson.ObjectIdHex(input.ID)}).One(&user)
	if err != nil {
		return &Profile{}, err
	}
	user.ID = bson.ObjectId(user.ID).Hex()
	return &user, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Profile(ctx context.Context, id string) (*Profile, error) {
	var user Profile

	if err := r.profiles.FindId(bson.ObjectIdHex(id)).One(&user); err != nil {
		return &Profile{}, err
	}

	user.ID = bson.ObjectId(user.ID).Hex()

	return &user, nil
}
func (r *queryResolver) Profiles(ctx context.Context) ([]*Profile, error) {
	var profiles []*Profile
	r.profiles.Find(bson.M{}).All(&profiles)
	return profiles, nil
}
