%aStruct-qLFzNHNsIk = type { i32, i64 }

@"%u " = global [4 x i8] c"%u \00"
@"%d " = global [4 x i8] c"%d \00"
@string.literal.axFpfngjKh = global [21 x i8] c"this shouldn't print\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.aYTarGVtRT = global [30 x i8] c"this should not print as well\00"
@string.literal.IUAyppssnW = global [25 x i8] c"this shouldnt print!!!!!\00"

define i32 @loadByPointer-qLFzNHNsIk(i32* %i-qLFzNHNsIk) {
loadByPointer-qLFzNHNsIk:
	%0 = alloca i32*
	store i32* %i-qLFzNHNsIk, i32** %0
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
	br i1 %6, label %if.then.NBVOVOSEsA, label %if.else.TLHjArvOxV

if.then.NBVOVOSEsA:
	%7 = load i32, i32* %1
	%8 = getelementptr inbounds [4 x i8], [4 x i8]* @"%u ", i32 0, i32 0
	%9 = call i32 (i8*, ...) @printf(i8* %8, i32 %7)
	%10 = load i32, i32* %1
	%11 = getelementptr inbounds [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%12 = call i32 (i8*, ...) @printf(i8* %11, i32 %10)
	br label %lastLeave.JwJhjZgkyx

if.else.TLHjArvOxV:
	br label %lastLeave.JwJhjZgkyx

lastLeave.JwJhjZgkyx:
	%13 = alloca i32
	store i32 3, i32* %13
	%14 = alloca i32*
	store i32* %13, i32** %14
	%15 = load i32*, i32** %14
	%16 = call i32 @loadByPointer-qLFzNHNsIk(i32* %15)
	%17 = alloca i32
	store i32 %16, i32* %17
	%18 = load i32, i32* %17
	%19 = icmp ne i32 %18, 3
	br i1 %19, label %if.then.LhFSilcxiw, label %if.else.oLqBTmbFXJ

if.then.LhFSilcxiw:
	%20 = getelementptr inbounds [21 x i8], [21 x i8]* @string.literal.axFpfngjKh, i32 0, i32 0
	%21 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%22 = call i32 (i8*, ...) @printf(i8* %21, i8* %20)
	br label %lastLeave.wtzpkdjjWS

if.else.oLqBTmbFXJ:
	br label %lastLeave.wtzpkdjjWS

lastLeave.wtzpkdjjWS:
	%23 = alloca %aStruct-qLFzNHNsIk
	%24 = getelementptr inbounds %aStruct-qLFzNHNsIk, %aStruct-qLFzNHNsIk* %23, i32 0, i32 0
	store i32 3, i32* %24
	%25 = alloca %aStruct-qLFzNHNsIk*
	store %aStruct-qLFzNHNsIk* %23, %aStruct-qLFzNHNsIk** %25
	%26 = load %aStruct-qLFzNHNsIk*, %aStruct-qLFzNHNsIk** %25
	%27 = bitcast %aStruct-qLFzNHNsIk* %26 to i32*
	%28 = alloca i32*
	store i32* %27, i32** %28
	%29 = load i32*, i32** %28
	%30 = getelementptr inbounds i32, i32* %29, i32 0
	%31 = load i32, i32* %30
	%32 = getelementptr inbounds %aStruct-qLFzNHNsIk, %aStruct-qLFzNHNsIk* %23, i32 0, i32 0
	%33 = load i32, i32* %32
	%34 = icmp ne i32 %31, %33
	br i1 %34, label %if.then.BYMRpNUWpC, label %if.else.rdMFeyGrGI

if.then.BYMRpNUWpC:
	%35 = getelementptr inbounds [30 x i8], [30 x i8]* @string.literal.aYTarGVtRT, i32 0, i32 0
	%36 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%37 = call i32 (i8*, ...) @printf(i8* %36, i8* %35)
	br label %lastLeave.BGlaWUFYQn

if.else.rdMFeyGrGI:
	br label %lastLeave.BGlaWUFYQn

lastLeave.BGlaWUFYQn:
	%38 = load i32*, i32** %28
	%39 = bitcast i32* %38 to %aStruct-qLFzNHNsIk*
	%40 = alloca %aStruct-qLFzNHNsIk
	%41 = load %aStruct-qLFzNHNsIk, %aStruct-qLFzNHNsIk* %39
	store %aStruct-qLFzNHNsIk %41, %aStruct-qLFzNHNsIk* %40
	%42 = getelementptr inbounds %aStruct-qLFzNHNsIk, %aStruct-qLFzNHNsIk* %40, i32 0, i32 0
	%43 = load i32, i32* %42
	%44 = getelementptr inbounds %aStruct-qLFzNHNsIk, %aStruct-qLFzNHNsIk* %23, i32 0, i32 0
	%45 = load i32, i32* %44
	%46 = icmp ne i32 %43, %45
	br i1 %46, label %if.then.QqEKSYGpsA, label %if.else.BDQiRHPTZl

if.then.QqEKSYGpsA:
	%47 = getelementptr inbounds [25 x i8], [25 x i8]* @string.literal.IUAyppssnW, i32 0, i32 0
	%48 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%49 = call i32 (i8*, ...) @printf(i8* %48, i8* %47)
	br label %lastLeave.uqrKiqEnVN

if.else.BDQiRHPTZl:
	br label %lastLeave.uqrKiqEnVN

lastLeave.uqrKiqEnVN:
	ret i32 0
}

declare ccc i32 @printf(i8* %0, ...)
