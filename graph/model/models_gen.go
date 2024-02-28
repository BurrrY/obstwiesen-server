// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Event struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
}

type Meadow struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Trees []*Tree `json:"trees"`
}

type Mutation struct {
}

type NewEvent struct {
	TreeID      string `json:"treeID"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type NewMeadow struct {
	Name string `json:"name"`
}

type NewTree struct {
	Name     string `json:"name"`
	MeadowID string `json:"meadowID"`
}

type Query struct {
}

type Tree struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Lat    *float64 `json:"lat,omitempty"`
	Lang   *float64 `json:"lang,omitempty"`
	Events []*Event `json:"events"`
}