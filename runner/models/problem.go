package models

// ... github.com/boj-vs-code/problem/server/models/problem.go
type Problem struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	InputDescription string `json:"inputDescription"`
	OutputDescription string `json:"outputDescription"`
	Testcases []string `json:"testcases"`
}
