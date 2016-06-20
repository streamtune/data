package parser

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/streamtune/data"
)

// DefaultPageParam is the default page HTTP parameter
// DefaultSizeParam is the default size HTTP parameter
// DefaultSortParam is the default sort HTTP parameter
const (
	DefaultPageParam string = "page"
	DefaultSizeParam string = "size"
	DefaultSortParam string = "sort"
	DefaultPage      int    = 0
	DefaultSize      int    = 10
)

// ErrWrongPageValues is returned when the wrong number of page parameter values are provided
// ErrInvalidPageValue is returned when an invalid page value is parsed
// ErrWrongSizeValues is returned when the wrong number of size parameter values are provided
// ErrWrongSortValue is returned when the wrong sort parameter is provided
var (
	ErrWrongPageValues  = errors.New("Wrong number of page parameter values, expected 1")
	ErrInvalidPageValue = errors.New("Page value must be numeric greater or equal than 0")
	ErrWrongSizeValues  = errors.New("Wrong number of size parameter values, expected 1")
	ErrInvalidSizeValue = errors.New("Size value must be numeric greater or equal than 1")
	ErrWrongSortValue   = errors.New("Wrong sort value provided: expected <p1>,<p2>,...,<pN>,<dir>")
)

// PageableParser is the object instance for page parameter parsing
type PageableParser struct {
	pageParam   string
	sizeParam   string
	sortParam   string
	defaultPage int
	defaultSize int
}

// NewDefaultParser will create a new PageParser with default parameters
func NewDefaultParser() *PageableParser {
	return NewParser(DefaultPageParam, DefaultSizeParam, DefaultSortParam, DefaultPage, DefaultSize)
}

// NewParser will create a new PageParser with provided parameters
func NewParser(pageParam, sizeParam, sortParam string, defaultPage, defaultSize int) *PageableParser {
	return &PageableParser{pageParam, sizeParam, sortParam, defaultPage, defaultSize}
}

// ParseRequest will parse the HTTP request object
func (parser *PageableParser) ParseRequest(req *http.Request) (*data.Pageable, error) {
	return parser.ParseURL(req.URL)
}

// ParseURL will parse the query parameters in HTTP URL
func (parser *PageableParser) ParseURL(url *url.URL) (*data.Pageable, error) {
	return parser.ParseValues(url.Query())
}

// ParseValues will parse a map of parameters
func (parser *PageableParser) ParseValues(params map[string][]string) (*data.Pageable, error) {
	page, err := parser.parsePage(params)
	if err != nil {
		return nil, err
	}
	size, err := parser.parseSize(params)
	if err != nil {
		return nil, err
	}
	sort, err := parser.parseSort(params)
	if err != nil {
		return nil, err
	}
	return data.NewSortedPageable(page, size, sort), nil
}

func (parser *PageableParser) parsePage(params map[string][]string) (int, error) {
	if value, ok := params[parser.pageParam]; ok {
		if len(value) != 1 {
			return parser.defaultPage, ErrWrongPageValues
		}
		page, err := strconv.Atoi(value[0])
		if err != nil || page < 0 {
			return parser.defaultPage, ErrInvalidPageValue
		}
		return page, nil
	}
	return parser.defaultPage, nil
}

func (parser *PageableParser) parseSize(params map[string][]string) (int, error) {
	if value, ok := params[parser.sizeParam]; ok {
		if len(value) != 1 {
			return parser.defaultSize, ErrWrongSizeValues
		}
		size, err := strconv.Atoi(value[0])
		if err != nil || size <= 0 {
			return parser.defaultSize, ErrInvalidSizeValue
		}
		return size, nil
	}
	return parser.defaultSize, nil
}

func (parser *PageableParser) parseSort(params map[string][]string) (*data.Sort, error) {
	if values, ok := params[parser.sortParam]; ok {
		var sort *data.Sort
		for i, v := range values {
			s, err := parser.parseOrder(v)
			if err != nil {
				return nil, err
			}
			if i == 0 {
				sort = s
			} else {
				sort = sort.And(s)
			}
		}
		return sort, nil
	}
	return nil, nil
}

func (parser *PageableParser) parseOrder(sort string) (*data.Sort, error) {
	parts := strings.Split(sort, ",")
	len := len(parts)
	if len == 1 {
		if parts[0] == "" {
			return nil, ErrWrongSortValue
		}
		return data.NewSort(data.OrderByProperty(parts[0])), nil
	}
	direction, err := data.ParseDirection(parts[len-1])
	if err != nil {
		return nil, err
	}
	orders := make([]data.Order, len-1)
	for i, o := range parts[:len-1] {
		orders[i] = data.OrderBy(o, direction)
	}
	return data.NewSort(orders[0], orders[1:]...), nil
}
