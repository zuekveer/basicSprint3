/*
Что нарушает принцип DRY: скопированный код или дублирующая бизнес-логика в разных функциях?
Оба


Упростите этот код, избегая дублирования (DRY + KISS):
func CalcPriceA(qty int) int {
	return qty*100 + qty*10
}
func CalcPriceB(qty int) int {
	return qty*100 + qty*10 + 5
}
*/

func calcPrice(qty, basePrice, percent, extra int) int {
    return qty*basePrice + qty*percent + extra
}

func CalcPriceA(qty int) int {
    return calcPrice(qty, 100, 10, 0)
}

func CalcPriceB(qty int) int {
    return calcPrice(qty, 100, 10, 5)
}