GOCMD=go #"$(which go)"
GIT_ROOT="$(git rev-parse --show-toplevel)"

.PHONY: hooks
hooks:
	mv .git/hooks .git/hooks.bak
	ln -s ../.githooks .git/hooks
	git config core.hooksPath .githooks


.PHONY: dep-install
dep-install:
	$(GOCMD) get


.PHONY: init
init: hooks dep-install
