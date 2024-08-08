package db

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"warehouse/internal/entity"
	"warehouse/internal/materials/repository"
	psql "warehouse/pkg/db/postgres"
)

// materialsRepo - Material Repository implementation
type materialsRepo struct {
	db *psql.Postgres
}

// NewMaterialsRepository - Material repository constructor
func NewMaterialsRepository(db *psql.Postgres) repository.Repository {
	return &materialsRepo{db: db}
}

// Create new material
func (m *materialsRepo) Create(ctx context.Context, material *entity.Material) error {
	typeID, err := m.getOrCreateType(ctx, material.Type)
	if err != nil {
		return err
	}

	statusID, err := m.getStatusID(ctx, material.Status)
	if err != nil {
		return err
	}

	query := `INSERT INTO materials (uuid, type_id, status_id, title, content, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = m.db.Conn.Exec(ctx, query, material.UUID, typeID, statusID, material.Name, material.Description, material.CreatedAt, material.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error inserting material: %w", err)
	}

	return nil
}

// GetByID
func (m *materialsRepo) GetByID(ctx context.Context, uuid int64) (*entity.Material, error) {
	var material entity.Material
	err := m.db.Conn.QueryRow(ctx, queryGetByID, uuid).Scan(&material)
	if err.Error() == "no rows in result set" {
		return nil, fmt.Errorf("material with id %s does not exists", uuid)
	}
	return &material, nil
}

func (m *materialsRepo) GetAll(ctx context.Context) ([]*entity.Material, error) {
	return nil, nil
}

// Update existing material except type field
func (m *materialsRepo) Update(ctx context.Context, material *entity.Material) (*entity.Material, error) {
	return &entity.Material{}, nil
}

// getOrCreateType - Проверка существования типа,если он не существует - создать новый в специальаной таблице
func (m *materialsRepo) getOrCreateType(ctx context.Context, typeName string) (int64, error) {
	var id int64

	// Проверка существования типа
	err := m.db.Conn.QueryRow(ctx, "SELECT id FROM material_types WHERE type_name = $1", typeName).Scan(&id)
	if err == nil {
		// Тип существует
		return id, nil
	}

	if err.Error() != "no rows in result set" {
		return 0, err
	}

	// Тип не существует, добавляем новый
	var lastInsertID int64
	err = m.db.Conn.QueryRow(ctx, "INSERT INTO material_types (type_name) VALUES ($1) RETURNING id", typeName).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (m *materialsRepo) getStatusID(ctx context.Context, statusName string) (int64, error) {
	var id int64

	// Проверка существования Статуса
	err := m.db.Conn.QueryRow(ctx, "SELECT id FROM material_statuses WHERE status = $1", statusName).Scan(&id)

	if err != nil {
		log.Err(err).Send()
		if err.Error() == "no rows in result set" {
			return 0, fmt.Errorf("status '%s' does not exists", statusName)
		}

		return 0, fmt.Errorf("db.getStatusID.fetching status id error")
	}
	// Статус существует
	return id, nil
}
