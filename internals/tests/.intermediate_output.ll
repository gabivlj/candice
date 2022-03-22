%aStruct-fTwCvMHKtQ = type { i32, i64 }

@"%u " = global [4 x i8] c"%u \00"
@"%d " = global [4 x i8] c"%d \00"
@string.literal.IjwgowylaE = global [21 x i8] c"this shouldn't print\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.aCqaPfgDdq = global [30 x i8] c"this should not print as well\00"
@string.literal.xvjQpdbTop = global [25 x i8] c"this shouldnt print!!!!!\00"

define i32 @loadByPointer-fTwCvMHKtQ(i32* %i-fTwCvMHKtQ) {
loadByPointer-fTwCvMHKtQ:
	%0 = alloca i32*
	store i32* %i-fTwCvMHKtQ, i32** %0
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
	br i1 %5, label %if.then.ljkqWIvIFC, label %if.else.RPGOxyebSh

if.then.ljkqWIvIFC:
	%6 = load i32, i32* %0
	%7 = getelementptr inbounds [4 x i8], [4 x i8]* @"%u ", i32 0, i32 0
	%8 = call i32 (i8*, ...) @printf(i8* %7, i32 %6)
	%9 = load i32, i32* %0
	%10 = getelementptr inbounds [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%11 = call i32 (i8*, ...) @printf(i8* %10, i32 %9)
	br label %lastLeave.WJMYiGFpgu

if.else.RPGOxyebSh:
	br label %lastLeave.WJMYiGFpgu

lastLeave.WJMYiGFpgu:
	%12 = alloca i32
	store i32 3, i32* %12
	%13 = alloca i32*
	store i32* %12, i32** %13
	%14 = load i32*, i32** %13
	%15 = call i32 @loadByPointer-fTwCvMHKtQ(i32* %14)
	%16 = alloca i32
	store i32 %15, i32* %16
	%17 = load i32, i32* %16
	%18 = icmp ne i32 %17, 3
	br i1 %18, label %if.then.CYascwMZdx, label %if.else.pzHZjlFqRe

if.then.CYascwMZdx:
	%19 = getelementptr inbounds [21 x i8], [21 x i8]* @string.literal.IjwgowylaE, i32 0, i32 0
	%20 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%21 = call i32 (i8*, ...) @printf(i8* %20, i8* %19)
	br label %lastLeave.fbOWnycIAX

if.else.pzHZjlFqRe:
	br label %lastLeave.fbOWnycIAX

lastLeave.fbOWnycIAX:
	%22 = alloca %aStruct-fTwCvMHKtQ
	%23 = getelementptr inbounds %aStruct-fTwCvMHKtQ, %aStruct-fTwCvMHKtQ* %22, i32 0, i32 0
	store i32 3, i32* %23
	%24 = alloca %aStruct-fTwCvMHKtQ
	%25 = load %aStruct-fTwCvMHKtQ, %aStruct-fTwCvMHKtQ* %22
	store %aStruct-fTwCvMHKtQ %25, %aStruct-fTwCvMHKtQ* %24
	%26 = alloca %aStruct-fTwCvMHKtQ*
	store %aStruct-fTwCvMHKtQ* %24, %aStruct-fTwCvMHKtQ** %26
	%27 = load %aStruct-fTwCvMHKtQ*, %aStruct-fTwCvMHKtQ** %26
	%28 = bitcast %aStruct-fTwCvMHKtQ* %27 to i32*
	%29 = alloca i32*
	store i32* %28, i32** %29
	%30 = load i32*, i32** %29
	%31 = getelementptr inbounds i32, i32* %30, i32 0
	%32 = load i32, i32* %31
	%33 = getelementptr inbounds %aStruct-fTwCvMHKtQ, %aStruct-fTwCvMHKtQ* %24, i32 0, i32 0
	%34 = load i32, i32* %33
	%35 = icmp ne i32 %32, %34
	br i1 %35, label %if.then.SbCvZSnYxG, label %if.else.GWzRZyZdQH

if.then.SbCvZSnYxG:
	%36 = getelementptr inbounds [30 x i8], [30 x i8]* @string.literal.aCqaPfgDdq, i32 0, i32 0
	%37 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%38 = call i32 (i8*, ...) @printf(i8* %37, i8* %36)
	br label %lastLeave.abdjqTWIpV

if.else.GWzRZyZdQH:
	br label %lastLeave.abdjqTWIpV

lastLeave.abdjqTWIpV:
	%39 = load i32*, i32** %29
	%40 = bitcast i32* %39 to %aStruct-fTwCvMHKtQ*
	%41 = alloca %aStruct-fTwCvMHKtQ
	%42 = load %aStruct-fTwCvMHKtQ, %aStruct-fTwCvMHKtQ* %40
	store %aStruct-fTwCvMHKtQ %42, %aStruct-fTwCvMHKtQ* %41
	%43 = getelementptr inbounds %aStruct-fTwCvMHKtQ, %aStruct-fTwCvMHKtQ* %41, i32 0, i32 0
	%44 = load i32, i32* %43
	%45 = getelementptr inbounds %aStruct-fTwCvMHKtQ, %aStruct-fTwCvMHKtQ* %24, i32 0, i32 0
	%46 = load i32, i32* %45
	%47 = icmp ne i32 %44, %46
	br i1 %47, label %if.then.kgAPnxAzjJ, label %if.else.MsoPpvvtoM

if.then.kgAPnxAzjJ:
	%48 = getelementptr inbounds [25 x i8], [25 x i8]* @string.literal.xvjQpdbTop, i32 0, i32 0
	%49 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%50 = call i32 (i8*, ...) @printf(i8* %49, i8* %48)
	br label %lastLeave.ohICRhETYc

if.else.MsoPpvvtoM:
	br label %lastLeave.ohICRhETYc

lastLeave.ohICRhETYc:
	ret i32 0
}

declare ccc i32 @printf(i8* %0, ...)
