%Point2 = type { i64, i64 }
%Point = type { i64, i64, %Point2 }

@"%d " = global [3 x i8] c"%d "

define i32 @main() {
_main:
	%0 = alloca %Point*
	%1 = alloca %Point
	%2 = getelementptr %Point, %Point* %1, i32 0, i32 0
	store i64 3, i64* %2
	%3 = getelementptr %Point, %Point* %1, i32 0, i32 1
	store i64 3, i64* %3
	%4 = alloca %Point2
	%5 = getelementptr %Point2, %Point2* %4, i32 0, i32 0
	store i64 3, i64* %5
	%6 = getelementptr %Point2, %Point2* %4, i32 0, i32 1
	store i64 3, i64* %6
	%7 = getelementptr %Point, %Point* %1, i32 0, i32 2
	%8 = load %Point2, %Point2* %4
	store %Point2 %8, %Point2* %7
	store %Point* %1, %Point** %0
	%9 = alloca i64
	%10 = load %Point*, %Point** %0
	%11 = getelementptr %Point, %Point* %10, i32 0, i32 2
	%12 = getelementptr %Point2, %Point2* %11, i32 0, i32 0
	%13 = load i64, i64* %12
	store i64 %13, i64* %9
	%14 = load i64, i64* %9
	%15 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%16 = call i32 (i8*, ...) @printf(i8* %15, i64 %14)
	ret i32 0
}

declare i32 @printf(i8* %0, ...)
