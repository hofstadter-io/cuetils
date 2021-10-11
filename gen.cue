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
	Name:    "cuetils"
	Package: "github.com/hofstadter-io/cuetils/cmd/cuetils"

	Usage:      "cuetils"
	Short:      "CUE Utilites for bulk ETL, diff, query, and other operations on data and config"
	Long:       Short

	OmitRun: true

	Commands: [
		#CountCommand,
		#DepthCommand,
		#DiffCommand,
		#PatchCommand,
		#PickCommand,
		#MaskCommand,
		#ReplaceCommand,
		#UpsertCommand,
		#TransformCommand,
		#ValidateCommand,
	]

	Pflags: [{
		Name:    "maxiter"
		Long:    "maxiter"
		Short:   "m"
		Type:    "int"
		Default: ""
		Help:    "maximum iterations for recursion"
	},{
		Name:    "out"
		Long:    "out"
		Type:    "string"
		Default: "\"cue\""
		Help:    "output encoding [cue,yaml,json]"
	}]

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

#CountCommand: schema.#Command & {
	Name:  "count"
	Usage: "count [globs...]"
	Short: "calculate the node count of a file or glob"
	Long:  Short

	Args: [{
		Name:     "globs"
		Type:     "[]string"
		Help:     "file glob to the operation"
		Rest:			true
	}]
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
		Name:     "globs"
		Type:     "[]string"
		Help:     "file glob to the operation"
		Rest:			true
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
		Name:     "globs"
		Type:     "[]string"
		Help:     "file glob to the operation"
		Rest:			true
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
		Name:     "globs"
		Type:     "[]string"
		Help:     "file glob to the operation"
		Rest:			true
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
		Name:     "globs"
		Type:     "[]string"
		Help:     "file glob to the operation"
		Rest:			true
	}]
}

#ReplaceCommand: schema.#Command & {
	Name:  "replace"
	Usage: "replace <orig> <glob>"
	Short: "apply the replace from the original to the glob file(s)"
	Long:  Short

	Args: [{
		Name:     "orig"
		Type:     "string"
		Required: true
		Help:     "original file to the operation"
	}, {
		Name:     "globs"
		Type:     "[]string"
		Help:     "file glob to the operation"
		Rest:			true
	}]
}

#UpsertCommand: schema.#Command & {
	Name:  "upsert"
	Usage: "upsert <orig> <glob>"
	Short: "apply the upsert from the original to the glob file(s)"
	Long:  Short

	Args: [{
		Name:     "orig"
		Type:     "string"
		Required: true
		Help:     "original file to the operation"
	}, {
		Name:     "globs"
		Type:     "[]string"
		Help:     "file glob to the operation"
		Rest:			true
	}]
}

#TransformCommand: schema.#Command & {
	Name:  "transform"
	Usage: "transform <orig> <glob>"
	Short: "apply the transform from the original to the glob file(s)"
	Long:  Short

	Args: [{
		Name:     "orig"
		Type:     "string"
		Required: true
		Help:     "original file to the operation"
	}, {
		Name:     "globs"
		Type:     "[]string"
		Help:     "file glob to the operation"
		Rest:			true
	}]
}

#ValidateCommand: schema.#Command & {
	Name:  "validate"
	Usage: "validate <orig> <glob>"
	Short: "validate with the original against the glob file(s)"
	Long:  Short

	Args: [{
		Name:     "orig"
		Type:     "string"
		Required: true
		Help:     "original file to the operation"
	}, {
		Name:     "globs"
		Type:     "[]string"
		Help:     "file glob to the operation"
		Rest:			true
	}]
}

