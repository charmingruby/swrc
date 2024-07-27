package example_repository

import "github.com/charmingruby/swrc/internal/example/domain/example_entity"

type ExampleRepository interface {
	Store(e *example_entity.Example) error
	FindByID(id string) (*example_entity.Example, error)
}
