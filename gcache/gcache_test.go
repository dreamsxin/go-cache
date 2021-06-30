package gcache_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gookit/cache"
	"github.com/gookit/cache/gcache"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func Example() {
	c := gcache.New(12)
	key := "name"

	// set
	c.Set(key, "cache value", cache.Seconds2)
	fmt.Println(c.Has(key))

	// get
	val := c.Get(key)
	fmt.Println(val)

	time.Sleep(2 * time.Second)

	// get expired
	val2 := c.Get(key)
	fmt.Println(val2)

	// del
	c.Del(key)
	fmt.Println(c.Has(key))

	// Output:
	// true
	// cache value
	// <nil>
	// false
}

func ExampleGCache_in_cachePkg() {
	cache.Register(gcache.Name, gcache.New(12))
	defer cache.UnregisterAll()

	key := "name1"

	// set
	cache.Set(key, "cache value", cache.Seconds2)
	fmt.Println(cache.Has(key))

	// get
	val := cache.Get(key)
	fmt.Println(val)

	time.Sleep(2 * time.Second)

	// get expired
	val2 := cache.Get(key)
	fmt.Println(val2)

	// del
	cache.Del(key)
	fmt.Println(cache.Has(key))

	// Output:
	// true
	// cache value
	// <nil>
	// false
}

func TestGCache_usage(t *testing.T) {
	is := assert.New(t)
	c := gcache.New(12)
	defer c.Clear()

	key := strutil.RandomCharsV2(12)
	is.False(c.Has(key))

	err := c.Set(key, "value", cache.Seconds3)
	is.NoError(err)
	is.True(c.Has(key))

	val := c.Get(key)
	is.Equal("value", val)

	// del
	err = c.Del(key)
	is.NoError(err)
	is.False(c.Has(key))
}

type user struct {
	Age int
	Name string
}

func TestGCache_object(t *testing.T) {
	is := assert.New(t)
	c := gcache.New(12)
	defer c.Clear()

	b1 := user {
		Age: 1,
		Name: "inhere",
	}

	key := strutil.RandomCharsV2(12)
	dump.P("cache key: " + key)
	is.False(c.Has(key))

	err := c.Set(key, b1, cache.Seconds3)
	is.NoError(err)
	is.True(c.Has(key))

	b2 := c.Get(key).(user)
	dump.P(b2)
	is.Equal("inhere", b2.Name)
}