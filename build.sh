BASE_PATH=/home/nur/go/src/test-grpc-gw

cd $BASE_PATH/cmd/gateway && go build
cd $BASE_PATH/cmd/hello-world && go build
cd $BASE_PATH/cmd/location && go build
