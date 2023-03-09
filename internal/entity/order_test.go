package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_I_Gets_An_Error_If_ID_Is_Blank(t *testing.T) {
	order := Order{}
	// err := order.Validate()
	// if err == nil {
	// 	t.Error("Expected an error")
	// }

	assert.Error(t, order.Validate(), "invalid id")
}

func Test_If_I_Gets_An_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order{ID: "123"}

	assert.Error(t, order.Validate(), "invalid id")
}

func Test_If_I_Gets_An_Error_If_Tax_Is_Blank(t *testing.T) {
	order := Order{ID: "123", Price: 39.90}

	assert.Error(t, order.Validate(), "invalid id")
}

func Test_With_All_Valid_Params(t *testing.T) {
	order := Order{ID: "123", Price: 39.90, Tax: 10.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 39.90, order.Price)
	assert.Equal(t, 10.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 43.89, order.FinalPrice)

}
