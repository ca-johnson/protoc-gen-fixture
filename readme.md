# protoc-gen-fixture

A protobuf plugin which writes the input to a plugin out as a file.

You can then use this as a fixture for debugging plugins as a standalone executable.

# Usage

Build protoc-gen-fixture and put it in your path

Run against your proto file(s): `protoc --fixture_out=. my_protofile.proto`

By default, protoc-gen-fixture writes to a file in the current working directory called `proto.fixture`

You can override this with the `filename` parameter: `--fixture_out=filename=path/to/my/file.fixture`

In your own plugins, add a flag to allow reading from a file instead of stdin e.g. `--fixture=proto.fixture`. You can then easily run your plugin in the debugger exactly as it would be run by protoc.