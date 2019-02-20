# Go-Format-Fix
Takes just about any builtin datatype and returns a formatted string as per your specification.

This is particularly useful for maintaining a consistent formatting for phone numbers, currencies, dates, and the like, but can be applied to any string that may need to be sanitized and formatted in a specific way.

# Usage

```
package main

// It's recommended to use an alias for the package,
// simply because it makes it easier to refer to after importing it.

import (
  formatter "github.com/GammaWatt/go-format-fix"
  "fmt"
  )


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

// x is converted to a string, then passed to your formatting function.
// It works with any width of int or float so you don't have to worry
// about type casting.

func testInt(x int8) string {
  return formatter.Format(x, myFormatter)
}

func testUint(x uint64) string {
  return formatter.Format(x, myFormatter)
}

func testF32(x float32) string {
  return formatter.Format(x, myFormatter)
}

func testF64(x float64) string {
  return formatter.Format(x, myFormatter)
}

func testByte(x []byte) string {
  return formatter.Format(x, myFormatter)
}

func testRune(x []rune) string {
  return formatter.Format(x, myFormatter)
}

func testString(x string) string {
  return formatter.Format(x, myFormatter)
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
