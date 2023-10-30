package store

import "grapgql/graph/model"

type Storer interface {
	AddStory(story model.Story) (*model.Story, error)
	UpdateStory(storyId string, input model.UpdateStoryInput) (*model.Story, error)
	FindAllStories() []*model.Story
	FindStoryById(storyId string) (*model.Story, error)
}

type Store struct {
	Storer
}

func NewStore(storer Storer) Store {

	return Store{Storer: storer}
}
