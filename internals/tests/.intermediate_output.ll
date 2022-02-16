%aStruct-GRvoEHBPEH = type { i32, i64 }

@"%u " = global [4 x i8] c"%u \00"
@"%d " = global [4 x i8] c"%d \00"
@string.literal.YoTMPPpTEE = global [21 x i8] c"this shouldn't print\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.kaLhRRYHWG = global [30 x i8] c"this should not print as well\00"
@string.literal.lgYhTsSVSB = global [25 x i8] c"this shouldnt print!!!!!\00"

declare ccc void @free(i8* %0, ...)

declare ccc i8* @realloc(i8* %0, i64 %1, ...)

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc i32 @loadByPointer-GRvoEHBPEH(i32* %i-GRvoEHBPEH) {
loadByPointer-GRvoEHBPEH:
	%0 = alloca i32*
	store i32* %i-GRvoEHBPEH, i32** %0
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
	br i1 %6, label %if.then.XcfYbFvPzo, label %if.else.EQYsxOIfds

if.then.XcfYbFvPzo:
	%7 = load i32, i32* %1
	%8 = getelementptr [4 x i8], [4 x i8]* @"%u ", i32 0, i32 0
	%9 = call i32 (i8*, ...) @printf(i8* %8, i32 %7)
	%10 = load i32, i32* %1
	%11 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%12 = call i32 (i8*, ...) @printf(i8* %11, i32 %10)
	br label %lastLeave.FOiSGFcPyM

if.else.EQYsxOIfds:
	br label %lastLeave.FOiSGFcPyM

lastLeave.FOiSGFcPyM:
	%13 = alloca i32
	store i32 3, i32* %13
	%14 = alloca i32*
	store i32* %13, i32** %14
	%15 = load i32*, i32** %14
	%16 = call i32 @loadByPointer-GRvoEHBPEH(i32* %15)
	%17 = alloca i32
	store i32 %16, i32* %17
	%18 = load i32, i32* %17
	%19 = icmp ne i32 %18, 3
	br i1 %19, label %if.then.ACOZBUbSBL, label %if.else.TwrBkAqufr

if.then.ACOZBUbSBL:
	%20 = getelementptr [21 x i8], [21 x i8]* @string.literal.YoTMPPpTEE, i32 0, i32 0
	%21 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%22 = call i32 (i8*, ...) @printf(i8* %21, i8* %20)
	br label %lastLeave.fNMoIEUxHC

if.else.TwrBkAqufr:
	br label %lastLeave.fNMoIEUxHC

lastLeave.fNMoIEUxHC:
	%23 = alloca %aStruct-GRvoEHBPEH
	%24 = getelementptr %aStruct-GRvoEHBPEH, %aStruct-GRvoEHBPEH* %23, i32 0, i32 0
	store i32 3, i32* %24
	%25 = alloca %aStruct-GRvoEHBPEH
	%26 = load %aStruct-GRvoEHBPEH, %aStruct-GRvoEHBPEH* %23
	store %aStruct-GRvoEHBPEH %26, %aStruct-GRvoEHBPEH* %25
	%27 = alloca %aStruct-GRvoEHBPEH*
	store %aStruct-GRvoEHBPEH* %25, %aStruct-GRvoEHBPEH** %27
	%28 = load %aStruct-GRvoEHBPEH*, %aStruct-GRvoEHBPEH** %27
	%29 = bitcast %aStruct-GRvoEHBPEH* %28 to i32*
	%30 = alloca i32*
	store i32* %29, i32** %30
	%31 = load i32*, i32** %30
	%32 = getelementptr i32, i32* %31, i32 0
	%33 = load i32, i32* %32
	%34 = getelementptr %aStruct-GRvoEHBPEH, %aStruct-GRvoEHBPEH* %25, i32 0, i32 0
	%35 = load i32, i32* %34
	%36 = icmp ne i32 %33, %35
	br i1 %36, label %if.then.GatcrpPpev, label %if.else.GClCxuFnei

if.then.GatcrpPpev:
	%37 = getelementptr [30 x i8], [30 x i8]* @string.literal.kaLhRRYHWG, i32 0, i32 0
	%38 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%39 = call i32 (i8*, ...) @printf(i8* %38, i8* %37)
	br label %lastLeave.eJIWPSnKfV

if.else.GClCxuFnei:
	br label %lastLeave.eJIWPSnKfV

lastLeave.eJIWPSnKfV:
	%40 = load i32*, i32** %30
	%41 = bitcast i32* %40 to %aStruct-GRvoEHBPEH*
	%42 = load %aStruct-GRvoEHBPEH, %aStruct-GRvoEHBPEH* %41
	%43 = alloca %aStruct-GRvoEHBPEH
	store %aStruct-GRvoEHBPEH %42, %aStruct-GRvoEHBPEH* %43
	%44 = getelementptr %aStruct-GRvoEHBPEH, %aStruct-GRvoEHBPEH* %43, i32 0, i32 0
	%45 = load i32, i32* %44
	%46 = getelementptr %aStruct-GRvoEHBPEH, %aStruct-GRvoEHBPEH* %25, i32 0, i32 0
	%47 = load i32, i32* %46
	%48 = icmp ne i32 %45, %47
	br i1 %48, label %if.then.qSEZFRcTxZ, label %if.else.JCcVXCbhle

if.then.qSEZFRcTxZ:
	%49 = getelementptr [25 x i8], [25 x i8]* @string.literal.lgYhTsSVSB, i32 0, i32 0
	%50 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%51 = call i32 (i8*, ...) @printf(i8* %50, i8* %49)
	br label %lastLeave.LpbKeprAgO

if.else.JCcVXCbhle:
	br label %lastLeave.LpbKeprAgO

lastLeave.LpbKeprAgO:
	ret i32 0
}
