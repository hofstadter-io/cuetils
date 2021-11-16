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
		#ExtendCommand,
		#TransformCommand,
		#ValidateCommand,
		#PipelineCommand,
	]

	Pflags: [{
		Name:    "maxiter"
		Long:    "maxiter"
		Short:   "m"
		Type:    "int"
		Default: ""
		Help:    "maximum iterations for recursion"
	},{
		Name:    "concrete"
		Long:    "concrete"
		Short:   "c"
		Type:    "bool"
		Default: "false"
		Help:    "enforce concrete outputs"
	},{
		Name:    "definitions"
		Long:    "definitions"
		Short:   "D"
		Type:    "bool"
		Default: "true"
		Help:    "process definitions in inputs and objects"
	},{
		Name:    "hidden"
		Long:    "hidden"
		Short:   "H"
		Type:    "bool"
		Default: "true"
		Help:    "process hidden fields in inputs and objects"
	},{
		Name:    "optional"
		Long:    "optional"
		Short:   "O"
		Type:    "bool"
		Default: "true"
		Help:    "process optional fields in inputs and objects"
	},{
		Name:    "load"
		Long:    "load"
		Type:    "bool"
		Default: "false"
		Help:    "use cue/load to support entrypoint and imports for args"
	},{
		Name:    "headers"
		Long:    "headers"
		Type:    "bool"
		Default: "false"
		Help:    "print the filename as a header during output"
	},{
		Name:    "accum"
		Long:    "accum"
		Type:    "string"
		Default: ""
		Help:    "accumulate operand results into a single value using accum as the label"
	},{
		Name:    "clean"
		Long:    "clean"
		Type:    "bool"
		Default: "false"
		Help:    "trim and unquote output, useful for basic lit output"
	},{
		Name:    "out"
		Long:    "out"
		Type:    "string"
		Default: "\"cue\""
		Help:    "output encoding [cue,yaml,json]"
	},{
		Name:    "outname"
		Long:    "outname"
		Type:    "string"
		Default: ""
		Help:    "output filename when being used"
	},{
		Name:    "overwrite"
		Long:    "overwrite"
		Type:    "bool"
		Default: "false"
		Help:    "overwrite files being processed"
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
	Usage: "count [files...]"
	Short: "count nodes in file(s)"
	Long:  Short

	Args: [{
		Name:     "globs"
		Type:     "[]string"
		Help:     "file globs to the operation"
		Rest:			true
	}]
}

#DepthCommand: schema.#Command & {
	Name:  "depth"
	Usage: "depth [files...]"
	Short: "calculate the depth of file(s)"
	Long:  Short

	Args: [{
		Name:     "globs"
		Type:     "[]string"
		Help:     "file globs to the operation"
		Rest:			true
	}]
}

#DiffCommand: schema.#Command & {
	Name:  "diff"
	Usage: "diff <orig> <next>"
	Short: "calculate the diff from orig to next file(s)"
	Long:  Short

	Args: [{
		Name:     "orig"
		Type:     "string"
		Help:     "orig file(s) to the operation"
		Required: true
	}, {
		Name:     "next"
		Type:     "string"
		Help:     "next file(s) to the operation"
		Required: true
	}]
}

#PatchCommand: schema.#Command & {
	Name:  "patch"
	Usage: "patch <patch> <orig>"
	Short: "apply pacth to orig file(s)"
	Long:  Short

	Args: [{
		Name:     "patch"
		Type:     "string"
		Help:     "the patch glob to apply"
		Required: true
	}, {
		Name:     "orig"
		Type:     "string"
		Help:     "file glob to the operation"
		Required: true
	}]
}

#PickCommand: schema.#Command & {
	Name:  "pick"
	Usage: "pick <code> [files...]"
	Short: "pick from file(s) with code"
	Long:  Short

	Args: [{
		Name:     "code"
		Type:     "string"
		Required: true
		Help:     "code for the operation"
	}, {
		Name:     "globs"
		Type:     "[]string"
		Help:     "file globs to the operation"
		Rest:			true
	}]
}

#MaskCommand: schema.#Command & {
	Name:  "mask"
	Usage: "mask <code> [files...]"
	Short: "mask from file(s) with code"
	Long:  Short

	Args: [{
		Name:     "code"
		Type:     "string"
		Required: true
		Help:     "code for the operation"
	}, {
		Name:     "globs"
		Type:     "[]string"
		Help:     "file globs to the operation"
		Rest:			true
	}]
}

#ReplaceCommand: schema.#Command & {
	Name:  "replace"
	Usage: "replace <code> [files...]"
	Short: "replace in file(s) with code (only if present)"
	Long:  Short

	Args: [{
		Name:     "code"
		Type:     "string"
		Required: true
		Help:     "code for the operation"
	}, {
		Name:     "globs"
		Type:     "[]string"
		Help:     "file globs to the operation"
		Rest:			true
	}]
}

#ExtendCommand: schema.#Command & {
	Name:  "extend"
	Usage: "upsert <code> [files...]"
	Short: "extend file(s) with code (only if not present)"
	Long:  Short

	Args: [{
		Name:     "code"
		Type:     "string"
		Required: true
		Help:     "code for the operation"
	}, {
		Name:     "globs"
		Type:     "[]string"
		Help:     "file globs to the operation"
		Rest:			true
	}]
}

#UpsertCommand: schema.#Command & {
	Name:  "upsert"
	Usage: "upsert <code> [files...]"
	Short: "upsert file(s) with code (extend and replace)"
	Long:  Short

	Args: [{
		Name:     "code"
		Type:     "string"
		Required: true
		Help:     "code for the operation"
	}, {
		Name:     "globs"
		Type:     "[]string"
		Help:     "file globs to the operation"
		Rest:			true
	}]
}

#TransformCommand: schema.#Command & {
	Name:  "transform"
	Usage: "transform <code> [files...]"
	Short: "transform file(s) with code"
	Long:  Short

	Args: [{
		Name:     "code"
		Type:     "string"
		Required: true
		Help:     "code for the operation"
	}, {
		Name:     "globs"
		Type:     "[]string"
		Help:     "file globs to the operation"
		Rest:			true
	}]
}

#ValidateCommand: schema.#Command & {
	Name:  "validate"
	Usage: "validate <schema> [files...]"
	Short: "validate file(s) with schema"
	Long:  Short

	Args: [{
		Name:     "schema"
		Type:     "string"
		Required: true
		Help:     "schema to validate with"
	}, {
		Name:     "globs"
		Type:     "[]string"
		Help:     "file globs to the operation"
		Rest:			true
	}]
}

#PipelineCommand: schema.#Command & {
	Name:  "pipeline"
	Usage: "pipeline <code> [files...]"
	Short: "run file(s) through a pipeline of operations"
	Long:  Short

	Args: [{
		Name:     "code"
		Type:     "string"
		Required: true
		Help:     "code for the operation"
	}, {
		Name:     "globs"
		Type:     "[]string"
		Help:     "file globs to the operation"
		Rest:			true
	}]
}

#CountHelp: ""
#DepthHelp: ""
#DiffHelp: ""
#PatchHelp: ""
#PickHelp: ""
#MaskHelp: ""
#ReplaceHelp: ""
#UpsertHelp: ""
#ExtendHelp: ""
#TransformHelp: ""
#ValidateHelp: ""
#PipelineHelp: ""
