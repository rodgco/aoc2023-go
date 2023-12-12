package util

func Min(x, y int) int {
 if x < y {
   return x
 }
 return y
}

func Max(x, y int) int {
 if x > y {
   return x
 }
 return y
}

func MinUint32(x, y uint32) uint32 {
 if x < y {
   return x
 }
 return y
}

func MaxUint32(x, y uint32) uint32 {
 if x > y {
   return x
 }
 return y
}

func MaxInt64(x, y int64) int64 {
 if x > y {
   return x
 }
 return y
}

func MinInt64(x, y int64) int64 {
 if x < y {
   return x
 }
 return y
}
