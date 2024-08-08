package materials

import (
	"context"
	"warehouse/internal/entity"
)

type UseCase interface {
	Create(ctx context.Context, material *entity.Material) (string, error)
}
