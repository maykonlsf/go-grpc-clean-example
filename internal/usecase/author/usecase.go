package author

import "github.com/maykonlf/go-grpc-clean-example/internal/domain/entities"

type UseCase interface {
	AddAuthor(*AddRequest) (*AddResponse, error)
}

func NewAuthorUseCase(authorService Service) UseCase {
	return &useCase{
		authorService: authorService,
	}
}

type useCase struct {
	authorService Service
}

func (u *useCase) AddAuthor(request *AddRequest) (*AddResponse, error) {
	id, err := u.authorService.Add(&entities.Author{
		FirstName: request.firstName,
		LastName:  request.lastName,
		Pseudonym: request.pseudonym,
	})

	if err != nil {
		return nil, err
	}

	return &AddResponse{id: id}, nil
}
