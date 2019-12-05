package views

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

func (r Runner) Post(json models.RunnerRequest) (gin.H, int) {
	fmt.Println(json)
	runnerResult := utils.Run(json)
	return gin.H {
		"code": runnerResult.Code,
		"success": runnerResult.Success,
		"failed": runnerResult.Failed,
		"message": runnerResult.Message,
	}, 200
}
