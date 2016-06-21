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

// Params contains the default parsing parameters
type Params struct {
	PageParam   string
	SizeParam   string
	SortParam   string
	DefaultPage int
	DefaultSize int
}

var defaultParams = Params{DefaultPageParam, DefaultSizeParam, DefaultSortParam, DefaultPage, DefaultSize}

// ParseHTTPRequest will parse the provided HTTP request with default parameters
func ParseHTTPRequest(req *http.Request) (*data.Pageable, error) {
	return ParseHTTPRequestWithParams(req, defaultParams)
}

// ParseHTTPRequestWithParams will parse the provided HTTP request given the provided parameters
func ParseHTTPRequestWithParams(req *http.Request, params Params) (*data.Pageable, error) {
	return ParseURLWithParams(req.URL, params)
}

// ParseURL will parse the provided URL with default parameters
func ParseURL(url *url.URL) (*data.Pageable, error) {
	return ParseURLWithParams(url, defaultParams)
}

// ParseURLWithParams will parse the provided URL given the provided parameters
func ParseURLWithParams(url *url.URL, params Params) (*data.Pageable, error) {
	return ParseValuesWithParams(url.Query(), params)
}

// ParseValues will parse the provided values with default parameters
func ParseValues(values map[string][]string) (*data.Pageable, error) {
	return ParseValuesWithParams(values, defaultParams)
}

// ParseValuesWithParams will parse the provided values with given parameters
func ParseValuesWithParams(values map[string][]string, params Params) (*data.Pageable, error) {

	page, err := parsePage(values, params)
	if err != nil {
		return nil, err
	}
	size, err := parseSize(values, params)
	if err != nil {
		return nil, err
	}
	sort, err := parseSort(values, params)
	if err != nil {
		return nil, err
	}
	return data.NewSortedPageable(page, size, sort), nil
}

func parsePage(values map[string][]string, params Params) (int, error) {
	if value, ok := values[params.PageParam]; ok {
		if len(value) != 1 {
			return params.DefaultPage, ErrWrongPageValues
		}
		page, err := strconv.Atoi(value[0])
		if err != nil || page < 0 {
			return params.DefaultPage, ErrInvalidPageValue
		}
		return page, nil
	}
	return params.DefaultPage, nil
}

func parseSize(values map[string][]string, params Params) (int, error) {
	if value, ok := values[params.SizeParam]; ok {
		if len(value) != 1 {
			return params.DefaultSize, ErrWrongSizeValues
		}
		size, err := strconv.Atoi(value[0])
		if err != nil || size <= 0 {
			return params.DefaultSize, ErrInvalidSizeValue
		}
		return size, nil
	}
	return params.DefaultSize, nil
}

func parseSort(values map[string][]string, params Params) (*data.Sort, error) {
	if values, ok := values[params.SortParam]; ok {
		sort := data.EmptySort()
		for _, v := range values {
			s, err := parseOrder(v)
			if err != nil {
				return nil, err
			}
			sort = sort.And(s)
		}
		return sort, nil
	}
	return nil, nil
}

func parseOrder(sort string) (*data.Sort, error) {
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
