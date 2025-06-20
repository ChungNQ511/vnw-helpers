# Go PGX Helpers

A comprehensive Go library providing utility functions and helpers for working with PostgreSQL using the `pgx` driver. This library simplifies common database operations, type conversions, and data handling when working with PostgreSQL in Go applications.

## Features

### 🗄️ Database Type Helpers
- **Type Conversion**: Convert Go types to PostgreSQL types (`pgtype.*`)
- **Type Reversion**: Convert PostgreSQL types back to Go types
- **Null Handling**: Proper handling of NULL values and validation
- **Generic Functions**: Type-safe conversions using Go generics

### 📅 Date & Time Utilities
- **Date Formatting**: Multiple predefined date/time formats
- **String Parsing**: Convert strings to PostgreSQL date/time types
- **Custom Formats**: Support for custom date/time formatting

### 🔧 Utility Functions
- **String Conversion**: Advanced string manipulation and conversion
- **Validation**: Null checking and data validation
- **Error Recovery**: Graceful error handling and recovery

## Installation

```bash
go get github.com/ChungNQ511/vnw-helpers
```

## Quick Start

```go
package main

import (
    "time"
    "github.com/ChungNQ511/vnw-helpers"
    "github.com/jackc/pgx/v5/pgtype"
)

func main() {
    // Convert Go types to PostgreSQL types
    textField := pgxhelpers.SetTextField("Hello World")
    intField := pgxhelpers.SetIntField[pgtype.Int4](42)
    dateField := pgxhelpers.SetDateField(time.Now())
    
    // Convert PostgreSQL types back to Go types
    text := pgxhelpers.RevertPgText(textField)
    number := pgxhelpers.RevertIntField(intField)
    date := pgxhelpers.RevertPgDate(dateField)
}
```

## Core Functions

### Type Conversion (To PostgreSQL)

#### Text Fields
```go
// Convert various types to pgtype.Text
text := pgxhelpers.SetTextField("Hello World")
text := pgxhelpers.SetTextField([]byte("Hello World"))
text := pgxhelpers.SetTextField(myStringer) // implements fmt.Stringer
```

#### Numeric Fields
```go
// Convert to pgtype.Float4 or pgtype.Float8
float4 := pgxhelpers.SetFloatField[pgtype.Float4](3.14)
float8 := pgxhelpers.SetFloatField[pgtype.Float8](3.14159)

// Convert to pgtype.Int2, pgtype.Int4, or pgtype.Int8
int4 := pgxhelpers.SetIntField[pgtype.Int4](42)
int8 := pgxhelpers.SetIntField[pgtype.Int8](123456789)
```

#### Date/Time Fields
```go
// Convert to pgtype.Date
date := pgxhelpers.SetDateField(time.Now())
date := pgxhelpers.SetDateField("2024-01-15")

// Convert to pgtype.Timestamp
timestamp := pgxhelpers.SetTimestampField(time.Now())
timestamp := pgxhelpers.SetTimestampField("2024-01-15 14:30:00")

// Convert to pgtype.Timestamptz
timestamptz := pgxhelpers.SetTimestamptzField(time.Now())
```

#### Boolean Fields
```go
// Convert to pgtype.Bool
boolField := pgxhelpers.SetBoolField(true)
boolField := pgxhelpers.PgBool("true")
```

#### Numeric Fields (Big Decimal)
```go
// Convert to pgtype.Numeric
numeric := pgxhelpers.SetNumericField("123.456")
numeric := pgxhelpers.SetNumericField(123.456)
```

### Type Reversion (From PostgreSQL)

#### Text Fields
```go
// Convert pgtype.Text back to string
text := pgxhelpers.RevertPgText(pgText)
```

#### Numeric Fields
```go
// Convert back to float64
float := pgxhelpers.RevertFloatField(pgFloat4)
float := pgxhelpers.RevertFloatField(pgFloat8)

// Convert back to int64
integer := pgxhelpers.RevertIntField(pgInt4)
integer := pgxhelpers.RevertIntField(pgInt8)
```

#### Date/Time Fields
```go
// Convert back to time.Time
date := pgxhelpers.RevertPgDate(pgDate)
timestamp := pgxhelpers.RevertPgTimestamp(pgTimestamp)
timestamptz := pgxhelpers.RevertPgTimestamptz(pgTimestamptz)
```

