package utils

import (
	"fmt"
	"github.com/boj-vs-code/code-runner/runner/models"
	"github.com/cosiner/argv"
	"log"
	"os"
	"os/exec"
	"bytes"
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

func Run(runnerRequest models.RunnerRequest) models.RunnerResult {
	problemId := runnerRequest.ProblemId
	code := runnerRequest.Code
	sourceFilename := GetSourceFilename()
	compile := GetCompileScript()
	script := GetExecutionScript()

	problem := models.GetProblemById(problemId)
	testCases := flatAndGroup(problem.Testcases)

	runnerResult := models.RunnerResult{}
	runnerResult.Code = code
	runnerResult.Failed = make([][3]string, 0)

	setUpSourceFile(code, sourceFilename)
	output, err := runCommand(compile)
	// In this context, output means output from stderr.
	if err != nil || output != "" {
		runnerResult.Message = fmt.Sprintf(
			"compile error: %s", output)
		return runnerResult
	}

	runnerResult.Success = true
	for i := 0; i < len(testCases); i++ {
		input, expected := testCases[i][0], testCases[i][1]
		success, actual := runTestCase(script, input, expected)
		runnerResult.Success = runnerResult.Success && success
		if !success {
			runnerResult.Failed = append(runnerResult.Failed, [3]string{input, expected, actual})
		}
	}

	if runnerResult.Success {
		runnerResult.Message = "success"
	} else {
		runnerResult.Message = "failed"
	}
	return runnerResult
}

func runCommand(command string) (string, error) {
	return runCommandWithInput(command, "")
}

func runCommandWithInput(command string, input string) (string, error) {
	log.Printf("Command \"%s\"", command)
	if length := len(input); length > 0 && input[length - 1] != '\n' {
		input += "\n"
	}

	argv, err := argv.Argv([]rune(command),	argv.ParseEnv(os.Environ()), argv.Run)
	if err != nil {
		return "", err
	}
	args := argv[0]  // it's one line
	cmd := exec.Command(args[0], args[1:]...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	stdin, err := cmd.StdinPipe()
	defer stdin.Close()
	if err != nil {
		return "", err
	}
	stdin.Write([]byte(input))

	if err = cmd.Run(); err != nil || stderr.Len() != 0 {
		return stderr.String(), err
	}
	return stdout.String(), nil
}

func runTestCase(command string, input string, expected string) (bool, string) {
	log.Printf("runTestCase called %s %s", input, expected)
	output, err := runCommandWithInput(command, input)
	if err != nil {
		log.Printf("runTestCase error %s", output)
	}
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
