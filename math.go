module "Math"
import "math"

func sin(a float64) float64 {
	return math.Sin(a);
}
func cos(a float64) float64 {
	return math.Cos(a);
}
func tan(a float64) float64 {
	return math.Tan(a);
}
func asin(a float64) float64 {
	return math.Asin(a);
}
func acos(a float64) float64 {
	return math.Acos(a);
}
func atan(a float64) float64 {
	return math.Atan(a);
}
