package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/pop/nulls"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Job struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	Name        string       `json:"name" db:"name"`
	Category    nulls.String `json:"category" db:"category"`
	Description string       `json:"description" db:"description"`
	Salary      string       `json:"salary" db:"salary"`
}

// String is not required by pop and may be deleted
func (j Job) String() string {
	jj, _ := json.Marshal(j)
	return string(jj)
}

// Jobs is not required by pop and may be deleted
type Jobs []Job

// String is not required by pop and may be deleted
func (j Jobs) String() string {
	jj, _ := json.Marshal(j)
	return string(jj)
}

// Validate gets run every time you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (j *Job) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: j.Name, Name: "Name"},
		&validators.StringIsPresent{Field: j.Description, Name: "Description"},
		&validators.StringIsPresent{Field: j.Salary, Name: "Salary"},
	), nil
}

// ValidateSave gets run every time you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (j *Job) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (j *Job) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
