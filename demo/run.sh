CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"
# go get gopkg.in/DATA-DOG/go-sqlmock.v1
# go get github.com/stretchr/testify/assert
gofmt -w src/
go test -v --cover ./...
go install main
export GOPATH="$OLDGOPATH"
echo 'finished'
