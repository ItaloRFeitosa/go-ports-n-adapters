package uid

import (
	"fmt"

	"github.com/google/uuid"
)

type UID string

func (u UID) String() string {
	return string(u)
}

func (u UID) Valid() bool {
	return Validate(u.String())
}

func Random() UID {
	return UID(uuid.NewString())
}

func New(uidAsStr string) (UID, error) {
	if !Validate(uidAsStr) {
		return "", fmt.Errorf("can't parse, invalid uid string") // TODO: replace with value error or sentinel error
	}

	return UID(uidAsStr), nil
}

func Validate(uidAsStr string) bool {
	_, err := uuid.Parse(uidAsStr)

	return err == nil
}
