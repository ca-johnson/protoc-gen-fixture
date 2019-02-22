// Copyright 2019 Carl Johnson. All Rights Reserved.
// See LICENSE for licensing terms.

// Simply dumps a fixture of stdin alongside the first protofile in the file descriptor
// This file then used to debug protobuf plugins, simulating the same protoc context

package main

import (
	"io/ioutil"
	"os"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

func main() {
	gen := generator.New()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		gen.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, gen.Request); err != nil {
		gen.Error(err, "parsing input proto")
	}

	if len(gen.Response.File) == 0 {
		gen.Error(nil, "no input files")
	}

	fixtureName := "helloworld"
	// fixtureName := strings.Replace(*gen.Response.File[0].Name, ".pb.go", ".fixture", -1)
	// Write out fixture for debugging/unit tests
	f, err := os.Create(fixtureName)
	if err != nil {
		gen.Error(err, "os.Create")
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		gen.Error(err, "f.Write")
	}
}
