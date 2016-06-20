package parser

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/streamtune/data"
	"github.com/stretchr/testify/assert"
)

func TestNewDefaultParser(t *testing.T) {
	fixture := NewDefaultParser()

	assert := assert.New(t)
	assert.NotNil(fixture)
	assert.Equal(DefaultPageParam, fixture.pageParam)
	assert.Equal(DefaultSizeParam, fixture.sizeParam)
	assert.Equal(DefaultSortParam, fixture.sortParam)
	assert.Equal(DefaultPage, fixture.defaultPage)
	assert.Equal(DefaultSize, fixture.defaultSize)
}

func TestNewParser(t *testing.T) {
	fixture := NewParser("p_page", "p_size", "p_sort", 99, 123)

	assert := assert.New(t)
	assert.NotNil(fixture)
	assert.Equal("p_page", fixture.pageParam)
	assert.Equal("p_size", fixture.sizeParam)
	assert.Equal("p_sort", fixture.sortParam)
	assert.Equal(99, fixture.defaultPage)
	assert.Equal(123, fixture.defaultSize)
}

func TestPageableParserParseRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/v1/api/list?page=1&size=25&sort=prop1,asc&sort=prop2,desc", nil)
	pageable, err := NewDefaultParser().ParseRequest(req)

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

func TestPageableParserParseURLWithParameters(t *testing.T) {
	url, _ := url.Parse("http://localhost/v1/api/list?page=1&size=25&sort=prop1,asc&sort=prop2,desc")
	pageable, err := NewDefaultParser().ParseURL(url)

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

func TestPageableParserParseValuesWithPage(t *testing.T) {
	values := map[string][]string{"page": []string{"5"}}
	pageable, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable)
	assert.Equal(5, pageable.Page)
	assert.Equal(DefaultSize, pageable.Size)
	assert.Nil(pageable.Sort)
}

func TestPageableParserParseValuesWithMultiplePageValues(t *testing.T) {
	values := map[string][]string{"page": []string{"5", "6"}}
	_, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrWrongPageValues, err)
}

func TestPageableParserParseValuesWithNonNumericPageValue(t *testing.T) {
	values := map[string][]string{"page": []string{"abc"}}
	_, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrInvalidPageValue, err)
}

func TestPageableParserParseValuesWithInvalidPageValue(t *testing.T) {
	values := map[string][]string{"page": []string{"-1"}}
	_, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrInvalidPageValue, err)
}

func TestPageableParserWithNoPageValue(t *testing.T) {
	values := map[string][]string{}
	pageable, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(DefaultPage, pageable.Page)
}

func TestPageableParserParseValuesWithSize(t *testing.T) {
	values := map[string][]string{"size": []string{"25"}}
	pageable, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable)
	assert.Equal(DefaultPage, pageable.Page)
	assert.Equal(25, pageable.Size)
	assert.Nil(pageable.Sort)
}

func TestPageableParserParseValuesWithMultipleSizeValues(t *testing.T) {
	values := map[string][]string{"size": []string{"15", "25"}}
	_, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrWrongSizeValues, err)
}

func TestPageableParserParseValuesWithNonNumericSizeValue(t *testing.T) {
	values := map[string][]string{"size": []string{"abc"}}
	_, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrInvalidSizeValue, err)
}

func TestPageableParserParseValuesWithInvalidSizeValue(t *testing.T) {
	values := map[string][]string{"size": []string{"0"}}
	_, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.Equal(ErrInvalidSizeValue, err)
}

func TestPageableParserWithNoSizeValue(t *testing.T) {
	values := map[string][]string{}
	pageable, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(DefaultSize, pageable.Size)
}

func TestPageableParserParserValuesWithSingleSortWithOneProperty(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1,desc"}}
	pageable, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable.Sort)
	assert.Equal(1, len(pageable.Sort.Orders))
	assert.Equal(data.Desc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
}

func TestPageableParserParserValuesWithSingleSortWithOnePropertyNoDirection(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1"}}
	pageable, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable.Sort)
	assert.Equal(1, len(pageable.Sort.Orders))
	assert.Equal(data.Asc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
}

func TestPageableParserParserValuesWithSingleSortWithTwoProperty(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1,prop2,desc"}}
	pageable, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable.Sort)
	assert.Equal(2, len(pageable.Sort.Orders))
	assert.Equal(data.Desc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
	assert.Equal(data.Desc, pageable.Sort.Orders[1].Direction)
	assert.Equal("prop2", pageable.Sort.Orders[1].Property)
}

func TestPageableParserParserValuesWithSingleSortWithTwoPropertyNoDirection(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1,prop2"}}
	_, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(data.ErrInvalidDirection, err)
}

func TestPageableParserParserValuesWithMultipleSortWithOneProperty(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1,desc", "prop2,asc"}}
	pageable, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable.Sort)
	assert.Equal(2, len(pageable.Sort.Orders))
	assert.Equal(data.Desc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
	assert.Equal(data.Asc, pageable.Sort.Orders[1].Direction)
	assert.Equal("prop2", pageable.Sort.Orders[1].Property)
}

func TestPageableParserParserValuesWithMultipleSortWithOnePropertyNoDirection(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1", "prop2"}}
	pageable, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(pageable.Sort)
	assert.Equal(2, len(pageable.Sort.Orders))
	assert.Equal(data.Asc, pageable.Sort.Orders[0].Direction)
	assert.Equal("prop1", pageable.Sort.Orders[0].Property)
	assert.Equal(data.Asc, pageable.Sort.Orders[1].Direction)
	assert.Equal("prop2", pageable.Sort.Orders[1].Property)
}

func TestPageableParserParserValuesWithMultipleSortWithTwoProperty(t *testing.T) {
	values := map[string][]string{"sort": []string{"prop1,prop2,desc", "prop3,prop4,asc"}}
	pageable, err := NewDefaultParser().ParseValues(values)

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

func TestPageableParserParserValuesWithNoSort(t *testing.T) {
	values := map[string][]string{}
	pageable, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Nil(pageable.Sort)
}

func TestPageableParserParserValuesWithSortNoValue(t *testing.T) {
	values := map[string][]string{"sort": []string{""}}
	_, err := NewDefaultParser().ParseValues(values)

	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(ErrWrongSortValue, err)
}
