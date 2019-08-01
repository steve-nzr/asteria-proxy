package srvbuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBufferSize(t *testing.T) {
	assert.Equal(t, (int64)(2048), getBufferSize())
}
