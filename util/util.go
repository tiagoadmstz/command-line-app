package util

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
)

//Map iterates slice and return another slice type
func Map[K comparable, V any](servers []K, function func(value K) V) []V {
	var strs []V
	for _, serve := range servers {
		strs = append(strs, function(serve))
	}
	return strs
}

//Search runs a function and prints values
func Search[V any](context *cli.Context, function func(host string) ([]V, error)) {
	host := context.String("host")
	values, err := function(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range values {
		fmt.Println(value)
	}
}
