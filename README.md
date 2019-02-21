# Go-Format-Fix
Takes just about any builtin datatype and returns a formatted string as per your specification.

This is particularly useful for maintaining a consistent formatting when receiving input for phone numbers, currencies, dates, and the like, but can be applied to any string that may need to be sanitized and formatted in a specific way.

**Note that this library does no number rounding when performing its processes.**

# Usage

The package was primarily designed to be used with interface values: i.e. functions of `func (interface{}) string` type.  
The functions passed as the formatters must always be of type `func (string) string`

```
// The formatter function you define must always be of type
// func (string) string
func myFormatter(y string) string {
  var runes []rune = []rune(y)
  var formattedResult []rune

  for i := range runes {
    formattedResult = append(formattedResult, runes[i])
    formattedResult = append(formattedResult, rune('^'))
  }
  return string(formattedResult[:len(formattedResult)-1])
}

func testInterface(x interface{}) string {
  return GoFormatFix.Format(x, myFormatter)
}

func main() {
    fmt.Println(testInterface(int8(127)))
    fmt.Println(testInterface(uint64(18446744073709551615)))
    fmt.Println(testInterface(float32(3.1415926)))
    fmt.Println(testInterface(float64(2.7182818284590459)))
    fmt.Println(testInterface([]byte{'a','b','c','d','e','f','g'}))
    fmt.Println(testInterface([]rune{'א','ב','ג','ד','ה','ו','ז'}))
    fmt.Println(testInterface(string("There be dragons")))
}
```
###### output
```
1^2^7
1^8^4^4^6^7^4^4^0^7^3^7^0^9^5^5^1^6^1^5
3^.^1^4^1^5^9^2^5
2^.^7^1^8^2^8^1^8^2^8^4^5^9^0^4^6
a^b^c^d^e^f^g
א^ב^ג^ד^ה^ו^ז
T^h^e^r^e^ ^b^e^ ^d^r^a^g^o^n^s
```

Although, if you choose, you could also do it with functions that require specific types as well without changing anything.

```
package main

// The package name itself is GoFormatFix, but aliasing the package import might also be helpful

import (
  "github.com/GammaWatt/go-format-fix"
  "fmt"
  )

func myFormatter(y string) string {
  var runes []rune = []rune(y)
  var formattedResult []rune

  for i := range runes {
    formattedResult = append(formattedResult, runes[i])
    formattedResult = append(formattedResult, rune('^'))
  }
  return string(formattedResult[:len(formattedResult)-1])
}

// x is converted to a string, then passed to your formatting function.
// It works with any width of int or float so you don't have to worry
// about type casting.

func testInt(x int8) string {
  return GoFormatFix.Format(x, myFormatter)
}

func testUint(x uint64) string {
  return GoFormatFix.Format(x, myFormatter)
}

func testF32(x float32) string {
  return GoFormatFix.Format(x, myFormatter)
}

func testF64(x float64) string {
  return GoFormatFix.Format(x, myFormatter)
}

func testByte(x []byte) string {
  return GoFormatFix.Format(x, myFormatter)
}

func testRune(x []rune) string {
  return GoFormatFix.Format(x, myFormatter)
}

func testString(x string) string {
  return GoFormatFix.Format(x, myFormatter)
}

func testInterface(x interface{}) string {
  return GoFormatFix.Format(x, myFormatter)
}

func main() {
  fmt.Println(testInt(int8(127)))
  fmt.Println(testUint(uint64(18446744073709551615)))
  // f32 only has 8-point precision... you'll get errors beyond that
  fmt.Println(testF32(float32(3.1415926)))
  // f64 only has 16-point precision... you'll get errors beyond that
  fmt.Println(testF64(float64(2.7182818284590459)))
  fmt.Println(testByte([]byte{'a','b','c','d','e','f','g'}))
  fmt.Println(testRune([]rune{'א','ב','ג','ד','ה','ו','ז'}))
  fmt.Println(testString(string("There be dragons")))
}
```

###### output

