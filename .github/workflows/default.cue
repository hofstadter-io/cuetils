package workflows

import "github.com/hofstadter-io/cuetils/.github/workflows/common"

common.#Workflow & {
	name: "default"
	on: ["push"]
	jobs: test: {
		steps: [ for step in common.#BuildSteps {step} ] + [{
			name: "Go tests"
			run: """
			go test -cover ./structural
			"""
		},{
			name: "CLI tests"
			run: """
			go test -cover ./test/cli
			"""
		},{
			name: "Cue tests"
			run: """
			for file in `ls test/cue/*.cue`; do
				cue eval $file > /dev/null
			done
			"""
		}]
	}
}
