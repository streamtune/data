package data

import (
	"errors"
	"math"
	"reflect"
)

// ErrInvalidContent is returned by NewPage when an invalid content is provided
var ErrInvalidContent = errors.New("Invalid content provided: expected Array or Slice")

// Page is the struct used tho hold a single page of data
type Page struct {
	Content       interface{} `json:"content"`
	Number        int         `json:"number"`
	Size          int         `json:"size"`
	TotalPages    int         `json:"totalPages"`
	TotalElements int         `json:"totalElements"`
}

// NewPage create a new Page object with provided content, pagination object and total number of elements
func NewPage(content interface{}, pageable *Pageable, totalElements int) (*Page, error) {
	contentType := reflect.TypeOf(content)
	kind := contentType.Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return nil, ErrInvalidContent
	}
	return &Page{
		Content:       content,
		Number:        pageable.Page,
		Size:          pageable.Size,
		TotalPages:    int(math.Ceil(float64(totalElements) / float64(pageable.Size))),
		TotalElements: totalElements,
	}, nil
}

// HasPrevious check if the page has a page before this one
func (page *Page) HasPrevious() bool {
	return page.Number > 0
}

// HasNext check if the page as a page over this one
func (page *Page) HasNext() bool {
	return page.Number < (page.TotalPages - 1)
}

// IsFirst check if the page is the first one
func (page *Page) IsFirst() bool {
	return page.Number == 0
}

// IsLast check if the page si the last one
func (page *Page) IsLast() bool {
	return (page.Number + 1) == page.TotalPages
}

// HasContent check if the page has content
func (page *Page) HasContent() bool {
	return reflect.ValueOf(page.Content).Len() > 0
}