```
1^2^7
1^8^4^4^6^7^4^4^0^7^3^7^0^9^5^5^1^6^1^5
3^.^1^4^1^5^9^2^5
2^.^7^1^8^2^8^1^8^2^8^4^5^9^0^4^6
a^b^c^d^e^f^g
א^ב^ג^ד^ה^ו^ז
T^h^e^r^e^ ^b^e^ ^d^r^a^g^o^n^s
```

# Extras

### Currency Formatting

An included default for currency formatting exists, but an option for custom arrangements exists as well:

* Again, type is nearly irrelevant, you can pass an int, string, float, []byte, or []rune and the result will be the same.

```
  // Default
  fmt.Println(GoFormatFix.FormatCurrency(1123456723456789))

  // Custom
  fmt.Println(GoFormatFix.FormatCurrencyCustom(1123456723456789, "~", " Bits", "-:-", "_|_", 2, 2))
  fmt.Println(GoFormatFix.FormatCurrencyCustom(1123456723456789, "<&>", " Bits", "-:-", "_|_", 2, 2))
  fmt.Println(GoFormatFix.FormatCurrencyCustom(1123456723456789, "<>", " Bits", "!", "-", 4, 22))
  fmt.Println(GoFormatFix.FormatCurrencyCustom(1123456723456789, "$ ", " Bits", "~", "..", 6, 7))
  fmt.Println(GoFormatFix.FormatCurrencyCustom("1123456723456789", "$ ", " Bits", "~", "..", 6, 8))
  fmt.Println(GoFormatFix.FormatCurrencyCustom(11234567.23456789, "$ ", " Bits", "~", "..", 6, 9))
  fmt.Println(GoFormatFix.FormatCurrencyCustom([]byte{1,1,2,3,4,5,6,7,2,3,4,5,6,7,8,9}, "$ ", " Bits", "~", "..", 6, 10))
  fmt.Println(GoFormatFix.FormatCurrencyCustom(0, "$ ", " Bits", "~", "..", 6, 10))
  fmt.Println(GoFormatFix.FormatCurrencyCustom(0, "$ ", " Bits", "~", "..", 6, 0))
  fmt.Println(GoFormatFix.FormatCurrencyCustom(0, "$ ", " Bits", "~", "..", 0, 2))
  fmt.Println(GoFormatFix.FormatCurrencyCustom(1234, "$ ", " Bits", "~", "..", 0, 4))
```
###### output

```
$ 11,234,567,234,567.89 USD
~11-:-23-:-45-:-67-:-23-:-45-:-67_|_89 Bits
<&>11-:-23-:-45-:-67-:-23-:-45-:-67_|_89 Bits
<>0-0000001123456723456789 Bits
$ 112~345672..3456789 Bits
$ 11~234567..23456789 Bits
$ 1~123456..723456789 Bits
$ 0..0000000000 Bits
$ 0..0000000000 Bits
$ 0 Bits
$ 0..00 Bits
$ 0..1234 Bits
```

Notice that for `FormatCurrencyCustom`:
* The first argument is the value to be formatted.  
* The second argument is the prefix to the number. (usually a dollar sign)
* The third argument is the prefix to the number. (usually a the currency abbreviation, or blank ("") )
* The fourth argument is the marker used to divide digit sets. (usually a comma)
* The fifth argument is the marker used to divide the decimal section from the whole numbers (usually a period, to denote cents) (wholeNumbers.rationalNumbers) ($1.99 == 1 whole / 99 rational)
* The sixth argument is used to specify how many digits denote a section. (usually 3) (123,456 == 2 sections of 3 numbers, divided by a comma (,) )
* The seventh argument is used to specify how many digits to the left you want the marker in the fifth argument to be placed. (i.e. the precision of the number) (if you want 1234 -> 1.234 <-- i.e. 3 digits to the left, this value would be 3. a value of four would produce 0.1234)

**Note: The values for the sixth and seventh arguments (digit spacing and precision) cannot be negative values. They are set as uint types for this reason and will overflow to a very large value or error out if a negative value is passed.**

# Contributing

* Bug fixes and suggestions are welcome.
