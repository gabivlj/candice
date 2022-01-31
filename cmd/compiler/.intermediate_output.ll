%Something- = type { i32 }

@string.literal.KOMHOYdjlS = global [13 x i8] c"hello world!\00"
@"%s " = global [4 x i8] c"%s \00"
@"%d " = global [4 x i8] c"%d \00"

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc void @outsideFunction-(%Something-* %s-) {
outsideFunction-:
	%0 = alloca %Something-*
	store %Something-* %s-, %Something-** %0
	%1 = load %Something-*, %Something-** %0
	%2 = getelementptr %Something-, %Something-* %1, i32 0, i32 0
	%3 = load %Something-*, %Something-** %0
	%4 = getelementptr %Something-, %Something-* %3, i32 0, i32 0
	%5 = load i32, i32* %4
	%6 = add i32 %5, 1
	store i32 %6, i32* %2
	ret void
}

define ccc %Something- @something-() {
something-:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.KOMHOYdjlS, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-
	%4 = getelementptr %Something-, %Something-* %3, i32 0, i32 0
	store i32 0, i32* %4
	%5 = alloca %Something-
	%6 = load %Something-, %Something-* %3
	store %Something- %6, %Something-* %5
	%7 = alloca %Something-*
	store %Something-* %5, %Something-** %7
	%8 = load %Something-*, %Something-** %7
	call void @outsideFunction-(%Something-* %8)
	%9 = load %Something-, %Something-* %5
	ret %Something- %9
}

define ccc void @main() {
main:
	%0 = call %Something- @something-()
	%1 = alloca %Something-
	store %Something- %0, %Something-* %1
	%2 = getelementptr %Something-, %Something-* %1, i32 0, i32 0
	%3 = load i32, i32* %2
	%4 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%5 = call i32 (i8*, ...) @printf(i8* %4, i32 %3)
	ret void
}
