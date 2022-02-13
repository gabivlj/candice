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

### Built-in Types

Candice has a variety of built-in types. Here is the list!

#### Bool

Can be '1' or '0'. If you make comparisons they produce this type.

- i1.

#### Signed Integers

They are natural numbers that can be negative or positive.

- i8.
- i16.
- i32.
- i64.

#### Unsigned Integers

They are natural numbers that can't be negative.

- u8.
- u16.
- u32.
- u64.

#### Floats

They are real numbers.

- f32.
- f64.

#### Pointers

Put \* in front of your type. We'll explain what they mean later. But if you come from C it just works like them.

#### Arrays

They are a static list of elements.

You use them like '[constant_integer]type'

For example:

```go
numbers := [5]i32{1, 2, 3, 4, 5}
empty_numbers := [5]i32{}
```

Remember that the value inside the [] should be constant, not a variable.

### Operations

You can make a lot of operations with numbers on candice.
'+', '-', '\*', '/'...

```go

a := 3
b := 5
c := 7
result := a * b / c - 1;
@print(result); // prints '1'

```

You can also use bitwise operations like:

- '&' AND
- '|' OR
- '^' XOR

### Ifs

Ifs are a must in all languages! When the condition they evaluate is true they will execute some code.

```go

number := 10

if number >= 5 && number <= 10
    @print("number is between 5 and 10");

```

As you can see && are used to check if both conditions are true.
You can even cover more cases at the same time.

```go
if number >= 5 && number <= 10 @print("number is between 5 and 10");
else if number > 0 || number < 0 {
    @print("atleast the number is more or less than 0");
} else {
    @print("oh no! candice is broken");
    @print("please fix...");
}
```

With these examples you are set to use ifs on candice!

### Loops

Loops are really simple. You are able to execute a block of code multiple times while a condition is true.

```go

value := 10

for value > 0 {
    @print("value is", value);
    value = value - 1;
}

```

This will print 'value is ...' 10 times.

You can also do this same code like this.

```go
for value := 10; value > 0; value = value - 1 {
    @print("value is", value);
}
```

You can even omit the braces.

```go
for value := 10; value > 0; value = value - 1
    @print("value is", value);

```

One way to use loops is to walk through a list.

```go

list := [100]i32{}

for i := 0; i < 100; i = i + 1 {
    // We are accessing the i th position of the list and setting it to i * i;
    list[i] = i * i;
}

for i := 0; i < 100;  i = i + 1 {
    @print("position", i, "is", list[i]);
}

```

Now you are ready to use for loops!
