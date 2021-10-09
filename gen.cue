package hof

import (
	"github.com/hofstadter-io/hofmod-cli/gen"
	"github.com/hofstadter-io/hofmod-cli/schema"
)

Cli: _ @gen(cli,st)
Cli: gen.#HofGenerator & {
  Outdir: "./"
	Cli: #CLI
}

#CLI: schema.#Cli & {
	Name:    "cuest"
	Package: "github.com/hofstadter-io/cuetils/cmd/cuest"

	Usage:      "cuest"
	Short:      "CUE Structural - compare and manipulate nested data, Yaml, and JSON"
	Long:       Short

	OmitRun: true

	Commands: [
		#DepthCommand,
		#DiffCommand,
		#PatchCommand,
		#PickCommand,
		#MaskCommand,
	]

	//
	// Addons
	//
	Releases: {
		Disable: false
		Draft:    false
		Author:   "Hofstadter, Inc"
		Homepage: "https://docs.hofstadter.io"

		GitHub: {
			Owner: "hofstadter-io"
			Repo:  "cuetils"
		}

		Docker: {
			Maintainer: "Hofstadter, Inc <open-source@hofstadter.io>"
			Repo:       "hofstadter"
		}
	}
	Updates:  true
	EnablePProf: true
}

#DepthCommand: schema.#Command & {
	Name:  "depth"
	Usage: "depth [globs...]"
	Short: "calculate the depth of a file or glob"
	Long:  Short

	Args: [{
		Name:     "globs"
		Type:     "[]string"
		Help:     "file glob to the operation"
		Rest:			true
	}]
}

#DiffCommand: schema.#Command & {
	Name:  "diff"
	Usage: "diff <orig> <glob>"
	Short: "calculate the diff from the original to the glob file(s)"
	Long:  Short

	Args: [{
		Name:     "orig"
		Type:     "string"
		Required: true
		Help:     "original file to the operation"
	}, {
		Name:     "glob"
		Type:     "string"
		Help:     "file glob to the operation"
	}]
}

#PatchCommand: schema.#Command & {
	Name:  "patch"
	Usage: "patch <patch> <glob>"
	Short: "apply the pacth to the glob file(s)"
	Long:  Short

	Args: [{
		Name:     "patch"
		Type:     "string"
		Required: true
		Help:     "the patch to apply"
	}, {
		Name:     "glob"
		Type:     "string"
		Help:     "file glob to apply the patch to"
	}]
}

#PickCommand: schema.#Command & {
	Name:  "pick"
	Usage: "pick <pick> <glob>"
	Short: "pick the original from the glob file(s)"
	Long:  Short

	Args: [{
		Name:     "pick"
		Type:     "string"
		Required: true
		Help:     "the pick to use"
	}, {
		Name:     "glob"
		Type:     "string"
		Help:     "file glob to apply the pick to"
	}]
}

#MaskCommand: schema.#Command & {
	Name:  "mask"
	Usage: "mask <mask> <glob>"
	Short: "mask the original from the glob file(s)"
	Long:  Short

	Args: [{
		Name:     "mask"
		Type:     "string"
		Required: true
		Help:     "the mask to apply"
	}, {
		Name:     "glob"
		Type:     "string"
		Help:     "file glob to apply the mask to"
	}]
}

