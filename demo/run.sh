CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"
# go get gopkg.in/DATA-DOG/go-sqlmock.v1
# go get github.com/stretchr/testify/assert
# go get golang.org/x/sys/cpu
# go get -u github.com/labstack/echo/...
gofmt -w src/
go test -v -cover -coverprofile=coverage.out  ./...
go tool cover -html=coverage.out
go install main
export GOPATH="$OLDGOPATH"
echo 'finished'
