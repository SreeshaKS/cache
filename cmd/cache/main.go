package main

import (
	"fmt"

	"github.com/cache-implementation/src"
)

func main() {
	cacheService := src.NewService(100, "lru")
	cacheService.Manager.Set("x", "y")
	fmt.Println(cacheService.Manager.Get("x"))

	cacheService.Manager.Set("x2", "y2")
	fmt.Println(cacheService.Manager.Get("x2"))

	cacheService.Manager.Set("x3", "y3")
	fmt.Println(cacheService.Manager.Get("x3"))

	cacheService.Manager.Set("x4", "y4")
	fmt.Println(cacheService.Manager.Get("x4"))
	fmt.Println(cacheService.Manager.Get("x3"))
	fmt.Println(cacheService.Manager.Get("x"))
}
