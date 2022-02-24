package build

import (
	"encoding/json"
	"io"
	"os"
)

type CompileKind string
type BinaryKind string

const (
	PureLLVM CompileKind = "llvm"
	CXX      CompileKind = "cxx"

	Object BinaryKind = "obj"
	Binary BinaryKind = "exe"
)

type ProjectConfiguration struct {
	Name          string      `json:"name"`
	EntryPoint    string      `json:"entrypoint"`
	CXX           string      `json:"cxx"`
	CompileKind   CompileKind `json:"kind"`
	Output        string      `json:"output"`
	CompilerFlags []string    `json:"flags"`
	BinaryKind    BinaryKind  `json:"binary"`
}

func ParseConfiguration(reader io.Reader) (ProjectConfiguration, error) {
	var configuration ProjectConfiguration
	err := json.NewDecoder(reader).Decode(&configuration)
	return configuration, err
}

func ParseConfigurationFile(filePath string) (ProjectConfiguration, error) {
	fd, err := os.Open(filePath)
	if err != nil {
		return ProjectConfiguration{}, err
	}

	return ParseConfiguration(fd)
}
