BUILD = go build -o world -tags

.PHONY: first
first:
	${BUILD} first

.PHONY: second
second:
	${BUILD} second