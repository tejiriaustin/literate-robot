package model

import (
	"time"

	"github.com/tejiriaustin/literate-robot/core/model"
)

type (
	User struct {
		model.Base         `gorm:",inline"`
		FirstName          string `gorm:"first_name"`
		LastName           string `gorm:"last_name"`
		VerificationStatus string `gorm:"verification_status"`
		Password           string `gorm:"password"`
	}

	UserProfile struct {
		model.Base  `gorm:",inline"`
		DateOfBirth time.Time `gorm:"date_of_birth"`
	}
)
