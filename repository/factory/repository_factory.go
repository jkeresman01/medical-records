package repositoryfactory

import (
	"reflect"
	"sync"

	"github.com/jkeresman01/medical-records/repository"
)

type factory struct {
	repositories map[reflect.Type]any
	mu           sync.Mutex
}

var Factory = &factory{
	repositories: make(map[reflect.Type]any),
}

func GetInstance[T any]() *repository.Repository[T] {
	Factory.mu.Lock()
	defer Factory.mu.Unlock()

	t := reflect.TypeOf((*T)(nil)).Elem()

	if repo, ok := Factory.repositories[t]; ok {
		return repo.(*repository.Repository[T])
	}

	newRepo := repository.NewRepository[T]()
	Factory.repositories[t] = newRepo
	return newRepo
}
