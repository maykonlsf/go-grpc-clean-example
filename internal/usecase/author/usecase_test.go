package author

import (
	"errors"
	"testing"

	"github.com/maykonlf/go-grpc-clean-example/internal/usecase/author/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddAuthor(t *testing.T) {
	serviceMock := &mocks.Service{}
	useCase := NewAuthorUseCase(serviceMock)

	t.Run("should return error when the service fails to add the author", func(t *testing.T) {
		serviceMock.On("Add", mock.Anything).Return("", errors.New("mock error")).Once()
		request := NewAddRequest().FirstName("name").LastName("lastName").Pseudonym("pseudonym")
		_, err := useCase.AddAuthor(request)

		assert.NotNil(t, err)
	})

	t.Run("should return the added author when succeed to register author to database", func(t *testing.T) {
		serviceMock.On("Add", mock.Anything).Return("my-mock-id-01", nil).Once()
		request := NewAddRequest().FirstName("name").LastName("lastName").Pseudonym("pseudonym")
		response, err := useCase.AddAuthor(request)

		assert.Nil(t, err)
		assert.Equal(t, "my-mock-id-01", response.GetID())
	})
}
