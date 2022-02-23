# Candice

Candice is a simple programming language to use. It uses LLVM under the hood as
the intermediate representation, and can be used on all machines that is compiled to, but you still need a linker/compiler
to link your generated objects, but Candice out of the box will manage it for you if you have one installed.

It targets to be a simple language to write C-like code, but with a different syntax and more features.
The syntax of Candice wants to be a mix of Golang, C and Rust. You'll find out as the tutorial goes on!

## Building and installing

- You need go 1.18 installed.
- If you want to build the compiler with LLVM included, so that we don't need Clang to compile programs,
  you need to install LLVM in your machine and add it in your PATH. If you don't know how search up how to install
  LLVM in your machine.
- If you don't have LLVM, Candice will default to clang usage.

Once you are set with those steps, go to cmd/build and execute:

```bash
go build .
```

That's it, now you can do with the generated executable whatever you want,
it's recommended that you set it on your machine's PATH so you can call it like this everywhere

```
candice
```

If it's installed and set fine, you should see the following output:

```bash
 Error Flags   Error retrieving flags: not enough arguments

        Usage:
                candice <mode> <path> <flags>
        Modes:
                run - Run the project in the desired path.
                build - Creates an executable of the project in the desired path.
                init - Creates a candice project
        Flags:
                --release - Create or runs an optimized build of the project.

```

Enjoy! Now we'll go on to writing some Candice.
