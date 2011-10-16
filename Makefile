include $(GOROOT)/src/Make.inc

TARG=DocJava
GOFILES=\
	DocJava.go\
	field.go\
	types.go\
	constructor.go\
	argument.go\
	class.go\
	mask.go\
	method.go\
	javadoc.go\
	interface.go

include $(GOROOT)/src/Make.cmd
