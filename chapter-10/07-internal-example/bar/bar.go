package bar

import "github.com/khoa/internal-example/foo/internal"  /** This is not allowed ** /

func FailUseDoubler(i int) int {
	return internal.Doubler(i)
}
