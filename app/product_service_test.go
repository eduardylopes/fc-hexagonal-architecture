package app_test

import (
	"testing"

	"github.com/eduardylopes/fc-hexagonal-architecture/app"
	mock_app "github.com/eduardylopes/fc-hexagonal-architecture/app/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockProductInterface(ctrl)
	persistence := mock_app.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := app.ProductService{Persistence: persistence}
	result, err := service.Get("")

	require.Nil(t, err)
	require.Equal(t, product, result)

}
