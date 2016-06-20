package parser

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/streamtune/data"
	"github.com/stretchr/testify/assert"
)

func TestNewDefaultParser(t *testing.T) {
	fixture := NewDefaultHttpPageable()

	assert := assert.New(t)
	assert.NotNil(fixture)
	assert.Equal(DefaultPageParam, fixture.pageParam)
	assert.Equal(DefaultSizeParam, fixture.sizeParam)
	assert.Equal(DefaultSortParam, fixture.sortParam)
	assert.Equal(DefaultPage, fixture.defaultPage)
	assert.Equal(DefaultSize, fixture.defaultSize)
}

func TestNewParser(t *testing.T) {
	fixture := NewHttpPageable("p_page", "p_size", "p_sort", 99, 123)

	assert := assert.New(t)
	assert.NotNil(fixture)
	assert.Equal("p_page", fixture.pageParam)
	assert.Equal("p_size", fixture.sizeParam)
	assert.Equal("p_sort", fixture.sortParam)
	assert.Equal(99, fixture.defaultPage)
	assert.Equal(123, fixture.defaultSize)
}

func TestPageableHttpParseRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/v1/api/list?page=1&size=25&sort=prop1,asc&sort=prop2,desc", nil)
	pageable, err := NewDefaultHttpPageable().ParseRequest(req)

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

func TestPageableHttpParseURLWithParameters(t *testing.T) {
	url, _ := url.Parse("http://localhost/v1/api/list?page=1&size=25&sort=prop1,asc&sort=prop2,desc")
	pageable, err := NewDefaultHttpPageable().ParseURL(url)

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

func TestPageableHttpParseValuesWithPage(t *testing.T) {
	values := map[string][]string{"page": []string{"5"}}
	pageable, err := NewDefaultHttpPageable().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable)
	assert.Equal(5, pageable.Page)
	assert.Equal(DefaultSize, pageable.Size)
	assert.Nil(pageable.Sort)
}

func TestPageableHttpParseValuesWithMultiplePageValues(t *testing.T) {
	values := map[string][]string{"page": []string{"5", "6"}}
	_, err := NewDefaultHttpPageable().ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrWrongPageValues, err)
}

func TestPageableHttpParseValuesWithNonNumericPageValue(t *testing.T) {
	values := map[string][]string{"page": []string{"abc"}}
	_, err := NewDefaultHttpPageable().ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrInvalidPageValue, err)
}

func TestPageableHttpParseValuesWithInvalidPageValue(t *testing.T) {
	values := map[string][]string{"page": []string{"-1"}}
	_, err := NewDefaultHttpPageable().ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrInvalidPageValue, err)
}

func TestPageableHttpWithNoPageValue(t *testing.T) {
	values := map[string][]string{}
	pageable, err := NewDefaultHttpPageable().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(DefaultPage, pageable.Page)
}

func TestPageableHttpParseValuesWithSize(t *testing.T) {
	values := map[string][]string{"size": []string{"25"}}
	pageable, err := NewDefaultHttpPageable().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable)
	assert.Equal(DefaultPage, pageable.Page)
	assert.Equal(25, pageable.Size)
	assert.Nil(pageable.Sort)
}

func TestPageableHttpParseValuesWithMultipleSizeValues(t *testing.T) {
	values := map[string][]string{"size": []string{"15", "25"}}
	_, err := NewDefaultHttpPageable().ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrWrongSizeValues, err)
}

func TestPageableHttpParseValuesWithNonNumericSizeValue(t *testing.T) {
	values := map[string][]string{"size": []string{"abc"}}
	_, err := NewDefaultHttpPageable().ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrInvalidSizeValue, err)
}

func TestPageableHttpParseValuesWithInvalidSizeValue(t *testing.T) {
	values := map[string][]string{"size": []string{"0"}}
	_, err := NewDefaultHttpPageable().ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrInvalidSizeValue, err)
}

func TestPageableHttpWithNoSizeValue(t *testing.T) {
	values := map[string][]string{}
	pageable, err := NewDefaultHttpPageable().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(DefaultSize, pageable.Size)
}

func TestPageableHttpParserValuesWithSingleSortWithOneProperty(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1,desc"}}
	pageable, err := NewDefaultHttpPageable().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable.Sort)
	assert.Equal(1, len(pageable.Sort.Orders))
	assert.Equal(data.Desc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
}

func TestPageableHttpParserValuesWithSingleSortWithOnePropertyNoDirection(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1"}}
	pageable, err := NewDefaultHttpPageable().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable.Sort)
	assert.Equal(1, len(pageable.Sort.Orders))
	assert.Equal(data.Asc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
}

func TestPageableHttpParserValuesWithSingleSortWithTwoProperty(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1,prop2,desc"}}
	pageable, err := NewDefaultHttpPageable().ParseValues(values)

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
	_, err := NewDefaultHttpPageable().ParseValues(values)

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(data.ErrInvalidDirection, err)
}

func TestPageableHttpParserValuesWithMultipleSortWithOneProperty(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1,desc", "prop2,asc"}}
	pageable, err := NewDefaultHttpPageable().ParseValues(values)

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
	pageable, err := NewDefaultHttpPageable().ParseValues(values)

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
	pageable, err := NewDefaultHttpPageable().ParseValues(values)

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
	pageable, err := NewDefaultHttpPageable().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Nil(pageable.Sort)
}

func TestPageableHttpParserValuesWithSortNoValue(t *testing.T) {
	values := map[string][]string{"sort": []string{""}}
	_, err := NewDefaultHttpPageable().ParseValues(values)

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(ErrWrongSortValue, err)
}
