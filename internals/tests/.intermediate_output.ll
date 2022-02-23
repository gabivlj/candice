%aStruct-wmDBQhqrKY = type { i32, i64 }

@"%u " = global [4 x i8] c"%u \00"
@"%d " = global [4 x i8] c"%d \00"
@string.literal.ximxVkCyIX = global [21 x i8] c"this shouldn't print\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.QRKhPWkuQy = global [30 x i8] c"this should not print as well\00"
@string.literal.LHkiAQSwMd = global [25 x i8] c"this shouldnt print!!!!!\00"

declare ccc void @free(i8* %0, ...)

declare ccc i8* @realloc(i8* %0, i64 %1, ...)

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc i32 @loadByPointer-wmDBQhqrKY(i32* %i-wmDBQhqrKY) {
loadByPointer-wmDBQhqrKY:
	%0 = alloca i32*
	store i32* %i-wmDBQhqrKY, i32** %0
	%1 = load i32*, i32** %0
	%2 = load i32, i32* %1
	ret i32 %2
}

define ccc i32 @main() {
main:
	%0 = alloca i32
	store i32 u0xFFFFFF3C, i32* %0
	%1 = load i32, i32* %0
	%2 = add i32 %1, 100
	store i32 %2, i32* %0
	%3 = load i32, i32* %0
	%4 = load i32, i32* %0
	%5 = icmp eq i32 %3, %4
	br i1 %5, label %if.then.CDuDOqtPxW, label %if.else.tEjMwpNjKk

if.then.CDuDOqtPxW:
	%6 = load i32, i32* %0
	%7 = getelementptr [4 x i8], [4 x i8]* @"%u ", i32 0, i32 0
	%8 = call i32 (i8*, ...) @printf(i8* %7, i32 %6)
	%9 = load i32, i32* %0
	%10 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%11 = call i32 (i8*, ...) @printf(i8* %10, i32 %9)
	br label %lastLeave.zcxCTYtOTU

if.else.tEjMwpNjKk:
	br label %lastLeave.zcxCTYtOTU

lastLeave.zcxCTYtOTU:
	%12 = alloca i32
	store i32 3, i32* %12
	%13 = alloca i32*
	store i32* %12, i32** %13
	%14 = load i32*, i32** %13
	%15 = call i32 @loadByPointer-wmDBQhqrKY(i32* %14)
	%16 = alloca i32
	store i32 %15, i32* %16
	%17 = load i32, i32* %16
	%18 = icmp ne i32 %17, 3
	br i1 %18, label %if.then.lSjLkJqciA, label %if.else.WWKtabrxzv

if.then.lSjLkJqciA:
	%19 = getelementptr [21 x i8], [21 x i8]* @string.literal.ximxVkCyIX, i32 0, i32 0
	%20 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%21 = call i32 (i8*, ...) @printf(i8* %20, i8* %19)
	br label %lastLeave.zqhYDxttZO

if.else.WWKtabrxzv:
	br label %lastLeave.zqhYDxttZO

lastLeave.zqhYDxttZO:
	%22 = alloca %aStruct-wmDBQhqrKY
	%23 = getelementptr %aStruct-wmDBQhqrKY, %aStruct-wmDBQhqrKY* %22, i32 0, i32 0
	store i32 3, i32* %23
	%24 = alloca %aStruct-wmDBQhqrKY
	%25 = load %aStruct-wmDBQhqrKY, %aStruct-wmDBQhqrKY* %22
	store %aStruct-wmDBQhqrKY %25, %aStruct-wmDBQhqrKY* %24
	%26 = alloca %aStruct-wmDBQhqrKY*
	store %aStruct-wmDBQhqrKY* %24, %aStruct-wmDBQhqrKY** %26
	%27 = load %aStruct-wmDBQhqrKY*, %aStruct-wmDBQhqrKY** %26
	%28 = bitcast %aStruct-wmDBQhqrKY* %27 to i32*
	%29 = alloca i32*
	store i32* %28, i32** %29
	%30 = load i32*, i32** %29
	%31 = getelementptr i32, i32* %30, i32 0
	%32 = load i32, i32* %31
	%33 = getelementptr %aStruct-wmDBQhqrKY, %aStruct-wmDBQhqrKY* %24, i32 0, i32 0
	%34 = load i32, i32* %33
	%35 = icmp ne i32 %32, %34
	br i1 %35, label %if.then.MnLCdImBan, label %if.else.wfdELTXhcO

if.then.MnLCdImBan:
	%36 = getelementptr [30 x i8], [30 x i8]* @string.literal.QRKhPWkuQy, i32 0, i32 0
	%37 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%38 = call i32 (i8*, ...) @printf(i8* %37, i8* %36)
	br label %lastLeave.hXOlbQjWbe

if.else.wfdELTXhcO:
	br label %lastLeave.hXOlbQjWbe

lastLeave.hXOlbQjWbe:
	%39 = load i32*, i32** %29
	%40 = bitcast i32* %39 to %aStruct-wmDBQhqrKY*
	%41 = load %aStruct-wmDBQhqrKY, %aStruct-wmDBQhqrKY* %40
	%42 = alloca %aStruct-wmDBQhqrKY
	store %aStruct-wmDBQhqrKY %41, %aStruct-wmDBQhqrKY* %42
	%43 = getelementptr %aStruct-wmDBQhqrKY, %aStruct-wmDBQhqrKY* %42, i32 0, i32 0
	%44 = load i32, i32* %43
	%45 = getelementptr %aStruct-wmDBQhqrKY, %aStruct-wmDBQhqrKY* %24, i32 0, i32 0
	%46 = load i32, i32* %45
	%47 = icmp ne i32 %44, %46
	br i1 %47, label %if.then.lqrUcBBtww, label %if.else.cITBiVWlIK

if.then.lqrUcBBtww:
	%48 = getelementptr [25 x i8], [25 x i8]* @string.literal.LHkiAQSwMd, i32 0, i32 0
	%49 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%50 = call i32 (i8*, ...) @printf(i8* %49, i8* %48)
	br label %lastLeave.qBBaiqwqcW

if.else.cITBiVWlIK:
	br label %lastLeave.qBBaiqwqcW

lastLeave.qBBaiqwqcW:
	ret i32 0
}
