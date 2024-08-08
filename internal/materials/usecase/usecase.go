package usecase

import (
	"context"
	"github.com/google/uuid"
	"time"
	"warehouse/internal/entity"
	"warehouse/internal/materials"
	"warehouse/internal/materials/repository"
)

// MaterialsUC - material use case implementation
type MaterialsUC struct {
	materialsRepo repository.Repository
}

// NewMaterialsUseCase - materials use case constructor
func NewMaterialsUseCase(materialsRepo repository.Repository) materials.UseCase {
	return &MaterialsUC{materialsRepo: materialsRepo}
}

// Create new material
func (m *MaterialsUC) Create(ctx context.Context, material *entity.Material) (string, error) {
	id, _ := uuid.NewUUID()

	material.UUID = id.String()
	material.CreatedAt = time.Now().Unix()
	material.UpdatedAt = time.Now().Unix()
	material.Status = "Активный"

	err := m.materialsRepo.Create(context.Background(), material)
	if err != nil {
		return "", err
	}

	return material.UUID, nil
}

func (m *MaterialsUC) GetByID(ctx context.Context, uuid int64) (*entity.Material, error) {

	return nil, nil
}

func (m *MaterialsUC) GetAll(ctx context.Context) ([]*entity.Material, error) {
	return nil, nil
}

// Update existing material except type field
func (m *MaterialsUC) Update(ctx context.Context, material *entity.Material) (*entity.Material, error) {
	return &entity.Material{}, nil
}
