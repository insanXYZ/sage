package throw

import (
	"errors"
	"fmt"
)

type limit string

const (
	Maximum limit = "maximum"
	Minimal limit = "minimal"
)

var (
	InvalidFile = errors.New("invalid image file")
)

func InvalidType(t string) error {
	return errors.New(fmt.Sprintf("this image not a %s type", t))
}

func InvalidTag(t string) error {
	return errors.New(fmt.Sprintf("invalid tag format %s", t))
}

func InvalidSize(size int64, l limit) error {
	return errors.New(fmt.Sprintf("image size is outside of %s allowed size of %d KB", l, size))

}
