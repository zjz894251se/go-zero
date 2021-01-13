package cache

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/mathx"
	"github.com/tal-tech/go-zero/core/stat"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/core/stores/redis/redistest"
	"github.com/tal-tech/go-zero/core/syncx"
)

var errTestNotFound = errors.New("not found")

func init() {
	logx.Disable()
	stat.SetReporter(nil)
}

func TestCacheNode_DelCache(t *testing.T) {
	store, clean, err := redistest.CreateRedis()
	assert.Nil(t, err)
	defer clean()

	cn := cacheNode{
		rds:            store,
		r:              rand.New(rand.NewSource(time.Now().UnixNano())),
		lock:           new(sync.Mutex),
		unstableExpiry: mathx.NewUnstable(expiryDeviation),
		stat:           NewCacheStat("any"),
		errNotFound:    errTestNotFound,
	}
	assert.Nil(t, cn.DelCache())
	assert.Nil(t, cn.DelCache([]string{}...))
	assert.Nil(t, cn.DelCache(make([]string, 0)...))
	cn.SetCache("first", "one")
	assert.Nil(t, cn.DelCache("first"))
	cn.SetCache("first", "one")
	cn.SetCache("second", "two")
	assert.Nil(t, cn.DelCache("first", "second"))
}

func TestCacheNode_InvalidCache(t *testing.T) {
	s, err := miniredis.Run()
	assert.Nil(t, err)
	defer s.Close()

	cn := cacheNode{
		rds:            redis.NewRedis(s.Addr(), redis.NodeType),
		r:              rand.New(rand.NewSource(time.Now().UnixNano())),
		lock:           new(sync.Mutex),
		unstableExpiry: mathx.NewUnstable(expiryDeviation),
		stat:           NewCacheStat("any"),
		errNotFound:    errTestNotFound,
	}
	s.Set("any", "value")
	var str string
	assert.NotNil(t, cn.GetCache("any", &str))
	assert.Equal(t, "", str)
	_, err = s.Get("any")
	assert.Equal(t, miniredis.ErrKeyNotFound, err)
}

func TestCacheNode_Take(t *testing.T) {
	store, clean, err := redistest.CreateRedis()
	assert.Nil(t, err)
	defer clean()

	cn := cacheNode{
		rds:            store,
		r:              rand.New(rand.NewSource(time.Now().UnixNano())),
		barrier:        syncx.NewSharedCalls(),
		lock:           new(sync.Mutex),
		unstableExpiry: mathx.NewUnstable(expiryDeviation),
		stat:           NewCacheStat("any"),
		errNotFound:    errTestNotFound,
	}
	var str string
	err = cn.Take(&str, "any", func(v interface{}) error {
		*v.(*string) = "value"
		return nil
	})
	assert.Nil(t, err)
	assert.Equal(t, "value", str)
	assert.Nil(t, cn.GetCache("any", &str))
	val, err := store.Get("any")
	assert.Nil(t, err)
	assert.Equal(t, `"value"`, val)
}

func TestCacheNode_TakeNotFound(t *testing.T) {
	store, clean, err := redistest.CreateRedis()
	assert.Nil(t, err)
	defer clean()

	cn := cacheNode{
		rds:            store,
		r:              rand.New(rand.NewSource(time.Now().UnixNano())),
		barrier:        syncx.NewSharedCalls(),
		lock:           new(sync.Mutex),
		unstableExpiry: mathx.NewUnstable(expiryDeviation),
		stat:           NewCacheStat("any"),
		errNotFound:    errTestNotFound,
	}
	var str string
	err = cn.Take(&str, "any", func(v interface{}) error {
		return errTestNotFound
	})
	assert.Equal(t, errTestNotFound, err)
	assert.Equal(t, errTestNotFound, cn.GetCache("any", &str))
	val, err := store.Get("any")
	assert.Nil(t, err)
	assert.Equal(t, `*`, val)

	store.Set("any", "*")
	err = cn.Take(&str, "any", func(v interface{}) error {
		return nil
	})
	assert.Equal(t, errTestNotFound, err)
	assert.Equal(t, errTestNotFound, cn.GetCache("any", &str))

	store.Del("any")
	var errDummy = errors.New("dummy")
	err = cn.Take(&str, "any", func(v interface{}) error {
		return errDummy
	})
	assert.Equal(t, errDummy, err)
}

func TestCacheNode_TakeWithExpire(t *testing.T) {
	store, clean, err := redistest.CreateRedis()
	assert.Nil(t, err)
	defer clean()

	cn := cacheNode{
		rds:            store,
		r:              rand.New(rand.NewSource(time.Now().UnixNano())),
		barrier:        syncx.NewSharedCalls(),
		lock:           new(sync.Mutex),
		unstableExpiry: mathx.NewUnstable(expiryDeviation),
		stat:           NewCacheStat("any"),
		errNotFound:    errors.New("any"),
	}
	var str string
	err = cn.TakeWithExpire(&str, "any", func(v interface{}, expire time.Duration) error {
		*v.(*string) = "value"
		return nil
	})
	assert.Nil(t, err)
	assert.Equal(t, "value", str)
	assert.Nil(t, cn.GetCache("any", &str))
	val, err := store.Get("any")
	assert.Nil(t, err)
	assert.Equal(t, `"value"`, val)
}

func TestCacheNode_String(t *testing.T) {
	store, clean, err := redistest.CreateRedis()
	assert.Nil(t, err)
	defer clean()

	cn := cacheNode{
		rds:            store,
		r:              rand.New(rand.NewSource(time.Now().UnixNano())),
		barrier:        syncx.NewSharedCalls(),
		lock:           new(sync.Mutex),
		unstableExpiry: mathx.NewUnstable(expiryDeviation),
		stat:           NewCacheStat("any"),
		errNotFound:    errors.New("any"),
	}
	assert.Equal(t, store.Addr, cn.String())
}

func TestCacheValueWithBigInt(t *testing.T) {
	store, clean, err := redistest.CreateRedis()
	assert.Nil(t, err)
	defer clean()

	cn := cacheNode{
		rds:            store,
		r:              rand.New(rand.NewSource(time.Now().UnixNano())),
		barrier:        syncx.NewSharedCalls(),
		lock:           new(sync.Mutex),
		unstableExpiry: mathx.NewUnstable(expiryDeviation),
		stat:           NewCacheStat("any"),
		errNotFound:    errors.New("any"),
	}

	const (
		key         = "key"
		value int64 = 323427211229009810
	)

	assert.Nil(t, cn.SetCache(key, value))
	var val interface{}
	assert.Nil(t, cn.GetCache(key, &val))
	assert.Equal(t, strconv.FormatInt(value, 10), fmt.Sprintf("%v", val))
}
