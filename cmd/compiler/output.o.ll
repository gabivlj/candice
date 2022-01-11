%HelloWorlds = type { i32 }
%FILE = type {}
%Point = type { i32, i32 }
%Points = type { [300 x %Point] }
%Player = type { %Point, i32, i8* }

@"%d " = global [4 x i8] c"%d \00"
@string.literal.OPOVxskFWk = global [8 x i8] c"./aFile\00"
@string.literal.uBnckAyoAl = global [3 x i8] c"a+\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.oIkQONXRCA = global [7 x i8] c"hello!\00"
@string.literal.PiswxNbCYw = global [12 x i8] c"Hello world\00"
@string.literal.zHIFmKFBKn = global [4 x i8] c"2ok\00"
@string.literal.yqcRXVyxTN = global [4 x i8] c"2ok\00"
@string.literal.QXFEnpUOLs = global [4 x i8] c"2ok\00"
@string.literal.HDRwRpyYaI = global [4 x i8] c"2ok\00"
@string.literal.XyYOZIQbyh = global [3 x i8] c"ok\00"
@string.literal.mYcKHNWzEm = global [5 x i8] c"nice\00"

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

declare ccc %HelloWorlds* @helloWorldsFunction()

declare ccc %FILE* @fopen(i8* %0, i8* %1)

declare ccc void @fwrite(i8* %0, i64 %1, i64 %2, %FILE* %3)

declare ccc void @fclose(%FILE* %0)

declare ccc i64 @fread(i8* %0, i64 %1, i64 %2, %FILE* %3)

declare ccc i64 @rewind(%FILE* %0)

declare ccc i8* @fgets(i8* %0, i32 %1, %FILE* %2)

declare ccc void @free(i8* %0)

define ccc void @doStuff() {
doStuff:
	%0 = call %HelloWorlds* @helloWorldsFunction()
	%1 = alloca %HelloWorlds*
	store %HelloWorlds* %0, %HelloWorlds** %1
	%2 = load %HelloWorlds*, %HelloWorlds** %1
	%3 = getelementptr %HelloWorlds, %HelloWorlds* %2, i32 0, i32 0
	%4 = load i32, i32* %3
	%5 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%6 = call i32 (i8*, ...) @printf(i8* %5, i32 %4)
	%7 = getelementptr [8 x i8], [8 x i8]* @string.literal.OPOVxskFWk, i32 0, i32 0
	%8 = getelementptr [3 x i8], [3 x i8]* @string.literal.uBnckAyoAl, i32 0, i32 0
	%9 = call %FILE* @fopen(i8* %7, i8* %8)
	%10 = alloca %FILE*
	store %FILE* %9, %FILE** %10
	%11 = load %FILE*, %FILE** %10
	%12 = call i64 @rewind(%FILE* %11)
	%13 = sext i32 300 to i64
	%14 = mul i64 %13, 1
	%15 = call i8* @malloc(i64 %14)
	%16 = bitcast i8* %15 to i8*
	%17 = alloca i8*
	store i8* %16, i8** %17
	br label %for.declaration.VgQdgOmfuv

leave.xhCFMadgPR:
	%18 = getelementptr [7 x i8], [7 x i8]* @string.literal.oIkQONXRCA, i32 0, i32 0
	%19 = sext i32 1 to i64
	%20 = sext i32 6 to i64
	%21 = load %FILE*, %FILE** %10
	call void @fwrite(i8* %18, i64 %19, i64 %20, %FILE* %21)
	%22 = load i8*, i8** %17
	call void @free(i8* %22)
	%23 = load %FILE*, %FILE** %10
	call void @fclose(%FILE* %23)
	ret void

for.declaration.VgQdgOmfuv:
	%24 = load i8*, i8** %17
	%25 = load %FILE*, %FILE** %10
	%26 = call i8* @fgets(i8* %24, i32 10, %FILE* %25)
	%27 = ptrtoint i8* %26 to i32
	%28 = icmp ne i32 %27, 0
	br i1 %28, label %for.block.UDygIlpGgC, label %leave.xhCFMadgPR

for.condition.iczaTaxsqF:
	%29 = load i8*, i8** %17
	%30 = load %FILE*, %FILE** %10
	%31 = call i8* @fgets(i8* %29, i32 10, %FILE* %30)
	%32 = ptrtoint i8* %31 to i32
	%33 = icmp ne i32 %32, 0
	br i1 %33, label %for.block.UDygIlpGgC, label %leave.xhCFMadgPR

for.block.UDygIlpGgC:
	%34 = load i8*, i8** %17
	%35 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%36 = call i32 (i8*, ...) @printf(i8* %35, i8* %34)
	br label %for.update.gBDBoHNviA

for.update.gBDBoHNviA:
	br label %for.condition.iczaTaxsqF
}

