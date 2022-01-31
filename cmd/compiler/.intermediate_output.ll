%Something-rhqrhBWRCi = type { i32 }

@string.literal.dOsdmeAAor = global [13 x i8] c"hello world!\00"
@"%s " = global [4 x i8] c"%s \00"
@"%d " = global [4 x i8] c"%d \00"

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc void @outsideFunction-rhqrhBWRCi(%Something-rhqrhBWRCi* %s-rhqrhBWRCi) {
outsideFunction-rhqrhBWRCi:
	%0 = alloca %Something-rhqrhBWRCi*
	store %Something-rhqrhBWRCi* %s-rhqrhBWRCi, %Something-rhqrhBWRCi** %0
	%1 = load %Something-rhqrhBWRCi*, %Something-rhqrhBWRCi** %0
	%2 = getelementptr %Something-rhqrhBWRCi, %Something-rhqrhBWRCi* %1, i32 0, i32 0
	%3 = load %Something-rhqrhBWRCi*, %Something-rhqrhBWRCi** %0
	%4 = getelementptr %Something-rhqrhBWRCi, %Something-rhqrhBWRCi* %3, i32 0, i32 0
	%5 = load i32, i32* %4
	%6 = add i32 %5, 1
	store i32 %6, i32* %2
	ret void
}

define ccc %Something-rhqrhBWRCi @something-rhqrhBWRCi() {
something-rhqrhBWRCi:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.dOsdmeAAor, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-rhqrhBWRCi
	%4 = getelementptr %Something-rhqrhBWRCi, %Something-rhqrhBWRCi* %3, i32 0, i32 0
	store i32 0, i32* %4
	%5 = alloca %Something-rhqrhBWRCi
	%6 = load %Something-rhqrhBWRCi, %Something-rhqrhBWRCi* %3
	store %Something-rhqrhBWRCi %6, %Something-rhqrhBWRCi* %5
	%7 = alloca %Something-rhqrhBWRCi*
	store %Something-rhqrhBWRCi* %5, %Something-rhqrhBWRCi** %7
	%8 = load %Something-rhqrhBWRCi*, %Something-rhqrhBWRCi** %7
	call void @outsideFunction-rhqrhBWRCi(%Something-rhqrhBWRCi* %8)
	%9 = load %Something-rhqrhBWRCi, %Something-rhqrhBWRCi* %5
	ret %Something-rhqrhBWRCi %9
}

define ccc i32 @something-tLwyhheCNT() {
something-tLwyhheCNT:
	ret i32 1
}

define ccc void @main() {
main:
	%0 = call %Something-rhqrhBWRCi @something-rhqrhBWRCi()
	%1 = alloca %Something-rhqrhBWRCi
	store %Something-rhqrhBWRCi %0, %Something-rhqrhBWRCi* %1
	%2 = alloca %Something-rhqrhBWRCi
	%3 = alloca %Something-rhqrhBWRCi
	%4 = load %Something-rhqrhBWRCi, %Something-rhqrhBWRCi* %2
	store %Something-rhqrhBWRCi %4, %Something-rhqrhBWRCi* %3
	%5 = call i32 @something-tLwyhheCNT()
	%6 = alloca i32
	store i32 %5, i32* %6
	%7 = getelementptr %Something-rhqrhBWRCi, %Something-rhqrhBWRCi* %1, i32 0, i32 0
	%8 = load i32, i32* %7
	%9 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%10 = call i32 (i8*, ...) @printf(i8* %9, i32 %8)
	%11 = load i32, i32* %6
	%12 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%13 = call i32 (i8*, ...) @printf(i8* %12, i32 %11)
	ret void
}
