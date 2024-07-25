package example_usecase

import (
	"github.com/charmingruby/swrc/internal/core"
	"github.com/charmingruby/swrc/internal/domain/example/example_entity"
)

func (s *ExampleUseCaseRegistry) GetExampleUseCase(id string) (*example_entity.Example, error) {
	example, err := s.exampleRepo.FindByID(id)
	if err != nil {
		return nil, core.NewNotFoundErr("example")
	}

	return example, nil
}
