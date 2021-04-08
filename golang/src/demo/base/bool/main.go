package main

func main() {
	c := 1
	// && 的优先级比 || 高（助记： && 对应逻辑乘法， || 对应逻辑加法，乘法比加法优先级要高）
	if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9' {
		// ...ASCII letter or digit...
	}
}

// 布尔值并不会隐式转换为数字值0或1，反之亦然。必须使用一个显式的if语句辅助转换
func itob(i int) bool { return i != 0 }
