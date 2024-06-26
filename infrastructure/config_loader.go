package infrastructure

import (
	"fmt"
	"os"

	"github.com/Rindrics/execute-script-with-merge/application"
	"github.com/Rindrics/execute-script-with-merge/domain"
)

func LoadConfig() (*application.Config, error) {
	requiredLabel := os.Getenv(domain.EnvVarRequiredLabel)
	if requiredLabel == "" {
		return nil, fmt.Errorf("environment variable %s is required", domain.EnvVarRequiredLabel)
	}

	baseBranch := os.Getenv(domain.EnvVarBaseBranch)
	if baseBranch == "" {
		return nil, fmt.Errorf("environment variable %s is required", domain.EnvVarBaseBranch)
	}

	targetScriptListDir := os.Getenv(domain.EnvVarTargetScriptListDir)
	if targetScriptListDir == "" {
		return nil, fmt.Errorf("environment variable %s is required", domain.EnvVarTargetScriptListDir)
	}

	return &application.Config{
		RequiredLabel:       requiredLabel,
		BaseBranch:          baseBranch,
		TargetScriptListDir: targetScriptListDir,
	}, nil
}
