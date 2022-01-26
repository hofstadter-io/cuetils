package hof

import (
	"github.com/hofstadter-io/hofmod-cli/gen"
	"github.com/hofstadter-io/hofmod-cli/schema"
)

Cli: gen.#HofGenerator & {
	@gen(cli,st)
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
		#InsertCommand,
		#UpsertCommand,
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
		Name:    "headers"
		Long:    "headers"
		Type:    "bool"
		Default: "false"
		Help:    "print the filename as a header during output"
	},{
		Name:    "accum"
		Long:    "accum"
		Short:   "a"
		Type:    "string"
		Default: ""
		Help:    "accumulate operand results into a single value using accum as the label"
	},{
		Name:    "clean"
		Long:    "clean"
		Short:   "r"
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
		Short:   "o"
		Type:    "string"
		Default: ""
		Help:    "output filename when being used"
	},{
		Name:    "overwrite"
		Long:    "overwrite"
		Short:   "F"
		Type:    "bool"
		Default: "false"
		Help:    "overwrite files being processed"
	},{
		Name:    "allTypeErrors"
		Long:    "type-errors"
		Short:   "E"
		Type:    "bool"
		Default: "false"
		Help:    "error when nodes or leafs have type mismatches"
	},{
		Name:    "nodeTypeErrors"
		Long:    "node-type-errors"
		Short:   "N"
		Type:    "bool"
		Default: "false"
		Help:    "error when nodes have type mismatches"
	},{
		Name:    "leafTypeErrors"
		Long:    "leaf-type-errors"
		Short:   "L"
		Type:    "bool"
		Default: "false"
		Help:    "error when leafs have type mismatches"
	},{
		Name:    "verbose"
		Long:    "verbose"
		Short:   "v"
		Type:    "bool"
		Default: "false"
		Help:    "verbose printing for some commands"
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
	Aliases: ["D"]
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
	Aliases: ["P"]
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
	Aliases: ["p"]
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
	Aliases: ["m"]
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
	Aliases: ["r"]
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

#InsertCommand: schema.#Command & {
	Name:  "insert"
	Aliases: ["i"]
	Usage: "insert <code> [files...]"
	Short: "insert into file(s) with code (only if not present)"
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
	Aliases: ["u"]
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
	Aliases: ["t"]
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
	Aliases: ["v"]
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
	Aliases: ["pipe", "dag"]
	Usage: "pipeline <code> [files...]"
	Short: "run file(s) through a pipeline of operations"
	Long:  Short

	Args: [{
		Name:     "globs"
		Type:     "[]string"
		Help:     "file globs to the operation"
		Rest:			true
	}]

	Flags: [{
		Name:    "list"
		Long:    "list"
		Short:   "l"
		Type:    "bool"
    Default: "false"
		Help:    "list available pipelines"
	},{
		Name:    "docs"
		Long:    "docs"
		Short:   "d"
		Type:    "bool"
    Default: "false"
		Help:    "print pipeline docs"
	},{
		Name:    "pipeline"
		Long:    "pipeline"
		Short:   "p"
		Type:    "[]string"
    Default: "nil"
		Help:    "pipeline labels to match and run"
	},{
		Name:    "tags"
		Long:    "tags"
		Short:   "t"
		Type:    "[]string"
    Default: "nil"
		Help:    "data tags to inject before run"
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
