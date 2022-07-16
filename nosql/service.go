package nosql

import (
	"context"

	"github.com/gmvbr/go-learning/nosql/entity"
)

type BookRepository interface {
	
	InsertOne(ctx context.Context, book *entity.Book) error
	
	InsertMany(ctx context.Context, book []*entity.Book) error

	FindById(ctx context.Context, id string) (*entity.Book, error)
	
	FindByName(ctx context.Context, name string, page, limit int) ([]*entity.Book, error)

	DeleteOne(ctx context.Context, book *entity.Book) (*entity.Book, error)

	DeleteMany(ctx context.Context, book []*entity.Book) error

	UpdateOne(ctx context.Context, book *entity.Book) error

	UpdateMany(ctx context.Context, book []*entity.Book) error

}
