package app_test

import (
	"testing"

	"github.com/eduardylopes/fc-hexagonal-architecture/app"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 10.0

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Status = app.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10.0
	err = product.Disable()
	require.Equal(t, "the price must be equal to zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}
	product.ID = uuid.NewString()
	product.Name = "Eduardy"
	product.Price = 14.99
	product.Status = app.ENABLED

	isValid, err := product.IsValid()
	require.True(t, isValid)
	require.Nil(t, err)

	product.Name = ""
	isValid, err = product.IsValid()
	require.False(t, isValid)
	require.Error(t, err)

	product.Status = "INVALID"
	isValid, err = product.IsValid()
	require.False(t, isValid)
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = app.ENABLED
	product.Name = "Eduardy"
	isValid, err = product.IsValid()
	require.True(t, isValid)
	require.Nil(t, err)

	product.Price = -10.0
	isValid, err = product.IsValid()
	require.False(t, isValid)
	require.Equal(t, "the price must be greater or equal to zero", err.Error())
}

func TestProduct_GetId(t *testing.T) {
	product := app.Product{}
	product.ID = "b359e818-4c80-4b44-9603-580da5f4945f"

	require.Equal(t, "b359e818-4c80-4b44-9603-580da5f4945f", product.ID)
}

func TestProduct_GetName(t *testing.T) {
	product := app.Product{}
	product.Name = "Eduardy Lopes"

	require.Equal(t, product.GetName(), "Eduardy Lopes")
}

func TestProduct_GetStatus(t *testing.T) {
	product := app.Product{}
	product.Status = app.ENABLED

	require.Equal(t, product.GetStatus(), "enabled")
}

func TestProduct_GetPrice(t *testing.T) {
	product := app.Product{}
	product.Price = 10.0

	require.Equal(t, product.GetPrice(), 10.0)
}
