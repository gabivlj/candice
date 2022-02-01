%Something-RrXtmEkgSV = type { i64 }
%Something-hDdwdxkkPY = type { i32 }
%Thing-MLrehTcaVM = type { i32 }

@string.literal.fLDiNqBTme = global [13 x i8] c"hello world!\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.emSlswNwtv = global [13 x i8] c"hello world!\00"
@"%d " = global [4 x i8] c"%d \00"

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc void @outsideFunction-RrXtmEkgSV(%Something-RrXtmEkgSV* %s-RrXtmEkgSV) {
outsideFunction-RrXtmEkgSV:
	%0 = alloca %Something-RrXtmEkgSV*
	store %Something-RrXtmEkgSV* %s-RrXtmEkgSV, %Something-RrXtmEkgSV** %0
	%1 = load %Something-RrXtmEkgSV*, %Something-RrXtmEkgSV** %0
	%2 = getelementptr %Something-RrXtmEkgSV, %Something-RrXtmEkgSV* %1, i32 0, i32 0
	%3 = load %Something-RrXtmEkgSV*, %Something-RrXtmEkgSV** %0
	%4 = getelementptr %Something-RrXtmEkgSV, %Something-RrXtmEkgSV* %3, i32 0, i32 0
	%5 = load i64, i64* %4
	%6 = sext i32 1 to i64
	%7 = add i64 %5, %6
	store i64 %7, i64* %2
	ret void
}

define ccc %Something-RrXtmEkgSV @something-RrXtmEkgSV() {
something-RrXtmEkgSV:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.fLDiNqBTme, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-RrXtmEkgSV
	%4 = sext i32 0 to i64
	%5 = getelementptr %Something-RrXtmEkgSV, %Something-RrXtmEkgSV* %3, i32 0, i32 0
	store i64 %4, i64* %5
	%6 = alloca %Something-RrXtmEkgSV
	%7 = load %Something-RrXtmEkgSV, %Something-RrXtmEkgSV* %3
	store %Something-RrXtmEkgSV %7, %Something-RrXtmEkgSV* %6
	%8 = alloca %Something-RrXtmEkgSV*
	store %Something-RrXtmEkgSV* %6, %Something-RrXtmEkgSV** %8
	%9 = load %Something-RrXtmEkgSV*, %Something-RrXtmEkgSV** %8
	call void @outsideFunction-RrXtmEkgSV(%Something-RrXtmEkgSV* %9)
	%10 = load %Something-RrXtmEkgSV, %Something-RrXtmEkgSV* %6
	ret %Something-RrXtmEkgSV %10
}

define ccc void @outsideFunction-hDdwdxkkPY(%Something-hDdwdxkkPY* %s-hDdwdxkkPY) {
outsideFunction-hDdwdxkkPY:
	%0 = alloca %Something-hDdwdxkkPY*
	store %Something-hDdwdxkkPY* %s-hDdwdxkkPY, %Something-hDdwdxkkPY** %0
	%1 = load %Something-hDdwdxkkPY*, %Something-hDdwdxkkPY** %0
	%2 = getelementptr %Something-hDdwdxkkPY, %Something-hDdwdxkkPY* %1, i32 0, i32 0
	%3 = load %Something-hDdwdxkkPY*, %Something-hDdwdxkkPY** %0
	%4 = getelementptr %Something-hDdwdxkkPY, %Something-hDdwdxkkPY* %3, i32 0, i32 0
	%5 = load i32, i32* %4
	%6 = add i32 %5, 1
	store i32 %6, i32* %2
	ret void
}

define ccc %Something-hDdwdxkkPY @something-hDdwdxkkPY() {
something-hDdwdxkkPY:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.emSlswNwtv, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-hDdwdxkkPY
	%4 = getelementptr %Something-hDdwdxkkPY, %Something-hDdwdxkkPY* %3, i32 0, i32 0
	store i32 0, i32* %4
	%5 = alloca %Something-hDdwdxkkPY
	%6 = load %Something-hDdwdxkkPY, %Something-hDdwdxkkPY* %3
	store %Something-hDdwdxkkPY %6, %Something-hDdwdxkkPY* %5
	%7 = alloca %Something-hDdwdxkkPY*
	store %Something-hDdwdxkkPY* %5, %Something-hDdwdxkkPY** %7
	%8 = load %Something-hDdwdxkkPY*, %Something-hDdwdxkkPY** %7
	call void @outsideFunction-hDdwdxkkPY(%Something-hDdwdxkkPY* %8)
	%9 = load %Something-hDdwdxkkPY, %Something-hDdwdxkkPY* %5
	ret %Something-hDdwdxkkPY %9
}

