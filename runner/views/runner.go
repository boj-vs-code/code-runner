package views

import (
	"fmt"
	"github.com/boj-vs-code/code-runner/runner/models"
	"github.com/boj-vs-code/code-runner/runner/utils"
	"github.com/hwangseonu/gin-restful"
)


type Runner struct {
	*gin_restful.Resource
}

func InitRunnerResource() Runner {
	return Runner{gin_restful.InitResource()}
}

func (r Runner) Post(json models.RunnerRequest) (*models.RunnerResult, int) {
	fmt.Println(json)
	return utils.Run(&json), 200
}
