# First command incase users forget to supply one
# cat self as help for simplicity
help:
	cat Makefile
.PHONY: help

gen:
	hof gen

test:
	hof test -t test

WORKFLOWS = default

.PHONY: workflow
workflows = $(addprefix workflow_, $(WORKFLOWS))
workflow: $(workflows)
$(workflows): workflow_%:
	@cue export --out yaml .github/workflows/$(subst workflow_,,$@).cue > .github/workflows/$(subst workflow_,,$@).yml

