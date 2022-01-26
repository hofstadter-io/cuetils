package flags

type PipelineFlagpole struct {
	List     bool
	Docs     bool
	Pipeline []string
	Tags     []string
}

var PipelineFlags PipelineFlagpole
