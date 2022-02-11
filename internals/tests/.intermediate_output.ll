%aStruct-eYykyHrkWf = type { i32, i64 }

@"%u " = global [4 x i8] c"%u \00"
@"%d " = global [4 x i8] c"%d \00"
@string.literal.oCYeSKpaOE = global [21 x i8] c"this shouldn't print\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.KxwgohiaSs = global [30 x i8] c"this should not print as well\00"
@string.literal.OYHLsLXmJh = global [25 x i8] c"this shouldnt print!!!!!\00"

declare ccc void @free(i8* %0, ...)

declare ccc i8* @realloc(i8* %0, i64 %1, ...)

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc i32 @loadByPointer-eYykyHrkWf(i32* %i-eYykyHrkWf) {
loadByPointer-eYykyHrkWf:
	%0 = alloca i32*
	store i32* %i-eYykyHrkWf, i32** %0
	%1 = load i32*, i32** %0
	%2 = load i32, i32* %1
	ret i32 %2
}

define ccc i32 @main() {
main:
	%0 = trunc i64 u0xFFFFFF3C to i32
	%1 = alloca i32
	store i32 %0, i32* %1
	%2 = load i32, i32* %1
	%3 = add i32 %2, 100
	store i32 %3, i32* %1
	%4 = load i32, i32* %1
	%5 = load i32, i32* %1
	%6 = icmp eq i32 %4, %5
	br i1 %6, label %if.then.vlHSYojuUi, label %if.else.qdASukkvJI

if.then.vlHSYojuUi:
	%7 = load i32, i32* %1
	%8 = getelementptr [4 x i8], [4 x i8]* @"%u ", i32 0, i32 0
	%9 = call i32 (i8*, ...) @printf(i8* %8, i32 %7)
	%10 = load i32, i32* %1
	%11 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%12 = call i32 (i8*, ...) @printf(i8* %11, i32 %10)
	br label %lastLeave.FrxaTmeljc

if.else.qdASukkvJI:
	br label %lastLeave.FrxaTmeljc

lastLeave.FrxaTmeljc:
	%13 = alloca i32
	store i32 3, i32* %13
	%14 = alloca i32*
	store i32* %13, i32** %14
	%15 = load i32*, i32** %14
	%16 = call i32 @loadByPointer-eYykyHrkWf(i32* %15)
	%17 = alloca i32
	store i32 %16, i32* %17
	%18 = load i32, i32* %17
	%19 = icmp ne i32 %18, 3
	br i1 %19, label %if.then.EzZVqkcBgl, label %if.else.oBrbNPUtKU

if.then.EzZVqkcBgl:
	%20 = getelementptr [21 x i8], [21 x i8]* @string.literal.oCYeSKpaOE, i32 0, i32 0
	%21 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%22 = call i32 (i8*, ...) @printf(i8* %21, i8* %20)
	br label %lastLeave.cwKxLphWkC

if.else.oBrbNPUtKU:
	br label %lastLeave.cwKxLphWkC

lastLeave.cwKxLphWkC:
	%23 = alloca %aStruct-eYykyHrkWf
	%24 = getelementptr %aStruct-eYykyHrkWf, %aStruct-eYykyHrkWf* %23, i32 0, i32 0
	store i32 3, i32* %24
	%25 = alloca %aStruct-eYykyHrkWf
	%26 = load %aStruct-eYykyHrkWf, %aStruct-eYykyHrkWf* %23
	store %aStruct-eYykyHrkWf %26, %aStruct-eYykyHrkWf* %25
	%27 = alloca %aStruct-eYykyHrkWf*
	store %aStruct-eYykyHrkWf* %25, %aStruct-eYykyHrkWf** %27
	%28 = load %aStruct-eYykyHrkWf*, %aStruct-eYykyHrkWf** %27
	%29 = bitcast %aStruct-eYykyHrkWf* %28 to i32*
	%30 = alloca i32*
	store i32* %29, i32** %30
	%31 = load i32*, i32** %30
	%32 = getelementptr i32, i32* %31, i32 0
	%33 = load i32, i32* %32
	%34 = getelementptr %aStruct-eYykyHrkWf, %aStruct-eYykyHrkWf* %25, i32 0, i32 0
	%35 = load i32, i32* %34
	%36 = icmp ne i32 %33, %35
	br i1 %36, label %if.then.hKslaepiZM, label %if.else.batEYuPpUL

if.then.hKslaepiZM:
	%37 = getelementptr [30 x i8], [30 x i8]* @string.literal.KxwgohiaSs, i32 0, i32 0
	%38 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%39 = call i32 (i8*, ...) @printf(i8* %38, i8* %37)
	br label %lastLeave.wndabUvLzG

if.else.batEYuPpUL:
	br label %lastLeave.wndabUvLzG

lastLeave.wndabUvLzG:
	%40 = load i32*, i32** %30
	%41 = bitcast i32* %40 to %aStruct-eYykyHrkWf*
	%42 = load %aStruct-eYykyHrkWf, %aStruct-eYykyHrkWf* %41
	%43 = alloca %aStruct-eYykyHrkWf
	store %aStruct-eYykyHrkWf %42, %aStruct-eYykyHrkWf* %43
	%44 = getelementptr %aStruct-eYykyHrkWf, %aStruct-eYykyHrkWf* %43, i32 0, i32 0
	%45 = load i32, i32* %44
	%46 = getelementptr %aStruct-eYykyHrkWf, %aStruct-eYykyHrkWf* %25, i32 0, i32 0
	%47 = load i32, i32* %46
	%48 = icmp ne i32 %45, %47
	br i1 %48, label %if.then.WPWZhveNhw, label %if.else.uficPPpgYI

if.then.WPWZhveNhw:
	%49 = getelementptr [25 x i8], [25 x i8]* @string.literal.OYHLsLXmJh, i32 0, i32 0
	%50 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%51 = call i32 (i8*, ...) @printf(i8* %50, i8* %49)
	br label %lastLeave.LvEyyJLovH

if.else.uficPPpgYI:
	br label %lastLeave.LvEyyJLovH

lastLeave.LvEyyJLovH:
	ret i32 0
}
