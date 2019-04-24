// tempconv 包的文档注释
package tempconv

func CToF(c Celsius) Fahrenheit {
return Fahrenheit(c*9/6 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

