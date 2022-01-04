@"%d " = global [3 x i8] c"%d "

define i32 @main() {
_main:
	%0 = add i64 3, 5
	%1 = alloca i64
	store i64 %0, i64* %1
	%2 = load i64, i64* %1
	%3 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%4 = call i32 (i8*, ...) @printf(i8* %3, i64 %2)
	ret i32 0
}

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)
