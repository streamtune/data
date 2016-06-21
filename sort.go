package data

import (
	"errors"
	"strings"
)

// ErrInvalidDirection is returned by ParseDirection function when the direction is not valid
// ErrInvalidNullHandling is returned by ParseNullHandling function when the null handling value is not valid
var (
	ErrInvalidDirection    = errors.New("Invalid value for order given! It has to be either 'asc' or 'desc' (case insensitive)")
	ErrInvalidNullHandling = errors.New("Invalid value for null handling given! It has to be 'native', 'nullsFirst' or 'nullsLast'")
)

// Sort options for queries
type Sort struct {
	Orders []Order
}

// EmptySort will create a new empty sort object
func EmptySort() *Sort {
	return &Sort{Orders: make([]Order, 0)}
}

// NewSort create a new Sort object with provided order clauses
func NewSort(order Order, orders ...Order) *Sort {
	target := make([]Order, 1, len(orders)+1)
	target[0] = order
	target = append(target, orders...)
	return &Sort{target}
}

// SortBy create a new Sort object with provided direction and properties
func SortBy(direction Direction, property string, properties ...string) *Sort {
	order := OrderBy(property, direction)
	orders := make([]Order, len(properties))
	for i, p := range properties {
		orders[i] = OrderBy(p, direction)
	}
	return NewSort(order, orders...)
}

// SortByProperties create a new Sort object with Asc direction for provided properties
func SortByProperties(property string, properties ...string) *Sort {
	return SortBy(Asc, property, properties...)
}

// And will join the provided Sort object with the target one and produces a new Sort instance
func (sort *Sort) And(other *Sort) *Sort {
	target := append(sort.Orders, other.Orders...)
	return &Sort{Orders: target}
}

// IsEmpty will check if the sort contains properties
func (sort *Sort) IsEmpty() bool {
	return len(sort.Orders) == 0
}

// Order it's the pairing of a property and a direction.
// It's used as input of Sort.
type Order struct {
	Property     string       `json:"property"`
	Direction    Direction    `json:"direction"`
	IgnoreCase   bool         `json:"ignoreCase"`
	NullHandling NullHandling `json:"nullHandling"`
}

// OrderBy will create a new Order instance for provided property and direction
func OrderBy(property string, direction Direction) Order {
	return Order{property, direction, false, Native}
}

// OrderByProperty will create a new Order instance for provided property and Asc direction
func OrderByProperty(property string) Order {
	return OrderBy(property, Asc)
}

// IsAscending check if the ordering of the Order Property should be ascending.
func (order Order) IsAscending() bool {
	return order.Direction == Asc
}

// IsDescending check if the ordering of the Order Property should be descending.
func (order Order) IsDescending() bool {
	return order.Direction == Desc
}

// WithDirection returns a new Order instance with given direction and every other property unchanged
func (order Order) WithDirection(direction Direction) Order {
	return Order{order.Property, direction, order.IgnoreCase, order.NullHandling}
}

// WithIgnoreCase returns a new Order instance with given ignore case flag
func (order Order) WithIgnoreCase() Order {
	return Order{order.Property, order.Direction, true, order.NullHandling}
}

// WithNullHandling returns a new Order instance with given null handling value
func (order Order) WithNullHandling(nullHandling NullHandling) Order {
	return Order{order.Property, order.Direction, order.IgnoreCase, nullHandling}
}

// NullsNative returns a new Order instance with Native null handling value
func (order Order) NullsNative() Order {
	return order.WithNullHandling(Native)
}

// NullsFirst returns a new Order instance with NullsFirst null handling value
func (order Order) NullsFirst() Order {
	return order.WithNullHandling(NullsFirst)
}

// NullsLast returns a new Order instance with NullsLast null handling value
func (order Order) NullsLast() Order {
	return order.WithNullHandling(NullsLast)
}

// Direction it's a type holding the sorting direction
type Direction string

// Asc is the ascending sorting direction
// Desc is the descending sorting direction
const (
	Asc  Direction = "asc"
	Desc Direction = "desc"
)

// ParseDirection will try to parse the direction, returning an error if it's not Asc or Desc
func ParseDirection(value string) (Direction, error) {
	switch Direction(strings.ToLower(value)) {
	case Asc:
		return Asc, nil
	case Desc:
		return Desc, nil
	default:
		return "", ErrInvalidDirection
	}
}

// MustParseDirection will parse the direction and panic if it's not Asc or Desc
func MustParseDirection(value string) Direction {
	direction, err := ParseDirection(value)
	if err != nil {
		panic(err)
	}
	return direction
}

// NullHandling it's a type holding the null handling setup
type NullHandling string

// Native let's the underlying data store to decide how to handle nulls
// NullsFirst will hint the data store to place the null values at the start of the list
// NullsLast will hint the data store to place the null values at the end of the list
const (
	Native     NullHandling = "native"
	NullsFirst NullHandling = "nullsFirst"
	NullsLast  NullHandling = "nullsLast"
)

// ParseNullHandling will try to parse the NullHandling instance, returning an error if it's not
// Native, NullsFirst, or NullsLast
func ParseNullHandling(value string) (NullHandling, error) {
	switch strings.ToLower(value) {
	case "native":
		return Native, nil
	case "nullsfirst":
		return NullsFirst, nil
	case "nullslast":
		return NullsLast, nil
	default:
		return "", ErrInvalidNullHandling
	}
}

// MustParseNullHandling will parse the null handling and panic if it's not Native, NullsFirst, or NullsLast
func MustParseNullHandling(value string) NullHandling {
	nullHandling, err := ParseNullHandling(value)
	if err != nil {
		panic(err)
	}
	return nullHandling
}
