package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/pkg/errors"
)

type AdminPriv struct {
	ID              uuid.UUID `json:"id" db:"id"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	AdministratorID uuid.UUID `json:"administrator_id" db:"administrator_id"`
	PrivilegeID     uuid.UUID `json:"privilege_id" db:"privilege_id" form:"privilege_id"`
}

// String is not required by pop and may be deleted
func (a AdminPriv) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// AdminPrivs is a helper struct for deleting and updating privs.
type AdminPrivs struct {
	AdministratorID uuid.UUID
	PrivilegeIDs    []uuid.UUID
}

// deletePrivileges is responsible for deleting all the current privs for the given Administrator.
// After calling this function, every item will be deleted from the admin_privs table (where administrator_id = ?)
// All deletion errors, will be concatenated into one string, and it will be returned as the errormsg.
// In case of successful deletion, it returnes nil.
func (a *AdminPrivs) deletePrivileges(tx *pop.Connection) error {
	// get all privs for the given admin and then delete one by one
	var currentPrivs []AdminPriv
	err := tx.Where("administrator_id = ?", a.AdministratorID).All(&currentPrivs)
	if err != nil {
		return err
	}
	derr := tx.Destroy(currentPrivs)
	if derr != nil {
		return derr
	}
	return nil
}

// Delete is responsible for deleting all the privs. for a given admin
func (a *AdminPrivs) Delete(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), a.deletePrivileges(tx)
}

// Update is responsible for updating privs. for a given admin.
// Due to the current solution (exists. check), we need to delete all the current ones
// and then insert all the new ones.
func (a *AdminPrivs) Update(tx *pop.Connection) (*validate.Errors, error) {
	deletionErr := a.deletePrivileges(tx)
	if deletionErr != nil {
		return validate.NewErrors(), errors.WithStack(deletionErr)
	}
	// insert all new privileges to the table. I only have privilege_id and administrator_id here
	errMsg := ""
	for _, privId := range a.PrivilegeIDs {
		adminPriv := &AdminPriv{
			AdministratorID: a.AdministratorID,
			PrivilegeID:     privId,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			ID:              uuid.Must(uuid.NewV4()),
		}
		_, insertError := tx.ValidateAndCreate(adminPriv)
		if insertError != nil {
			errMsg = errMsg + " | " + insertError.Error()
		}
	}
	if errMsg != "" {
		return validate.NewErrors(), errors.New(errMsg)
	}
	return validate.NewErrors(), nil
}

// String is not required by pop and may be deleted
func (a AdminPrivs) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *AdminPriv) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *AdminPriv) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *AdminPriv) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
