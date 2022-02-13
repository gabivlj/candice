package build

const CandiceJSONDefault = `{
	"name": "project",
	"entrypoint": "main.cd",
	"cxx": "clang",
	"kind": "CXX",
	"flags": ["-m64"],
	"output": "program"
}
`

const CandiceProgramDefault = `func main() {
	@print("Hello world!");
}

`
