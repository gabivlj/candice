%HelloWorlds = type { i32 }
%FILE = type {}
%Point = type { i32, i32 }
%Points = type { [300 x %Point] }
%Player = type { %Point, i32, i8* }

@"%d " = global [4 x i8] c"%d \00"
@string.literal.PueXlFWiio = global [8 x i8] c"./aFile\00"
@string.literal.vTknroeTDW = global [3 x i8] c"a+\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.GhCAwTurpF = global [7 x i8] c"hello!\00"
@string.literal.DSRXXeXXtX = global [12 x i8] c"Hello world\00"
@string.literal.hwSKbVHNWp = global [6 x i8] c"notok\00"
@string.literal.gEjvmGoACf = global [6 x i8] c"notok\00"
@string.literal.ZJYAruHZWS = global [6 x i8] c"notok\00"
@string.literal.bLQxECJlqI = global [6 x i8] c"notok\00"
@string.literal.mbponPxlhI = global [3 x i8] c"ok\00"
@string.literal.IPwFRgvcMF = global [5 x i8] c"nice\00"

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
	%7 = getelementptr [8 x i8], [8 x i8]* @string.literal.PueXlFWiio, i32 0, i32 0
	%8 = getelementptr [3 x i8], [3 x i8]* @string.literal.vTknroeTDW, i32 0, i32 0
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
	br label %for.declaration.EcAJLvHZTD

leave.YkmPesIYik:
	%18 = getelementptr [7 x i8], [7 x i8]* @string.literal.GhCAwTurpF, i32 0, i32 0
	%19 = sext i32 1 to i64
	%20 = sext i32 6 to i64
	%21 = load %FILE*, %FILE** %10
	call void @fwrite(i8* %18, i64 %19, i64 %20, %FILE* %21)
	%22 = load i8*, i8** %17
	call void @free(i8* %22)
	%23 = load %FILE*, %FILE** %10
	call void @fclose(%FILE* %23)
	ret void

for.declaration.EcAJLvHZTD:
	%24 = load i8*, i8** %17
	%25 = load %FILE*, %FILE** %10
	%26 = call i8* @fgets(i8* %24, i32 10, %FILE* %25)
	%27 = ptrtoint i8* %26 to i32
	%28 = icmp ne i32 %27, 0
	br i1 %28, label %for.block.avkyXiRAlP, label %leave.YkmPesIYik

for.condition.QHuMeuncii:
	%29 = load i8*, i8** %17
	%30 = load %FILE*, %FILE** %10
	%31 = call i8* @fgets(i8* %29, i32 10, %FILE* %30)
	%32 = ptrtoint i8* %31 to i32
	%33 = icmp ne i32 %32, 0
	br i1 %33, label %for.block.avkyXiRAlP, label %leave.YkmPesIYik

for.block.avkyXiRAlP:
	%34 = load i8*, i8** %17
	%35 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%36 = call i32 (i8*, ...) @printf(i8* %35, i8* %34)
	br label %for.update.YpJKZitjTU

for.update.YpJKZitjTU:
	br label %for.condition.QHuMeuncii
}

