import "strings"

tasks: {
  @pipeline(stdin)
  input: { #Msg: "enter text: " } @task(os/stdin)
  final: { #O: strings.ToUpper(input.Contents) } @task(os/stdout)
}
