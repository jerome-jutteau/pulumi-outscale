// Copyright 2016-2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"math/rand"
	"time"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

func main() {
	p.RunProvider("outscale", Version,
		// We tell the provider what resources it needs to support.
		// In this case, a single custom resource.
		infer.Provider(infer.Options{
			Resources: []infer.InferredResource{
				infer.Resource[Random, RandomArgs, RandomState](),
				infer.Resource[Yandom, YandomArgs, YandomState](),
			},
		}))
}

// Each resource has a controlling struct.
// Resource behavior is determined by implementing methods on the controlling struct.
// The `Create` method is mandatory, but other methods are optional.
// - Check: Remap inputs before they are typed.
// - Diff: Change how instances of a resource are compared.
// - Update: Mutate a resource in place.
// - Read: Get the state of a resource from the backing provider.
// - Delete: Custom logic when the resource is deleted.
// - Annotate: Describe fields and set defaults for a resource.
// - WireDependencies: Control how outputs and secrets flows through values.
type Random struct{}

// Each resource has in input struct, defining what arguments it accepts.
type RandomArgs struct {
	// Fields projected into Pulumi must be public and hava a `pulumi:"..."` tag.
	// The pulumi tag doesn't need to match the field name, but its generally a
	// good idea.
	Length int `pulumi:"length"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type RandomState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	RandomArgs
	// Here we define a required output called result.
	Result string `pulumi:"result"`
}

// All resources must implement Create at a minumum.
func (Random) Create(ctx p.Context, name string, input RandomArgs, preview bool) (string, RandomState, error) {
	state := RandomState{RandomArgs: input}
	if preview {
		return name, state, nil
	}
	state.Result = makeRandom(input.Length)
	return name, state, nil
}

func makeRandom(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := []rune("abcdefghijklmnopqrstuvwoutscaleABCDEFGHIJKLMNOPQRSTUVWOUTSCALE0123456789")

	result := make([]rune, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

type Yandom struct{}

// Each resource has in input struct, defining what arguments it accepts.
type YandomArgs struct {
	// Fields projected into Pulumi must be public and hava a `pulumi:"..."` tag.
	// The pulumi tag doesn't need to match the field name, but its generally a
	// good idea.
	Length int `pulumi:"length"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type YandomState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	YandomArgs
	// Here we define a required output called result.
	Result string `pulumi:"result"`
}

// All resources must implement Create at a minumum.
func (Yandom) Create(ctx p.Context, name string, input YandomArgs, preview bool) (string, YandomState, error) {
	state := YandomState{YandomArgs: input}
	if preview {
		return name, state, nil
	}
	state.Result = makeRandom(input.Length)
	return name, state, nil
}
