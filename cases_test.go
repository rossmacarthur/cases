package cases_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/rossmacarthur/cases"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	in        string
	snakeCase string
	camelCase string
}

var tests []testCase = []testCase{
	{},

	{in: "Test",
		snakeCase: "test",
		camelCase: "test",
	},

	{in: "test case",
		snakeCase: "test_case",
		camelCase: "testCase",
	},

	{in: " test case",
		snakeCase: "test_case",
		camelCase: "testCase",
	},

	{in: "test case ",
		snakeCase: "test_case",
		camelCase: "testCase",
	},

	{in: "Test Case",
		snakeCase: "test_case",
		camelCase: "testCase",
	},

	{in: " Test Case",
		snakeCase: "test_case",
		camelCase: "testCase",
	},

	{in: "camelCase",
		snakeCase: "camel_case",
		camelCase: "camelCase",
	},

	{in: "PascalCase",
		snakeCase: "pascal_case",
		camelCase: "pascalCase",
	},

	{in: "snake_case",
		snakeCase: "snake_case",
		camelCase: "snakeCase",
	},

	{in: "SCREAMING_SNAKE_CASE",
		snakeCase: "screaming_snake_case",
		camelCase: "screamingSnakeCase",
	},

	{in: "kebab-case",
		snakeCase: "kebab_case",
		camelCase: "kebabCase",
	},

	{in: "SCREAMING-KEBAB-CASE",
		snakeCase: "screaming_kebab_case",
		camelCase: "screamingKebabCase",
	},

	{in: "Title Case ",
		snakeCase: "title_case",
		camelCase: "titleCase",
	},

	{in: "Train-Case ",
		snakeCase: "train_case",
		camelCase: "trainCase",
	},

	{in: "This is a Test case.",
		snakeCase: "this_is_a_test_case",
		camelCase: "thisIsATestCase",
	},

	{in: "MixedUP CamelCase, with some Spaces",
		snakeCase: "mixed_up_camel_case_with_some_spaces",
		camelCase: "mixedUpCamelCaseWithSomeSpaces",
	},

	{in: "mixed_up_ snake_case with some _spaces",
		snakeCase: "mixed_up_snake_case_with_some_spaces",
		camelCase: "mixedUpSnakeCaseWithSomeSpaces",
	},

	{in: "this-contains_ ALLKinds OfWord_Boundaries",
		snakeCase: "this_contains_all_kinds_of_word_boundaries",
		camelCase: "thisContainsAllKindsOfWordBoundaries",
	},

	{in: "XΣXΣ baﬄe",
		snakeCase: "xσxσ_baﬄe",
		camelCase: "xσxσBaﬄe",
	},

	{in: "XMLHttpRequest",
		snakeCase: "xml_http_request",
		camelCase: "xmlHttpRequest",
	},

	{in: "FIELD_NAME11",
		snakeCase: "field_name11",
		camelCase: "fieldName11",
	},

	{in: "99BOTTLES",
		snakeCase: "99bottles",
		camelCase: "99bottles",
	},

	{in: "FieldNamE11",
		snakeCase: "field_nam_e11",
		camelCase: "fieldNamE11",
	},

	{in: "abc123def456",
		snakeCase: "abc123def456",
		camelCase: "abc123def456",
	},

	{in: "abc123DEF456",
		snakeCase: "abc123_def456",
		camelCase: "abc123Def456",
	},

	{in: "abc123Def456",
		snakeCase: "abc123_def456",
		camelCase: "abc123Def456",
	},
	{in: "abc123DEf456",
		snakeCase: "abc123_d_ef456",
		camelCase: "abc123DEf456",
	},

	{in: "ABC123def456",
		snakeCase: "abc123def456",
		camelCase: "abc123def456",
	},

	{in: "ABC123DEF456",
		snakeCase: "abc123def456",
		camelCase: "abc123def456",
	},

	{in: "ABC123Def456",
		snakeCase: "abc123_def456",
		camelCase: "abc123Def456",
	},

	{in: "ABC123DEf456",
		snakeCase: "abc123d_ef456",
		camelCase: "abc123dEf456",
	},

	{in: "ABC123dEEf456FOO",
		snakeCase: "abc123d_e_ef456_foo",
		camelCase: "abc123dEEf456Foo",
	},

	{in: "abcDEF",
		snakeCase: "abc_def",
		camelCase: "abcDef",
	},

	{in: "ABcDE",
		snakeCase: "a_bc_de",
		camelCase: "aBcDe",
	},
}

func TestToCamel(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			got := cases.ToCamel(tc.in)
			require.Equal(t, tc.camelCase, got, fmt.Sprintf("'%s'", tc.in))
		})
	}
}

func TestToPascal(t *testing.T) {
	result := cases.ToPascal("test case")
	require.Equal(t, "TestCase", result)
}

func TestToSnake(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			got := cases.ToSnake(tc.in)
			require.Equal(t, tc.snakeCase, got, fmt.Sprintf("'%s'", tc.in))
		})
	}
}

func TestToScreamingSnake(t *testing.T) {
	result := cases.ToScreamingSnake("test case")
	require.Equal(t, "TEST_CASE", result)
}

func TestToKebab(t *testing.T) {
	result := cases.ToKebab("test case")
	require.Equal(t, "test-case", result)
}

func TestToScreamingKebab(t *testing.T) {
	result := cases.ToScreamingKebab("test case")
	require.Equal(t, "TEST-CASE", result)
}

func TestToTitle(t *testing.T) {
	result := cases.ToTitle("test case")
	require.Equal(t, "Test Case", result)
}

func TestToTrain(t *testing.T) {
	result := cases.ToTrain("test case")
	require.Equal(t, "Test-Case", result)
}

func BenchmarkToSnake(b *testing.B) {
	s := strings.Repeat("ThisIs a/Test_case.", 100)
	for i := 0; i < b.N; i++ {
		cases.ToSnake(s)
	}
}
