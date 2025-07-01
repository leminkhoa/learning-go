package foo

import "github.com/khoa/internal-example/foo/internal"

func UseDoubler(i int) int {
	return internal.Doubler(i)
}
