%aStruct-YegqrkKzmc = type { i32, i64 }

@"%u " = global [4 x i8] c"%u \00"
@"%d " = global [4 x i8] c"%d \00"
@string.literal.TyKZxAakRI = global [21 x i8] c"this shouldn't print\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.XGhemsyFsF = global [30 x i8] c"this should not print as well\00"
@string.literal.tHOGxDTceR = global [25 x i8] c"this shouldnt print!!!!!\00"

define i32 @loadByPointer-YegqrkKzmc(i32* %i-YegqrkKzmc) {
loadByPointer-YegqrkKzmc:
	%0 = alloca i32*
	store i32* %i-YegqrkKzmc, i32** %0
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
	br i1 %5, label %if.then.MjLVIoYIkz, label %if.else.aiifkpYpLW

if.then.MjLVIoYIkz:
	%6 = load i32, i32* %0
	%7 = getelementptr [4 x i8], [4 x i8]* @"%u ", i32 0, i32 0
	%8 = call i32 (i8*, ...) @printf(i8* %7, i32 %6)
	%9 = load i32, i32* %0
	%10 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%11 = call i32 (i8*, ...) @printf(i8* %10, i32 %9)
	br label %lastLeave.xCFjEKwNIl

if.else.aiifkpYpLW:
	br label %lastLeave.xCFjEKwNIl

lastLeave.xCFjEKwNIl:
	%12 = alloca i32
	store i32 3, i32* %12
	%13 = alloca i32*
	store i32* %12, i32** %13
	%14 = load i32*, i32** %13
	%15 = call i32 @loadByPointer-YegqrkKzmc(i32* %14)
	%16 = alloca i32
	store i32 %15, i32* %16
	%17 = load i32, i32* %16
	%18 = icmp ne i32 %17, 3
	br i1 %18, label %if.then.dhHQYWWIVa, label %if.else.rPiQndrkbg

if.then.dhHQYWWIVa:
	%19 = getelementptr [21 x i8], [21 x i8]* @string.literal.TyKZxAakRI, i32 0, i32 0
	%20 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%21 = call i32 (i8*, ...) @printf(i8* %20, i8* %19)
	br label %lastLeave.InJkNEPYbA

if.else.rPiQndrkbg:
	br label %lastLeave.InJkNEPYbA

lastLeave.InJkNEPYbA:
	%22 = alloca %aStruct-YegqrkKzmc
	%23 = getelementptr %aStruct-YegqrkKzmc, %aStruct-YegqrkKzmc* %22, i32 0, i32 0
	store i32 3, i32* %23
	%24 = alloca %aStruct-YegqrkKzmc
	%25 = load %aStruct-YegqrkKzmc, %aStruct-YegqrkKzmc* %22
	store %aStruct-YegqrkKzmc %25, %aStruct-YegqrkKzmc* %24
	%26 = alloca %aStruct-YegqrkKzmc*
	store %aStruct-YegqrkKzmc* %24, %aStruct-YegqrkKzmc** %26
	%27 = load %aStruct-YegqrkKzmc*, %aStruct-YegqrkKzmc** %26
	%28 = bitcast %aStruct-YegqrkKzmc* %27 to i32*
	%29 = alloca i32*
	store i32* %28, i32** %29
	%30 = load i32*, i32** %29
	%31 = getelementptr i32, i32* %30, i32 0
	%32 = load i32, i32* %31
	%33 = getelementptr %aStruct-YegqrkKzmc, %aStruct-YegqrkKzmc* %24, i32 0, i32 0
	%34 = load i32, i32* %33
	%35 = icmp ne i32 %32, %34
	br i1 %35, label %if.then.AmTeNPFtaZ, label %if.else.fuwVVUZzfE

if.then.AmTeNPFtaZ:
	%36 = getelementptr [30 x i8], [30 x i8]* @string.literal.XGhemsyFsF, i32 0, i32 0
	%37 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%38 = call i32 (i8*, ...) @printf(i8* %37, i8* %36)
	br label %lastLeave.GEKNmncDoK

if.else.fuwVVUZzfE:
	br label %lastLeave.GEKNmncDoK

lastLeave.GEKNmncDoK:
	%39 = load i32*, i32** %29
	%40 = bitcast i32* %39 to %aStruct-YegqrkKzmc*
	%41 = load %aStruct-YegqrkKzmc, %aStruct-YegqrkKzmc* %40
	%42 = alloca %aStruct-YegqrkKzmc
	store %aStruct-YegqrkKzmc %41, %aStruct-YegqrkKzmc* %42
	%43 = getelementptr %aStruct-YegqrkKzmc, %aStruct-YegqrkKzmc* %42, i32 0, i32 0
	%44 = load i32, i32* %43
	%45 = getelementptr %aStruct-YegqrkKzmc, %aStruct-YegqrkKzmc* %24, i32 0, i32 0
	%46 = load i32, i32* %45
	%47 = icmp ne i32 %44, %46
	br i1 %47, label %if.then.BtQsfsmaEP, label %if.else.htfBNXghGK

if.then.BtQsfsmaEP:
	%48 = getelementptr [25 x i8], [25 x i8]* @string.literal.tHOGxDTceR, i32 0, i32 0
	%49 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%50 = call i32 (i8*, ...) @printf(i8* %49, i8* %48)
	br label %lastLeave.YSzdcNWfbr

if.else.htfBNXghGK:
	br label %lastLeave.YSzdcNWfbr

lastLeave.YSzdcNWfbr:
	ret i32 0
}

declare ccc i32 @printf(i8* %0, ...)
