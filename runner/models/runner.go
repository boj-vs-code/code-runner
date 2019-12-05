package models


type RunnerRequest struct {
	ProblemId rune `json:"problem_id" validate:"required"`
	Code string `json:"code" validate:"required"`
}


type RunnerResult struct {
	Code string `json:code`
	Success bool `json:success`
	Failed [][3]string `json:failed`
	Message string `json:message`
}
