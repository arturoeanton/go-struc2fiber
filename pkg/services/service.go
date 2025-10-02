package services

import "github.com/arturoeanton/go-struc2fiber/pkg/repositories"

type IService[T any] interface {
	GetAll() ([]*T, int64, error)
	GetByID(id interface{}) (*T, error)
	GetByCriteria(criteria string, args ...interface{}) ([]*T, int64, error)
	Create(item *T) (int64, error)
	Update(item *T) (int64, error)
	Delete(id interface{}) (int64, error)
}

type Service[T any] struct {
	repo repositories.IRepository[T]
}

func NewService[T any](repo repositories.IRepository[T]) *Service[T] {
	return &Service[T]{
		repo: repo,
	}
}

func (r *Service[T]) GetAll() ([]*T, int64, error) {
	items, c, err := r.repo.GetAll()
	if err != nil {
		return nil, c, err
	}
	return items, c, nil
}

func (r *Service[T]) GetByID(id interface{}) (*T, error) {
	user, err := r.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Service[T]) GetByCriteria(criteria string, args ...interface{}) ([]*T, int64, error) {
	items, c, err := r.repo.GetByCriteria(criteria, args...)
	if err != nil {
		return nil, c, err
	}
	return items, c, nil
}

func (r *Service[T]) Create(item *T) (int64, error) {
	id, err := r.repo.Create(item)
	if err != nil {
		return 0, err
	}

	return *id, nil
}

func (r *Service[T]) Update(item *T) (int64, error) {
	c, err := r.repo.Update(item)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (r *Service[T]) Delete(id interface{}) (int64, error) {
	c, err := r.repo.Delete(id)
	if err != nil {
		return c, err
	}
	return c, nil
}
