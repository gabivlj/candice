@"%d %d %d " = global [10 x i8] c"%d %d %d \00"

define i32 @main() {
main:
	%0 = alloca i32
	store i32 3, i32* %0
	%1 = alloca i32
	store i32 4, i32* %1
	%2 = load i32, i32* %0
	%3 = load i32, i32* %1
	%4 = call i32 asm "mov $0, $1\0A mov $0, $2", "=r,r,r,~{dirflag},~{fpsr},~{flags}"(i32 %2, i32 %3)
	%5 = alloca i32
	store i32 %4, i32* %5
	%6 = load i32, i32* %5
	%7 = load i32, i32* %0
	%8 = load i32, i32* %1
	%9 = getelementptr [10 x i8], [10 x i8]* @"%d %d %d ", i32 0, i32 0
	%10 = call i32 (i8*, ...) @printf(i8* %9, i32 %6, i32 %7, i32 %8)
	ret i32 0
}

declare ccc i32 @printf(i8* %0, ...)
