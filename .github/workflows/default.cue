package workflows

import "github.com/hofstadter-io/hof/.github/workflows/common"

common.#Workflow & {
	name: "default"
	on: ["push"]
	jobs: test: {
		steps: [ for step in common.#BuildSteps {step} ] + [{
			name: "Run tests"
			run: """
			hof test -t test
			"""
		}]
	}
}
