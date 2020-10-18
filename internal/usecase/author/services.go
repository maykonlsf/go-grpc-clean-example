package author

import "github.com/maykonlf/go-grpc-clean-example/internal/domain/entities"

//go:generate mockery --name=Service
type Service interface {
	Add(author *entities.Author) (string, error)
}
