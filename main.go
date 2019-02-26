// Copyright 2019 Carl Johnson. All Rights Reserved.

package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

var fFixture = flag.String("fixture", "", "read from this file instead of stdin")

// Optionally output to custom filename, specified by filename param to this plugin
// i.e. fixture_out=filename=myfixture:.
func getFixtureOutName(genParams string) string {
	for _, parameter := range strings.Split(genParams, ",") {
		kvp := strings.SplitN(parameter, "=", 2)
		if len(kvp) != 2 || kvp[0] != "filename" {
			continue
		}
		return kvp[1]
	}

	return "proto.fixture"
}

func main() {
	flag.Parse()
	gen := generator.New()

	var data []byte
	var err error
	if *fFixture != "" {
		data, err = ioutil.ReadFile(*fFixture)
	} else {
		data, err = ioutil.ReadAll(os.Stdin)
	}
	if err != nil {
		panic(err)
	}
	if err := proto.Unmarshal(data, gen.Request); err != nil {
		panic(err)
	}

	if len(gen.Request.FileToGenerate) == 0 {
		panic("no files to generate")
	}

	f, err := os.Create(getFixtureOutName(gen.Request.GetParameter()))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}
}
