generate() {
  mkdir ../../$1/protobufs
  mkdir ../../$1/protobufs/general
  protoc --go_opt=paths=source_relative --go_out=../../$1/protobufs/general \
          --go-grpc_out=../../$1/protobufs/general --go-grpc_opt=paths=source_relative \
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
    echo $s;
    generate ${s};
  done
else
  generate $service;
fi

