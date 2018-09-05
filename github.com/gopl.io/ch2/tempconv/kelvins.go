package tempconv

type Kelvins float64

func CToT(c Celsius) Kelvins{
	return Kelvins(c+273.15)
} 