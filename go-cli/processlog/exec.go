package main

import (
	"io"
	"path/filepath"
	"os"
)

func execution(commandName string, args [] string, stdout, stderr io.Writer) {
	if logDir == "" {
		stdout = os.Stdout
		stderr = os.Stderr
	} else {
		ts := time.Now().Unix()

		stdoutFileName := fmt.Sprintf("%s-%v-stdout.log", commandName, ts)
		stdoutFile, err := os.Create(filepath.Join(logDir, stdoutFileName))
		if err != nil {
			return nil, nil, err
		}
		stdout = io.MultiWriter(os.Stdout, stdoutFile)

		stderrFileName := fmt.Sprintf("%s-%v-stderr.log", commandName, ts)
		stderrFile, err := os.Create(filepath.Join(logDir, stderrFileName))

		if err != nil {
			return nil, nil, err
		}
		stderr = io.MultiWriter(os.Stderr, stderrFile)
	}
	return
}