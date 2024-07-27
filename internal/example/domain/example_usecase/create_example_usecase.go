package example_usecase

import (
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/example/domain/example_dto"
	"github.com/charmingruby/swrc/internal/example/domain/example_entity"
)

func (s *ExampleUseCaseRegistry) CreateExampleUseCase(dto example_dto.CreateExampleUseCaseDTO) error {
	example, err := example_entity.NewExample(dto.Name)
	if err != nil {
		return err
	}

	if err := s.exampleRepo.Store(example); err != nil {
		return core.NewInternalErr("create example store")
	}

	return nil
}
