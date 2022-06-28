%aStruct-AtjApNeCmQ = type { i32, i64 }

@string.literal.xeeQEFnZXm = global [1 x i8] c"\00"
@"%u %s" = global [6 x i8] c"%u %s\00"
@string.literal.HCbUNRUAKc = global [1 x i8] c"\00"
@"%d %s" = global [6 x i8] c"%d %s\00"
@string.literal.HkHgHpLbfX = global [21 x i8] c"this shouldn't print\00"
@"%s" = global [3 x i8] c"%s\00"
@string.literal.QxYslOgoBC = global [30 x i8] c"this should not print as well\00"
@string.literal.RVoQhxXblh = global [25 x i8] c"this shouldnt print!!!!!\00"

define i32 @loadByPointer-AtjApNeCmQ(i32* %i-AtjApNeCmQ) {
loadByPointer-AtjApNeCmQ:
	%0 = alloca i32*
	store i32* %i-AtjApNeCmQ, i32** %0
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
	br i1 %6, label %if.then.EeRDlKSedw, label %if.else.TswFZYlWrr

if.then.EeRDlKSedw:
	%7 = load i32, i32* %1
	%8 = getelementptr inbounds [1 x i8], [1 x i8]* @string.literal.xeeQEFnZXm, i32 0, i32 0
	%9 = getelementptr inbounds [6 x i8], [6 x i8]* @"%u %s", i32 0, i32 0
	%10 = call i32 (i8*, ...) @printf(i8* %9, i32 %7, i8* %8)
	%11 = load i32, i32* %1
	%12 = getelementptr inbounds [1 x i8], [1 x i8]* @string.literal.HCbUNRUAKc, i32 0, i32 0
	%13 = getelementptr inbounds [6 x i8], [6 x i8]* @"%d %s", i32 0, i32 0
	%14 = call i32 (i8*, ...) @printf(i8* %13, i32 %11, i8* %12)
	br label %lastLeave.jvFySGenHV

if.else.TswFZYlWrr:
	br label %lastLeave.jvFySGenHV

lastLeave.jvFySGenHV:
	%15 = alloca i32
	store i32 3, i32* %15
	%16 = alloca i32*
	store i32* %15, i32** %16
	%17 = load i32*, i32** %16
	%18 = call i32 @loadByPointer-AtjApNeCmQ(i32* %17)
	%19 = alloca i32
	store i32 %18, i32* %19
	%20 = load i32, i32* %19
	%21 = icmp ne i32 %20, 3
	br i1 %21, label %if.then.SlWpmWawUr, label %if.else.qZEegurczU

if.then.SlWpmWawUr:
	%22 = getelementptr inbounds [21 x i8], [21 x i8]* @string.literal.HkHgHpLbfX, i32 0, i32 0
	%23 = getelementptr inbounds [3 x i8], [3 x i8]* @"%s", i32 0, i32 0
	%24 = call i32 (i8*, ...) @printf(i8* %23, i8* %22)
	br label %lastLeave.MhYqsZNUYM

if.else.qZEegurczU:
	br label %lastLeave.MhYqsZNUYM

lastLeave.MhYqsZNUYM:
	%25 = alloca %aStruct-AtjApNeCmQ
	%26 = getelementptr inbounds %aStruct-AtjApNeCmQ, %aStruct-AtjApNeCmQ* %25, i32 0, i32 0
	store i32 3, i32* %26
	%27 = alloca %aStruct-AtjApNeCmQ*
	store %aStruct-AtjApNeCmQ* %25, %aStruct-AtjApNeCmQ** %27
	%28 = load %aStruct-AtjApNeCmQ*, %aStruct-AtjApNeCmQ** %27
	%29 = bitcast %aStruct-AtjApNeCmQ* %28 to i32*
	%30 = alloca i32*
	store i32* %29, i32** %30
	%31 = load i32*, i32** %30
	%32 = getelementptr inbounds i32, i32* %31, i32 0
	%33 = load i32, i32* %32
	%34 = getelementptr inbounds %aStruct-AtjApNeCmQ, %aStruct-AtjApNeCmQ* %25, i32 0, i32 0
	%35 = load i32, i32* %34
	%36 = icmp ne i32 %33, %35
	br i1 %36, label %if.then.STBSDtXxnu, label %if.else.pANGgCCSXJ

if.then.STBSDtXxnu:
	%37 = getelementptr inbounds [30 x i8], [30 x i8]* @string.literal.QxYslOgoBC, i32 0, i32 0
	%38 = getelementptr inbounds [3 x i8], [3 x i8]* @"%s", i32 0, i32 0
	%39 = call i32 (i8*, ...) @printf(i8* %38, i8* %37)
	br label %lastLeave.CliRnLMZwW

if.else.pANGgCCSXJ:
	br label %lastLeave.CliRnLMZwW

lastLeave.CliRnLMZwW:
	%40 = load i32*, i32** %30
	%41 = bitcast i32* %40 to %aStruct-AtjApNeCmQ*
	%42 = alloca %aStruct-AtjApNeCmQ
	%43 = load %aStruct-AtjApNeCmQ, %aStruct-AtjApNeCmQ* %41
	store %aStruct-AtjApNeCmQ %43, %aStruct-AtjApNeCmQ* %42
	%44 = getelementptr inbounds %aStruct-AtjApNeCmQ, %aStruct-AtjApNeCmQ* %42, i32 0, i32 0
	%45 = load i32, i32* %44
	%46 = getelementptr inbounds %aStruct-AtjApNeCmQ, %aStruct-AtjApNeCmQ* %25, i32 0, i32 0
	%47 = load i32, i32* %46
	%48 = icmp ne i32 %45, %47
	br i1 %48, label %if.then.DsyMcNtWNx, label %if.else.GZqAJBVcXP

if.then.DsyMcNtWNx:
	%49 = getelementptr inbounds [25 x i8], [25 x i8]* @string.literal.RVoQhxXblh, i32 0, i32 0
	%50 = getelementptr inbounds [3 x i8], [3 x i8]* @"%s", i32 0, i32 0
	%51 = call i32 (i8*, ...) @printf(i8* %50, i8* %49)
	br label %lastLeave.nbdVyMKUlp

if.else.GZqAJBVcXP:
	br label %lastLeave.nbdVyMKUlp

lastLeave.nbdVyMKUlp:
	ret i32 0
}

declare ccc i32 @printf(i8* %0, ...)
