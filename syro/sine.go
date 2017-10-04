package syro

var sinTable = [...]int16{
	0, 23169, 32767, 23169, 0, -23169, -32767, -23169,
}

func Sin(phase int, bData bool) int16 {
	sin := int32(sinTable[phase])

	if bData {
		if sin > 0 {
			sin = 32767 - sin
			sin = (sin * sin) / 32767
			sin = 32767 - sin
		} else if sin < 0 {
			sin += 32767
			sin = (sin * sin) / 32767
			sin -= 32767
		}
	}

	return int16(sin)
}
