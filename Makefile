build:
	go build .

gen: build
	(protoc  \
	--plugin="protoc-gen-fixture" \
	--proto_path=test/. \
	--fixture_out=filename=test/proto.fixture:. \
	test/test.proto)	