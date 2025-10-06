# Matchgen - Struct Matcher Generator

Matchgen is a CLI tool that automatically generates matcher structs for struct-based input types, making it easy to create complex matchers for testing and validation.

## Installation

```bash
go build -o matchgen ./matchgen/
```

## Usage

```bash
./matchgen -out <output_file.go> -match_type <StructName> -out_type <MatcherName> [-out_package <PackageName>]
```

### Flags

- `-out`: The name of the .go file to generate
- `-match_type`: The name of the Go type to match against  
- `-out_type`: The name of the Go matcher type to generate
- `-out_package`: (Optional) Package name for the generated matcher. If not specified, uses the same package as the match_type

## Example

Given this input struct:

```go
package match

type User struct {
    ID       int64
    Username string
    Email    string
    IsActive bool
}
```

Run the generator:

```bash
./matchgen -out user_matcher.go -match_type User -out_type UserMatcher
```

Or to generate in a different package:

```bash
./matchgen -out user_matcher.go -match_type User -out_type UserMatcher -out_package matchers
```

This will generate:

```go
package match

import (
    "github.com/krelinga/go-match/matchfmt"
)

type UserMatcher struct {
    ID       Matcher[int64]
    Username Matcher[string]
    Email    Matcher[string]
    IsActive Matcher[bool]
}

func (m *UserMatcher) Match(got User) (bool, string) {
    var details []string
    allMatched := true

    if m.ID != nil {
        matched, explanation := m.ID.Match(got.ID)
        if !matched {
            allMatched = false
        }
        details = append(details, explanation)
    }

    if m.Username != nil {
        matched, explanation := m.Username.Match(got.Username)
        if !matched {
            allMatched = false
        }
        details = append(details, explanation)
    }

    if m.Email != nil {
        matched, explanation := m.Email.Match(got.Email)
        if !matched {
            allMatched = false
        }
        details = append(details, explanation)
    }

    if m.IsActive != nil {
        matched, explanation := m.IsActive.Match(got.IsActive)
        if !matched {
            allMatched = false
        }
        details = append(details, explanation)
    }

    return allMatched, matchfmt.Explain(allMatched, "UserMatcher", details...)
}
```

## Usage in Tests

```go
func TestUser(t *testing.T) {
    user := User{
        ID:       123,
        Username: "johndoe",
        Email:    "john@example.com",
        IsActive: true,
    }

    matcher := &UserMatcher{
        ID:       Equal(123),
        Username: StringContains("john"),
        Email:    StringHasSuffix("@example.com"),
        IsActive: Equal(true),
    }

    matched, explanation := matcher.Match(user)
    if !matched {
        t.Errorf("User should match: %s", explanation)
    }
}
```

## Features

- **Automatic field detection**: Only exported struct fields are included
- **Type preservation**: Field types are correctly preserved in the generated matcher
- **Nil field handling**: Fields set to nil are ignored during matching
- **Logical AND operation**: All non-nil matchers must pass for overall success
- **Rich explanations**: Uses matchfmt.Explain() for detailed match results
- **Import detection**: Automatically includes necessary imports (e.g., "time" package)
- **Cross-package support**: Can generate matchers in different packages with proper import handling

## Supported Field Types

- Basic types: `int`, `string`, `bool`, `float64`, etc.
- Pointers: `*string`, `*int`, etc.
- Slices: `[]string`, `[]int`, etc.
- Maps: `map[string]interface{}`, etc.
- Time: `time.Time`
- Custom structs: `UserProfile`, etc.
- Complex nested types: `map[string][]CustomType`