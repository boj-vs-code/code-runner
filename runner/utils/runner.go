package utils

import (
	"github.com/boj-vs-code/code-runner/runner/models"
	"log"
	"os"
	"os/exec"
	"strings"
)

func GetExecutionScript() string {
	return os.Getenv("RUNNER_SCRIPT")
}

func GetCompileScript() string {
	return os.Getenv("RUNNER_COMPILE")
}

func GetSourceFilename() string {
	return os.Getenv("RUNNER_SOURCE_FILENAME")
}

func Run(runnerRequest *models.RunnerRequest) *models.RunnerResult {
	problemId := runnerRequest.ProblemId
	code := runnerRequest.Code
	sourceFilename := GetSourceFilename()
	compile := GetCompileScript()
	script := GetExecutionScript()

	problem := models.GetProblemById(problemId)
	testCases := flatAndGroup(problem.Testcases)

	var runnerResult models.RunnerResult
	runnerResult.Code = code
	runnerResult.Failed = make([][3]string, 0)
	runnerResult.Success = true

	setUpSourceFile(code, sourceFilename)
	runCommand(compile, "")
	for i := 0; i < len(testCases); i++ {
		input, expected := testCases[i][0], testCases[i][1]
		success, actual := runTestCase(script, input, expected)
		runnerResult.Success = runnerResult.Success && success
		if !success {
			log.Printf("failed %s %s %s\n", input, expected, actual)
			runnerResult.Failed = append(runnerResult.Failed, [3]string{input, expected, actual})
		}
	}
	return &runnerResult
}

func runCommand(command string, input string) string {
	if length := len(input); length > 0 && input[length - 1] != '\n' {
		input += "\n"
	}

	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	stdin, _ := cmd.StdinPipe()
	defer stdin.Close()

	stdin.Write([]byte(input))

	output, _ := cmd.Output()
	return string(output)
}

func runTestCase(command string, input string, expected string) (bool, string) {
	log.Printf("runTestCase called %s %s", input, expected)
	output := runCommand(command, input)
	if output == expected || output == expected + "\n" {
		return true, output
	} else {
		return false, output
	}
}

func setUpSourceFile(code string, filename string) {
	fo, err := os.Create(filename)
	defer fo.Close()
	if err != nil {
		panic(err)
	}

	_, err = fo.WriteString(code)
	if err != nil {
		panic(err)
	}
}

func flatAndGroup(xs []string) [][2]string {
	testCases := make([][2]string, len(xs)/2)
	for i := 0; i < len(xs); i += 2 {
		testCases[i/2] = [2]string {xs[i], xs[i+1]}
	}
	return testCases
}
