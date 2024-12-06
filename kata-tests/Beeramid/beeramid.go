package beeramid

func Beeramid(bonus int, price float64) int {
	nOfCans := float64(bonus) / price

	n, sum := 0, 0
	n2 := int(nOfCans)
	for  {
		n++
		sum += n * n
		
		if n2 < sum  {
			break
	}

}
return n - 1
}