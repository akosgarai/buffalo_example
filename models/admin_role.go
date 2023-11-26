package models

import (
	"github.com/gobuffalo/uuid"
)

type AdminRole struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Label       string    `json:"label" db:"label"`
	Description string    `json:"description" db:"description"`
}

// AdminRoles is not required by pop and may be deleted
type AdminRoles []AdminRole
