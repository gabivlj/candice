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
