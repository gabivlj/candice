%aStruct = type { i32, i64 }

@"%u " = global [4 x i8] c"%u \00"
@"%d " = global [4 x i8] c"%d \00"
@string.literal.WEmaSGafXz = global [21 x i8] c"this shouldn't print\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.kcBqqRvAcQ = global [30 x i8] c"this should not print as well\00"
@string.literal.iYzYIJFYJI = global [25 x i8] c"this shouldnt print!!!!!\00"

define i32 @main() {
_main:
	%0 = alloca i32
	store i32 u0xFFFFFF3C, i32* %0
	%1 = load i32, i32* %0
	%2 = add i32 %1, 100
	store i32 %2, i32* %0
	%3 = load i32, i32* %0
	%4 = load i32, i32* %0
	%5 = icmp eq i32 %3, %4
	br i1 %5, label %if.then.WepMYCOMBM, label %if.else.ZmTGScXLPi

if.then.WepMYCOMBM:
	%6 = load i32, i32* %0
	%7 = getelementptr [4 x i8], [4 x i8]* @"%u ", i32 0, i32 0
	%8 = call i32 (i8*, ...) @printf(i8* %7, i32 %6)
	%9 = load i32, i32* %0
	%10 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%11 = call i32 (i8*, ...) @printf(i8* %10, i32 %9)
	br label %leave.UpubIemDMP

if.else.ZmTGScXLPi:
	br label %leave.UpubIemDMP

leave.UpubIemDMP:
	%12 = alloca i32
	store i32 3, i32* %12
	%13 = alloca i32*
	store i32* %12, i32** %13
	%14 = load i32*, i32** %13
	%15 = call i32 @loadByPointer(i32* %14)
	%16 = alloca i32
	store i32 %15, i32* %16
	%17 = load i32, i32* %16
	%18 = icmp ne i32 %17, 3
	br i1 %18, label %if.then.ZPJisfZgKV, label %if.else.PLMjixdTPX

if.then.ZPJisfZgKV:
	%19 = getelementptr [21 x i8], [21 x i8]* @string.literal.WEmaSGafXz, i32 0, i32 0
	%20 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%21 = call i32 (i8*, ...) @printf(i8* %20, i8* %19)
	br label %leave.AweJlROqcF

if.else.PLMjixdTPX:
	br label %leave.AweJlROqcF

leave.AweJlROqcF:
	%22 = alloca %aStruct
	%23 = getelementptr %aStruct, %aStruct* %22, i32 0, i32 0
	store i32 3, i32* %23
	%24 = alloca %aStruct*
	store %aStruct* %22, %aStruct** %24
	%25 = load %aStruct*, %aStruct** %24
	%26 = bitcast %aStruct* %25 to i32*
	%27 = alloca i32*
	store i32* %26, i32** %27
	%28 = load i32*, i32** %27
	%29 = getelementptr i32, i32* %28, i32 0
	%30 = load i32, i32* %29
	%31 = getelementptr %aStruct, %aStruct* %22, i32 0, i32 0
	%32 = load i32, i32* %31
	%33 = icmp ne i32 %30, %32
	br i1 %33, label %if.then.RqiMERvUea, label %if.else.mUrSzufsQq

if.then.RqiMERvUea:
	%34 = getelementptr [30 x i8], [30 x i8]* @string.literal.kcBqqRvAcQ, i32 0, i32 0
	%35 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%36 = call i32 (i8*, ...) @printf(i8* %35, i8* %34)
	br label %leave.jQIKcqhsIL

if.else.mUrSzufsQq:
	br label %leave.jQIKcqhsIL

leave.jQIKcqhsIL:
	%37 = load i32*, i32** %27
	%38 = bitcast i32* %37 to %aStruct*
	%39 = load %aStruct, %aStruct* %38
	%40 = alloca %aStruct
	store %aStruct %39, %aStruct* %40
	%41 = getelementptr %aStruct, %aStruct* %40, i32 0, i32 0
	%42 = load i32, i32* %41
	%43 = getelementptr %aStruct, %aStruct* %22, i32 0, i32 0
	%44 = load i32, i32* %43
	%45 = icmp ne i32 %42, %44
	br i1 %45, label %if.then.eXPXJinfEd, label %if.else.EFrKctbshv

if.then.eXPXJinfEd:
	%46 = getelementptr [25 x i8], [25 x i8]* @string.literal.iYzYIJFYJI, i32 0, i32 0
	%47 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%48 = call i32 (i8*, ...) @printf(i8* %47, i8* %46)
	br label %leave.BIkBxhPEtb

if.else.EFrKctbshv:
	br label %leave.BIkBxhPEtb

leave.BIkBxhPEtb:
	ret i32 0
}

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc i32 @loadByPointer(i32* %i) {
loadByPointer:
	%0 = alloca i32*
	store i32* %i, i32** %0
	%1 = load i32*, i32** %0
	%2 = load i32, i32* %1
	ret i32 %2
}
