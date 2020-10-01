## Kedro gRPC Go Client
--------------------------

1. ListPipelines
2. Run Trigger
3. Streaming Run Status Response

Generate gRPC code for client
----------------------------------------

We are going to use gRPC to generate libraries for Go.
To generate the Go code, you'll need to install  protoc_.

.. _protoc: https://github.com/google/protobuf/#protocol-compiler-installation

.. code-block:: bash

 # Go Client
 $ protoc -I protobuf/ --go_out=plugins=grpc:protobuf_kedro/ protobuf/kedro.proto

The latter will generate kedro.pb.go.

Run the client, simply run:

.. code-block:: bash

  go run client.go


Kedro gRPC Server installation:

`https://github.com/mmchougule/kedro-grpc-server`
