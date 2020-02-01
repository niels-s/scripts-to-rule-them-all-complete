package main

import (
	"github.com/posener/complete/v2"
	"github.com/posener/complete/v2/predict"
)

func main() {
	environments := predict.Set{"development", "staging", "production", "approval", "sandbox"}

	deploy := &complete.Command{
		Args: environments,
	}
	deploy.Complete("script/deploy")
}
