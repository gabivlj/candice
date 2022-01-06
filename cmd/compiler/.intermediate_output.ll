@"%d " = global [4 x i8] c"%d \00"

define i32 @main() {
_main:
	%0 = call i32 @fibonacci(i32 45)
	%1 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i32 %0)
	ret i32 0
}

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc i32 @fibonacci(i32 %current) {
fibonacci:
	%0 = alloca i32
	store i32 %current, i32* %0
	%1 = load i32, i32* %0
	%2 = icmp slt i32 %1, 2
	br i1 %2, label %if.then.xkezMbivNL, label %if.else.cbFfuzKhBg

if.then.xkezMbivNL:
	%3 = load i32, i32* %0
	ret i32 %3

if.else.cbFfuzKhBg:
	br label %leave.oUUjvQMktq

leave.oUUjvQMktq:
	%4 = load i32, i32* %0
	%5 = sub i32 %4, 1
	%6 = call i32 @fibonacci(i32 %5)
	%7 = load i32, i32* %0
	%8 = sub i32 %7, 2
	%9 = call i32 @fibonacci(i32 %8)
	%10 = add i32 %6, %9
	ret i32 %10
}
