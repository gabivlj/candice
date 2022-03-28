%aStruct-WsrqsERJWE = type { i32, i64 }

@"%u " = global [4 x i8] c"%u \00"
@"%d " = global [4 x i8] c"%d \00"
@string.literal.WyCvXBDqbb = global [21 x i8] c"this shouldn't print\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.ryWHMNOrWg = global [30 x i8] c"this should not print as well\00"
@string.literal.pVfACIBOoe = global [25 x i8] c"this shouldnt print!!!!!\00"

define i32 @loadByPointer-WsrqsERJWE(i32* %i-WsrqsERJWE) {
loadByPointer-WsrqsERJWE:
	%0 = alloca i32*
	store i32* %i-WsrqsERJWE, i32** %0
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
	br i1 %5, label %if.then.hNKnvRSCdr, label %if.else.JpbnYpkKLW

if.then.hNKnvRSCdr:
	%6 = load i32, i32* %0
	%7 = getelementptr inbounds [4 x i8], [4 x i8]* @"%u ", i32 0, i32 0
	%8 = call i32 (i8*, ...) @printf(i8* %7, i32 %6)
	%9 = load i32, i32* %0
	%10 = getelementptr inbounds [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%11 = call i32 (i8*, ...) @printf(i8* %10, i32 %9)
	br label %lastLeave.LVOBEiMfxy

if.else.JpbnYpkKLW:
	br label %lastLeave.LVOBEiMfxy

lastLeave.LVOBEiMfxy:
	%12 = alloca i32
	store i32 3, i32* %12
	%13 = alloca i32*
	store i32* %12, i32** %13
	%14 = load i32*, i32** %13
	%15 = call i32 @loadByPointer-WsrqsERJWE(i32* %14)
	%16 = alloca i32
	store i32 %15, i32* %16
	%17 = load i32, i32* %16
	%18 = icmp ne i32 %17, 3
	br i1 %18, label %if.then.RbjlMqrgme, label %if.else.qelQtEUGrK

if.then.RbjlMqrgme:
	%19 = getelementptr inbounds [21 x i8], [21 x i8]* @string.literal.WyCvXBDqbb, i32 0, i32 0
	%20 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%21 = call i32 (i8*, ...) @printf(i8* %20, i8* %19)
	br label %lastLeave.fkjflVQJUi

if.else.qelQtEUGrK:
	br label %lastLeave.fkjflVQJUi

lastLeave.fkjflVQJUi:
	%22 = alloca %aStruct-WsrqsERJWE
	%23 = getelementptr inbounds %aStruct-WsrqsERJWE, %aStruct-WsrqsERJWE* %22, i32 0, i32 0
	store i32 3, i32* %23
	%24 = alloca %aStruct-WsrqsERJWE*
	store %aStruct-WsrqsERJWE* %22, %aStruct-WsrqsERJWE** %24
	%25 = load %aStruct-WsrqsERJWE*, %aStruct-WsrqsERJWE** %24
	%26 = bitcast %aStruct-WsrqsERJWE* %25 to i32*
	%27 = alloca i32*
	store i32* %26, i32** %27
	%28 = load i32*, i32** %27
	%29 = getelementptr inbounds i32, i32* %28, i32 0
	%30 = load i32, i32* %29
	%31 = getelementptr inbounds %aStruct-WsrqsERJWE, %aStruct-WsrqsERJWE* %22, i32 0, i32 0
	%32 = load i32, i32* %31
	%33 = icmp ne i32 %30, %32
	br i1 %33, label %if.then.YSdvqyNXCQ, label %if.else.WvglXQXper

if.then.YSdvqyNXCQ:
	%34 = getelementptr inbounds [30 x i8], [30 x i8]* @string.literal.ryWHMNOrWg, i32 0, i32 0
	%35 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%36 = call i32 (i8*, ...) @printf(i8* %35, i8* %34)
	br label %lastLeave.prBWEBNRHN

if.else.WvglXQXper:
	br label %lastLeave.prBWEBNRHN

lastLeave.prBWEBNRHN:
	%37 = load i32*, i32** %27
	%38 = bitcast i32* %37 to %aStruct-WsrqsERJWE*
	%39 = alloca %aStruct-WsrqsERJWE
	%40 = load %aStruct-WsrqsERJWE, %aStruct-WsrqsERJWE* %38
	store %aStruct-WsrqsERJWE %40, %aStruct-WsrqsERJWE* %39
	%41 = getelementptr inbounds %aStruct-WsrqsERJWE, %aStruct-WsrqsERJWE* %39, i32 0, i32 0
	%42 = load i32, i32* %41
	%43 = getelementptr inbounds %aStruct-WsrqsERJWE, %aStruct-WsrqsERJWE* %22, i32 0, i32 0
	%44 = load i32, i32* %43
	%45 = icmp ne i32 %42, %44
	br i1 %45, label %if.then.zaKmeiUnuc, label %if.else.TUJMkzyHfe

if.then.zaKmeiUnuc:
	%46 = getelementptr inbounds [25 x i8], [25 x i8]* @string.literal.pVfACIBOoe, i32 0, i32 0
	%47 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%48 = call i32 (i8*, ...) @printf(i8* %47, i8* %46)
	br label %lastLeave.qXsDDAHott

if.else.TUJMkzyHfe:
	br label %lastLeave.qXsDDAHott

lastLeave.qXsDDAHott:
	ret i32 0
}

declare ccc i32 @printf(i8* %0, ...)
