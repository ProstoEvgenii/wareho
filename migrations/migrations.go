package migrations

import (
	"context"
	"fmt"
	"log"
	psql "warehouse/pkg/db/postgres"
)

const tableMaterialStatuses string = `CREATE TABLE IF NOT EXISTS material_statuses (id SERIAL PRIMARY KEY,status VARCHAR(50) UNIQUE NOT NULL);`

const insertStatusActive string = `INSERT INTO material_statuses (status) VALUES ('Архивный') ON CONFLICT (status) DO NOTHING;`
const insertStatusArchived string = `INSERT INTO material_statuses (status) VALUES ('Активный')ON CONFLICT (status) DO NOTHING;`

const tableMaterialTypes string = `CREATE TABLE IF NOT EXISTS material_types (
id SERIAL PRIMARY KEY,
type_name VARCHAR(50) UNIQUE NOT NULL
);`

const tableMaterialMaterials string = `CREATE TABLE IF NOT EXISTS materials (
uuid TEXT PRIMARY KEY,
type_id INT NOT NULL REFERENCES material_types(id),
status_id INT NOT NULL REFERENCES material_statuses(id),
title VARCHAR(255) NOT NULL,
content TEXT NOT NULL,
created_at NUMERIC NOT NULL,
updated_at NUMERIC NOT NULL
);`

func Run(ctx context.Context, db *psql.Postgres) error {
	var err error
	_, err = db.Conn.Exec(ctx, tableMaterialStatuses)
	_, err = db.Conn.Exec(ctx, insertStatusActive)
	_, err = db.Conn.Exec(ctx, insertStatusArchived)
	_, err = db.Conn.Exec(ctx, tableMaterialTypes)
	_, err = db.Conn.Exec(ctx, tableMaterialMaterials)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}
	log.Println("Migrations executed successfully")
	return nil
}
