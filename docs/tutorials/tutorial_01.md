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

You will encounter a type mismatch error.
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
    // you can also do --value.
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

You can also use the `continue` and `break` keywords, as they work like other languages.

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

### Creating object files for libraries

Do you want to support C calling conventions? Or generate object files for your own libraries?
Set on candice.json the following flag:

```js
{
    //...
    "binary": "obj"
}
```

And when you create your library functions you should define them as:

```go
pub func myLibraryFunction() {
    //...
}

```

That should work just fine! When you do `candice build .` it will generate an object file with
myLibraryFunction defined as a symbol for your programs to call.

## Strings

We can concatenate, but be careful with memory leaks!

```go
func main() {
    s := "hello";
    c := "world";

    // 🚨🚨🚨🚨🚨🚨🚨🚨
    // Be careful!!! You are dynamically allocating memory here, you have to free it!
    @print(s + " " + c + "!");

    // This is correct ✅
    string := s + " ";
    string2 := string + c
    string3 := string2 + "!";
    @print(string); // hello world!
    @free(string);
    @free(string2);
    @free(string3);
}
```

Or compare them

```go

func main() {
    if "hello world" != "hello world" {
        @print("this shouldn't be printing...");
    }
}
```

## Unreachable

Sometimes you might wanna indicate the compiler that some code is unreachable, and it will infer that
it's ok to have some kind of return types

```go

func infinite_loop() i32 {
    for 1 {
        @print("hello world!");
    }

    @unreachable();
}

func loop_that_runs_once() i32 {
    for 1 {
        @print("hello world!");
        return 1;
    }

    @unreachable();
}


```

Be careful though, because if unreachable block runs, it will crash!

## Union Types

On Candice there are not enums (yet?), but we have union types, where you can represent your polymorphic data. Here
is an example.

```go
union NumberData {
    bit   i1
    byte  i8
    short i16
    int   i32
    long  i64
}

BIT := 0 as i8
BYTE := 1 as i8
SHORT := 2 as i8
INT := 3 as i8
LONG := 4 as i8

struct Number {
    kind i8
    data NumberData
}

func newInteger(integer i32) Number {
    return @Number{
        kind: INT,
        data: integer,
    }
}

func newByte(byte i8) Number {
    return @Number{
        kind: BYTE,
        data: byte,
    }
}

// etc...

func main() {
    integer := newInteger(10);
    if integer.kind == INT {
        @print("I am the integer", integer.data.int , "nice to meet you!")
    } else if integer.kind == BIT {
        @print("I am a bit...")
    }
}
```

Be careful with accessing unions that you don't know the type of!

## Calling functions related to a type (EXPERIMENTAL)

This functionality is still really experimental and not efficient but cool nonetheless.

Consider the following code:

```go
func min(i i32, j i32) i32 {
	if i <= j {
		return i;
	}

	return j;
}

func max(i i32, j i32) i32 {
	if i >= j {
		return i;
	}

	return j;
}

func main() {
	element := 100;
	element2 := 150;
	minMax := min(max(element2, element), 130);
}
```

It would be really cool to chain that min max thing in a better way right? Now you can simplify that to the following:

```go
    minMax := element.max(element2).min(130);
```

How it works is that Candice will look up a function called 'max' that matches the type with your element on the left!
This also can work on any type as well.

```go
// arr.cd
struct Array {
    ptr *T
    length i32
    capacity i32
}

func get(arr Array, i i32) *T {
    return &arr.ptr[i];
}

func with_capacity(capacity i32) Array {
    ptr := @alloc(T, capacity);
    return @Array {
        ptr: ptr,
        length: 0,
        capacity: capacity,
    };
}

// main.cd
import arr, i32, "./arr.cd";
func main() {
    // init array and whatever
    array := arr.with_capacity(10)
    *array.get(0) = 3;
}


```

## Constants

This feature is a really useful one! Declaring constant variables is a key feature of a language.
Usually you can make all the normal integer operations and declare string literals on constants, but you can't
call functions or do things that you aren't able todo on compile time.

```go
const SOMETHING := 3;
func main() {
    const SOMETHING2 := 3 + SOMETHING;
    @print(SOMETHING2);
}

```

## Switch statements

Switch statements in Candice are like other languages, keep in mind that case expressions need to be constant.

```go
element := 3;
switch element {
    case 4 {
        @print("badly implemented!");
    }

    case 3 {
        @print("ok!");
    }

    default {
        @print("this won't run for sure...");
    }
}
```

## Conditional compile time flags

```go

    #if WINDOWS {
        // instead of putting it on the .json config file, put it here!
        @add_compiler_flag("-m64");
    }

```

## Multiple return values and declarations

On Candice, this is totally valid code:

```go
    // a is 3, b is 4, c is 5
    a, b, c := 3, 4, 5;
    // Now a is 4 and b is 3!
    a, b = b, a;
    // You can do it with any kind of variables...
    w, array, array2 := @Whatever{}, [1]i32{1}, [1]i32{1}
    w.a, w.b, w.c, array[0], array2 = 1, 2, &3, 4, [1]i32{4};
    w.a, w.b, w.c, array[0], array2 = 1, 2, &3, 4, [1]i32{4};
```

The only limitation is that you can't use the const keyword with
multiple values.

You also can make that a function returns multiple values...

```go

func multipleReturner() (i32, i64) {
    return 1, 2 as i64;
}

func main() {
    value1, value2 := multipleReturner();
}

```

But you can't do this.

```go
// 🚫
value, value1, value2 := 1, multipleReturner();
```

## Blocks that return values

You can create block expressions that return values like in Rust or similar.

```go
    someValue := :i32 if value == 3 {
        return value + 1;
    } else {
        return value + 2;
    }

    otherValue := :i64 {
        return 0 as i64;
    }
```

If some path of the block don't return a value, it will throw an error similar to the ones you might find if you didn't return on functions.

## Variable capturing on Anonymous functions

You can capture variables with an anonymous function. But they work differently
compared to other languages.

- The reason is that to support passing functions to external C functions, we can't do fancy stuff with anonymous ones. As they don't differ with functions declared globally.

```go

a := 3;
lambda := func() {
    // captures by value
    @print(a);
};

aRef := &a;
lambdaRefCapture := func() {
    // Captures by reference
    @print(*aRef);
};

lambda(); // 3
lambdaRefCapture() // 3
a = 4;
lambda(); // 3
lambdaRefCapture() // 4

```

Once you create an anonymous function in Candice, it will be the same across its lifetime even though it captures
a local variable, internally it creates a global variable that it accesses once it runs the function, and when
you capture multiple times in a for loop for example it will ovewrite that value.

```go

lambdas := [100]func(){}

for i := 0; i < 100; i++ {
    lambdas[i] = func() { // internally creates a global private value that will access to read 'i', 'i' is copied everytime to it
        @print(i);
    };
}

for i := 0; i < 100; i++ {
    lambdas[i](); // '99' being printed always
}

```

If you want different instances on each lambda you should consider doing another thing like
creating your custom struct with a function attribute that accepts it the struct as a parameter.

## Problems?

If you encounter any kind of bug or problem while following this tutorial, feel free to open a issue on this repository!

```

```
