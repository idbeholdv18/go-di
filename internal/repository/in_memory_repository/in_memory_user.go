package inmemoryrepository

import "github/idbeholdv18/go-di/internal/repository"

type InMemoryUserRepository struct {
	db []*repository.User
}

func NewInMemoryUserRepository(db []*repository.User) *InMemoryUserRepository {
	return &InMemoryUserRepository{
		db: db,
	}
}

func (r *InMemoryUserRepository) Create(name string, age int) *repository.User {
	candidate := &repository.User{
		Name: name,
		Age:  age,
	}
	r.db = append(r.db, candidate)
	return candidate
}

func (r *InMemoryUserRepository) FindByName(name string) *repository.User {
	for _, u := range r.db {
		if u.Name == name {
			return u
		}
	}
	return nil
}

func (r *InMemoryUserRepository) DeleteByName(name string) *repository.User {
	var candidate *repository.User
	var candidatePos int
	for i, u := range r.db {
		if u.Name == name {
			candidate = u
			candidatePos = i
		}
	}
	if candidate == nil {
		return nil
	}

	if candidatePos == 0 {
		r.db = r.db[1:]
	} else if candidatePos == len(r.db)-1 {
		r.db = r.db[:len(r.db)]
	} else {
		r.db = r.db[0:candidatePos:len(r.db)]
	}
	return candidate

}
