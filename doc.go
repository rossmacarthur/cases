// Package cases provides functions for converting strings between different
// cases.
//
// The currently supported cases are:
//
// | Function            | Output               |
// | ------------------- | -------------------- |
// | ToCamel(s)          | camelCase            |
// | ToPascal(s)         | PascalCase           |
// | ToSnake(s)          | snake_case           |
// | ToScreamingSnake(s) | SCREAMING_SNAKE_CASE |
// | ToKebab(s)          | kebab-case           |
// | ToScreamingKebab(s) | SCREAMING-KEBAB-CASE |
// | ToTitle(s)          | Title Case           |
// | ToTrain(s)          | Train-Case           |
//
// Word boundaries are defined as follows:
// - A set consecutive Unicode spaces, underscores or hyphens
//   e.g. "foo _bar" is two words (foo and bar)
// - A transition from a lowercase letter to an uppercase letter
//   e.g. fooBar is two words (foo and Bar)
// - The second last uppercase letter in a word with multiple uppercase letters
//   e.g. FOOBar is two words (FOO and Bar)

package cases