define ccc %Thing-MLrehTcaVM* @create-CWobWCZbkz() {
create-CWobWCZbkz:
	%0 = sext i32 1 to i64
	%1 = mul i64 %0, 4
	%2 = call i8* @malloc(i64 %1)
	%3 = bitcast i8* %2 to %Thing-MLrehTcaVM*
	%4 = alloca %Thing-MLrehTcaVM*
	store %Thing-MLrehTcaVM* %3, %Thing-MLrehTcaVM** %4
	%5 = load %Thing-MLrehTcaVM*, %Thing-MLrehTcaVM** %4
	ret %Thing-MLrehTcaVM* %5
}

define ccc i32 @call-TqcnFiFmrz(i32 (i32)* %f-TqcnFiFmrz) {
call-TqcnFiFmrz:
	%0 = alloca i32 (i32)*
	store i32 (i32)* %f-TqcnFiFmrz, i32 (i32)** %0
	%1 = load i32 (i32)*, i32 (i32)** %0
	%2 = call i32 %1(i32 1)
	ret i32 %2
}

define ccc %Something-RrXtmEkgSV @something-MLrehTcaVM() {
something-MLrehTcaVM:
	%0 = alloca %Something-RrXtmEkgSV
	%1 = sext i32 3 to i64
	%2 = getelementptr %Something-RrXtmEkgSV, %Something-RrXtmEkgSV* %0, i32 0, i32 0
	store i64 %1, i64* %2
	%3 = load %Something-RrXtmEkgSV, %Something-RrXtmEkgSV* %0
	ret %Something-RrXtmEkgSV %3
}

define ccc i32 @wow-MLrehTcaVM(i32 %i-MLrehTcaVM) {
wow-MLrehTcaVM:
	%0 = alloca i32
	store i32 %i-MLrehTcaVM, i32* %0
	%1 = load i32, i32* %0
	ret i32 %1
}

define ccc void @main() {
main:
	%0 = call i32 @call-TqcnFiFmrz(i32 (i32)* @wow-MLrehTcaVM)
	%1 = alloca i32
	store i32 %0, i32* %1
	%2 = load i32, i32* %1
	%3 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%4 = call i32 (i8*, ...) @printf(i8* %3, i32 %2)
	%5 = call %Thing-MLrehTcaVM* @create-CWobWCZbkz()
	%6 = alloca %Thing-MLrehTcaVM*
	store %Thing-MLrehTcaVM* %5, %Thing-MLrehTcaVM** %6
	%7 = load %Thing-MLrehTcaVM*, %Thing-MLrehTcaVM** %6
	%8 = getelementptr %Thing-MLrehTcaVM, %Thing-MLrehTcaVM* %7, i32 0, i32 0
	%9 = load i32, i32* %8
	%10 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%11 = call i32 (i8*, ...) @printf(i8* %10, i32 %9)
	%12 = call %Something-RrXtmEkgSV @something-RrXtmEkgSV()
	%13 = alloca %Something-RrXtmEkgSV
	store %Something-RrXtmEkgSV %12, %Something-RrXtmEkgSV* %13
	%14 = call %Something-hDdwdxkkPY @something-hDdwdxkkPY()
	%15 = alloca %Something-hDdwdxkkPY
	store %Something-hDdwdxkkPY %14, %Something-hDdwdxkkPY* %15
	%16 = call %Something-RrXtmEkgSV @something-MLrehTcaVM()
	%17 = alloca %Something-RrXtmEkgSV
	store %Something-RrXtmEkgSV %16, %Something-RrXtmEkgSV* %17
	%18 = getelementptr %Something-hDdwdxkkPY, %Something-hDdwdxkkPY* %15, i32 0, i32 0
	%19 = load i32, i32* %18
	%20 = getelementptr %Something-RrXtmEkgSV, %Something-RrXtmEkgSV* %13, i32 0, i32 0
	%21 = load i64, i64* %20
	%22 = trunc i64 %21 to i32
	%23 = add i32 %19, %22
	%24 = getelementptr %Something-RrXtmEkgSV, %Something-RrXtmEkgSV* %17, i32 0, i32 0
	%25 = load i64, i64* %24
	%26 = trunc i64 %25 to i32
	%27 = add i32 %23, %26
	%28 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%29 = call i32 (i8*, ...) @printf(i8* %28, i32 %27)
	ret void
}
