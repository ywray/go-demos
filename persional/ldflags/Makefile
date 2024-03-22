LDFLAGS = -X "main.buildStamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'`"
LDFLAGS += -X "main.testLog=testlogInMakefile"
LDFLAGS += -X "go-demos/persional/ldflags/test_import.TestVersion=testversion"
#LDFLAGS += "-X main.githash=`git describe --long --dirty --abbrev=14`"
build:
	go build -ldflags "${LDFLAGS}" -x -o build-version main.go
