package parser

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/streamtune/data"
	"github.com/stretchr/testify/assert"
)

func TestParseHTTPRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/v1/api/list?page=1&size=25&sort=prop1,asc&sort=prop2,desc", nil)
	pageable, err := ParseHTTPRequest(req)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(1, pageable.Page)
	assert.Equal(25, pageable.Size)
	assert.Equal(2, len(pageable.Sort.Orders))
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
	assert.Equal(data.Asc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop2", pageable.Sort.Orders[1].Property)
	assert.Equal(data.Desc, pageable.Sort.Orders[1].Direction)
}

func TestParseHTTPRequestWithParams(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/v1/api/list?p_page=1&p_size=25&p_sort=prop1,asc&p_sort=prop2,desc", nil)
	pageable, err := ParseHTTPRequestWithParams(req, Params{
		PageParam:   "p_page",
		SizeParam:   "p_size",
		SortParam:   "p_sort",
		DefaultPage: 1,
		DefaultSize: 50,
	})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(1, pageable.Page)
	assert.Equal(25, pageable.Size)
	assert.Equal(2, len(pageable.Sort.Orders))
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
	assert.Equal(data.Asc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop2", pageable.Sort.Orders[1].Property)
	assert.Equal(data.Desc, pageable.Sort.Orders[1].Direction)
}

func TestParseURL(t *testing.T) {
	url, _ := url.Parse("http://localhost/v1/api/list?page=1&size=25&sort=prop1,asc&sort=prop2,desc")
	pageable, err := ParseURL(url)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(1, pageable.Page)
	assert.Equal(25, pageable.Size)
	assert.Equal(2, len(pageable.Sort.Orders))
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
	assert.Equal(data.Asc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop2", pageable.Sort.Orders[1].Property)
	assert.Equal(data.Desc, pageable.Sort.Orders[1].Direction)
}

func TestParseURLWithParams(t *testing.T) {
	url, _ := url.Parse("http://localhost/v1/api/list?p_page=1&p_size=25&p_sort=prop1,asc&p_sort=prop2,desc")
	pageable, err := ParseURLWithParams(url, Params{
		PageParam:   "p_page",
		SizeParam:   "p_size",
		SortParam:   "p_sort",
		DefaultPage: 1,
		DefaultSize: 50,
	})

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(1, pageable.Page)
	assert.Equal(25, pageable.Size)
	assert.Equal(2, len(pageable.Sort.Orders))
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
	assert.Equal(data.Asc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop2", pageable.Sort.Orders[1].Property)
	assert.Equal(data.Desc, pageable.Sort.Orders[1].Direction)
}

func TestParseValues(t *testing.T) {
	values := map[string][]string{"page": []string{"5"}}
	pageable, err := ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable)
	assert.Equal(5, pageable.Page)
	assert.Equal(DefaultSize, pageable.Size)
	assert.Nil(pageable.Sort)
}

func TestParseValuesWithMultiplePageValues(t *testing.T) {
	values := map[string][]string{"page": []string{"5", "6"}}
	_, err := ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrWrongPageValues, err)
}

func TestParseValuesWithNonNumericPageValue(t *testing.T) {
	values := map[string][]string{"page": []string{"abc"}}
	_, err := ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrInvalidPageValue, err)
}

func TestParseValuesWithInvalidPageValue(t *testing.T) {
	values := map[string][]string{"page": []string{"-1"}}
	_, err := ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrInvalidPageValue, err)
}

func TestParseValuesWithNoPageValue(t *testing.T) {
	values := map[string][]string{}
	pageable, err := ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(DefaultPage, pageable.Page)
}

func TestParseValuesWithSize(t *testing.T) {
	values := map[string][]string{"size": []string{"25"}}
	pageable, err := ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable)
	assert.Equal(DefaultPage, pageable.Page)
	assert.Equal(25, pageable.Size)
	assert.Nil(pageable.Sort)
}

func TestParseValuesWithMultipleSizeValues(t *testing.T) {
	values := map[string][]string{"size": []string{"15", "25"}}
	_, err := ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrWrongSizeValues, err)
}

func TestParseValuesWithNonNumericSizeValue(t *testing.T) {
	values := map[string][]string{"size": []string{"abc"}}
	_, err := ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrInvalidSizeValue, err)
}

func TestParseValuesWithInvalidSizeValue(t *testing.T) {
	values := map[string][]string{"size": []string{"0"}}
	_, err := ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrInvalidSizeValue, err)
}

func TestPageableHttpWithNoSizeValue(t *testing.T) {
	values := map[string][]string{}
	pageable, err := ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(DefaultSize, pageable.Size)
}

func TestPageableHttpParserValuesWithSingleSortWithOneProperty(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1,desc"}}
	pageable, err := ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable.Sort)
	assert.Equal(1, len(pageable.Sort.Orders))
	assert.Equal(data.Desc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
}

func TestPageableHttpParserValuesWithSingleSortWithOnePropertyNoDirection(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1"}}
	pageable, err := ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable.Sort)
	assert.Equal(1, len(pageable.Sort.Orders))
	assert.Equal(data.Asc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
}

func TestPageableHttpParserValuesWithSingleSortWithTwoProperty(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1,prop2,desc"}}
	pageable, err := ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable.Sort)
	assert.Equal(2, len(pageable.Sort.Orders))
	assert.Equal(data.Desc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
	assert.Equal(data.Desc, pageable.Sort.Orders[1].Direction)
	assert.Equal("prop2", pageable.Sort.Orders[1].Property)
}

func TestPageableHttpParserValuesWithSingleSortWithTwoPropertyNoDirection(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1,prop2"}}
	_, err := ParseValues(values)

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(data.ErrInvalidDirection, err)
}

func TestPageableHttpParserValuesWithMultipleSortWithOneProperty(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1,desc", "prop2,asc"}}
	pageable, err := ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable.Sort)
	assert.Equal(2, len(pageable.Sort.Orders))
	assert.Equal(data.Desc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
	assert.Equal(data.Asc, pageable.Sort.Orders[1].Direction)
	assert.Equal("prop2", pageable.Sort.Orders[1].Property)
}

func TestPageableHttpParserValuesWithMultipleSortWithOnePropertyNoDirection(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1", "prop2"}}
	pageable, err := ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable.Sort)
	assert.Equal(2, len(pageable.Sort.Orders))
	assert.Equal(data.Asc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
	assert.Equal(data.Asc, pageable.Sort.Orders[1].Direction)
	assert.Equal("prop2", pageable.Sort.Orders[1].Property)
}

func TestPageableHttpParserValuesWithMultipleSortWithTwoProperty(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1,prop2,desc", "prop3,prop4,asc"}}
	pageable, err := ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable.Sort)
	assert.Equal(4, len(pageable.Sort.Orders))
	assert.Equal(data.Desc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
	assert.Equal(data.Desc, pageable.Sort.Orders[1].Direction)
	assert.Equal("prop2", pageable.Sort.Orders[1].Property)
	assert.Equal(data.Asc, pageable.Sort.Orders[2].Direction)
	assert.Equal("prop3", pageable.Sort.Orders[2].Property)
	assert.Equal(data.Asc, pageable.Sort.Orders[3].Direction)
	assert.Equal("prop4", pageable.Sort.Orders[3].Property)
}

func TestPageableHttpParserValuesWithNoSort(t *testing.T) {
	values := map[string][]string{}
	pageable, err := ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Nil(pageable.Sort)
}

func TestPageableHttpParserValuesWithSortNoValue(t *testing.T) {
	values := map[string][]string{"sort": []string{""}}
	_, err := ParseValues(values)

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(ErrWrongSortValue, err)
}
