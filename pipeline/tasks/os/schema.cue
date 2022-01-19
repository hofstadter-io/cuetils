package os

Exec: {
  @task(os.Exec)

	cmd: string | [string, ...string]

	// dir specifies the working directory of the command.
	// The default is the current working directory.
	dir?: string

	// env defines the environment variables to use for this system.
	// If the value is a list, the entries mus be of the form key=value,
	// where the last value takes precendence in the case of multiple
	// occurrances of the same key.
	env: [string]: string | [...=~"="]

	// stdout captures the output from stdout if it is of type bytes or string.
	// The default value of null indicates it is redirected to the stdout of the
	// current process.
	stdout: *null | string | bytes

	// stderr is like stdout, but for errors.
	stderr: *null | string | bytes

	// stdin specifies the input for the process. If stdin is null, the stdin
	// of the current process is redirected to this command (the default).
	// If it is of typ bytes or string, that input will be used instead.
	stdin: *null | string | bytes

	// success is set to true when the process terminates with with a zero exit
	// code or false otherwise. The user can explicitly specify the value
	// force a fatal error if the desired success code is not reached.
	success: bool

  // the exit code of the command
  exitcode: int
}

ReadFile: {
  @task(os.ReadFile)
  filename: string
  f: filename

  // filled by cuetils
  contents: string | bytes
}

WriteFile: {
  @task(os.WriteFile)
  filename: string
  f: filename

  // filled by cuetils
  contents: string | bytes
  mode: int | *0o666
}

// A Value are all possible values allowed in flags.
// A null value unsets an environment variable.
Value: bool | number | *string | null

// Name indicates a valid flag name.
Name: !="" & !~"^[$]"

// Setenv defines a set of command line flags, the values of which will be set
// at run time. The doc comment of the flag is presented to the user in help.
//
// To define a shorthand, define the shorthand as a new flag referring to
// the flag of which it is a shorthand.
Setenv: {
	@task(os.Setenv)

	{[Name]: Value}
}

// Getenv gets and parses the specific command line variables.
Getenv: {
	@task(os.Getenv)

	{[Name]: Value}
}

// Environ populates a struct with all environment variables.
Environ: {
	@task(os.Environ)

	// A map of all populated values.
	// Individual entries may be specified ahead of time to enable
	// validation and parsing. Values that are marked as required
	// will fail the task if they are not found.
	{[Name]: Value}
}

// Clearenv clears environment variables.
Clearenv: {
	@task(os.Clearenv)

  // Var Names to clear, if empty clears all
	vars: [...Name]
}
