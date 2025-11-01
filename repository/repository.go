package repository

import (
	"github.com/jkeresman01/medical-records/db"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any]() *Repository[T] {
	return &Repository[T]{db: db.DB}
}

func (r *Repository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *Repository[T]) FindByID(id uint) (*T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	return &entity, err
}

func (r *Repository[T]) FindAll() ([]T, error) {
	var entities []T
	err := r.db.Find(&entities).Error
	return entities, err
}

func (r *Repository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *Repository[T]) DeleteByID(id uint) error {
	var entity T
	return r.db.Delete(&entity, id).Error
}

func (r *Repository[T]) FindByIDWithPreloads(id uint, preloads ...string) (*T, error) {
	var entity T
	query := r.db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	err := query.First(&entity, id).Error
	return &entity, err
}

func (r *Repository[T]) FindAllWithPreloads(preloads ...string) ([]T, error) {
	var entities []T
	query := r.db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	err := query.Find(&entities).Error
	return entities, err
}
