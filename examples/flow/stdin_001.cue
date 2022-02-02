import "strings"

tasks: {
  @flow()
  input: { msg: "enter text: " } @task(os.Stdin)
  final: { text: strings.ToUpper(input.contents) } @task(os.Stdout)
}
