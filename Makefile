build:
	go build .

# install:
# 	export PATH=$PATH:.

gen: build
	(protoc  \
	--plugin="protoc-gen-fixtures" \
	--proto_path=. \
	--fixtures_out=test/. \
	test.proto)	