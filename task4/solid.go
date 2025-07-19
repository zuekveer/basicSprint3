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

type Report struct {
	Title string
	Body  string
}

type Printer struct{}

func (p Printer) Print(report Report) {
    fmt.Printf("Отчёт: %s\n%s\n", report.Title, report.Body)
}

type JSONExporter struct{}

func (e JSONExporter) ExportJSON(report Report) []byte {
    data := map[string]string{
        "title": report.Title,
        "body":  report.Body,
    }
    jsonData, _ := json.Marshal(data)
    return jsonData
}