%aStruct-ERIDDHUFbI = type { i32, i64 }

@"%u " = global [4 x i8] c"%u \00"
@"%d " = global [4 x i8] c"%d \00"
@string.literal.QFcfnwQTuK = global [21 x i8] c"this shouldn't print\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.TQVxERNlAF = global [30 x i8] c"this should not print as well\00"
@string.literal.fBVlOsPypL = global [25 x i8] c"this shouldnt print!!!!!\00"

define i32 @loadByPointer-ERIDDHUFbI(i32* %i-ERIDDHUFbI) {
loadByPointer-ERIDDHUFbI:
	%0 = alloca i32*
	store i32* %i-ERIDDHUFbI, i32** %0
	%1 = load i32*, i32** %0
	%2 = load i32, i32* %1
	ret i32 %2
}

define i32 @main() {
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
	br i1 %6, label %if.then.sLJvVlSpnq, label %if.else.GmPFUeLMOJ

if.then.sLJvVlSpnq:
	%7 = load i32, i32* %1
	%8 = getelementptr inbounds [4 x i8], [4 x i8]* @"%u ", i32 0, i32 0
	%9 = call i32 (i8*, ...) @printf(i8* %8, i32 %7)
	%10 = load i32, i32* %1
	%11 = getelementptr inbounds [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%12 = call i32 (i8*, ...) @printf(i8* %11, i32 %10)
	br label %lastLeave.FvACuKJjmR

if.else.GmPFUeLMOJ:
	br label %lastLeave.FvACuKJjmR

lastLeave.FvACuKJjmR:
	%13 = alloca i32
	store i32 3, i32* %13
	%14 = alloca i32*
	store i32* %13, i32** %14
	%15 = load i32*, i32** %14
	%16 = call i32 @loadByPointer-ERIDDHUFbI(i32* %15)
	%17 = alloca i32
	store i32 %16, i32* %17
	%18 = load i32, i32* %17
	%19 = icmp ne i32 %18, 3
	br i1 %19, label %if.then.fSzgbSetXc, label %if.else.DTOOhMakoD

if.then.fSzgbSetXc:
	%20 = getelementptr inbounds [21 x i8], [21 x i8]* @string.literal.QFcfnwQTuK, i32 0, i32 0
	%21 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%22 = call i32 (i8*, ...) @printf(i8* %21, i8* %20)
	br label %lastLeave.oCnisETddk

if.else.DTOOhMakoD:
	br label %lastLeave.oCnisETddk

lastLeave.oCnisETddk:
	%23 = alloca %aStruct-ERIDDHUFbI
	%24 = getelementptr inbounds %aStruct-ERIDDHUFbI, %aStruct-ERIDDHUFbI* %23, i32 0, i32 0
	store i32 3, i32* %24
	%25 = alloca %aStruct-ERIDDHUFbI*
	store %aStruct-ERIDDHUFbI* %23, %aStruct-ERIDDHUFbI** %25
	%26 = load %aStruct-ERIDDHUFbI*, %aStruct-ERIDDHUFbI** %25
	%27 = bitcast %aStruct-ERIDDHUFbI* %26 to i32*
	%28 = alloca i32*
	store i32* %27, i32** %28
	%29 = load i32*, i32** %28
	%30 = getelementptr inbounds i32, i32* %29, i32 0
	%31 = load i32, i32* %30
	%32 = getelementptr inbounds %aStruct-ERIDDHUFbI, %aStruct-ERIDDHUFbI* %23, i32 0, i32 0
	%33 = load i32, i32* %32
	%34 = icmp ne i32 %31, %33
	br i1 %34, label %if.then.RfLReRJFiT, label %if.else.ESKkrrOgnw

if.then.RfLReRJFiT:
	%35 = getelementptr inbounds [30 x i8], [30 x i8]* @string.literal.TQVxERNlAF, i32 0, i32 0
	%36 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%37 = call i32 (i8*, ...) @printf(i8* %36, i8* %35)
	br label %lastLeave.XQffbYowHC

if.else.ESKkrrOgnw:
	br label %lastLeave.XQffbYowHC

lastLeave.XQffbYowHC:
	%38 = load i32*, i32** %28
	%39 = bitcast i32* %38 to %aStruct-ERIDDHUFbI*
	%40 = alloca %aStruct-ERIDDHUFbI
	%41 = load %aStruct-ERIDDHUFbI, %aStruct-ERIDDHUFbI* %39
	store %aStruct-ERIDDHUFbI %41, %aStruct-ERIDDHUFbI* %40
	%42 = getelementptr inbounds %aStruct-ERIDDHUFbI, %aStruct-ERIDDHUFbI* %40, i32 0, i32 0
	%43 = load i32, i32* %42
	%44 = getelementptr inbounds %aStruct-ERIDDHUFbI, %aStruct-ERIDDHUFbI* %23, i32 0, i32 0
	%45 = load i32, i32* %44
	%46 = icmp ne i32 %43, %45
	br i1 %46, label %if.then.QiCdxBLdPZ, label %if.else.JLPZBaQngf

if.then.QiCdxBLdPZ:
	%47 = getelementptr inbounds [25 x i8], [25 x i8]* @string.literal.fBVlOsPypL, i32 0, i32 0
	%48 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%49 = call i32 (i8*, ...) @printf(i8* %48, i8* %47)
	br label %lastLeave.ZSCsyRdYHg

if.else.JLPZBaQngf:
	br label %lastLeave.ZSCsyRdYHg

lastLeave.ZSCsyRdYHg:
	ret i32 0
}

declare ccc i32 @printf(i8* %0, ...)
