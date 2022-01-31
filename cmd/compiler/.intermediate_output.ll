%T-uRrWAnDfsz = type i64
%Something-uRrWAnDfsz = type { %T-uRrWAnDfsz }

@string.literal.QKAtzYXqRo = global [13 x i8] c"hello world!\00"
@"%s " = global [4 x i8] c"%s \00"
@0 = global [1 x i8] c"\00"
@"%d " = global [4 x i8] c"%d \00"

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc void @outsideFunction-uRrWAnDfsz(%Something-uRrWAnDfsz* %s-uRrWAnDfsz) {
outsideFunction-uRrWAnDfsz:
	%0 = alloca %Something-uRrWAnDfsz*
	store %Something-uRrWAnDfsz* %s-uRrWAnDfsz, %Something-uRrWAnDfsz** %0
	%1 = load %Something-uRrWAnDfsz*, %Something-uRrWAnDfsz** %0
	%2 = getelementptr %Something-uRrWAnDfsz, %Something-uRrWAnDfsz* %1, i32 0, i32 0
	%3 = load %Something-uRrWAnDfsz*, %Something-uRrWAnDfsz** %0
	%4 = getelementptr %Something-uRrWAnDfsz, %Something-uRrWAnDfsz* %3, i32 0, i32 0
	%5 = load %T-uRrWAnDfsz, %T-uRrWAnDfsz* %4
	%6 = sext i32 1 to i64
	%7 = add %T-uRrWAnDfsz %5, %6
	store %T-uRrWAnDfsz %7, %T-uRrWAnDfsz* %2
	ret void
}

define ccc %Something-uRrWAnDfsz @something-uRrWAnDfsz() {
something-uRrWAnDfsz:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.QKAtzYXqRo, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-uRrWAnDfsz
	%4 = sext i32 0 to i64
	%5 = getelementptr %Something-uRrWAnDfsz, %Something-uRrWAnDfsz* %3, i32 0, i32 0
	store i64 %4, %T-uRrWAnDfsz* %5
	%6 = alloca %Something-uRrWAnDfsz
	%7 = load %Something-uRrWAnDfsz, %Something-uRrWAnDfsz* %3
	store %Something-uRrWAnDfsz %7, %Something-uRrWAnDfsz* %6
	%8 = alloca %Something-uRrWAnDfsz*
	store %Something-uRrWAnDfsz* %6, %Something-uRrWAnDfsz** %8
	%9 = load %Something-uRrWAnDfsz*, %Something-uRrWAnDfsz** %8
	call void @outsideFunction-uRrWAnDfsz(%Something-uRrWAnDfsz* %9)
	%10 = load %Something-uRrWAnDfsz, %Something-uRrWAnDfsz* %6
	ret %Something-uRrWAnDfsz %10
}

define ccc i32 @something-tOOmZGhrrc() {
something-tOOmZGhrrc:
	ret i32 1
}

define ccc void @main() {
main:
	%0 = call %Something-uRrWAnDfsz @something-uRrWAnDfsz()
	%1 = alloca %Something-uRrWAnDfsz
	store %Something-uRrWAnDfsz %0, %Something-uRrWAnDfsz* %1
	%2 = alloca %Something-uRrWAnDfsz
	%3 = alloca %Something-uRrWAnDfsz
	%4 = load %Something-uRrWAnDfsz, %Something-uRrWAnDfsz* %2
	store %Something-uRrWAnDfsz %4, %Something-uRrWAnDfsz* %3
	%5 = call i32 @something-tOOmZGhrrc()
	%6 = alloca i32
	store i32 %5, i32* %6
	%7 = getelementptr %Something-uRrWAnDfsz, %Something-uRrWAnDfsz* %1, i32 0, i32 0
	%8 = load %T-uRrWAnDfsz, %T-uRrWAnDfsz* %7
	%9 = getelementptr [1 x i8], [1 x i8]* @0, i32 0, i32 0
	%10 = call i32 (i8*, ...) @printf(i8* %9, %T-uRrWAnDfsz %8)
	%11 = load i32, i32* %6
	%12 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%13 = call i32 (i8*, ...) @printf(i8* %12, i32 %11)
	ret void
}
