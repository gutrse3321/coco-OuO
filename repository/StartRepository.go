package repository

import (
	"coco/common/datasource"
	"fmt"
)

type StartRepo struct {
	Source datasource.IDb `inject:""`
}

func (s *StartRepo) Speak(message string) string {
	return fmt.Sprintf("[Repository] speak: %s", message)
}
