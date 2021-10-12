package hof

import "strings"

//
////// Defined (partially) test configuration
//

#GoBaseTest: {
	skip: bool | *false

	sysenv: bool | *false
	env?: [string]: string
	args?: [...string]
	verbose?: bool | int

	dir: string
	...
}

#GoBashTest: #GoBaseTest & {
	dir: string
	script: string | *"""
	rm -rf .workdir
	go test -cover ./
	"""
	...
}

#GoBashCover: #GoBaseTest & {
	dir: string
	back: strings.Repeat("../", strings.Count(dir, "/") + 1)
	script: string | *"""
	rm -rf .workdir
	go test -cover ./ -coverprofile cover.out -json > tests.json
	"""
	...
}

//
////// Actual test configuration
//


cli: _ @test(suite,cli)
cli: {

	test_all: #GoBashTest @test(bash,test)
	test_all: {
		dir: "test/cli"
	}
	cover_all: #GoBashCover @test(bash,cover)
	cover_all: {
		dir: "test/cli"
	}

}