define ccc void @main() {
main:
	%0 = alloca %Player
	%1 = alloca %Point
	%2 = getelementptr %Point, %Point* %1, i32 0, i32 0
	store i32 3, i32* %2
	%3 = getelementptr %Point, %Point* %1, i32 0, i32 1
	store i32 3, i32* %3
	%4 = getelementptr %Player, %Player* %0, i32 0, i32 0
	%5 = load %Point, %Point* %1
	store %Point %5, %Point* %4
	%6 = getelementptr %Player, %Player* %0, i32 0, i32 1
	store i32 35, i32* %6
	%7 = getelementptr [12 x i8], [12 x i8]* @string.literal.PiswxNbCYw, i32 0, i32 0
	%8 = getelementptr %Player, %Player* %0, i32 0, i32 2
	store i8* %7, i8** %8
	%9 = getelementptr %Player, %Player* %0, i32 0, i32 2
	%10 = load i8*, i8** %9
	%11 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%12 = call i32 (i8*, ...) @printf(i8* %11, i8* %10)
	%13 = getelementptr %Player, %Player* %0, i32 0, i32 0
	%14 = getelementptr %Point, %Point* %13, i32 0, i32 0
	%15 = getelementptr %Player, %Player* %0, i32 0, i32 0
	%16 = getelementptr %Point, %Point* %15, i32 0, i32 1
	%17 = load i32, i32* %16
	%18 = add i32 %17, 5
	store i32 %18, i32* %14
	%19 = getelementptr %Player, %Player* %0, i32 0, i32 0
	%20 = getelementptr %Point, %Point* %19, i32 0, i32 0
	%21 = load i32, i32* %20
	%22 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%23 = call i32 (i8*, ...) @printf(i8* %22, i32 %21)
	br label %for.declaration.raerakxfpX

leave.CgMbnduOAO:
	%24 = icmp ne i32 1, 0
	br i1 %24, label %if.then.VcxmNpBPdU, label %if.else.jhrxIGlzAI

for.declaration.raerakxfpX:
	%25 = alloca i32
	store i32 0, i32* %25
	%26 = load i32, i32* %25
	%27 = icmp sle i32 %26, 100
	br i1 %27, label %for.block.AQxWIRmRTL, label %leave.CgMbnduOAO

for.condition.lZvmFgIMTU:
	%28 = load i32, i32* %25
	%29 = icmp sle i32 %28, 100
	br i1 %29, label %for.block.AQxWIRmRTL, label %leave.CgMbnduOAO

for.block.AQxWIRmRTL:
	%30 = load i32, i32* %25
	%31 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%32 = call i32 (i8*, ...) @printf(i8* %31, i32 %30)
	br label %for.declaration.WetdVUhxYv

for.update.SayYCbxcqP:
	%33 = load i32, i32* %25
	%34 = add i32 %33, 1
	store i32 %34, i32* %25
	br label %for.condition.lZvmFgIMTU

leave.zRFkgyvkHw:
	br label %for.update.SayYCbxcqP

for.declaration.WetdVUhxYv:
	%35 = alloca i32
	store i32 0, i32* %35
	%36 = load i32, i32* %35
	%37 = icmp sle i32 %36, 100
	br i1 %37, label %for.block.LjzuwOcQCG, label %leave.zRFkgyvkHw

for.condition.SvqeskxECk:
	%38 = load i32, i32* %35
	%39 = icmp sle i32 %38, 100
	br i1 %39, label %for.block.LjzuwOcQCG, label %leave.zRFkgyvkHw

for.block.LjzuwOcQCG:
	%40 = load i32, i32* %25
	%41 = load i32, i32* %35
	%42 = add i32 %40, %41
	%43 = icmp sle i32 %42, 10
	br i1 %43, label %if.then.HuVRCvXWcq, label %if.else.bQVQHrcmut

for.update.yiVVoqsTTo:
	%44 = load i32, i32* %35
	%45 = add i32 %44, 1
	store i32 %45, i32* %35
	br label %for.condition.SvqeskxECk

if.then.HuVRCvXWcq:
	%46 = load i32, i32* %25
	%47 = load i32, i32* %35
	%48 = add i32 %46, %47
	%49 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%50 = call i32 (i8*, ...) @printf(i8* %49, i32 %48)
	br label %lastLeave.QUylTHHqnW

if.else.bQVQHrcmut:
	br label %lastLeave.QUylTHHqnW

lastLeave.QUylTHHqnW:
	br label %for.update.yiVVoqsTTo

if.then.VcxmNpBPdU:
	%51 = icmp ne i32 0, 0
	%52 = icmp ne i32 0, 0
	%53 = icmp ne i32 0, 0
	%54 = icmp ne i32 1, 0
	br i1 %51, label %if.then.pudgqEEkfI, label %leave.NjffhwifxD

if.then.pudgqEEkfI:
	%55 = getelementptr [4 x i8], [4 x i8]* @string.literal.zHIFmKFBKn, i32 0, i32 0
	%56 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%57 = call i32 (i8*, ...) @printf(i8* %56, i8* %55)
	br label %lastLeave.lsfolrBPgC

if.else.skXiCLNQHg:
	%58 = getelementptr [4 x i8], [4 x i8]* @string.literal.yqcRXVyxTN, i32 0, i32 0
	%59 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%60 = call i32 (i8*, ...) @printf(i8* %59, i8* %58)
	br label %lastLeave.lsfolrBPgC

elseif.then.zTReCPqRCq:
	%61 = getelementptr [4 x i8], [4 x i8]* @string.literal.QXFEnpUOLs, i32 0, i32 0
	%62 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%63 = call i32 (i8*, ...) @printf(i8* %62, i8* %61)
	br label %lastLeave.lsfolrBPgC

elseif.then.bsDILNIAQG:
	%64 = getelementptr [4 x i8], [4 x i8]* @string.literal.HDRwRpyYaI, i32 0, i32 0
	%65 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%66 = call i32 (i8*, ...) @printf(i8* %65, i8* %64)
	br label %lastLeave.lsfolrBPgC

elseif.then.EmMxEFdBAt:
	%67 = getelementptr [3 x i8], [3 x i8]* @string.literal.XyYOZIQbyh, i32 0, i32 0
	%68 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%69 = call i32 (i8*, ...) @printf(i8* %68, i8* %67)
	br label %lastLeave.lsfolrBPgC

leave.NjffhwifxD:
	br i1 %52, label %elseif.then.zTReCPqRCq, label %leave.uvgdZksMIn

leave.uvgdZksMIn:
	br i1 %53, label %elseif.then.bsDILNIAQG, label %leave.nyJmxHspPy

leave.nyJmxHspPy:
	br i1 %54, label %elseif.then.EmMxEFdBAt, label %if.else.skXiCLNQHg

lastLeave.lsfolrBPgC:
	%70 = getelementptr [5 x i8], [5 x i8]* @string.literal.mYcKHNWzEm, i32 0, i32 0
	%71 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%72 = call i32 (i8*, ...) @printf(i8* %71, i8* %70)
	br label %lastLeave.uCIMcicchU

if.else.jhrxIGlzAI:
	br label %lastLeave.uCIMcicchU

lastLeave.uCIMcicchU:
	ret void
}
