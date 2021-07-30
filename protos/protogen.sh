protoc \
  --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  --js_out=import_style=commonjs,binary:../client/_proto \
  --ts_out=service=grpc-web:../client/_proto \
  orderWatcher.proto

