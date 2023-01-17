package validator

import (
	"errors"

	"github.com/PhononDAO/phonon-core/pkg/model"
)

var ErrMissingPubKey = errors.New("phonon missing public key")

// Validates that a phonon's presented public key represents an actual crypto asset
type Validator interface {
	Validate(phonon *model.Phonon) (valid bool, err error)
}
