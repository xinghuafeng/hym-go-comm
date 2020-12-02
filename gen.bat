cd app/service/go-micro/services/proto
protoc --micro_out=. --go_out=. *.proto
protoc-go-inject-tag -input=test.pb.go
cd .. && cd .. && cd .. && cd .. && cd ..
