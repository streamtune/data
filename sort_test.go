package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNullHandlingNativeShouldReturnTheValue(t *testing.T) {
	value, err := ParseNullHandling("native")

	assert.Nil(t, err)
	assert.Equal(t, Native, value)
}

func TestParseNullHandlingNullsFirstShouldReturnTheValue(t *testing.T) {
	value, err := ParseNullHandling("nullsFirst")

	assert.Nil(t, err)
	assert.Equal(t, NullsFirst, value)
}

func TestParseNullHandlingNullsLastShouldReturnTheValue(t *testing.T) {
	value, err := ParseNullHandling("nullsLast")

	assert.Nil(t, err)
	assert.Equal(t, NullsLast, value)
}

func TestParseNullHandlingInvalidValueShouldReturnError(t *testing.T) {
	_, err := ParseNullHandling("invalid")

	assert.Equal(t, ErrInvalidNullHandling, err)
}

func TestMustParseNullHandlingNativeShouldReturnTheValue(t *testing.T) {
	value := MustParseNullHandling("native")

	assert.Equal(t, Native, value)
}

func TestMustParseNullHandlingInvalidShouldPanic(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()

	MustParseNullHandling("invalid")
}

func TestParseDirectionAscShouldReturnTheValue(t *testing.T) {
	value, err := ParseDirection("asc")

	assert.Nil(t, err)
	assert.Equal(t, Asc, value)
}

func TestParseDirectionDescShouldReturnTheValue(t *testing.T) {
	value, err := ParseDirection("desc")

	assert.Nil(t, err)
	assert.Equal(t, Desc, value)
}

func TestParseDirectionInvalidValueShouldReturnError(t *testing.T) {
	_, err := ParseDirection("invalid")

	assert.Equal(t, ErrInvalidDirection, err)
}

func TestMustParseDirectionAscShouldReturnTheValue(t *testing.T) {
	value := MustParseDirection("asc")

	assert.Equal(t, Asc, value)
}

func TestMustParseDirectionInvalidShouldPanic(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()

	MustParseDirection("invalid")
}

func TestOrderByCreateNewOrderInstance(t *testing.T) {
	order := OrderBy("prop1", Desc)

	assert := assert.New(t)
	assert.NotZero(order)
	assert.Equal("prop1", order.Property)
	assert.Equal(Desc, order.Direction)
	assert.False(order.IgnoreCase)
	assert.Equal(Native, order.NullHandling)
}

func TestOrderByPropertyCreateNewOrderInstanceWithDefaultDirection(t *testing.T) {
	order := OrderByProperty("prop1")

	assert := assert.New(t)
	assert.NotZero(order)
	assert.Equal("prop1", order.Property)
	assert.Equal(Asc, order.Direction)
	assert.False(order.IgnoreCase)
	assert.Equal(Native, order.NullHandling)
}

func TestOrderIsAscendingWillReturnTrue(t *testing.T) {
	order := OrderBy("prop1", Asc)

	assert.True(t, order.IsAscending())
}

func TestOrderIsAscendingWillReturnFalse(t *testing.T) {
	order := OrderBy("prop1", Desc)

	assert.False(t, order.IsAscending())
}

func TestOrderIsDescendingWillReturnTrue(t *testing.T) {
	order := OrderBy("prop1", Desc)

	assert.True(t, order.IsDescending())
}

func TestOrderIsDescendingWillReturnFalse(t *testing.T) {
	order := OrderBy("prop1", Asc)

	assert.False(t, order.IsDescending())
}

func TestWithDirectionWillCreateNewOrderWithProvidedDirection(t *testing.T) {
	order := OrderBy("prop1", Asc).WithDirection(Desc)

	assert := assert.New(t)
	assert.NotZero(order)
	assert.Equal("prop1", order.Property)
	assert.Equal(Desc, order.Direction)
	assert.False(order.IgnoreCase)
	assert.Equal(Native, order.NullHandling)
}

func TestWithIgnoreCaseWillCreateNewOrderWithIgnoreCaseFlagSet(t *testing.T) {
	order := OrderBy("prop1", Asc).WithIgnoreCase()

	assert := assert.New(t)
	assert.NotZero(order)
	assert.Equal("prop1", order.Property)
	assert.Equal(Asc, order.Direction)
	assert.True(order.IgnoreCase)
	assert.Equal(Native, order.NullHandling)
}

func TestWithNullHandlingWillCreateNewOrderWithProvidedNullHandling(t *testing.T) {
	order := OrderBy("prop1", Asc).WithNullHandling(NullsFirst)

	assert := assert.New(t)
	assert.NotZero(order)
	assert.Equal("prop1", order.Property)
	assert.Equal(Asc, order.Direction)
	assert.False(order.IgnoreCase)
	assert.Equal(NullsFirst, order.NullHandling)
}

func TestNullsNativeWillCreateNewOrderWithNativeNullHandling(t *testing.T) {
	order := OrderBy("prop1", Asc).NullsNative()

	assert := assert.New(t)
	assert.NotZero(order)
	assert.Equal("prop1", order.Property)
	assert.Equal(Asc, order.Direction)
	assert.False(order.IgnoreCase)
	assert.Equal(Native, order.NullHandling)
}

func TestNullsFirstWillCreateNewOrderWithNullsFirstNullHandling(t *testing.T) {
	order := OrderBy("prop1", Asc).NullsFirst()

	assert := assert.New(t)
	assert.NotZero(order)
	assert.Equal("prop1", order.Property)
	assert.Equal(Asc, order.Direction)
	assert.False(order.IgnoreCase)
	assert.Equal(NullsFirst, order.NullHandling)
}

func TestNullsLastWillCreateNewOrderWithNullsLastNullHandling(t *testing.T) {
	order := OrderBy("prop1", Asc).NullsLast()

	assert := assert.New(t)
	assert.NotZero(order)
	assert.Equal("prop1", order.Property)
	assert.Equal(Asc, order.Direction)
	assert.False(order.IgnoreCase)
	assert.Equal(NullsLast, order.NullHandling)
}

func TestSortByPropertiesWithOneProperty(t *testing.T) {
	sort := SortByProperties("prop1")

	assert := assert.New(t)
	assert.NotNil(sort)
	assert.Equal(1, len(sort.Orders))
	assert.Equal("prop1", sort.Orders[0].Property)
	assert.Equal(Asc, sort.Orders[0].Direction)
}

func TestSortByPropertiesWithMorePropertoes(t *testing.T) {
	sort := SortByProperties("prop1", "prop2", "prop3")

	assert := assert.New(t)
	assert.NotNil(sort)
	assert.Equal(3, len(sort.Orders))
	assert.Equal("prop1", sort.Orders[0].Property)
	assert.Equal(Asc, sort.Orders[0].Direction)
	assert.Equal("prop2", sort.Orders[1].Property)
	assert.Equal(Asc, sort.Orders[1].Direction)
	assert.Equal("prop3", sort.Orders[2].Property)
	assert.Equal(Asc, sort.Orders[2].Direction)
}
