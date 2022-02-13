# Tutorial 1

This is the introduction tutorial for Candice.
To start a new Candice project on your current folder you should do the following:

```bash
candice init .
```

This will create a new candice.json and an entry point main.cd with the following content:

```javascript
{
  // Name of the project
  "name": "project",
  // File that contains the main function
  "entrypoint": "main.cd",
  // The compiling mode of the project
  // it accepts either LLVM or CXX.
  // CXX = Means that will use a LLVM acceptable compiler like clang that is set on 'cxx'.
  // LLVM = Means that it won't use any external compiler to compile the LLVM code, but it still needs a linker, which you will set on 'cxx'.
  // Both ways are good, it's preferable to use CXX if the developers have clang installed, if not you should use LLVM.
  "kind": "CXX",
  // Compiler that will be used for linking if you are setting 'kind' = 'LLVM' or compiling if you are using 'kind' = 'CXX'
  "cxx": "clang",
  // Flags that will be passed on to the compiler or linker, you can set here external libraries, compiler flags etc.
  "flags": ["-m64"],
  // Name of the executable that will be created.
  "output": "program"
}
```

## Writing your first candice program

On your main.cd you should see the following.

```go
func main() {
	@print("Hello world!");
}
```

If you do:

```bash
candice run .
```

You should see a Hello world! appear.

You can also do:

```bash
candice run . --release
```

It will also make appear a Hello world, but when you want to have a more optimized binary
you should run candice with that flag.

### Variables

A variable declaration looks like this on candice:

```go
func main() {
    // ... snip ...

    variable : i32 = 1;
}
```

You can also omit the type, because Candice infers it for you.

```go
func main() {
    // ... snip ...

    variable := 1;
}
```

By default numbers are signed 32 bit integers. You can cast them to other sizes like this:

```go
func main() {
    // ... snip ...

    variable2 := 2 as i64;
}
```

Remember that if you try to operate with different integer types you will find an error:

```go
func main() {
    // ... snip ...

    @print(variable + variable2);
}
```

You will encounter with the following error:

```bash
 Error Analyzing   error analyzing on 6:19 (at +): _(variable+variable2)_ :: mismatched types, expected=i64, got=i32
```

You can fix it by casting one of the variables, preferably the one with the lesser size.

```go

//...snip
@print(variable as i64 + variable2);

```

If you run again you will find '3' being printed.
