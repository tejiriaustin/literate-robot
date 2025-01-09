package model

import (
	"time"

	"github.com/google/uuid"
)

var _ Models = (*Base)(nil)

type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"_id"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"deleted_at"`
	Version   uint       `json:"version" gorm:"_version"`
}

type Models interface {
	Initialize(id uuid.UUID, now time.Time)
	GetId() string
	SetID(id uuid.UUID)
	SetUpdatedAt()
	GetVersion() uint
	SetVersion(v uint)
}

func (m Base) GetId() string {
	return m.ID.String()
}

func (m Base) SetID(id uuid.UUID) {
	m.ID = id
}

func (m Base) GetVersion() uint {
	return m.Version
}

func (m Base) SetVersion(v uint) {
	m.Version = v
}

func (m Base) SetUpdatedAt() {
	t := time.Now().UTC()
	m.UpdatedAt = &t
}

func (m Base) Initialize(id uuid.UUID, now time.Time) {
	m.ID = id
	t := now.UTC()
	m.CreatedAt = &t
}
