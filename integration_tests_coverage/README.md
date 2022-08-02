### build test binary
go test -c ./ -cover -covermode=count -coverpkg=./...

### invoke server run using test mode
./integration_tests_coverage.test -test.coverprofile my-service.out

### use http client to send request to server
curl http://localhost:8080

### collect coverage data from my-service.out
less my-service.out