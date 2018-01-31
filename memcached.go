package cachego

import (
    "github.com/bradfitz/gomemcache/memcache"
    "time"
)

type (
    // Memcached it's a wrap around the memcached driver
    Memcached struct {
        driver *memcache.Client
    }
)

// NewMemcached - Create an instance of Memcached
func NewMemcached(driver *memcache.Client) *Memcached {
    return &Memcached{driver}
}

// Check if cached key exists in Memcached storage
func (m *Memcached) Contains(key string) bool {
    if _, err := m.Fetch(key); err != nil {
        return false
    }

    return true
}

// Delete the cached key from Memcached storage
func (m *Memcached) Delete(key string) error {
    return m.driver.Delete(key)
}

// Retrieve the cached value from key of the Memcached storage
func (m *Memcached) Fetch(key string) (string, error) {
    item, err := m.driver.Get(key)

    if err != nil {
        return "", err
    }

    value := string(item.Value[:])

    return value, nil
}

// Retrieve multiple cached value from keys of the Memcached storage
func (m *Memcached) FetchMulti(keys []string) map[string]string {
    result := make(map[string]string)

    items, err := m.driver.GetMulti(keys)

    if err != nil {
        return result
    }

    for _, i := range items {
        result[i.Key] = string(i.Value[:])
    }

    return result
}

// Remove all cached keys in Memcached storage
func (m *Memcached) Flush() error {
    return m.driver.FlushAll()
}

// Save a value in Memcached storage by key
func (m *Memcached) Save(key string, value string, lifeTime time.Duration) error {
    err := m.driver.Set(
        &memcache.Item{
            Key:        key,
            Value:      []byte(value),
            Expiration: int32(lifeTime.Seconds()),
        },
    )

    return err
}
