package captcha

import (
	"github.com/chris1678/go-run/cache"
)

type cacheStore struct {
	expiration int
}

// Set sets the digits for the captcha id.
func (e *cacheStore) Set(id string, value string) error {
	//fmt.Println("cacheStore=====", id)
	return cache.Set(id, value, e.expiration)
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
func (e *cacheStore) Get(id string, clear bool) string {
	v, err := cache.Get(id)
	if err == nil {
		if clear {
			_ = cache.Del(id)
		}
		return v
	}
	return ""
}

//Verify captcha's answer directly
func (e *cacheStore) Verify(id, answer string, clear bool) bool {
	return e.Get(id, clear) == answer
}
