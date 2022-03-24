%aStruct-lcdzRdlmDm = type { i32, i64 }

@"%u " = global [4 x i8] c"%u \00"
@"%d " = global [4 x i8] c"%d \00"
@string.literal.EYhwZySCWT = global [21 x i8] c"this shouldn't print\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.eTRwDylQIo = global [30 x i8] c"this should not print as well\00"
@string.literal.CgwbdXFXPr = global [25 x i8] c"this shouldnt print!!!!!\00"

define i32 @loadByPointer-lcdzRdlmDm(i32* %i-lcdzRdlmDm) {
loadByPointer-lcdzRdlmDm:
	%0 = alloca i32*
	store i32* %i-lcdzRdlmDm, i32** %0
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
	br i1 %5, label %if.then.zCrHcoiHRf, label %if.else.nFMbAYyKlq

if.then.zCrHcoiHRf:
	%6 = load i32, i32* %0
	%7 = getelementptr inbounds [4 x i8], [4 x i8]* @"%u ", i32 0, i32 0
	%8 = call i32 (i8*, ...) @printf(i8* %7, i32 %6)
	%9 = load i32, i32* %0
	%10 = getelementptr inbounds [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%11 = call i32 (i8*, ...) @printf(i8* %10, i32 %9)
	br label %lastLeave.HAVeKtePsB

if.else.nFMbAYyKlq:
	br label %lastLeave.HAVeKtePsB

lastLeave.HAVeKtePsB:
	%12 = alloca i32
	store i32 3, i32* %12
	%13 = alloca i32*
	store i32* %12, i32** %13
	%14 = load i32*, i32** %13
	%15 = call i32 @loadByPointer-lcdzRdlmDm(i32* %14)
	%16 = alloca i32
	store i32 %15, i32* %16
	%17 = load i32, i32* %16
	%18 = icmp ne i32 %17, 3
	br i1 %18, label %if.then.rtOBXbRlkg, label %if.else.DMNJVHwlAb

if.then.rtOBXbRlkg:
	%19 = getelementptr inbounds [21 x i8], [21 x i8]* @string.literal.EYhwZySCWT, i32 0, i32 0
	%20 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%21 = call i32 (i8*, ...) @printf(i8* %20, i8* %19)
	br label %lastLeave.NiaGHCzphY

if.else.DMNJVHwlAb:
	br label %lastLeave.NiaGHCzphY

lastLeave.NiaGHCzphY:
	%22 = alloca %aStruct-lcdzRdlmDm
	%23 = getelementptr inbounds %aStruct-lcdzRdlmDm, %aStruct-lcdzRdlmDm* %22, i32 0, i32 0
	store i32 3, i32* %23
	%24 = alloca %aStruct-lcdzRdlmDm*
	store %aStruct-lcdzRdlmDm* %22, %aStruct-lcdzRdlmDm** %24
	%25 = load %aStruct-lcdzRdlmDm*, %aStruct-lcdzRdlmDm** %24
	%26 = bitcast %aStruct-lcdzRdlmDm* %25 to i32*
	%27 = alloca i32*
	store i32* %26, i32** %27
	%28 = load i32*, i32** %27
	%29 = getelementptr inbounds i32, i32* %28, i32 0
	%30 = load i32, i32* %29
	%31 = getelementptr inbounds %aStruct-lcdzRdlmDm, %aStruct-lcdzRdlmDm* %22, i32 0, i32 0
	%32 = load i32, i32* %31
	%33 = icmp ne i32 %30, %32
	br i1 %33, label %if.then.onHejoMMvs, label %if.else.WFhqhlPIfQ

if.then.onHejoMMvs:
	%34 = getelementptr inbounds [30 x i8], [30 x i8]* @string.literal.eTRwDylQIo, i32 0, i32 0
	%35 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%36 = call i32 (i8*, ...) @printf(i8* %35, i8* %34)
	br label %lastLeave.PhKpleJxXB

if.else.WFhqhlPIfQ:
	br label %lastLeave.PhKpleJxXB

lastLeave.PhKpleJxXB:
	%37 = load i32*, i32** %27
	%38 = bitcast i32* %37 to %aStruct-lcdzRdlmDm*
	%39 = alloca %aStruct-lcdzRdlmDm
	%40 = load %aStruct-lcdzRdlmDm, %aStruct-lcdzRdlmDm* %38
	store %aStruct-lcdzRdlmDm %40, %aStruct-lcdzRdlmDm* %39
	%41 = getelementptr inbounds %aStruct-lcdzRdlmDm, %aStruct-lcdzRdlmDm* %39, i32 0, i32 0
	%42 = load i32, i32* %41
	%43 = getelementptr inbounds %aStruct-lcdzRdlmDm, %aStruct-lcdzRdlmDm* %22, i32 0, i32 0
	%44 = load i32, i32* %43
	%45 = icmp ne i32 %42, %44
	br i1 %45, label %if.then.jcmqtgnQeD, label %if.else.VmnEoEtUMY

if.then.jcmqtgnQeD:
	%46 = getelementptr inbounds [25 x i8], [25 x i8]* @string.literal.CgwbdXFXPr, i32 0, i32 0
	%47 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%48 = call i32 (i8*, ...) @printf(i8* %47, i8* %46)
	br label %lastLeave.MPObkhTATa

if.else.VmnEoEtUMY:
	br label %lastLeave.MPObkhTATa

lastLeave.MPObkhTATa:
	ret i32 0
}

declare ccc i32 @printf(i8* %0, ...)
