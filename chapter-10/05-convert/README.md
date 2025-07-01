## Viewing Documentation with `go doc`

To view documentation for the `convert` package and its symbols, use the following commands from the parent directory of `convert.go`:

- View all documentation for the package:
  ```sh
  go doc
  ```

- View documentation for the `Money` type:
  ```sh
  go doc Money
  ```

- View documentation for the `Convert` function:
  ```sh
  go doc Convert
  ```

**Note:**  
Make sure your file is named `convert.go` and is inside a directory named `convert`, and that the package declaration at the top of the file is `package convert`. Run these commands from the directory containing the `convert` folder (not inside the `convert` folder itself).

This assumes that you're running the command within this directory `05-go-doc`.

To run it from outside, need to specify the directory name. For example, `go doc 05-go-doc.Money`

```
(base) khoale@khoale-Nitro-AN515-57:~/study/book-practice/learning-go/chapter-10$ go doc 05-go-doc.Money
package convert // import "chapter-10/05-go-doc"

type Money struct {
        Value    float64
        Currency string
}
    Money represents the combination of an amount of money and the currency the
    money is in.

    The value is stored using a github.com/shopspring/decimal.Decimal

func Convert(from Money, to string) (Money, error)
```

## Viewing Documentation with `HTML format`

If you want to preview your documentation’s HTML formatting before it is published
on the web, use pkgsite. This is the same program that powers pkg.go.dev (which I
will talk about later in the chapter). To install pkgsite, use this command:
`$ go install golang.org/x/pkgsite/cmd/pkgsite@latest`

