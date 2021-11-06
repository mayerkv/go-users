package repository

import (
	"github.com/mayerkv/go-users/domain"
	"sync"
)

type InMemoryUserRepository struct {
	sync.Mutex
	items map[string]domain.User
}

func NewInMemoryUserRepository() domain.UserRepository {
	return &InMemoryUserRepository{
		items: map[string]domain.User{},
	}
}

func (r *InMemoryUserRepository) Save(user *domain.User) error {
	r.Lock()
	defer r.Unlock()

	r.items[user.Id] = *user

	return nil
}

func (r *InMemoryUserRepository) FindById(id string) (*domain.User, error) {
	r.Lock()
	defer r.Unlock()

	if user, ok := r.items[id]; ok {
		return &user, nil
	}

	return nil, nil
}

func (r *InMemoryUserRepository) FindByEmail(email string) (*domain.User, error) {
	r.Lock()
	defer r.Unlock()

	for _, item := range r.items {
		if item.Email == email {
			return &item, nil
		}
	}

	return nil, nil
}

func (r *InMemoryUserRepository) Delete(user *domain.User) error {
	r.Lock()
	defer r.Unlock()

	delete(r.items, user.Id)

	return nil
}
