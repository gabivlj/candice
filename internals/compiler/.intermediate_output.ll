@"%d " = global [3 x i8] c"%d "

define i32 @main() {
_main:
	%0 = alloca i64
	%1 = add i64 3, 5
	store i64 %1, i64* %0
	%2 = load i64, i64* %0
	%3 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%4 = call i32 (i8*, ...) @printf(i8* %3, i64 %2)
	ret i32 0
}

declare i32 @printf(i8* %0, ...)
