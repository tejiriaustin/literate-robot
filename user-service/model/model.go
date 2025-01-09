package model

import (
	"github.com/tejiriaustin/literate-robot/core/model"
)

type (
	User struct {
		model.Base `gorm:",inline"`
		FirstName  string `gorm:"first_name"`
		LastName   string `gorm:"last_name"`
	}
)
