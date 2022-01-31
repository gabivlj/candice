%Something = type { i32 }

@string.literal.PVQafpfHTq = global [13 x i8] c"hello world!\00"
@"%s " = global [4 x i8] c"%s \00"
@"%d " = global [4 x i8] c"%d \00"

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc %Something @something() {
something:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.PVQafpfHTq, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something
	%4 = getelementptr %Something, %Something* %3, i32 0, i32 0
	store i32 0, i32* %4
	%5 = load %Something, %Something* %3
	ret %Something %5
}

define ccc void @main() {
main:
	%0 = call %Something @something()
	%1 = alloca %Something
	store %Something %0, %Something* %1
	%2 = getelementptr %Something, %Something* %1, i32 0, i32 0
	%3 = load i32, i32* %2
	%4 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%5 = call i32 (i8*, ...) @printf(i8* %4, i32 %3)
	ret void
}
