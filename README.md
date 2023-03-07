# cases

[![Go Reference](https://pkg.go.dev/badge/rossmacarthur/cases/format.svg)](https://pkg.go.dev/github.com/rossmacarthur/cases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/rossmacarthur/cases/build.yaml?branch=trunk)](https://github.com/rossmacarthur/cases/actions/workflows/build.yaml)

A case conversion library for Go.

The currently supported cases are:

| Function                              | Output                 |
| :------------------------------------ | :--------------------- |
| `cases.ToCamel(s)`                    | `camelCase`            |
| `cases.ToPascal(s)`                   | `PascalCase`           |
| `cases.ToSnake(s)`                    | `snake_case`           |
| `cases.ToScreamingSnake(s)`           | `SCREAMING_SNAKE_CASE` |
| `cases.ToKebab(s)`                    | `kebab-case`           |
| `cases.ToScreamingKebab(s)`           | `SCREAMING-KEBAB-CASE` |
| `cases.ToTrain(s)`                    | `Train-Case`           |
| `cases.ToLower(s)`                    | `lower case`           |
| `cases.ToTitle(s)`                    | `Title Case`           |
| `cases.ToUpper(s)`                    | `UPPER CASE`           |
| `cases.Transform(s, wordFn, delimFn)` | *your own case here*   |

Word boundaries are defined as follows:
- A set of consecutive Unicode non-letter/number/symbol e.g. `foo _bar` is two
  words (`foo` and `bar`)
- A transition from a lowercase letter to an uppercase letter e.g. `fooBar` is
  two words (`foo` and `Bar`)
- The second last uppercase letter in a word with multiple uppercase letters
  e.g. `FOOBar` is two words (`FOO` and `Bar`)

## Getting started

Install using

```sh
go get -u github.com/rossmacarthur/cases
```

Now convert a string using the relevant function.

```go
import "github.com/rossmacarthur/cases"

cases.ToSnake("XMLHttpRequest") // returns "xml_http_request"
```

## Customizing

This library also exposes a `Transform` function which allows flexible
customization of the output.

For example if you wanted `dotted.snake.case` you could do the following.

```go
import (
    "strings"
    "github.com/rossmacarthur/cases"
)

func delimDot(s *strings.Builder) {
    s.WriteRune('.')
}

cases.Transform("XmlHttpRequest", cases.ToLower, delimDot) // returns xml.http.request
```

Here is a more involved example in order to handle acronyms in `PascalCase`.

```go
import (
    "strings"
    "github.com/rossmacarthur/cases"
)

// The default ToPascal function has no understanding of acronyms
cases.ToPascal("xml_http_request") // returns "XmlHttpRequest"

// We can instead use Transform directly
writeFn := func(s *strings.Builder, word string) {
    w := strings.ToUpper(asLower)
    if w == "XML" || w == "HTTP" {
        s.WriteString(w)
    } else {
        // fallback to default
        cases.WriteTitle(s, word)
    }
}
cases.Transform("xml_http_request", writeFn, nil) // returns "XMLHTTPRequest"
```

## License

This project is distributed under the terms of both the MIT license and the
Apache License (Version 2.0).

See [LICENSE-APACHE](LICENSE-APACHE) and [LICENSE-MIT](LICENSE-MIT) for details.
