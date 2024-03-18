package knowledge

import (
	"fmt"
	"sync"
	"time"
)

var cache = &sync.Map{}

func Caching_01() {
	arr := [8]int{1, 1, 2, 3, 4, 2, 3, 5}
	for _, v := range arr {
		//Check if the result is already in the cache
		if cacheResult, ok := cache.Load(v); ok {
			fmt.Println("Caching_01 Cache Hit!")
			fmt.Println("Caching_01 Cache result ", cacheResult)
		} else {
			fmt.Println("Caching_01 Cache Miss!")
			result := calculate(v)
			//store the result in the cache with a time to live (TTL)
			cache.Store(v, result)
			go func(num int) {
				//Remove the entry from cache after 2 minutes (120 seconds)
				time.Sleep(2 * time.Minute)
				cache.Delete(num)
			}(v)
		}
	}
}

func calculate(num int) int {
	return num * 10
}
