%aStruct-htLuYkEYcC = type { i32, i64 }

@"%u " = global [4 x i8] c"%u \00"
@"%d " = global [4 x i8] c"%d \00"
@string.literal.grSkMJooFv = global [21 x i8] c"this shouldn't print\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.ZjYvHfUQkj = global [30 x i8] c"this should not print as well\00"
@string.literal.jjXqgTlRjI = global [25 x i8] c"this shouldnt print!!!!!\00"

define i32 @loadByPointer-htLuYkEYcC(i32* %i-htLuYkEYcC) {
loadByPointer-htLuYkEYcC:
	%0 = alloca i32*
	store i32* %i-htLuYkEYcC, i32** %0
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
	br i1 %6, label %if.then.UuXYRLiOyO, label %if.else.KFGTBNFPgB

if.then.UuXYRLiOyO:
	%7 = load i32, i32* %1
	%8 = getelementptr inbounds [4 x i8], [4 x i8]* @"%u ", i32 0, i32 0
	%9 = call i32 (i8*, ...) @printf(i8* %8, i32 %7)
	%10 = load i32, i32* %1
	%11 = getelementptr inbounds [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%12 = call i32 (i8*, ...) @printf(i8* %11, i32 %10)
	br label %lastLeave.gLDKxoBufh

if.else.KFGTBNFPgB:
	br label %lastLeave.gLDKxoBufh

lastLeave.gLDKxoBufh:
	%13 = alloca i32
	store i32 3, i32* %13
	%14 = alloca i32*
	store i32* %13, i32** %14
	%15 = load i32*, i32** %14
	%16 = call i32 @loadByPointer-htLuYkEYcC(i32* %15)
	%17 = alloca i32
	store i32 %16, i32* %17
	%18 = load i32, i32* %17
	%19 = icmp ne i32 %18, 3
	br i1 %19, label %if.then.uoCxeUoWAx, label %if.else.LSrfDOuLFB

if.then.uoCxeUoWAx:
	%20 = getelementptr inbounds [21 x i8], [21 x i8]* @string.literal.grSkMJooFv, i32 0, i32 0
	%21 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%22 = call i32 (i8*, ...) @printf(i8* %21, i8* %20)
	br label %lastLeave.SatQIaXTic

if.else.LSrfDOuLFB:
	br label %lastLeave.SatQIaXTic

lastLeave.SatQIaXTic:
	%23 = alloca %aStruct-htLuYkEYcC
	%24 = getelementptr inbounds %aStruct-htLuYkEYcC, %aStruct-htLuYkEYcC* %23, i32 0, i32 0
	store i32 3, i32* %24
	%25 = alloca %aStruct-htLuYkEYcC*
	store %aStruct-htLuYkEYcC* %23, %aStruct-htLuYkEYcC** %25
	%26 = load %aStruct-htLuYkEYcC*, %aStruct-htLuYkEYcC** %25
	%27 = bitcast %aStruct-htLuYkEYcC* %26 to i32*
	%28 = alloca i32*
	store i32* %27, i32** %28
	%29 = load i32*, i32** %28
	%30 = getelementptr inbounds i32, i32* %29, i32 0
	%31 = load i32, i32* %30
	%32 = getelementptr inbounds %aStruct-htLuYkEYcC, %aStruct-htLuYkEYcC* %23, i32 0, i32 0
	%33 = load i32, i32* %32
	%34 = icmp ne i32 %31, %33
	br i1 %34, label %if.then.tLJnWwfugv, label %if.else.wKnDojlcHc

if.then.tLJnWwfugv:
	%35 = getelementptr inbounds [30 x i8], [30 x i8]* @string.literal.ZjYvHfUQkj, i32 0, i32 0
	%36 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%37 = call i32 (i8*, ...) @printf(i8* %36, i8* %35)
	br label %lastLeave.lOqYNAWDoU

if.else.wKnDojlcHc:
	br label %lastLeave.lOqYNAWDoU

lastLeave.lOqYNAWDoU:
	%38 = load i32*, i32** %28
	%39 = bitcast i32* %38 to %aStruct-htLuYkEYcC*
	%40 = alloca %aStruct-htLuYkEYcC
	%41 = load %aStruct-htLuYkEYcC, %aStruct-htLuYkEYcC* %39
	store %aStruct-htLuYkEYcC %41, %aStruct-htLuYkEYcC* %40
	%42 = getelementptr inbounds %aStruct-htLuYkEYcC, %aStruct-htLuYkEYcC* %40, i32 0, i32 0
	%43 = load i32, i32* %42
	%44 = getelementptr inbounds %aStruct-htLuYkEYcC, %aStruct-htLuYkEYcC* %23, i32 0, i32 0
	%45 = load i32, i32* %44
	%46 = icmp ne i32 %43, %45
	br i1 %46, label %if.then.szSiLVDYPb, label %if.else.zubQcgZabV

if.then.szSiLVDYPb:
	%47 = getelementptr inbounds [25 x i8], [25 x i8]* @string.literal.jjXqgTlRjI, i32 0, i32 0
	%48 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%49 = call i32 (i8*, ...) @printf(i8* %48, i8* %47)
	br label %lastLeave.tAMWbsunOw

if.else.zubQcgZabV:
	br label %lastLeave.tAMWbsunOw

lastLeave.tAMWbsunOw:
	ret i32 0
}

declare ccc i32 @printf(i8* %0, ...)
