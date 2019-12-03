package main

import (
	"github.com/boj-vs-code/code-runner/runner/views"
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-restful"
)

func main() {
	r := gin.Default()
	v1 := gin_restful.NewApi(r, "/")
	views.RegisterViews(v1)
	_ = r.Run(":5000")
}
