package pkg1


var logMessage = "[LOG]"
// Version of the calculator
var Version = "1.0"

func internalSum(number int) int {
	return number - 1
}

// 计算，加起来
func Sum1(number1, number2 int) int {
	return number1 + number2
}