#### Boolean Fields
```go
// Convert back to bool
boolean := pgxhelpers.RevertPgBool(pgBool)
```

## Date/Time Utilities

### Predefined Formats
```go
import "github.com/ChungNQ511/vnw-helpers/datecvx"

// Available formats
datecvx.Date_DDMMYYYY        // "02/01/2006"
datecvx.Date_YYYYMMDD        // "2006/01/02"
datecvx.Date_DDMMYYYY_HHMM   // "02/01/2006 15:04"
datecvx.Date_DDMMYYYY_HHMMSS // "02/01/2006 15:04:05"
datecvx.Time_HHMMSS          // "15:04:05"
datecvx.Time_HHMM            // "15:04"
datecvx.DateTime_RFC3339     // RFC3339 format
```

### Formatting Functions
```go
// Format time.Time to string
formatted := datecvx.FormatTime(time.Now(), datecvx.Date_DDMMYYYY)
// Output: "15/01/2024"

// Format with custom format
formatted := datecvx.FormatTimeCustom(time.Now(), datecvx.Date_DDMMYYYY_HHMMSS)
// Output: "15/01/2024 14:30:25"
```

## String Conversion Utilities

### Array/Slice Conversion
```go
import "github.com/ChungNQ511/vnw-helpers/strconvx"

// Convert PostgreSQL array string to Go slice
numbers := strconvx.ConvertToSlice[int]("{1,2,3,4,5}")
strings := strconvx.ConvertToSlice[string]("{hello,world,test}")

// Handle empty/null arrays
empty := strconvx.ConvertToSlice[int]("{}")     // Returns empty slice
empty := strconvx.ConvertToSlice[int]("null")   // Returns empty slice
```

## Utility Functions

### Null Checking
```go
import "github.com/ChungNQ511/vnw-helpers/funcvx"

// Check if value is not null/empty
isValid := funcvx.NotNull("hello")     // true
isValid := funcvx.NotNull("")          // false
isValid := funcvx.NotNull("null")      // false
isValid := funcvx.NotNull(0)           // false
isValid := funcvx.NotNull(42)          // true
isValid := funcvx.NotNull(time.Time{}) // false
```

## Usage Examples

### Database Insert Example
```go
type User struct {
    ID       int64
    Name     string
    Age      int
    Email    string
    Created  time.Time
    IsActive bool
}

func InsertUser(db *pgxpool.Pool, user User) error {
    query := `
        INSERT INTO users (name, age, email, created_at, is_active)
        VALUES ($1, $2, $3, $4, $5)
    `
    
    _, err := db.Exec(context.Background(), query,
        pgxhelpers.SetTextField(user.Name),
        pgxhelpers.SetIntField[pgtype.Int4](user.Age),
        pgxhelpers.SetTextField(user.Email),
        pgxhelpers.SetTimestampField(user.Created),
        pgxhelpers.SetBoolField(user.IsActive),
    )
    
    return err
}
```

### Database Query Example
```go
func GetUser(db *pgxpool.Pool, id int64) (*User, error) {
    query := `SELECT name, age, email, created_at, is_active FROM users WHERE id = $1`
    
    var name pgtype.Text
    var age pgtype.Int4
    var email pgtype.Text
    var created pgtype.Timestamp
    var isActive pgtype.Bool
    
    err := db.QueryRow(context.Background(), query, id).Scan(
        &name, &age, &email, &created, &isActive,
    )
    if err != nil {
        return nil, err
    }
    
    user := &User{
        ID:       id,
        Name:     pgxhelpers.RevertPgText(name),
        Age:      int(pgxhelpers.RevertIntField(age)),
        Email:    pgxhelpers.RevertPgText(email),
        Created:  pgxhelpers.RevertPgTimestamp(created),
        IsActive: pgxhelpers.RevertPgBool(isActive),
    }
    
    return user, nil
}
```

## Requirements

- Go 1.23.0 or higher
- `github.com/jackc/pgx/v5` v5.7.5 or higher
- `golang.org/x/sync` v0.15.0 or higher

## License

This project is licensed under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Support

If you encounter any issues or have questions, please open an issue on the GitHub repository. 