package cuetils

import "embed"

//go:embed recurse/*.cue structural/*.cue
var CueEmbeds embed.FS
