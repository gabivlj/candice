%Something-QyQOGHHKMW = type { i32 }

@string.literal.rvbuTUqFfI = global [13 x i8] c"hello world!\00"
@"%s " = global [4 x i8] c"%s \00"
@"%d " = global [4 x i8] c"%d \00"

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc void @outsideFunction-QyQOGHHKMW(%Something-QyQOGHHKMW* %s-QyQOGHHKMW) {
outsideFunction-QyQOGHHKMW:
	%0 = alloca %Something-QyQOGHHKMW*
	store %Something-QyQOGHHKMW* %s-QyQOGHHKMW, %Something-QyQOGHHKMW** %0
	%1 = load %Something-QyQOGHHKMW*, %Something-QyQOGHHKMW** %0
	%2 = getelementptr %Something-QyQOGHHKMW, %Something-QyQOGHHKMW* %1, i32 0, i32 0
	%3 = load %Something-QyQOGHHKMW*, %Something-QyQOGHHKMW** %0
	%4 = getelementptr %Something-QyQOGHHKMW, %Something-QyQOGHHKMW* %3, i32 0, i32 0
	%5 = load i32, i32* %4
	%6 = add i32 %5, 1
	store i32 %6, i32* %2
	ret void
}

define ccc %Something-QyQOGHHKMW @something-QyQOGHHKMW() {
something-QyQOGHHKMW:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.rvbuTUqFfI, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-QyQOGHHKMW
	%4 = getelementptr %Something-QyQOGHHKMW, %Something-QyQOGHHKMW* %3, i32 0, i32 0
	store i32 0, i32* %4
	%5 = alloca %Something-QyQOGHHKMW
	%6 = load %Something-QyQOGHHKMW, %Something-QyQOGHHKMW* %3
	store %Something-QyQOGHHKMW %6, %Something-QyQOGHHKMW* %5
	%7 = alloca %Something-QyQOGHHKMW*
	store %Something-QyQOGHHKMW* %5, %Something-QyQOGHHKMW** %7
	%8 = load %Something-QyQOGHHKMW*, %Something-QyQOGHHKMW** %7
	call void @outsideFunction-QyQOGHHKMW(%Something-QyQOGHHKMW* %8)
	%9 = load %Something-QyQOGHHKMW, %Something-QyQOGHHKMW* %5
	ret %Something-QyQOGHHKMW %9
}

define ccc i32 @something-cPvbIEDgqH() {
something-cPvbIEDgqH:
	ret i32 1
}

define ccc void @main() {
main:
	%0 = call %Something-QyQOGHHKMW @something-QyQOGHHKMW()
	%1 = alloca %Something-QyQOGHHKMW
	store %Something-QyQOGHHKMW %0, %Something-QyQOGHHKMW* %1
	%2 = call i32 @something-cPvbIEDgqH()
	%3 = alloca i32
	store i32 %2, i32* %3
	%4 = getelementptr %Something-QyQOGHHKMW, %Something-QyQOGHHKMW* %1, i32 0, i32 0
	%5 = load i32, i32* %4
	%6 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%7 = call i32 (i8*, ...) @printf(i8* %6, i32 %5)
	%8 = load i32, i32* %3
	%9 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%10 = call i32 (i8*, ...) @printf(i8* %9, i32 %8)
	ret void
}
