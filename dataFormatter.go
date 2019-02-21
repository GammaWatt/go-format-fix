package GoFormatFix

import (
  "strconv"
  "reflect"
)

func normalizeToUint64(x interface{}) uint64 {
  ui8ToUi64 := func (x interface{}) uint64 {
    r := x.(uint8)
    var f uint64
    f = uint64(r)
    return f
  }

  ui16ToUi64 := func (x interface{}) uint64 {
    r, l := x.(uint16)
  	if !l {
      return ui8ToUi64(x)
  	} else {
  	 var f uint64
     f = uint64(r)
     return f
  	}
  }

  ui32ToUi64 := func (x interface{}) uint64 {
    r, l := x.(uint32)
  	if !l {
  	 return ui16ToUi64(x)
  	} else {
  	 var f uint64
  	 f = uint64(r)
     return f
   }
 }

  ui64ToUi64 := func (x interface{}) uint64 {
    r, l := x.(uint64)
  	if !l {
  	 return ui32ToUi64(x)
  	} else {
     return r
  	}
  }

  r, l := x.(uint)
	if !l {
	 return ui64ToUi64(x)
	} else {
	 var f uint64
	 f = uint64(r)
   return f
  }
}


func normalizeToInt64(x interface{}) int64 {
  i8Toi64 := func (x interface{}) int64 {
    r := x.(int8)
    var f int64
    f = int64(r)
    return f
  }

  i16Toi64 := func (x interface{}) int64 {
    r, l := x.(int16)
  	if !l {
      return i8Toi64(x)
  	} else {
  	 var f int64
     f = int64(r)
     return f
  	}
  }

  i32Toi64 := func (x interface{}) int64 {
    r, l := x.(int32)
  	if !l {
  	 return i16Toi64(x)
  	} else {
  	 var f int64
  	 f = int64(r)
     return f
   }
 }

  i64Toi64 := func (x interface{}) int64 {
    r, l := x.(int64)
  	if !l {
  	 return i32Toi64(x)
  	} else {
     return r
  	}
  }

  r, l := x.(int)
	if !l {
	 return i64Toi64(x)
	} else {
	 var f int64
	 f = int64(r)
   return f
  }
}

func normalizeDataType(x interface{}) string {
  if reflect.TypeOf(x).String()[:3] == "int" {
    y := strconv.FormatInt(normalizeToInt64(x), 10)
    return y
  } else if reflect.TypeOf(x).String()[:4] == "uint" {
    y := strconv.FormatUint(normalizeToUint64(x), 10)
    return y
  } else if dataType := reflect.TypeOf(x).String(); dataType[:5] == "float" {
    bitSize, _ := strconv.Atoi(dataType[5:7])
    var y string
    if bitSize == 64 {
      y = strconv.FormatFloat(x.(float64), 'f', -1, bitSize)
    } else {
      // strconv only accepts float64 arg
      y = strconv.FormatFloat(float64(x.(float32)), 'f', -1, bitSize)
    }
    return y
  } else if reflect.TypeOf(x).String() == "string" {
    y := x.(string)
    return y
    // []byte is evaluates as []uint8 by reflect
  } else if reflect.TypeOf(x).String() == "[]uint8" {
    y := string(x.([]byte))
    return y
    // []rune is evaluated as []int32 by reflect
  } else if reflect.TypeOf(x).String() == "[]int32" {
    y := string(x.([]rune))
    return y
  } else {
    panic(
      reflect.ValueOf(x).String() +
      " must be of type float, int, uint string, []byte, or []rune...")
  }
}

type formatterFunction func(string) string

func Format(x interface{}, y formatterFunction) string {
  n := normalizeDataType(x)
  return y(n)
}
// Format takes x as string and plugs it into y to be formatted as you wish
// always returns a string as formatted by y
// Acceptable types are any float, any int, any uint, string, []byte, or []rune.
