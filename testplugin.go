package main

import (
	configurator "github.com/OpenCHAMI/configurator/pkg"
	"github.com/OpenCHAMI/configurator/pkg/generator"
	"github.com/OpenCHAMI/configurator/pkg/util"
)

type TestGenerator struct{}

func (g *TestGenerator) GetName() string        { return "test" }
func (g *TestGenerator) GetVersion() string     { return "v1.0.0" }
func (g *TestGenerator) GetDescription() string { return "This is a plugin creating for running tests." }
func (g *TestGenerator) Generate(config *configurator.Config, opts ...util.Option) (generator.FileMap, error) {
	return generator.FileMap{"test": []byte("test")}, nil
}
var Generator TestGenerator
