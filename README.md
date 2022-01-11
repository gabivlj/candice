## Candice
Dead simple programming language.
Not expected to be prepared for a real developer environment but can
be fun to work with.

## Philosophy
* Simple syntax.
* Low level capabilities.
* C interoperability.
* Compiled and strongly typed.

````go
// Compile like: ./compiler -name output main.cd
// Write structs like this
struct Person {
    // C like strings
    name   *i8

    age    u8
    isCool i1

    // You can have recursive structs
    nextPerson *Person

    // Static arrays
    ID [10]i32
}

func DefaultPerson(isCool i1) *Person {
    // allocate in heap one person
    person := @alloc(Person, 1);

    if isCool {
        person.age = @cast(u8, 21);
        person.name = "James";
    } else {
        person.age = @cast(u8, 66);
        person.name = "Pedro";
    }

    person.isCool = isCool;
    // Explicit null pointers (cast integer to pointer)
    person.nextPerson = @cast(*Person, 0);

    person.ID = [10]i32{1, 1, 1, 1};

    return person;
}

// write external functions from C like this.
extern func free(*i8) void;

// Write your entry code here
func main() {
    person := DefaultPerson(@cast(i1, 1));
    @println(person.name);

    persons := @alloc(Person, 300);

    // or
    // persons :*Person = @alloc(Person, 300);

    for i := 0; i < 300; i = i + 1 {
        // copy
        persons[i] = *DefaultPerson(i >= 100);
    }

    // alternative for loops
    i := 0;
    for i < 300 {
        if i >= 299 @println("Ifs work without braces!");
        i = i + 1;
    }

    free(@cast(*i8, persons));
}

````