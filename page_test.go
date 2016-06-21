package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPageWithArray(t *testing.T) {
	content := [3]int{1, 2, 3}
	page, err := NewPage(content, NewPageable(1, 3), 10)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(page)
	assert.Equal(content, page.Content)
	assert.Equal(1, page.Number)
	assert.Equal(3, page.Size)
	assert.Equal(10, page.TotalElements)
	assert.Equal(4, page.TotalPages)
}

func TestNewPageWithSlice(t *testing.T) {
	content := make([]string, 3, 10)
	page, err := NewPage(content, NewPageable(1, 3), 10)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(page)
	assert.Equal(content, page.Content)
	assert.Equal(1, page.Number)
	assert.Equal(3, page.Size)
	assert.Equal(10, page.TotalElements)
	assert.Equal(4, page.TotalPages)
}

func TestNewPageWithInvalidContent(t *testing.T) {
	page, err := NewPage("wrong", NewPageable(1, 3), 10)

	assert := assert.New(t)
	assert.Nil(page)
	assert.NotNil(err)
	assert.Equal(ErrInvalidContent, err)
}

func TestHasPreviousOnFirstPage(t *testing.T) {
	page, _ := NewPage([]int{1, 2, 3}, NewPageable(0, 10), 25)

	assert.False(t, page.HasPrevious())
}

func TestHasPrevoiusOnOtherPage(t *testing.T) {
	page, _ := NewPage([]int{1, 2, 3}, NewPageable(1, 10), 25)

	assert.True(t, page.HasPrevious())
}

func TestHasNextOnLastPage(t *testing.T) {
	page, _ := NewPage([]int{1, 2, 3}, NewPageable(1, 10), 13)

	assert.False(t, page.HasNext())
}

func TestHasNextOnOtherPage(t *testing.T) {
	page, _ := NewPage([]int{1, 2, 3}, NewPageable(0, 10), 13)

	assert.True(t, page.HasNext())
}

func TestHasNextBeyondLastPage(t *testing.T) {
	page, _ := NewPage([]int{1, 2, 3}, NewPageable(2, 10), 13)

	assert.False(t, page.HasNext())
}

func TestIsFirstOnFirstPage(t *testing.T) {
	page, _ := NewPage([]int{1, 2, 3}, NewPageable(0, 10), 13)

	assert.True(t, page.IsFirst())
}

func TestIsFirstBeyondFirstPage(t *testing.T) {
	page, _ := NewPage([]int{1, 2, 3}, NewPageable(1, 10), 13)

	assert.False(t, page.IsFirst())
}

func TestIsLastOnLastPage(t *testing.T) {
	page, _ := NewPage([]int{1, 2, 3}, NewPageable(1, 10), 13)

	assert.True(t, page.IsLast())
}

func TestIsLastBeforeFirstPage(t *testing.T) {
	page, _ := NewPage([]int{1, 2, 3}, NewPageable(1, 3), 13)

	assert.False(t, page.IsLast())
}

func TestIsLastBeyondLastPage(t *testing.T) {
	page, _ := NewPage([]int{1, 2, 3}, NewPageable(2, 10), 13)

	assert.False(t, page.IsLast())
}

func TestHasContentWithEmptyArray(t *testing.T) {
	page, _ := NewPage([0]int{}, NewPageable(1, 10), 10)

	assert.False(t, page.HasContent())
}

func TestHasContentWithArray(t *testing.T) {
	page, _ := NewPage([3]int{1, 2, 3}, NewPageable(1, 10), 10)

	assert.True(t, page.HasContent())
}

func TestHasContentWithEmptySlice(t *testing.T) {
	page, _ := NewPage(make([]int, 0), NewPageable(1, 10), 10)

	assert.False(t, page.HasContent())
}

func TestHasContentWithSlice(t *testing.T) {
	page, _ := NewPage(make([]int, 3), NewPageable(1, 10), 10)

	assert.True(t, page.HasContent())
}
