/*
Что нарушает принцип DRY: скопированный код или дублирующая бизнес-логика в разных функциях?

Упростите этот код, избегая дублирования (DRY + KISS):
func CalcPriceA(qty int) int {
	return qty*100 + qty*10
}
func CalcPriceB(qty int) int {
	return qty*100 + qty*10 + 5
}
*/