package main

import (
	"os"

	"hydectl/cmd"
	"hydectl/internal/logger"
	"hydectl/internal/plugin"

	"github.com/adrg/xdg"
)

func main() {
	logger.SetupLogging()

	cmd.ScriptPaths = []string{
		xdg.ConfigHome + "/lib/hydectl/scripts",
		// os.Getenv("HOME") + "/.local/lib/hyde",
		os.Getenv("HOME") + "/.local/lib/hydectl/scripts",
		"/usr/local/lib/hydectl/scripts",
		"/usr/lib/hydectl/scripts",
	}

	if len(os.Args) > 1 {
		scriptName := os.Args[1]
		scripts, err := plugin.FindAllScripts(cmd.ScriptPaths)
		if err != nil {
			logger.Errorf("Error finding scripts: %v", err)
			cmd.Execute()
			return
		}

		if scriptPath, ok := scripts[scriptName]; ok {
			logger.Debugf("Executing script: %s", scriptPath)
			plugin.ExecuteScript(scriptPath, os.Args[2:])
			return
		}
	}

	cmd.Execute()
}
