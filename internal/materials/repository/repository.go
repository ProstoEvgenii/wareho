package repository

import (
	"context"
	"warehouse/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, material *entity.Material) error
	GetByID(ctx context.Context, uuid int64) (*entity.Material, error)
	GetAll(ctx context.Context) ([]*entity.Material, error)
	Update(ctx context.Context, material *entity.Material) (*entity.Material, error)
}
