@"%d " = global [3 x i8] c"%d "

define i32 @main() {
_main:
	%0 = add i64 3, 3
	%1 = sub i64 332, 1
	%2 = sdiv i64 %1, 3
	%3 = add i64 %0, %2
	%4 = mul i64 %3, 5
	%5 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%6 = call i32 (i8*, ...) @printf(i8* %5, i64 %4)
	ret i32 0
}

declare i32 @printf(i8* %0, ...)
