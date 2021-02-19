package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// League stores a collection of players.
type League []Player

// NewLeague creates a league from JSON.
func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}

// Find tries to return a player from a league
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
