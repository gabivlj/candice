define i32 @main() {
_main:
	%0 = alloca i32*
	%1 = mul i64 5, 4
	%2 = call i8* @malloc(i64 %1)
	%3 = bitcast i8* %2 to i32*
	store i32* %3, i32** %0
	ret i32 0
}

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)
