@"%d " = global [3 x i8] c"%d "

define i32 @main() {
_main:
	%0 = xor i64 3322323, 51231212
	%1 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i64 %0)
	ret i32 0
}

declare i32 @printf(i8* %0, ...)
