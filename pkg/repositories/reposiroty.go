package repositories

import (
	"context"
	"reflect"

	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	FlagLog bool
)

type IRepository[T any] interface {
	GetAll() ([]*T, int64, error)
	GetByID(id interface{}) (*T, error)
	GetByCriteria(criteria string, args ...interface{}) ([]*T, int64, error)
	Create(item *T) (*int64, error)
	Update(item *T) (int64, error)
	Delete(id interface{}) (int64, error)
	GetTx() *gorm.DB
	SetTx(tx *gorm.DB)
	SetPreloads(preloads ...string)
}

type Repository[T any] struct {
	tx       *gorm.DB
	ctx      context.Context
	preloads []string
}

func NewRepository[T any]() *Repository[T] {
	return NewRepositoryWithContext[T](context.Background())
}

func (r *Repository[T]) SetPreloads(preloads ...string) {
	r.preloads = preloads
}

func NewRepositoryWithContext[T any](ctx context.Context) *Repository[T] {

	r := &Repository[T]{}
	r.ctx = ctx
	r.tx = DB.WithContext(r.ctx)

	return r
}

func (r *Repository[T]) GetTx() *gorm.DB {
	return r.tx
}
func (r *Repository[T]) SetTx(tx *gorm.DB) {
	r.tx = tx
}

func (r *Repository[T]) GetAll() ([]*T, int64, error) {
	items := []*T{}
	db := r.tx
	if len(r.preloads) > 0 {
		db = db.Preload(r.preloads[0])
		for _, preload := range r.preloads[1:] {
			db = db.Preload(preload)
		}
	}

	result := db.Find(&items)
	return items, result.RowsAffected, result.Error
}

func (r *Repository[T]) GetByCriteria(criteria string, args ...interface{}) ([]*T, int64, error) {
	items := []*T{}
	db := r.tx
	if len(r.preloads) > 0 {
		db = db.Preload(r.preloads[0])
		for _, preload := range r.preloads[1:] {
			db = db.Preload(preload)
		}
	}

	result := db.Where(criteria, args...).Find(&items)
	return items, result.RowsAffected, result.Error
}

func (r *Repository[T]) GetByID(id interface{}) (*T, error) {
	item := CreateNewElement[T]()
	db := r.tx
	if len(r.preloads) > 0 {
		db = db.Preload(r.preloads[0])
		for _, preload := range r.preloads[1:] {
			db = db.Preload(preload)
		}
	}

	result := db.First(item, "id = ?", id)
	return item, result.Error
}

func (r *Repository[T]) Create(item *T) (*int64, error) {
	result := r.tx.Create(item)
	id := reflect.ValueOf(item).Elem().FieldByName("ID").Interface().(int64)
	return &id, result.Error
}

func (r *Repository[T]) Update(item *T) (int64, error) {
	result := r.tx.Save(item)
	return result.RowsAffected, result.Error
}

func (r *Repository[T]) Delete(id interface{}) (int64, error) {
	item := CreateNewElement[T]()
	result := r.tx.Delete(item, "id = ?", id)
	return result.RowsAffected, result.Error
}

func CreateNewElement[T any]() *T {
	t := reflect.TypeOf((*T)(nil)).Elem()
	v := reflect.New(t).Elem()
	return v.Addr().Interface().(*T)
}
