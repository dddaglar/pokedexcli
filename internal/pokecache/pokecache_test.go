package pokecache

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	test1 := `{
		"count": 1089,
		"next": "https://pokeapi.co/api/v2/location-area/?offset=21&limit=1",
		"previous": "https://pokeapi.co/api/v2/location-area/?offset=19&limit=1",
		"results": [
			{
			"name": "mt-coronet-1f-route-216",
			"url": "https://pokeapi.co/api/v2/location-area/21/"
			}
		]
	}`
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area/?offset=20&limit=1",
			val: []byte(test1),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if !bytes.Equal(val, c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}
