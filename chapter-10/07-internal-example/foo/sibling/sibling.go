package sibling

import "github.com/khoa/internal-example/foo/internal"

func FailUseDoubler(i int) int {
	return internal.Doubler(i)
}
