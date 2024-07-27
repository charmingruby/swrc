package example_usecase

import (
	"github.com/charmingruby/swrc/internal/example/domain/example_dto"
	"github.com/charmingruby/swrc/internal/example/domain/example_entity"
	"github.com/charmingruby/swrc/internal/example/domain/example_repository"
)

type ExampleUseCase interface {
	CreateExampleUseCase(dto example_dto.CreateExampleUseCaseDTO) error
	GetExampleUseCase(id string) (*example_entity.Example, error)
}

func NewExampleUseCaseRegistry(exampleRepo example_repository.ExampleRepository) *ExampleUseCaseRegistry {
	return &ExampleUseCaseRegistry{
		exampleRepo: exampleRepo,
	}
}

type ExampleUseCaseRegistry struct {
	exampleRepo example_repository.ExampleRepository
}
