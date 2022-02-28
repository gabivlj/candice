package build

const CandiceJSONDefault = `{
	"name": "project",
	"entrypoint": "main.cd",
	"cxx": "clang",
	"kind": "cxx",
	"flags": ["-m64"],
	"output": "program"
}
`

const CandiceProgramDefault = `func main() {
	@print("Hello world!");
}

`
