   
#!/usr/bin/bash

# Proto file name
PROTO_FILE_NAME="todo.proto"

# Output folder path for generated files
OUTPUT_PATH="./src/proto/todo"

# Folder containing *.proto files
PROTO_PATH="./proto"

# Path to this plugin, Note this must be an absolute path on Windows (see #15)
PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"

# Path to the grpc_node_plugin
PROTOC_GEN_GRPC_PATH="./node_modules/.bin/grpc_tools_node_protoc_plugin"

# Generate JS code
yarn run grpc_tools_node_protoc \
  --plugin="protoc-gen-grpc=${PROTOC_GEN_GRPC_PATH}" \
  --js_out="import_style=commonjs,binary:$OUTPUT_PATH" \
  --grpc_out=$OUTPUT_PATH \
  --proto_path=$PROTO_PATH \
  ${PROTO_FILE_NAME}

# Generate TS code
yarn run grpc_tools_node_protoc \
  --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
  --ts_out=$OUTPUT_PATH \
  --proto_path=$PROTO_PATH \
  ${PROTO_FILE_NAME}
