Sometimes you want to share a function, type, or constant among packages in your
module, but you don’t want to make it part of your API. Go supports this via the
special internal package name.

When you create a package called internal, the exported identifiers in that package
and its subpackages are accessible only to the direct parent package of internal and
the sibling packages of internal. Let’s look at an example to see how this works.


**Note**: <u>When you create a package called internal, the exported identifiers in that package
and its subpackages are accessible only to the direct parent package of internal and
the sibling packages of internal. </u> Be aware that attempting to use the internal function from bar.go in the bar package
or from example.go in the root package results in a compilation error:

```
$ go build ./...
package github.com/khoa/internal-example
example.go:3:8: use of internal package
github.com/khoa/internal-example/foo/internal not allowed

package github.com/khoa/internal-example/bar
bar/bar.go:3:8: use of internal package
github.com/khoa/internal-example/foo/internal not allowed
```