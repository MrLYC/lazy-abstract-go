BUILD = go build -o greeting -tags

.PHONY: first
first:
	${BUILD} first

.PHONY: second
second:
	${BUILD} second

.PHONY: third
third:
	${BUILD} third

.PHONY: fourth
fourth:
	${BUILD} fourth
	
.PHONY: fifth
fifth:
	${BUILD} fifth