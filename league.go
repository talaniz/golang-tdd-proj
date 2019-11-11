package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// League manages players and scores
type League []Player

// NewLeague parses json and returns a list of players
func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}

// Find locates returns a player in a League
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
