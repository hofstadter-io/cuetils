package api

#Request: {
	Method: *"GET" | "POST" | "PUT" | "DELETE" | "OPTIONS" | "HEAD" | "CONNECT" | "TRACE" | "PATCH"
	Host:   string
	Path:   string | *""
	Auth?:  string
	Headers?: [string]: string
	Query?: [string]:   string
	Data?:    string | {...}
	Timeout?: string
	Retry?: {
		Count?: int
		Timer?: string
		Codes: [...int]
	}
}

#Response: {
	Status?: int
	Headers?: [string]: string
	Body?: string
	Value: _
	// latency?: float
}
