package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/tejiriaustin/literate-robot/core/database"
	"github.com/tejiriaustin/literate-robot/core/model"
)

type (
	Repository[T model.Models] struct {
		db *database.Client
	}
)

func NewRepository[T model.Models](client *database.Client) *Repository[T] {
	return &Repository[T]{db: client}
}

func (r *Repository[T]) Create(ctx context.Context, data T) (T, error) {
	data.SetVersion(1) // Initialize version to 1 for new records
	data.Initialize(uuid.New(), time.Now())
	fmt.Println(data)
	result := r.db.DB.WithContext(ctx).Create(&data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

func (r *Repository[T]) FindOne(ctx context.Context, queryFilter *Query) (T, error) {
	var result T
	db := r.db.DB.WithContext(ctx)

	if queryFilter != nil {
		db = db.DB.Where(queryFilter.query, queryFilter.args...)
	}

	if err := db.First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, ErrNoDocumentsFound
		}
		return result, err
	}
	return result, nil
}

func (r *Repository[T]) FindMany(ctx context.Context, queryFilter *Query) ([]T, error) {
	var results []T
	db := r.db.DB.WithContext(ctx)

	if queryFilter != nil {
		db = db.Where(queryFilter.query, queryFilter.args...)
	}

	if err := db.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *Repository[T]) DeleteMany(ctx context.Context, queryFilter *Query) error {
	db := r.db.DB.WithContext(ctx)

	if queryFilter != nil {
		db = db.Where(queryFilter.query, queryFilter.args...)
	}

	var model T
	if err := db.Delete(&model).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository[T]) Update(ctx context.Context, dataObject T) (T, error) {
	currentVersion := dataObject.GetVersion()
	newVersion := currentVersion + 1
	dataObject.SetVersion(newVersion)

	err := r.db.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&dataObject).
			Where("_version = ?", currentVersion).
			Updates(dataObject)

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return ErrConcurrentModification
		}

		return nil
	})

	if err != nil {
		// Revert version change on error
		dataObject.SetVersion(currentVersion)
		return dataObject, err
	}

	return dataObject, nil
}

func (r *Repository[T]) Select(ctx context.Context, target interface{}, query string, args ...interface{}) error {
	return r.db.DB.WithContext(ctx).Select(query, args...).Scan(target).Error
}