define ccc void @main() {
main:
	%0 = alloca float
	store float 0x3FF3333320000000, float* %0
	%1 = load float, float* %0
	%2 = fptosi float %1 to i16
	%3 = alloca i16
	store i16 %2, i16* %3
	%4 = load i16, i16* %3
	%5 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%6 = call i32 (i8*, ...) @printf(i8* %5, i16 %4)
	%7 = alloca %Player
	%8 = alloca %Point
	%9 = getelementptr %Point, %Point* %8, i32 0, i32 0
	store i32 3, i32* %9
	%10 = getelementptr %Point, %Point* %8, i32 0, i32 1
	store i32 3, i32* %10
	%11 = getelementptr %Player, %Player* %7, i32 0, i32 0
	%12 = load %Point, %Point* %8
	store %Point %12, %Point* %11
	%13 = getelementptr %Player, %Player* %7, i32 0, i32 1
	store i32 35, i32* %13
	%14 = getelementptr [12 x i8], [12 x i8]* @string.literal.DSRXXeXXtX, i32 0, i32 0
	%15 = getelementptr %Player, %Player* %7, i32 0, i32 2
	store i8* %14, i8** %15
	%16 = getelementptr %Player, %Player* %7, i32 0, i32 2
	%17 = load i8*, i8** %16
	%18 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%19 = call i32 (i8*, ...) @printf(i8* %18, i8* %17)
	%20 = getelementptr %Player, %Player* %7, i32 0, i32 0
	%21 = getelementptr %Point, %Point* %20, i32 0, i32 0
	%22 = getelementptr %Player, %Player* %7, i32 0, i32 0
	%23 = getelementptr %Point, %Point* %22, i32 0, i32 1
	%24 = load i32, i32* %23
	%25 = add i32 %24, 5
	store i32 %25, i32* %21
	%26 = getelementptr %Player, %Player* %7, i32 0, i32 0
	%27 = getelementptr %Point, %Point* %26, i32 0, i32 0
	%28 = load i32, i32* %27
	%29 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%30 = call i32 (i8*, ...) @printf(i8* %29, i32 %28)
	br label %for.declaration.UmsRhzlXCe

leave.XKvdDIzNQp:
	%31 = icmp ne i32 1, 0
	br i1 %31, label %if.then.rOzarfmQwz, label %if.else.eawvPNdqAC

for.declaration.UmsRhzlXCe:
	%32 = alloca i32
	store i32 0, i32* %32
	%33 = load i32, i32* %32
	%34 = icmp sle i32 %33, 100
	br i1 %34, label %for.block.nIobQHqads, label %leave.XKvdDIzNQp

for.condition.saWSNyWaSX:
	%35 = load i32, i32* %32
	%36 = icmp sle i32 %35, 100
	br i1 %36, label %for.block.nIobQHqads, label %leave.XKvdDIzNQp

for.block.nIobQHqads:
	%37 = load i32, i32* %32
	%38 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%39 = call i32 (i8*, ...) @printf(i8* %38, i32 %37)
	br label %for.declaration.wdhAVljmOK

for.update.eLxDXVblqI:
	%40 = load i32, i32* %32
	%41 = add i32 %40, 1
	store i32 %41, i32* %32
	br label %for.condition.saWSNyWaSX

leave.PvnIgBFlvD:
	br label %for.update.eLxDXVblqI

for.declaration.wdhAVljmOK:
	%42 = alloca i32
	store i32 0, i32* %42
	%43 = load i32, i32* %42
	%44 = icmp sle i32 %43, 100
	br i1 %44, label %for.block.oeIktuVcQX, label %leave.PvnIgBFlvD

for.condition.caLbJNEtXG:
	%45 = load i32, i32* %42
	%46 = icmp sle i32 %45, 100
	br i1 %46, label %for.block.oeIktuVcQX, label %leave.PvnIgBFlvD

for.block.oeIktuVcQX:
	%47 = load i32, i32* %32
	%48 = load i32, i32* %42
	%49 = add i32 %47, %48
	%50 = icmp sle i32 %49, 10
	br i1 %50, label %if.then.gMBjtHEwGr, label %if.else.FjiBjRIGSA

for.update.YvgBrjXxxL:
	%51 = load i32, i32* %42
	%52 = add i32 %51, 1
	store i32 %52, i32* %42
	br label %for.condition.caLbJNEtXG

if.then.gMBjtHEwGr:
	%53 = load i32, i32* %32
	%54 = load i32, i32* %42
	%55 = add i32 %53, %54
	%56 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%57 = call i32 (i8*, ...) @printf(i8* %56, i32 %55)
	br label %lastLeave.HEosnTOERo

if.else.FjiBjRIGSA:
	br label %lastLeave.HEosnTOERo

lastLeave.HEosnTOERo:
	br label %for.update.YvgBrjXxxL

if.then.rOzarfmQwz:
	%58 = icmp ne i32 0, 0
	%59 = icmp ne i32 0, 0
	%60 = icmp ne i32 0, 0
	%61 = icmp ne i32 1, 0
	br i1 %58, label %if.then.AkFAWzrRFP, label %leave.BdacOwDVTB

if.then.AkFAWzrRFP:
	%62 = getelementptr [6 x i8], [6 x i8]* @string.literal.hwSKbVHNWp, i32 0, i32 0
	%63 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%64 = call i32 (i8*, ...) @printf(i8* %63, i8* %62)
	br label %lastLeave.GGMYJOpUiJ

if.else.KyqXCBmnMq:
	%65 = getelementptr [6 x i8], [6 x i8]* @string.literal.gEjvmGoACf, i32 0, i32 0
	%66 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%67 = call i32 (i8*, ...) @printf(i8* %66, i8* %65)
	br label %lastLeave.GGMYJOpUiJ

elseif.then.AnUPOSbteW:
	%68 = getelementptr [6 x i8], [6 x i8]* @string.literal.ZJYAruHZWS, i32 0, i32 0
	%69 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%70 = call i32 (i8*, ...) @printf(i8* %69, i8* %68)
	br label %lastLeave.GGMYJOpUiJ

elseif.then.kCfRNVPqBe:
	%71 = getelementptr [6 x i8], [6 x i8]* @string.literal.bLQxECJlqI, i32 0, i32 0
	%72 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%73 = call i32 (i8*, ...) @printf(i8* %72, i8* %71)
	br label %lastLeave.GGMYJOpUiJ

elseif.then.cNHDwiIWXX:
	%74 = getelementptr [3 x i8], [3 x i8]* @string.literal.mbponPxlhI, i32 0, i32 0
	%75 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%76 = call i32 (i8*, ...) @printf(i8* %75, i8* %74)
	br label %lastLeave.GGMYJOpUiJ

leave.BdacOwDVTB:
	br i1 %59, label %elseif.then.AnUPOSbteW, label %leave.sJizqMDXVz

leave.sJizqMDXVz:
	br i1 %60, label %elseif.then.kCfRNVPqBe, label %leave.SAsmafeUgb

leave.SAsmafeUgb:
	br i1 %61, label %elseif.then.cNHDwiIWXX, label %if.else.KyqXCBmnMq

lastLeave.GGMYJOpUiJ:
	%77 = getelementptr [5 x i8], [5 x i8]* @string.literal.IPwFRgvcMF, i32 0, i32 0
	%78 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%79 = call i32 (i8*, ...) @printf(i8* %78, i8* %77)
	br label %lastLeave.DAiDWuNNuS

if.else.eawvPNdqAC:
	br label %lastLeave.DAiDWuNNuS

lastLeave.DAiDWuNNuS:
	%80 = alloca [300 x [300 x i32]]
	br label %for.declaration.lSJNDNYuDM

leave.NQzJrVUTIT:
	%81 = getelementptr [300 x [300 x i32]], [300 x [300 x i32]]* %80, i32 0, i32 200
	%82 = getelementptr [300 x i32], [300 x i32]* %81, i32 0, i32 200
	%83 = load i32, i32* %82
	%84 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%85 = call i32 (i8*, ...) @printf(i8* %84, i32 %83)
	%86 = getelementptr [300 x [300 x i32]], [300 x [300 x i32]]* %80, i32 0, i32 201
	%87 = getelementptr [300 x i32], [300 x i32]* %86, i32 0, i32 201
	%88 = load i32, i32* %87
	%89 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%90 = call i32 (i8*, ...) @printf(i8* %89, i32 %88)
	ret void

for.declaration.lSJNDNYuDM:
	%91 = alloca i32
	store i32 0, i32* %91
	%92 = load i32, i32* %91
	%93 = icmp slt i32 %92, 300
	br i1 %93, label %for.block.KdIYPmZlEX, label %leave.NQzJrVUTIT

for.condition.hIoFPgNWMP:
	%94 = load i32, i32* %91
	%95 = icmp slt i32 %94, 300
	br i1 %95, label %for.block.KdIYPmZlEX, label %leave.NQzJrVUTIT

for.block.KdIYPmZlEX:
	br label %for.declaration.rdXZAfRaZx

for.update.aHhsXtZDHc:
	%96 = load i32, i32* %91
	%97 = add i32 %96, 1
	store i32 %97, i32* %91
	br label %for.condition.hIoFPgNWMP

leave.JVoudsvLyj:
	br label %for.update.aHhsXtZDHc

for.declaration.rdXZAfRaZx:
	%98 = alloca i32
	store i32 0, i32* %98
	%99 = load i32, i32* %98
	%100 = icmp slt i32 %99, 300
	br i1 %100, label %for.block.ayhuhrgcvt, label %leave.JVoudsvLyj

for.condition.VgeIQOVXmn:
	%101 = load i32, i32* %98
	%102 = icmp slt i32 %101, 300
	br i1 %102, label %for.block.ayhuhrgcvt, label %leave.JVoudsvLyj

for.block.ayhuhrgcvt:
	%103 = load i32, i32* %91
	%104 = icmp eq i32 %103, 200
	br i1 %104, label %if.then.TVFsTrCHBs, label %if.else.boJMkwSnnG

for.update.RlyggMQKuq:
	%105 = load i32, i32* %98
	%106 = add i32 %105, 1
	store i32 %106, i32* %98
	br label %for.condition.VgeIQOVXmn

if.then.TVFsTrCHBs:
	%107 = load i32, i32* %91
	%108 = getelementptr [300 x [300 x i32]], [300 x [300 x i32]]* %80, i32 0, i32 %107
	%109 = load i32, i32* %98
	%110 = getelementptr [300 x i32], [300 x i32]* %108, i32 0, i32 %109
	store i32 1, i32* %110
	br label %lastLeave.uGiNviqtHC

if.else.boJMkwSnnG:
	%111 = load i32, i32* %91
	%112 = getelementptr [300 x [300 x i32]], [300 x [300 x i32]]* %80, i32 0, i32 %111
	%113 = load i32, i32* %98
	%114 = getelementptr [300 x i32], [300 x i32]* %112, i32 0, i32 %113
	%115 = load i32, i32* %91
	%116 = load i32, i32* %98
	%117 = add i32 %115, %116
	store i32 %117, i32* %114
	br label %lastLeave.uGiNviqtHC

lastLeave.uGiNviqtHC:
	br label %for.update.RlyggMQKuq
}
