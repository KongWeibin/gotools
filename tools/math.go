package tools

//Abs 求一个int64绝对值
func Abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
