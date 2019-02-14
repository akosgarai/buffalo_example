package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Administrator struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`

	Pwd             string `json:"-" db:"-"`
	PwdConfirmation string `json:"-" db:"-"`
}

// Create wraps up the pattern of encrypting the password and
// running validations. Useful when writing tests.
func (a *Administrator) Create(tx *pop.Connection) (*validate.Errors, error) {
	a.Email = strings.ToLower(a.Email)
	ph, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return validate.NewErrors(), errors.WithStack(err)
	}
	a.Password = string(ph)
	return tx.ValidateAndCreate(a)
}

// String is not required by pop and may be deleted
func (a Administrator) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Administrators is not required by pop and may be deleted
type Administrators []Administrator

// String is not required by pop and may be deleted
func (a Administrators) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
func (a *Administrator) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: a.Name, Name: "Name"},
		&validators.StringIsPresent{Field: a.Username, Name: "Username"},
		&validators.StringIsPresent{Field: a.Password, Name: "Password"},
		&validators.StringIsPresent{Field: a.Email, Name: "Email"},
		// check to see if the email address is already taken:
		&validators.FuncValidator{
			Field:   a.Email,
			Name:    "Email",
			Message: "%s is already taken",
			Fn: func() bool {
				var b bool
				q := tx.Where("email = ?", a.Email)
				if a.ID != uuid.Nil {
					q = q.Where("id != ?", a.ID)
				}
				b, err = q.Exists(a)
				if err != nil {
					return false
				}
				return !b
			},
		},
	), err
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
func (a *Administrator) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: a.Pwd, Name: "Pwd"},
		&validators.StringsMatch{Name: "Pwd", Field: a.Pwd, Field2: a.PwdConfirmation, Message: "Password does not match confirmation"},
	), err
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Administrator) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
