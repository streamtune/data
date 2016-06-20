package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPageableWillCreateUnsortedPageable(t *testing.T) {
	pageable := NewPageable(0, 10)

	assert := assert.New(t)
	assert.NotNil(pageable)
	assert.Equal(0, pageable.Page)
	assert.Equal(10, pageable.Size)
	assert.Equal(EmptySort(), pageable.Sort)
}

func TestNewSortedPageableWillCreateSortedPageable(t *testing.T) {
	pageable := NewSortedPageable(1, 25, SortBy(Asc, "dateTime"))

	assert := assert.New(t)
	assert.NotNil(pageable)
	assert.Equal(1, pageable.Page)
	assert.Equal(25, pageable.Size)
	assert.NotNil(pageable.Sort)
}

func TestPageableOffssetIsPageBySize(t *testing.T) {
	pageable := NewPageable(3, 25)

	assert.Equal(t, 75, pageable.Offset())
}

func TestHasPreviousOnFirstPageShouldReturnFalse(t *testing.T) {
	pageable := NewPageable(0, 10)

	assert.False(t, pageable.HasPrevious())
}

func TestHasPreviousOnSeondPageShouldReturnTrue(t *testing.T) {
	pageable := NewPageable(1, 10)

	assert.True(t, pageable.HasPrevious())
}

func TestPreviousOrFirstWillReturnPreviousPage(t *testing.T) {
	pageable := NewPageable(2, 10).PreviousOrFirst()

	assert.Equal(t, 1, pageable.Page)
}

func TestPreviousOrFirstWillReturnFirstPage(t *testing.T) {
	pageable := NewPageable(0, 10).PreviousOrFirst()

	assert.Equal(t, 0, pageable.Page)
}

func TestNextWillReturnNextPage(t *testing.T) {
	pageable := NewPageable(5, 10).Next()

	assert.Equal(t, 6, pageable.Page)
}

func TestPreviousWillReturnPreviousPage(t *testing.T) {
	pageable := NewPageable(2, 10).Previous()

	assert.Equal(t, 1, pageable.Page)
}

func TestPreviousWithFirstPageWillReturnTheFirstPage(t *testing.T) {
	pageable := NewPageable(0, 10).Previous()

	assert.Equal(t, 0, pageable.Page)
}

func TestFirstWillReturnFirstPage(t *testing.T) {
	pageable := NewPageable(5, 10).First()

	assert.Equal(t, 0, pageable.Page)
}
