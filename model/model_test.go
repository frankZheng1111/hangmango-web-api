package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseToLimitAndOffset(t *testing.T) {
	paginate := new(Paginate)
	limit, offset := paginate.ParseToLimitAndOffset()
	assert.Equal(t, 0, offset)
	assert.Equal(t, 30, limit)

	paginate.Page = 2
	paginate.PageSize = 10
	limit, offset = paginate.ParseToLimitAndOffset()
	assert.Equal(t, 10, offset)
	assert.Equal(t, 10, limit)
}
