package util

import (
	"math"

)

var KB = uint64(math.Pow(2,10))
var MB = uint64(math.Pow(2,20))
var GB = uint64(math.Pow(2,30))
var TB = uint64(math.Pow(2,40))

func TempToFarenhiet(c int)int{
	return c*9/5 +32
}

func BytesToKB(b uint64) float64{
	return float64(b)/float64(KB)
}
func BytesToMB(b uint64)float64  {
	return float64(b)/float64(MB)
}
func BytesToGB(b uint64)float64  {
	return float64(b)/float64(GB)
}
func BytesToTB(b uint64)float64{
	return float64(b)/float64(TB)
}
