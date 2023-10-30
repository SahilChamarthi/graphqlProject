package mstore

import (
	"errors"
	"grapgql/graph/model"
)

type Service struct {
	storyStore map[string]*model.Story
}

func NewService() Service {

	return Service{storyStore: make(map[string]*model.Story)}
}

func (s Service) AddStory(story *model.Story) (*model.Story, error) {

	s.storyStore[story.ID] = story
	return story, nil

}

func (s Service) UpdateStory(storyId string, input model.UpdateStoryInput) (*model.Story, error) {
	val, ok := s.storyStore[storyId]

	if !ok {
		return nil, errors.New("story not found")
	}

	val.Content = *input.Content
	val.Title = *input.Title

	return val, nil
}

func (s Service) FindAllStories() []*model.Story {

	var stories []*model.Story

	for _, v := range s.storyStore {

		stories = append(stories, v)
	}

	return stories
}

func (s Service) FindStoryById(storyId string) (*model.Story, error) {

	story, ok := s.storyStore[storyId]

	if !ok {
		return nil, errors.New("story not found")
	}

	return story, nil

}
