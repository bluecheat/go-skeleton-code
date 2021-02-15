package driver

import (
	"crypto/sha256"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/patrickmn/go-cache"
	"reflect"
	"time"
)

type MemDB struct {
	cache *cache.Cache
}

func NewMemDB() *MemDB {
	cache := cache.New(5*time.Minute, 100*time.Minute)
	return &MemDB{
		cache: cache,
	}
}

func (m *MemDB) Get(model interface{}) error {
	name := typeofStruct(model)
	value, found := m.cache.Get(name)
	if !found {
		return fmt.Errorf("not found model = %v", name)
	}
	model = &value
	return nil
}

func (m MemDB) Set(model interface{}) error {
	name := typeofStruct(model)
	m.cache.Set(name, model, cache.DefaultExpiration)
	return nil
}

func (m MemDB) Update(model interface{}) error {
	name := typeofStruct(model)
	m.cache.Set(name, model, cache.DefaultExpiration)
	return nil
}

func (m MemDB) Delete(model interface{}) error {
	name := typeofStruct(model)
	m.cache.Delete(name)
	return nil
}

func (m MemDB) Close() error {
	return nil
}

func typeofStruct(x interface{}) string {
	return reflect.TypeOf(x).String()
}

func hash(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashedTable(x interface{}) string {
	return typeofStruct(x) + ":" + hash(x)
}
