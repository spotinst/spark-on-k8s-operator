package sparkapplication

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSubmissionCacheExist(t *testing.T) {
	cache := NewSubmissionCache(1 * time.Second)
	assert.False(t, cache.Exist("key1"))
	cache.Set("key1")
	assert.True(t, cache.Exist("key1"))
	time.Sleep(2 * time.Second)
	assert.False(t, cache.Exist("key1"))
}
