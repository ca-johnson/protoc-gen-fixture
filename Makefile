build:
	go build .

# For allowing others to easily add --fixtures_out to their protoc cmd
# Add to our bazel packages?

# install:
# 	export PATH=$PATH:.

gen: build
	(protoc  \
	--plugin="protoc-gen-fixture" \
	--proto_path=test/. \
	--fixture_out=filename=test/proto.fixture:. \
	test/test.proto)	