package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/BurrrY/obstwiesen-server/graph/model"
	gonanoid "github.com/matoous/go-nanoid/v2"
	log "github.com/sirupsen/logrus"
)

// Trees is the resolver for the trees field.
func (r *meadowResolver) Trees(ctx context.Context, obj *model.Meadow) ([]*model.Tree, error) {
	var meadows []*model.Tree
	meadows, err := storage.GetTreesOfMeadow(obj.ID)
	return meadows, err
}

// CreateMeadow is the resolver for the createMeadow field.
func (r *mutationResolver) CreateMeadow(ctx context.Context, input model.NewMeadow) (*model.Meadow, error) {
	id, _ := gonanoid.New()

	meadow := &model.Meadow{
		ID:    id,
		Name:  input.Name,
		Trees: nil,
	}
	storage.StoreMeadow(meadow)
	return meadow, nil
}

// CreateTree is the resolver for the createTree field.
func (r *mutationResolver) CreateTree(ctx context.Context, input model.NewTree) (*model.Tree, error) {
	id, _ := gonanoid.New()

	tree := &model.Tree{
		ID:   id,
		Name: input.Name,
	}

	storage.AddTree(tree, input.MeadowID)
	return tree, nil
}

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, input model.NewEvent) (*model.Event, error) {
	eventID, _ := gonanoid.New()

	log.Info(input.Files)

	elemnt := &model.Event{
		ID:          eventID,
		Title:       input.Title,
		Description: *input.Description,
		Timestamp:   input.Timestamp,
	}
	err := storage.AddEvent(elemnt, input.ParentID)

	for _, file := range input.Files {
		fileID, _ := gonanoid.New()
		err = filestore.StoreFile(file, eventID, fileID)
		//	err := storage.FileToEvent(eventID, fileID)
	}

	return elemnt, err
}

// SingleUpload is the resolver for the singleUpload field.
func (r *mutationResolver) SingleUpload(ctx context.Context, parentID string, file graphql.Upload) (*model.File, error) {
	log.Info("SingleUpload - singleUpload")
	log.Info(file)
	log.Info(parentID)
	return &model.File{
		ParentID: "0",
		Path:     file.Filename,
	}, nil
}

// MultipleUpload is the resolver for the multipleUpload field.
func (r *mutationResolver) MultipleUpload(ctx context.Context, parentID string, files []*graphql.Upload) ([]*model.File, error) {
	log.Info("MultipleUpload - multipleUpload")

	log.Info(files)
	log.Info(parentID)

	return []*model.File{
		&model.File{
			ParentID: "0",
			Path:     "file.Filename",
		},
	}, nil
}

// Meadow is the resolver for the meadow field.
func (r *queryResolver) Meadow(ctx context.Context, meadowID string) (*model.Meadow, error) {
	var meadow *model.Meadow
	meadow, err := storage.GetMeadowByID(meadowID)
	return meadow, err
}

// Meadows is the resolver for the meadows field.
func (r *queryResolver) Meadows(ctx context.Context) ([]*model.Meadow, error) {
	var meadows []*model.Meadow
	meadows, err := storage.GetMeadows()
	return meadows, err
}

// Trees is the resolver for the trees field.
func (r *queryResolver) Trees(ctx context.Context, meadowID string) ([]*model.Tree, error) {
	var trees []*model.Tree
	trees, err := storage.GetTreesOfMeadow(meadowID)
	return trees, err
}

// Tree is the resolver for the tree field.
func (r *queryResolver) Tree(ctx context.Context, treeID string) (*model.Tree, error) {
	var tree *model.Tree
	log.Info("Tree requestd", treeID)
	tree, err := storage.GetTreeByID(treeID)
	return tree, err
}

// Events is the resolver for the events field.
func (r *queryResolver) Events(ctx context.Context, treeID string) ([]*model.Event, error) {
	var elem []*model.Event
	elem, err := storage.GetEventsOfTree(treeID)

	for _, event := range elem {
		filesList, _ := filestore.GetFiles(event.ID)
		event.Files = filesList
	}

	return elem, err
}

// Events is the resolver for the events field.
func (r *treeResolver) Events(ctx context.Context, obj *model.Tree) ([]*model.Event, error) {
	var elem []*model.Event
	elem, err := storage.GetEventsOfTree(obj.ID)
	for _, event := range elem {
		filesList, _ := filestore.GetFiles(event.ID)
		event.Files = filesList
	}
	return elem, err
}

// Meadow returns MeadowResolver implementation.
func (r *Resolver) Meadow() MeadowResolver { return &meadowResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Tree returns TreeResolver implementation.
func (r *Resolver) Tree() TreeResolver { return &treeResolver{r} }

type meadowResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type treeResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) SingleUpload(ctx context.Context, file graphql.Upload) (bool, error) {
	log.Info("received file ", file.File)
	return true, nil
}
func (r *queryResolver) MultipleUpload(ctx context.Context, parentID string, req []*model.UploadFile) ([]*model.File, error) {
	log.Info("received file ", req)
	return nil, nil
}
