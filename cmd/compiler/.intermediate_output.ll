%T-bNhQjtiGSj = type i64
%Something-bNhQjtiGSj = type { i64 }
%T-SHxznMerpt = type i32
%Something-SHxznMerpt = type { i32 }

@string.literal.LtTYvBRlGM = global [13 x i8] c"hello world!\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.LhQKRIoHvO = global [13 x i8] c"hello world!\00"
@"%d " = global [4 x i8] c"%d \00"

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc void @outsideFunction-bNhQjtiGSj(%Something-bNhQjtiGSj* %s-bNhQjtiGSj) {
outsideFunction-bNhQjtiGSj:
	%0 = alloca %Something-bNhQjtiGSj*
	store %Something-bNhQjtiGSj* %s-bNhQjtiGSj, %Something-bNhQjtiGSj** %0
	%1 = load %Something-bNhQjtiGSj*, %Something-bNhQjtiGSj** %0
	%2 = getelementptr %Something-bNhQjtiGSj, %Something-bNhQjtiGSj* %1, i32 0, i32 0
	%3 = load %Something-bNhQjtiGSj*, %Something-bNhQjtiGSj** %0
	%4 = getelementptr %Something-bNhQjtiGSj, %Something-bNhQjtiGSj* %3, i32 0, i32 0
	%5 = load i64, i64* %4
	%6 = sext i32 1 to i64
	%7 = add i64 %5, %6
	store i64 %7, i64* %2
	ret void
}

define ccc %Something-bNhQjtiGSj @something-bNhQjtiGSj() {
something-bNhQjtiGSj:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.LtTYvBRlGM, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-bNhQjtiGSj
	%4 = sext i32 0 to i64
	%5 = getelementptr %Something-bNhQjtiGSj, %Something-bNhQjtiGSj* %3, i32 0, i32 0
	store i64 %4, i64* %5
	%6 = alloca %Something-bNhQjtiGSj
	%7 = load %Something-bNhQjtiGSj, %Something-bNhQjtiGSj* %3
	store %Something-bNhQjtiGSj %7, %Something-bNhQjtiGSj* %6
	%8 = alloca %Something-bNhQjtiGSj*
	store %Something-bNhQjtiGSj* %6, %Something-bNhQjtiGSj** %8
	%9 = load %Something-bNhQjtiGSj*, %Something-bNhQjtiGSj** %8
	call void @outsideFunction-bNhQjtiGSj(%Something-bNhQjtiGSj* %9)
	%10 = load %Something-bNhQjtiGSj, %Something-bNhQjtiGSj* %6
	ret %Something-bNhQjtiGSj %10
}

define ccc void @outsideFunction-SHxznMerpt(%Something-SHxznMerpt* %s-SHxznMerpt) {
outsideFunction-SHxznMerpt:
	%0 = alloca %Something-SHxznMerpt*
	store %Something-SHxznMerpt* %s-SHxznMerpt, %Something-SHxznMerpt** %0
	%1 = load %Something-SHxznMerpt*, %Something-SHxznMerpt** %0
	%2 = getelementptr %Something-SHxznMerpt, %Something-SHxznMerpt* %1, i32 0, i32 0
	%3 = load %Something-SHxznMerpt*, %Something-SHxznMerpt** %0
	%4 = getelementptr %Something-SHxznMerpt, %Something-SHxznMerpt* %3, i32 0, i32 0
	%5 = load i32, i32* %4
	%6 = add i32 %5, 1
	store i32 %6, i32* %2
	ret void
}

define ccc %Something-SHxznMerpt @something-SHxznMerpt() {
something-SHxznMerpt:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.LhQKRIoHvO, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-SHxznMerpt
	%4 = getelementptr %Something-SHxznMerpt, %Something-SHxznMerpt* %3, i32 0, i32 0
	store i32 0, i32* %4
	%5 = alloca %Something-SHxznMerpt
	%6 = load %Something-SHxznMerpt, %Something-SHxznMerpt* %3
	store %Something-SHxznMerpt %6, %Something-SHxznMerpt* %5
	%7 = alloca %Something-SHxznMerpt*
	store %Something-SHxznMerpt* %5, %Something-SHxznMerpt** %7
	%8 = load %Something-SHxznMerpt*, %Something-SHxznMerpt** %7
	call void @outsideFunction-SHxznMerpt(%Something-SHxznMerpt* %8)
	%9 = load %Something-SHxznMerpt, %Something-SHxznMerpt* %5
	ret %Something-SHxznMerpt %9
}

define ccc i32 @something-yztBzJToDx() {
something-yztBzJToDx:
	ret i32 1
}

define ccc void @main() {
main:
	%0 = call %Something-bNhQjtiGSj @something-bNhQjtiGSj()
	%1 = alloca %Something-bNhQjtiGSj
	store %Something-bNhQjtiGSj %0, %Something-bNhQjtiGSj* %1
	%2 = call %Something-SHxznMerpt @something-SHxznMerpt()
	%3 = alloca %Something-SHxznMerpt
	store %Something-SHxznMerpt %2, %Something-SHxznMerpt* %3
	%4 = alloca %Something-bNhQjtiGSj
	%5 = alloca %Something-bNhQjtiGSj
	%6 = load %Something-bNhQjtiGSj, %Something-bNhQjtiGSj* %4
	store %Something-bNhQjtiGSj %6, %Something-bNhQjtiGSj* %5
	%7 = call i32 @something-yztBzJToDx()
	%8 = alloca i32
	store i32 %7, i32* %8
	%9 = getelementptr %Something-SHxznMerpt, %Something-SHxznMerpt* %3, i32 0, i32 0
	%10 = load i32, i32* %9
	%11 = getelementptr %Something-bNhQjtiGSj, %Something-bNhQjtiGSj* %1, i32 0, i32 0
	%12 = load i64, i64* %11
	%13 = trunc i64 %12 to i32
	%14 = add i32 %10, %13
	%15 = load i32, i32* %8
	%16 = add i32 %14, %15
	%17 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%18 = call i32 (i8*, ...) @printf(i8* %17, i32 %16)
	ret void
}
