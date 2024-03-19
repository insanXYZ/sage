package sage

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/insanXYZ/sage/throw"
)

type Sage struct{}

func New() *Sage {
	return &Sage{}
}

var (
	tagSupport   = []string{"png", "jpeg", "gif", "jpg"}
	tagWithValue = []string{"minsize", "maxsize"}
)

func (s *Sage) Validate(file interface{}, tag ...string) error {

	switch t := file.(type) {
	case *os.File:
		stat, _ := t.Stat()
		t.Seek(0, 0)
		return valid(t, tag, stat.Size())
	case *multipart.FileHeader:
		open, err := t.Open()
		if err != nil {
			return err
		}
		defer open.Close()
		defer open.Seek(0, 0)
		return valid(open, tag, t.Size)
	}

	return throw.InvalidFile
}

func valid(file io.Reader, tag []string, size int64) error {
	var t string

	read, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	t = http.DetectContentType(read)

	splitType := strings.Split(t, "/")
	if len(tag) == 0 {
		if splitType[0] != "image" {
			return throw.InvalidFile
		}
	}

	if len(tag) > 0 {
		for _, s := range tag {
			ls := strings.ToLower(s)

			err := validTag(ls)
			if err != nil {
				return err
			}

			if slices.Contains(tagSupport, ls) {
				if splitType[1] != ls {
					return throw.InvalidType(ls)
				}
			}

			if strings.Contains(ls, "=") && slices.ContainsFunc(tagWithValue, func(s string) bool {
				return strings.Contains(ls, s)
			}) {
				s := strings.Split(ls, "=")
				atoi, err := strconv.Atoi(s[1])
				if err != nil {
					return err
				}

				switch strings.ToLower(s[0]) {
				case "minsize":
					if int(size/1024) < atoi {
						return throw.InvalidSize(int64(atoi), throw.Minimal)
					}
				case "maxsize":
					if int(size/1024) > atoi {
						return throw.InvalidSize(int64(atoi), throw.Maximum)
					}
				}
			}

		}
	}

	return nil
}

func validTag(tag string) error {
	if slices.Contains(tagSupport, tag) || slices.Contains(tagWithValue, tag) {
		return nil
	}
	return throw.InvalidTag(tag)
}
