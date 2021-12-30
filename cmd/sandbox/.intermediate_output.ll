%Point = type { i64, i64, %Point* }

@"%d " = global [3 x i8] c"%d "

define i32 @main() {
_main:
	%0 = alloca %Point
	%1 = alloca %Point
	%2 = getelementptr %Point, %Point* %1, i32 0, i32 0
	store i64 43, i64* %2
	%3 = getelementptr %Point, %Point* %1, i32 0, i32 1
	store i64 55, i64* %3
	%4 = mul i64 33, 0
	%5 = call i8* @malloc(i64 %4)
	%6 = bitcast i8* %5 to %Point*
	%7 = alloca %Point*
	store %Point* %6, %Point** %7
	%8 = getelementptr %Point, %Point* %1, i32 0, i32 2
	%9 = load %Point*, %Point** %7
	store %Point* %9, %Point** %8
	%10 = load %Point, %Point* %1
	store %Point %10, %Point* %0
	%11 = getelementptr %Point, %Point* %0, i32 0, i32 0
	store i64 3, i64* %11
	%12 = alloca i64
	%13 = getelementptr %Point, %Point* %0, i32 0, i32 0
	%14 = load i64, i64* %13
	store i64 %14, i64* %12
	%15 = load i64, i64* %12
	%16 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%17 = call i32 (i8*, ...) @printf(i8* %16, i64 %15)
	ret i32 0
}

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)
