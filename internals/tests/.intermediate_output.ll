%aStruct-JXrvDSxPmo = type { i32, i64 }

@"%u " = global [4 x i8] c"%u \00"
@"%d " = global [4 x i8] c"%d \00"
@string.literal.sLMyPAXKOM = global [21 x i8] c"this shouldn't print\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.YLetUWFvrA = global [30 x i8] c"this should not print as well\00"
@string.literal.xdhHUUsgUG = global [25 x i8] c"this shouldnt print!!!!!\00"

define i32 @loadByPointer-JXrvDSxPmo(i32* %i-JXrvDSxPmo) {
loadByPointer-JXrvDSxPmo:
	%0 = alloca i32*
	store i32* %i-JXrvDSxPmo, i32** %0
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
	br i1 %5, label %if.then.GkATznYmzD, label %if.else.zwQQvRsakP

if.then.GkATznYmzD:
	%6 = load i32, i32* %0
	%7 = getelementptr inbounds [4 x i8], [4 x i8]* @"%u ", i32 0, i32 0
	%8 = call i32 (i8*, ...) @printf(i8* %7, i32 %6)
	%9 = load i32, i32* %0
	%10 = getelementptr inbounds [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%11 = call i32 (i8*, ...) @printf(i8* %10, i32 %9)
	br label %lastLeave.nnCVrlzINq

if.else.zwQQvRsakP:
	br label %lastLeave.nnCVrlzINq

lastLeave.nnCVrlzINq:
	%12 = alloca i32
	store i32 3, i32* %12
	%13 = alloca i32*
	store i32* %12, i32** %13
	%14 = load i32*, i32** %13
	%15 = call i32 @loadByPointer-JXrvDSxPmo(i32* %14)
	%16 = alloca i32
	store i32 %15, i32* %16
	%17 = load i32, i32* %16
	%18 = icmp ne i32 %17, 3
	br i1 %18, label %if.then.ouWdofGqpT, label %if.else.HypzDhHcnN

if.then.ouWdofGqpT:
	%19 = getelementptr inbounds [21 x i8], [21 x i8]* @string.literal.sLMyPAXKOM, i32 0, i32 0
	%20 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%21 = call i32 (i8*, ...) @printf(i8* %20, i8* %19)
	br label %lastLeave.XhWZeEwTob

if.else.HypzDhHcnN:
	br label %lastLeave.XhWZeEwTob

lastLeave.XhWZeEwTob:
	%22 = alloca %aStruct-JXrvDSxPmo
	%23 = getelementptr inbounds %aStruct-JXrvDSxPmo, %aStruct-JXrvDSxPmo* %22, i32 0, i32 0
	store i32 3, i32* %23
	%24 = alloca %aStruct-JXrvDSxPmo*
	store %aStruct-JXrvDSxPmo* %22, %aStruct-JXrvDSxPmo** %24
	%25 = load %aStruct-JXrvDSxPmo*, %aStruct-JXrvDSxPmo** %24
	%26 = bitcast %aStruct-JXrvDSxPmo* %25 to i32*
	%27 = alloca i32*
	store i32* %26, i32** %27
	%28 = load i32*, i32** %27
	%29 = getelementptr inbounds i32, i32* %28, i32 0
	%30 = load i32, i32* %29
	%31 = getelementptr inbounds %aStruct-JXrvDSxPmo, %aStruct-JXrvDSxPmo* %22, i32 0, i32 0
	%32 = load i32, i32* %31
	%33 = icmp ne i32 %30, %32
	br i1 %33, label %if.then.fBrHwuwvaP, label %if.else.EtVabyNNuz

if.then.fBrHwuwvaP:
	%34 = getelementptr inbounds [30 x i8], [30 x i8]* @string.literal.YLetUWFvrA, i32 0, i32 0
	%35 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%36 = call i32 (i8*, ...) @printf(i8* %35, i8* %34)
	br label %lastLeave.eoXdkNQEss

if.else.EtVabyNNuz:
	br label %lastLeave.eoXdkNQEss

lastLeave.eoXdkNQEss:
	%37 = load i32*, i32** %27
	%38 = bitcast i32* %37 to %aStruct-JXrvDSxPmo*
	%39 = alloca %aStruct-JXrvDSxPmo
	%40 = load %aStruct-JXrvDSxPmo, %aStruct-JXrvDSxPmo* %38
	store %aStruct-JXrvDSxPmo %40, %aStruct-JXrvDSxPmo* %39
	%41 = getelementptr inbounds %aStruct-JXrvDSxPmo, %aStruct-JXrvDSxPmo* %39, i32 0, i32 0
	%42 = load i32, i32* %41
	%43 = getelementptr inbounds %aStruct-JXrvDSxPmo, %aStruct-JXrvDSxPmo* %22, i32 0, i32 0
	%44 = load i32, i32* %43
	%45 = icmp ne i32 %42, %44
	br i1 %45, label %if.then.MdIVHrdTBE, label %if.else.cvYchFYTLZ

if.then.MdIVHrdTBE:
	%46 = getelementptr inbounds [25 x i8], [25 x i8]* @string.literal.xdhHUUsgUG, i32 0, i32 0
	%47 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%48 = call i32 (i8*, ...) @printf(i8* %47, i8* %46)
	br label %lastLeave.BcuEmGHnTj

if.else.cvYchFYTLZ:
	br label %lastLeave.BcuEmGHnTj

lastLeave.BcuEmGHnTj:
	ret i32 0
}

declare ccc i32 @printf(i8* %0, ...)
