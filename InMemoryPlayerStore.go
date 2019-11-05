package main

// NewInMemoryPlayerStore generates a new InMemoryPlayerStore instance
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// InMemoryPlayerStore is a storage implementation
type InMemoryPlayerStore struct {
	store map[string]int
}

// GetPlayerScore retrieves a player score
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// RecordWin processes wins, implements PlayerStore interface
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// GetLeague returns a slice of players that belong to a league
func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player

	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
