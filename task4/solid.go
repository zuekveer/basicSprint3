/*
Приведи пример нарушения принципа Single Responsibility.

Переделай структуру Report так, чтобы она отвечала только за данные, а печать и экспорт были вынесены:
type Report struct {
	Title string
	Body  string
}
func (r Report) Print() {}       // ← вынести
func (r Report) ExportJSON() {}  // ← вынести

*/