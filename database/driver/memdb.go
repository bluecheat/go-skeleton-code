package driver

import (
	"crypto/sha256"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/patrickmn/go-cache"
	"reflect"
	"strconv"
	"time"
)

type MemDB struct {
	cache *cache.Cache
	count map[string]int
}

func NewMemDB() *MemDB {

	cache := cache.New(5*time.Minute, 100*time.Minute)
	return &MemDB{
		cache: cache,
		count: make(map[string]int),
	}
}

func (m *MemDB) Get(model interface{}, where ...interface{}) (interface{}, error) {
	if len(where) == 0 {
		return nil, fmt.Errorf("no valid where")
	}
	value, found := m.cache.Get(key(model, where[0].(uint64)))
	if !found {
		return nil, fmt.Errorf("not found model = %v", where[0].(uint64))
	}
	return value, nil
}

func (m *MemDB) Set(model interface{}, where ...interface{}) error {
	if len(where) == 0 {
		return fmt.Errorf("no valid where")
	}

	m.cache.Set(key(model, where[0].(uint64)), model, cache.DefaultExpiration)
	m.count[typeofStruct(model)]++
	return nil
}

func (m MemDB) Update(model interface{}, where ...interface{}) error {
	if len(where) == 0 {
		return fmt.Errorf("no valid where")
	}

	m.cache.Set(key(model, where[0].(uint64)), model, cache.DefaultExpiration)
	return nil
}

func (m MemDB) Delete(model interface{}, where ...interface{}) error {
	if len(where) == 0 {
		return fmt.Errorf("no valid where")
	}

	m.cache.Delete(key(model, where[0].(uint64)))
	m.count[typeofStruct(model)]--
	return nil
}

func (m MemDB) Count(model interface{}, where ...interface{}) (int, error) {
	return m.count[typeofStruct(model)], nil
}

func (m MemDB) Close() error {
	return nil
}

func typeofStruct(x interface{}) string {
	return reflect.TypeOf(x).String()
}

func key(x interface{}, id uint64) string {
	return typeofStruct(x) + ":" + strconv.FormatUint(id, 10)
}

func hash(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashedTable(x interface{}) string {
	return typeofStruct(x) + ":" + hash(x)
}
