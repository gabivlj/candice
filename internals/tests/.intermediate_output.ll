%aStruct-GckvFVSbWq = type { i32, i64 }

@"%u " = global [4 x i8] c"%u \00"
@"%d " = global [4 x i8] c"%d \00"
@string.literal.tzIYwlZUFb = global [21 x i8] c"this shouldn't print\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.beaAVxwBQQ = global [30 x i8] c"this should not print as well\00"
@string.literal.BOyheddmvp = global [25 x i8] c"this shouldnt print!!!!!\00"

define i32 @loadByPointer-GckvFVSbWq(i32* %i-GckvFVSbWq) {
loadByPointer-GckvFVSbWq:
	%0 = alloca i32*
	store i32* %i-GckvFVSbWq, i32** %0
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
	br i1 %5, label %if.then.PRsociasXe, label %if.else.NgJRwPxQRA

if.then.PRsociasXe:
	%6 = load i32, i32* %0
	%7 = getelementptr [4 x i8], [4 x i8]* @"%u ", i32 0, i32 0
	%8 = call i32 (i8*, ...) @printf(i8* %7, i32 %6)
	%9 = load i32, i32* %0
	%10 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%11 = call i32 (i8*, ...) @printf(i8* %10, i32 %9)
	br label %lastLeave.xJvRVTBcIk

if.else.NgJRwPxQRA:
	br label %lastLeave.xJvRVTBcIk

lastLeave.xJvRVTBcIk:
	%12 = alloca i32
	store i32 3, i32* %12
	%13 = alloca i32*
	store i32* %12, i32** %13
	%14 = load i32*, i32** %13
	%15 = call i32 @loadByPointer-GckvFVSbWq(i32* %14)
	%16 = alloca i32
	store i32 %15, i32* %16
	%17 = load i32, i32* %16
	%18 = icmp ne i32 %17, 3
	br i1 %18, label %if.then.gPzzWIZoVG, label %if.else.guAlHGojKW

if.then.gPzzWIZoVG:
	%19 = getelementptr [21 x i8], [21 x i8]* @string.literal.tzIYwlZUFb, i32 0, i32 0
	%20 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%21 = call i32 (i8*, ...) @printf(i8* %20, i8* %19)
	br label %lastLeave.EjPYzXZPVO

if.else.guAlHGojKW:
	br label %lastLeave.EjPYzXZPVO

lastLeave.EjPYzXZPVO:
	%22 = alloca %aStruct-GckvFVSbWq
	%23 = getelementptr %aStruct-GckvFVSbWq, %aStruct-GckvFVSbWq* %22, i32 0, i32 0
	store i32 3, i32* %23
	%24 = alloca %aStruct-GckvFVSbWq
	%25 = load %aStruct-GckvFVSbWq, %aStruct-GckvFVSbWq* %22
	store %aStruct-GckvFVSbWq %25, %aStruct-GckvFVSbWq* %24
	%26 = alloca %aStruct-GckvFVSbWq*
	store %aStruct-GckvFVSbWq* %24, %aStruct-GckvFVSbWq** %26
	%27 = load %aStruct-GckvFVSbWq*, %aStruct-GckvFVSbWq** %26
	%28 = bitcast %aStruct-GckvFVSbWq* %27 to i32*
	%29 = alloca i32*
	store i32* %28, i32** %29
	%30 = load i32*, i32** %29
	%31 = getelementptr i32, i32* %30, i32 0
	%32 = load i32, i32* %31
	%33 = getelementptr %aStruct-GckvFVSbWq, %aStruct-GckvFVSbWq* %24, i32 0, i32 0
	%34 = load i32, i32* %33
	%35 = icmp ne i32 %32, %34
	br i1 %35, label %if.then.UbavPfUPve, label %if.else.QQzoozuAIw

if.then.UbavPfUPve:
	%36 = getelementptr [30 x i8], [30 x i8]* @string.literal.beaAVxwBQQ, i32 0, i32 0
	%37 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%38 = call i32 (i8*, ...) @printf(i8* %37, i8* %36)
	br label %lastLeave.QGHDfGIhQW

if.else.QQzoozuAIw:
	br label %lastLeave.QGHDfGIhQW

lastLeave.QGHDfGIhQW:
	%39 = load i32*, i32** %29
	%40 = bitcast i32* %39 to %aStruct-GckvFVSbWq*
	%41 = load %aStruct-GckvFVSbWq, %aStruct-GckvFVSbWq* %40
	%42 = alloca %aStruct-GckvFVSbWq
	store %aStruct-GckvFVSbWq %41, %aStruct-GckvFVSbWq* %42
	%43 = getelementptr %aStruct-GckvFVSbWq, %aStruct-GckvFVSbWq* %42, i32 0, i32 0
	%44 = load i32, i32* %43
	%45 = getelementptr %aStruct-GckvFVSbWq, %aStruct-GckvFVSbWq* %24, i32 0, i32 0
	%46 = load i32, i32* %45
	%47 = icmp ne i32 %44, %46
	br i1 %47, label %if.then.dwNdyCxxwA, label %if.else.hyiYzIYfkd

if.then.dwNdyCxxwA:
	%48 = getelementptr [25 x i8], [25 x i8]* @string.literal.BOyheddmvp, i32 0, i32 0
	%49 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%50 = call i32 (i8*, ...) @printf(i8* %49, i8* %48)
	br label %lastLeave.KXhrNDIuRn

if.else.hyiYzIYfkd:
	br label %lastLeave.KXhrNDIuRn

lastLeave.KXhrNDIuRn:
	ret i32 0
}

declare ccc i32 @printf(i8* %0, ...)
