package unitconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }

func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

func FeetToMeter(f Feet) Meter { return Meter(f * 0.3048) }

func MeterToFeet(m Meter) Feet { return Feet(m / 0.3048) }

func PoundToKilogram(p Pound) Kilogram { return Kilogram(p * 0.453592) }

func KilogramToPound(k Kilogram) Pound { return Pound(k / 0.453592) }
