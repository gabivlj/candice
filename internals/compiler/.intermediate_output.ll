%Point = type { i64, i64 }

define i32 @main() {
_main:
	%0 = alloca %Point
	%1 = alloca %Point
	%2 = getelementptr %Point, %Point* %1, i32 0, i32 0
	store i64 3, i64* %2
	%3 = getelementptr %Point, %Point* %1, i32 0, i32 1
	store i64 3, i64* %3
	%4 = load %Point, %Point* %1
	store %Point %4, %Point* %0
	ret i32 0
}

declare i32 @printf(i8* %0, ...)
