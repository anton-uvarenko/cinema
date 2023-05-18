generate() {
  mkdir ../../$1/protobufs
  mkdir ../../$1/protobufs/auth
  protoc --go_opt=paths=source_relative --go_out=../../$1/protobufs/auth \
          --go-grpc_out=../../$1/protobufs/auth --go-grpc_opt=paths=source_relative \
          --proto_path=../general \
          --proto_path=. \
          ../general/*.proto \
           *.proto
}

service=$1

if [[ -z "$service" ]]; then
  echo "No service specified"
  exit 1
fi

if [[ $service == "all" ]]; then
  for s in "broker" "authorization-service" "user-service"
  do
    generate ${s};
  done
else
  generate $service;
fi

