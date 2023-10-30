package mstore

import (
	"errors"
	"grapgql/graph/model"
)

type Service struct {
	StoryStore map[string]*model.Story
}

func NewService() Service {

	return Service{StoryStore: make(map[string]*model.Story)}
}

func (s Service) AddStory(story *model.Story) (*model.Story, error) {

	s.StoryStore[story.ID] = story
	return story, nil

}

func (s Service) UpdateStory(storyId string) (*model.Story, error) {
	val, ok := s.StoryStore[storyId]

	if !ok {
		return nil, errors.New("story not found")
	}

	return val, nil
}

func (s Service) FindAllStories() []*model.Story {

	var stories []*model.Story

	for _, v := range s.StoryStore {

		stories = append(stories, v)
	}

	return stories
}

func (s Service) FindStoryByID(storyId string) (*model.Story, error) {

	story, ok := s.StoryStore[storyId]

	if !ok {
		return nil, errors.New("story not found")
	}

	return story, nil

}
