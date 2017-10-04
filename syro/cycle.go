package syro

const (
	KORGSYRO_NUM_OF_CHANNEL   = 2
	KORGSYRO_QAM_CYCLE        = 8
	KORGSYRO_NUM_OF_CYCLE     = 2
	KORGSYRO_NUM_OF_CYCLE_BUF = KORGSYRO_QAM_CYCLE * KORGSYRO_NUM_OF_CYCLE
)

type Channel struct {
	CycleSample [KORGSYRO_NUM_OF_CYCLE_BUF]int16
	LastPhase   int
	LpfZ        int32
}

func SingleCycle(sc *Channel, writePage int, dat uint8, block bool) {
	var phaseOrg, phase, vol, dlt, writePos, writePosLast int
	var dat1, dat2 int32

	writePos = writePage * 8
	writePosLast = (8 * 2) - 1
	if writePos != 0 {
		writePosLast = writePos - 1
	}

	phaseOrg = int((dat >> uint64(1)) & 3)
	phase = phaseOrg * (8 / 4)
	vol = int(dat & 1)
	if vol != 0 {
		vol = 16
	} else {
		vol = 4
	}

	for i := uint8(0); i < KORGSYRO_QAM_CYCLE; i++ {
		dat1 = int32(Sin(phase, block))
		dat1 = (dat1 * int32(vol)) / 24
		if i == 0 {
			if phaseOrg != sc.LastPhase {
				if (sc.LastPhase&1)&(phaseOrg&1)|((sc.LastPhase+1)&3) == phaseOrg {
					dat2 = int32(sc.CycleSample[writePosLast])
					dlt = int(dat1 - dat2)
					dlt /= 3
					dat1 -= int32(dlt)
					dat2 += int32(dlt)
					sc.CycleSample[writePosLast] = int16(dat2)
				}
			}
		}

		sc.CycleSample[writePos] = int16(dat1)
		writePos++

		if (phase + 1) == KORGSYRO_QAM_CYCLE {
			phase = 0
		}
	}
	sc.LastPhase = phaseOrg
}

func SmoothStartMark(sc *Channel, writePage int) {
	var writePos, writePosLast int
	var dat1, dat2, dat3, avg int32

	writePos = writePage * 8
	writePosLast = func() int {
		if writePos != 0 {
			return writePos - 1
		} else {
			return (8 * 2) - 1
		}
	}()

	dat1 = int32(sc.CycleSample[writePosLast])
	dat2 = int32(sc.CycleSample[writePos])
	dat3 = int32(sc.CycleSample[writePos+1])

	avg = (dat1 + dat2 + dat3) / 3

	dat1 = (dat1 + avg) / 2
	dat2 = (dat2 + avg) / 2
	dat3 = (dat3 + avg) / 2

	sc.CycleSample[writePosLast] = int16(dat1)
	sc.CycleSample[writePos] = int16(dat2)
	sc.CycleSample[writePos+1] = int16(dat3)
}

func Gap(scs []*Channel, writePage int) {
	for ch := uint8(0); ch < KORGSYRO_NUM_OF_CHANNEL; ch++ {
		SingleCycle(scs[ch], writePage, 1, false)
	}
}

func StartMark(scs []*Channel, writePage int) {
	for ch := uint8(0); ch < KORGSYRO_NUM_OF_CHANNEL; ch++ {
		SingleCycle(scs[ch], writePage, uint8(5), false)
		SmoothStartMark(scs[ch], writePage)
	}
}

func ChannelInfo(scs []*Channel, writePage int) {
	for ch := uint8(0); ch < KORGSYRO_NUM_OF_CHANNEL; ch++ {
		SingleCycle(scs[ch], writePage, ch, true)
	}
}
