# cases

[![Go Reference](https://pkg.go.dev/badge/rossmacarthur/cases/format.svg)](https://pkg.go.dev/github.com/rossmacarthur/cases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/rossmacarthur/cases/build.yaml?branch=trunk)](https://github.com/rossmacarthur/cases/actions/workflows/build.yaml)

A case conversion library for Go.

The currently supported cases are:

| Function                  | Output               |
| ------------------------- | -------------------- |
| cases.ToCamel(s)          | camelCase            |
| cases.ToPascal(s)         | PascalCase           |
| cases.ToSnake(s)          | snake_case           |
| cases.ToScreamingSnake(s) | SCREAMING_SNAKE_CASE |
| cases.ToKebab(s)          | kebab-case           |
| cases.ToScreamingKebab(s) | SCREAMING-KEBAB-CASE |
| cases.ToTitle(s)          | Title Case           |
| cases.ToTrain(s)          | Train-Case           |

Word boundaries are defined as follows:
- A set consecutive Unicode spaces, underscores or hyphens
  e.g. "foo _bar" is two words (foo and bar)
- A transition from a lowercase letter to an uppercase letter
  e.g. fooBar is two words (foo and Bar)
- The second last uppercase letter in a word with multiple uppercase letters
  e.g. FOOBar is two words (FOO and Bar)

## License

This project is distributed under the terms of both the MIT license and the
Apache License (Version 2.0).

See [LICENSE-APACHE](LICENSE-APACHE) and [LICENSE-MIT](LICENSE-MIT) for details.
