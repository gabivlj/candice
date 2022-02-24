%aStruct-BzjsOmawei = type { i32, i64 }

@"%u " = global [4 x i8] c"%u \00"
@"%d " = global [4 x i8] c"%d \00"
@string.literal.oOHLmNaChJ = global [21 x i8] c"this shouldn't print\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.RKcpdciWib = global [30 x i8] c"this should not print as well\00"
@string.literal.vjLYeGIzLO = global [25 x i8] c"this shouldnt print!!!!!\00"

declare ccc void @free(i8* %0, ...)

declare ccc i8* @realloc(i8* %0, i64 %1, ...)

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define i32 @loadByPointer-BzjsOmawei(i32* %i-BzjsOmawei) {
loadByPointer-BzjsOmawei:
	%0 = alloca i32*
	store i32* %i-BzjsOmawei, i32** %0
	%1 = load i32*, i32** %0
	%2 = load i32, i32* %1
	ret i32 %2
}

define i32 @main() {
main:
	%0 = alloca i32
	store i32 u0xFFFFFF3C, i32* %0
	%1 = load i32, i32* %0
	%2 = add i32 %1, 100
	store i32 %2, i32* %0
	%3 = load i32, i32* %0
	%4 = load i32, i32* %0
	%5 = icmp eq i32 %3, %4
	br i1 %5, label %if.then.tjYBGelBLI, label %if.else.mjEmdsEGdM

if.then.tjYBGelBLI:
	%6 = load i32, i32* %0
	%7 = getelementptr [4 x i8], [4 x i8]* @"%u ", i32 0, i32 0
	%8 = call i32 (i8*, ...) @printf(i8* %7, i32 %6)
	%9 = load i32, i32* %0
	%10 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%11 = call i32 (i8*, ...) @printf(i8* %10, i32 %9)
	br label %lastLeave.KtWtPnYTuz

if.else.mjEmdsEGdM:
	br label %lastLeave.KtWtPnYTuz

lastLeave.KtWtPnYTuz:
	%12 = alloca i32
	store i32 3, i32* %12
	%13 = alloca i32*
	store i32* %12, i32** %13
	%14 = load i32*, i32** %13
	%15 = call i32 @loadByPointer-BzjsOmawei(i32* %14)
	%16 = alloca i32
	store i32 %15, i32* %16
	%17 = load i32, i32* %16
	%18 = icmp ne i32 %17, 3
	br i1 %18, label %if.then.YzERIHpyMN, label %if.else.YorkmeLOwZ

if.then.YzERIHpyMN:
	%19 = getelementptr [21 x i8], [21 x i8]* @string.literal.oOHLmNaChJ, i32 0, i32 0
	%20 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%21 = call i32 (i8*, ...) @printf(i8* %20, i8* %19)
	br label %lastLeave.RHNbUtDzlc

if.else.YorkmeLOwZ:
	br label %lastLeave.RHNbUtDzlc

lastLeave.RHNbUtDzlc:
	%22 = alloca %aStruct-BzjsOmawei
	%23 = getelementptr %aStruct-BzjsOmawei, %aStruct-BzjsOmawei* %22, i32 0, i32 0
	store i32 3, i32* %23
	%24 = alloca %aStruct-BzjsOmawei
	%25 = load %aStruct-BzjsOmawei, %aStruct-BzjsOmawei* %22
	store %aStruct-BzjsOmawei %25, %aStruct-BzjsOmawei* %24
	%26 = alloca %aStruct-BzjsOmawei*
	store %aStruct-BzjsOmawei* %24, %aStruct-BzjsOmawei** %26
	%27 = load %aStruct-BzjsOmawei*, %aStruct-BzjsOmawei** %26
	%28 = bitcast %aStruct-BzjsOmawei* %27 to i32*
	%29 = alloca i32*
	store i32* %28, i32** %29
	%30 = load i32*, i32** %29
	%31 = getelementptr i32, i32* %30, i32 0
	%32 = load i32, i32* %31
	%33 = getelementptr %aStruct-BzjsOmawei, %aStruct-BzjsOmawei* %24, i32 0, i32 0
	%34 = load i32, i32* %33
	%35 = icmp ne i32 %32, %34
	br i1 %35, label %if.then.hLlZGQTuEt, label %if.else.kpqyxIaZSm

if.then.hLlZGQTuEt:
	%36 = getelementptr [30 x i8], [30 x i8]* @string.literal.RKcpdciWib, i32 0, i32 0
	%37 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%38 = call i32 (i8*, ...) @printf(i8* %37, i8* %36)
	br label %lastLeave.damgAuAkzf

if.else.kpqyxIaZSm:
	br label %lastLeave.damgAuAkzf

lastLeave.damgAuAkzf:
	%39 = load i32*, i32** %29
	%40 = bitcast i32* %39 to %aStruct-BzjsOmawei*
	%41 = load %aStruct-BzjsOmawei, %aStruct-BzjsOmawei* %40
	%42 = alloca %aStruct-BzjsOmawei
	store %aStruct-BzjsOmawei %41, %aStruct-BzjsOmawei* %42
	%43 = getelementptr %aStruct-BzjsOmawei, %aStruct-BzjsOmawei* %42, i32 0, i32 0
	%44 = load i32, i32* %43
	%45 = getelementptr %aStruct-BzjsOmawei, %aStruct-BzjsOmawei* %24, i32 0, i32 0
	%46 = load i32, i32* %45
	%47 = icmp ne i32 %44, %46
	br i1 %47, label %if.then.elOMbquAAk, label %if.else.SdpHpDyDBG

if.then.elOMbquAAk:
	%48 = getelementptr [25 x i8], [25 x i8]* @string.literal.vjLYeGIzLO, i32 0, i32 0
	%49 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%50 = call i32 (i8*, ...) @printf(i8* %49, i8* %48)
	br label %lastLeave.dWHMBcXpvy

if.else.SdpHpDyDBG:
	br label %lastLeave.dWHMBcXpvy

lastLeave.dWHMBcXpvy:
	ret i32 0
}
