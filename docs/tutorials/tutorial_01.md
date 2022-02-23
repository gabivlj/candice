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

If you just want a binary to deploy somewhere, do:

```bash
candice build . --release
```

This will generate a binary of your current OS.

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

// You can also create numbers like this:
numberHexa := 0xFFF;
numberBin := 0b11010110;

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

for i := 0; i < 100; i = i + 1 {
    @print("position", i, "is", list[i]);
}

```

Now you are ready to use for loops!

### Allocating dynamic memory

On candice allocating dynamic memory is straightforward.

```go

memory : *i32 = @alloc(i32, 100);
// or
// memory := @alloc(i32, 100);

```

This just created 100 i32 integers.
To access the first element of this chunk of memory you can do this:

```go
memory[0] = 10;
@print(memory[0]); // 10
```

If you want to resize the memory you created you can do this:

```go
// the size now is 200.
memory = @realloc(memory, 200);
```

Every chunk of memory you create, you obviously have to free it when you are not using it anymore.

```go
@free(memory);
```

Now you know how to allocate memory on the heap with Candice!

### Structs

Structs on Candice are pretty simple.

```go

struct Point {
    x f64
    y f64
}

struct Points {
    points *Point
    number_of_points u32
}

func main() {
    points := @Points{
		points: @alloc(Point, 3),
		number_of_points: 10 as u32
	};

	@print(@sizeof(Points)); // 16 bytes

	points.points[0] = @Point{
		x: 0.0 as f64,
		y: 0.0 as f64,
	};

	points.points[0].x = points.points[0].x + (10.0 as f64);

	@print(points.points[0].x); // 10.0
}

```

You can also allocate or have list of your own structs.

```go

pointsOfPoints := [100]Points{points}
pointsOfPointsAlloc := @alloc(Points, 100);

```

### Functions

Candice also contains functions.

```go


// Simple function that prints hello
func hello() {
    @print("hello");
}

// Function that adds two i32.
func add(x i32, y i32) i32 {
    return x + y;
}

// Calls a simple function
func callMyFunction(myFunction func() void) {
    myFunction();
}

func main() {
    // You can call those functions like this:
    hello(); // prints hello
    @print(add(3, 3)); // prints 6
    callMyFunction(hello); // Calls hello.
    // You can even create functions like this!
    callMyFunction(func() {
        @print("prettier hello");
    });

	localFunction := func() {
		@print("local function");
	};

	callMyFunction(localFunction)

	result := func () i32 {
		@print("tricky!");
		return 1;
	}();
	@print(result);
}

```

### Importing files

On candice you can import functions or structs from other files.

```go
// cool.cd

struct Cool {
    value i32
}

func Create() Cool {
    return @Cool {
        value: 0
    };
}

func Print(cool *Cool) {
    @print("Cool { value:", cool.value, "}");
}

```

Then you can import it in another file, for example in your main.cd:

```go
import cool, "cool.cd";

func main() {
	coolObject := cool.Create();
	cool.Print(&coolObject); // Cool { value: 0 }
}

```

### Generic files

On candice there is a new concept that we call generic files, where you can define generic types on top of the file and
when you import the file you can pass the types as parameters.
For example imagine we want to create a generic struct called 'Pair'.

```go
// generic.cd
type T;
type C;

struct Pair {
    first T
    second C
}

func New(first T, second C) Pair {
    return @Pair {
        first: first,
        second: second,
    };
}
```

```go
// main.cd
import pairnumbers, i32, *i8, "generic.cd";

func main() {
	pair := pairnumbers.New(1, "hello");
	@print(pair.first, pair.second);
}
```

As we can see, we pass an i32 and \*i8 as type parameters, and the imported file is compiled with those types.

We can do even crazier stuff like creating a generic callback struct.

```go
type T;
type C;
type V;

type Funk = func(T, V) C;

struct Callback {
    funk Funk
    value V
}

func F(value V, f Funk) Callback {
    return @Callback {
        funk: f,
        value: value,
    };
}

func Run(param T, callback Callback) C {
    return callback.funk(param, callback.value);
}

```

We accept 3 parameters, the first one is the type of the parameter when calling the callback,
the second one is the type of the parameter that the callback is storing, and the third one is the type that it returns.
We can use it like this.

```go
import callback, i32, i64, i32, "callback.cd";

func main() {
	i := 1
	c := callback.F(i, func(x i32, y i32) i64 {
		return (x + y) as i64;
	});
	result := callback.Run(32, c);
	@print(result);
}

```

As we can see, we can get really creative with generic files.
We can even pass to generic files our own types.

```go

type F32 = f32

struct Point {
	x f32
	y f32
}

import callback, F32, Point, Point, "callback.cd";

func main() {
	c := callback.F(@Point{x: 1.0}, func(x F32, p Point) Point {
		return @Point {
			x: x,
			y: p.x,
		}
	});
	result := callback.Run(0.0, c);
	@print(result.x, result.y);
}
```

### Importing external functions from C

We can import files from C like this.

```go
extern func thefunction(...params...) theReturnType;
```

For example we can use the file library from the std like this.

```go

type FILE = i8


extern func fopen(*i8, *i8) *FILE;
extern func fwrite(*i8, i64, i64, *FILE) void;
extern func fclose(*FILE) void;
extern func fread(*i8, i64, i64, *FILE) i64;
extern func rewind(*FILE) i64;
extern func fgets(*i8, i32, *FILE) *i8;


func main() {
    theFile := fopen("./some_file", "a+");
    rewind(theFile);
    output := @alloc(i8, 300);
    for fgets(output, 10, theFile) as i32 != 0 {
        @print(output);
    }
    fwrite("hello!", 1 as i64, 6 as i64, theFile);
    @free(output);
    fclose(theFile)
}

```

### Compile time Ifs

You can compile specific code only if a condition meets, like if you are on a specific platform or architecture.

```go
#if LINUX {
    func platform() *i8 {
        return "linux"
    }
}


#if WINDOWS {
    func platform() *i8 {
        return "windows"
    }
}

#if MACOS {
    func platform() *i8 {
        return "darwin"
    }
}

#if X64 {
    func arch() *i8 {
        return "amd64"
    }
}

#if ARM64 {
    func arch() *i8 {
        return "arm64"
    }
}

func main() {
    @print(platform())
    @print(arch())
}

```

### Blocks and scope

You can create blocks of code as well.

```go

{
    variable := 0
    {
        variable2 := variable + 5;
    } // You can't use variable2 anymore
} // You can't use variable anymore

```

### Order of declaration

On candice you can declare in the order you want functions and structs.

### Global Variables

Global variables can only be initalised with constant expressions, at the moment constant expressions in candice
are string literals, numbers, and number casts.

### Variable shadowing

This is totally valid Candice code.

```go

func thing(x i32) {
    // redeclare x as a i64
    x := 0 as i64;
    {
        // redeclare x in this block as a i16
        x := 1 as i16
        @print(x) // 1
    }
    @print(x) // 0
}

```

### Undefined number of parameters on C functions

You can declare printf C functions that accept and undefined number of parameters like this:

```go
extern func printf(*i8, ..);
```

Candice doesn't have any plans of supporting infinite number of variables on candice functions.

## Problems?

If you encounter any kind of bug or problem while following this tutorial, feel free to open a issue on this repository!
