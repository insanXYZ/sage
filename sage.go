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

var (
	tagSupport   = []string{"png", "jpeg", "gif", "jpg", "bmp"}
	tagWithValue = []string{"minsize", "maxsize"}
)

func Validate(file interface{}, tag ...string) error {
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

func valid(file io.Reader, tags []string, size int64) error {
	var t string

	read, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	t = http.DetectContentType(read)

	splitType := strings.Split(t, "/")
	if len(tags) == 0 {
		if splitType[0] != "image" {
			return throw.InvalidFile
		}
	}

	if len(tags) > 0 {
		for _, singleTag := range tags {
			lowerStringSingleTag := strings.ToLower(singleTag)

			errInvalidTag := validTag(lowerStringSingleTag)
			if errInvalidTag != nil {
				return errInvalidTag
			}

			if slices.Contains(tagSupport, lowerStringSingleTag) {
				if splitType[1] != lowerStringSingleTag {
					return throw.InvalidType(lowerStringSingleTag)
				}
			}

			sliceContainWithValue := slices.ContainsFunc(tagWithValue, func(s string) bool {
				return strings.Contains(lowerStringSingleTag, s)
			})

			if strings.Contains(lowerStringSingleTag, "=") && sliceContainWithValue {
				str := strings.Split(lowerStringSingleTag, "=")
				atoi, errAtoi := strconv.Atoi(str[1])
				if errAtoi != nil {
					return errAtoi
				}

				switch strings.ToLower(str[0]) {
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

func validTag(tags string) error {
	if strings.Contains(tags, "=") {
		split := strings.Split(tags, "=")
		if !slices.Contains(tagWithValue, split[0]) {
			return throw.InvalidTag(tags)
		}
		tags = split[0]
	}

	if !slices.Contains(tagSupport, tags) || !slices.Contains(tagWithValue, tags) {
		return throw.InvalidTag(tags)
	}

	return nil
}
