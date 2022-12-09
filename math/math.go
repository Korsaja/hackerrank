package math

func findPoint(px int32, py int32, qx int32, qy int32) []int32 {
	return []int32{2*qx - px, 2*qy - py}
}

func maximumDraws(n int32) int32 {
	return n + 1
}

func handshake(n int32) int32 {
	return n * (n - 1) >> 1
}

func lowestTriangle(triangle int32, area int32) int32 {
	s := (2 * area) / triangle
	if (2*area)%triangle != 0 {
		s++
	}
	return s
}
