package GoFormatFix

import (
  "github.com/GammaWatt/Go-Seq-Util"
)

func paddedCents(x string, pennySize uint) string {
  var cents string
  for i := uint(0); i < pennySize; i++ {
    cents += "0"
  }
  result := cents + x
  return result[uint(len(result))-pennySize:]
}

func calculateCurrencyFormatDivisions(y []byte, spacing uint) uint {
  if spacing == uint(0) {
    return 0
  } else {
    var points uint = uint(len(y)) / spacing
    if uint(len(y)) % spacing == 0 {
      // we don't want 2 points if spacing is 3 and there are 6 digits...
      points--
    }
    return points
  }
}

func insertCurrencyMarkers(x, separator, decimalMarker string, spacing, pennySize uint) string {
  // pennySize and spacing aren't allowed to be negative
  var cents string = paddedCents(x, pennySize)
  var numberFinish string = string(decimalMarker) + cents
  if uint(len(x)) < 1 + pennySize {
    return "0" + numberFinish
  }  else {
    if pennySize == uint(0) {
      numberFinish = ""
    }
    y := []byte(x[:uint(len(x))-pennySize])
    var points uint = calculateCurrencyFormatDivisions(y, spacing)
    var formattedNumber []byte = make([]byte, uint(len(y)) + (points * uint(len(separator))))
    if uint(len(y)) > spacing {
      breakPoint := uint(1)
      for i, cursor := uint(len(y)), 0; i > 0; i-- {
        formattedNumber[cursor] = y[uint(len(y))-i]
        cursor++
        if points > 0 {
          if i % spacing == breakPoint && i > spacing {
            for o := range separator {
              formattedNumber[cursor] = separator[o]
              cursor++
            }
          }
        }
      }
    } else {
      formattedNumber = y
    }
    var result string
    if len(formattedNumber) < 1 {
      result = "Check number val... It's blank after formatting..."
    } else {
      result = string(formattedNumber)
    }
    return result + numberFinish
  }
}

func prependPrefix(prefix, x string) string {
    return prefix + x
}

func appendSuffix(suffix, x string) string {
  return x + suffix
}

// Remember, GoFormatFix.Format takes a function reference as an second argument.
// It automatically plugs a stringified x into it and runs it
func currencyFormat(prefix, suffix, separator, decimalMarker string, spacing, pennySize uint) func(string) string {
  return func (x string) string {
    var thousandsSeparator string = separator
    var decimalSeparator string = decimalMarker
    y := GoSeqUtil.RemoveNonNumericChars(x)
    y = insertCurrencyMarkers(y, thousandsSeparator, decimalSeparator, spacing, pennySize)
    y = prependPrefix(prefix, y)
    y = appendSuffix(suffix, y)
    return y
  }
}

func FormatCurrencyCustom(x interface{}, prefix, suffix, separator, decimalMarker string, spacing, pennySize uint) string {
  return Format(x, currencyFormat(prefix, suffix, separator, decimalMarker, spacing, pennySize))
}

func FormatCurrency(x interface{}) string {
  return FormatCurrencyCustom(x, "$ ", " USD", ",", ".", 3, 2)
}

func FormatCurrencyCO(x interface{}) string {
  return FormatCurrencyCustom(x, "$ ", "", ".", ",", 3, 2)
}
