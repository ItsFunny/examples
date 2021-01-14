package hash_conflitct

import (
	"github.com/influxdata/influxdb/pkg/testing/assert"
	"testing"
)

func TestHashOpenAddress_Find_Error(t *testing.T) {
	hashOpenAddress := NewHashOpenAddress()
	value, e := hashOpenAddress.Find(0)
	if nil != e {
		panic(e)
	}
	assert.Equal(t, value, true)
}

func TestHashOpenAddress_Find_Normal(t *testing.T) {
	hashOpenAddress := NewHashOpenAddress()
	b, e := hashOpenAddress.Put(0)
	if nil != e {
		panic(e)
	}
	assert.Equal(t, b, true)

	find, e := hashOpenAddress.Find(0)
	if nil != e {
		panic(e)
	}
	assert.Equal(t, find, true)
}
