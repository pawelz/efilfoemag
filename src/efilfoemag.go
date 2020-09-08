// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	inputCap = 10240
)

var (
	inputFileName = flag.String("input", "", fmt.Sprintf("Path to the input .elif file. Must be smaller than %dB.", inputCap))
	// outputDir = flag.String("output", "", "Path to the output directory. Must not exist.")
)

func main() {
	flag.Parse()
	if a := flag.Args(); len(flag.Args()) != 0 {
		log.Fatalf("Invalid non-flag arguments %v.\n", a)
	}

	if *inputFileName != "" {
		log.Fatalf("Missing mandatory flag --input.")
	}

	inputFile, err := os.Open(*inputFileName)
	if err != nil {
		log.Fatalf("Failed to open the input file %q: %v.", *inputFileName, err)
	}
	defer inputFile.Close()

	inputData := make([]byte, 10240)
	bytesRead, err := inputFile.Read(inputData)
	if err != nil {
		log.Fatalf("Failed to read the input file %q: %v.", *inputFileName, err)
	}
	if bytesRead == inputCap {
		log.Fatalf("The input file %q is too large. Must be smaller than %dB.", inputCap)
	}

	//c = grid.Parse(inputData)
}
