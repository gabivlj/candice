%aStruct-cKcEYAdvZb = type { i32, i64 }

@"%u " = global [4 x i8] c"%u \00"
@"%d " = global [4 x i8] c"%d \00"
@string.literal.zYATsiwJhP = global [21 x i8] c"this shouldn't print\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.cMoOEzkaFk = global [30 x i8] c"this should not print as well\00"
@string.literal.YeQCyAufMg = global [25 x i8] c"this shouldnt print!!!!!\00"

declare ccc void @free(i8* %0, ...)

declare ccc i8* @realloc(i8* %0, i64 %1, ...)

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc i32 @loadByPointer-cKcEYAdvZb(i32* %i-cKcEYAdvZb) {
loadByPointer-cKcEYAdvZb:
	%0 = alloca i32*
	store i32* %i-cKcEYAdvZb, i32** %0
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
	br i1 %6, label %if.then.FEdtQpXygQ, label %if.else.nmQtgUOkof

if.then.FEdtQpXygQ:
	%7 = load i32, i32* %1
	%8 = getelementptr [4 x i8], [4 x i8]* @"%u ", i32 0, i32 0
	%9 = call i32 (i8*, ...) @printf(i8* %8, i32 %7)
	%10 = load i32, i32* %1
	%11 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%12 = call i32 (i8*, ...) @printf(i8* %11, i32 %10)
	br label %lastLeave.miXaUtNMWQ

if.else.nmQtgUOkof:
	br label %lastLeave.miXaUtNMWQ

lastLeave.miXaUtNMWQ:
	%13 = alloca i32
	store i32 3, i32* %13
	%14 = alloca i32*
	store i32* %13, i32** %14
	%15 = load i32*, i32** %14
	%16 = call i32 @loadByPointer-cKcEYAdvZb(i32* %15)
	%17 = alloca i32
	store i32 %16, i32* %17
	%18 = load i32, i32* %17
	%19 = icmp ne i32 %18, 3
	br i1 %19, label %if.then.mQxVnKsBea, label %if.else.TpyRRGGvAg

if.then.mQxVnKsBea:
	%20 = getelementptr [21 x i8], [21 x i8]* @string.literal.zYATsiwJhP, i32 0, i32 0
	%21 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%22 = call i32 (i8*, ...) @printf(i8* %21, i8* %20)
	br label %lastLeave.xVhWAQincM

if.else.TpyRRGGvAg:
	br label %lastLeave.xVhWAQincM

lastLeave.xVhWAQincM:
	%23 = alloca %aStruct-cKcEYAdvZb
	%24 = getelementptr %aStruct-cKcEYAdvZb, %aStruct-cKcEYAdvZb* %23, i32 0, i32 0
	store i32 3, i32* %24
	%25 = alloca %aStruct-cKcEYAdvZb
	%26 = load %aStruct-cKcEYAdvZb, %aStruct-cKcEYAdvZb* %23
	store %aStruct-cKcEYAdvZb %26, %aStruct-cKcEYAdvZb* %25
	%27 = alloca %aStruct-cKcEYAdvZb*
	store %aStruct-cKcEYAdvZb* %25, %aStruct-cKcEYAdvZb** %27
	%28 = load %aStruct-cKcEYAdvZb*, %aStruct-cKcEYAdvZb** %27
	%29 = bitcast %aStruct-cKcEYAdvZb* %28 to i32*
	%30 = alloca i32*
	store i32* %29, i32** %30
	%31 = load i32*, i32** %30
	%32 = getelementptr i32, i32* %31, i32 0
	%33 = load i32, i32* %32
	%34 = getelementptr %aStruct-cKcEYAdvZb, %aStruct-cKcEYAdvZb* %25, i32 0, i32 0
	%35 = load i32, i32* %34
	%36 = icmp ne i32 %33, %35
	br i1 %36, label %if.then.iptKzkItMk, label %if.else.vwdqklrNoU

if.then.iptKzkItMk:
	%37 = getelementptr [30 x i8], [30 x i8]* @string.literal.cMoOEzkaFk, i32 0, i32 0
	%38 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%39 = call i32 (i8*, ...) @printf(i8* %38, i8* %37)
	br label %lastLeave.HVyZycVnAp

if.else.vwdqklrNoU:
	br label %lastLeave.HVyZycVnAp

lastLeave.HVyZycVnAp:
	%40 = load i32*, i32** %30
	%41 = bitcast i32* %40 to %aStruct-cKcEYAdvZb*
	%42 = load %aStruct-cKcEYAdvZb, %aStruct-cKcEYAdvZb* %41
	%43 = alloca %aStruct-cKcEYAdvZb
	store %aStruct-cKcEYAdvZb %42, %aStruct-cKcEYAdvZb* %43
	%44 = getelementptr %aStruct-cKcEYAdvZb, %aStruct-cKcEYAdvZb* %43, i32 0, i32 0
	%45 = load i32, i32* %44
	%46 = getelementptr %aStruct-cKcEYAdvZb, %aStruct-cKcEYAdvZb* %25, i32 0, i32 0
	%47 = load i32, i32* %46
	%48 = icmp ne i32 %45, %47
	br i1 %48, label %if.then.JZLDYdaIAs, label %if.else.FBrqnLYjPe

if.then.JZLDYdaIAs:
	%49 = getelementptr [25 x i8], [25 x i8]* @string.literal.YeQCyAufMg, i32 0, i32 0
	%50 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%51 = call i32 (i8*, ...) @printf(i8* %50, i8* %49)
	br label %lastLeave.eGeayeAeBb

if.else.FBrqnLYjPe:
	br label %lastLeave.eGeayeAeBb

lastLeave.eGeayeAeBb:
	ret i32 0
}
