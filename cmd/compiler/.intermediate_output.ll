%HelloWorlds = type { i32 }
%FILE = type {}
%Point = type { i32, i32 }
%Points = type { [300 x %Point] }
%Player = type { %Point, i32, i8* }

@"%d " = global [4 x i8] c"%d \00"
@string.literal.rwLeGRRqdN = global [8 x i8] c"./aFile\00"
@string.literal.CXCfXizldm = global [3 x i8] c"a+\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.nPACqrtGtL = global [7 x i8] c"hello!\00"
@string.literal.tGWoHgNBiU = global [12 x i8] c"Hello world\00"
@string.literal.XjbevbPhCY = global [6 x i8] c"notok\00"
@string.literal.kkPyWsgDzl = global [6 x i8] c"notok\00"
@string.literal.LJtxwxJJct = global [6 x i8] c"notok\00"
@string.literal.lXtQaNiNkW = global [6 x i8] c"notok\00"
@string.literal.hRDifGCSTU = global [3 x i8] c"ok\00"
@string.literal.hbOaItKuZX = global [5 x i8] c"nice\00"
@string.literal.vLUAUgFKTN = global [2 x i8] c"w\00"
@string.literal.StlfCYcfNH = global [2 x i8] c"q\00"
@string.literal.AhkjoSVsJR = global [2 x i8] c"e\00"

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
	%7 = getelementptr [8 x i8], [8 x i8]* @string.literal.rwLeGRRqdN, i32 0, i32 0
	%8 = getelementptr [3 x i8], [3 x i8]* @string.literal.CXCfXizldm, i32 0, i32 0
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
	br label %for.declaration.cQRkmuOaHf

leave.ZgSNmRugqO:
	%18 = getelementptr [7 x i8], [7 x i8]* @string.literal.nPACqrtGtL, i32 0, i32 0
	%19 = sext i32 1 to i64
	%20 = sext i32 6 to i64
	%21 = load %FILE*, %FILE** %10
	call void @fwrite(i8* %18, i64 %19, i64 %20, %FILE* %21)
	%22 = load i8*, i8** %17
	call void @free(i8* %22)
	%23 = load %FILE*, %FILE** %10
	call void @fclose(%FILE* %23)
	ret void

for.declaration.cQRkmuOaHf:
	%24 = load i8*, i8** %17
	%25 = load %FILE*, %FILE** %10
	%26 = call i8* @fgets(i8* %24, i32 10, %FILE* %25)
	%27 = ptrtoint i8* %26 to i32
	%28 = icmp ne i32 %27, 0
	br i1 %28, label %for.block.soVeOJyLqg, label %leave.ZgSNmRugqO

for.condition.EdGQqtGSMK:
	%29 = load i8*, i8** %17
	%30 = load %FILE*, %FILE** %10
	%31 = call i8* @fgets(i8* %29, i32 10, %FILE* %30)
	%32 = ptrtoint i8* %31 to i32
	%33 = icmp ne i32 %32, 0
	br i1 %33, label %for.block.soVeOJyLqg, label %leave.ZgSNmRugqO

for.block.soVeOJyLqg:
	%34 = load i8*, i8** %17
	%35 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%36 = call i32 (i8*, ...) @printf(i8* %35, i8* %34)
	br label %for.update.GBRaUFiqMn

for.update.GBRaUFiqMn:
	br label %for.condition.EdGQqtGSMK
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
	%14 = getelementptr [12 x i8], [12 x i8]* @string.literal.tGWoHgNBiU, i32 0, i32 0
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
	br label %for.declaration.HgJCanRxmI

leave.gEunVREtRD:
	%31 = icmp ne i32 1, 0
	br i1 %31, label %if.then.xsvwdTRJXI, label %if.else.DpqraBusdP

for.declaration.HgJCanRxmI:
	%32 = alloca i32
	store i32 0, i32* %32
	%33 = load i32, i32* %32
	%34 = icmp sle i32 %33, 100
	br i1 %34, label %for.block.NpXjpzvGVw, label %leave.gEunVREtRD

for.condition.pYlyOqrUIe:
	%35 = load i32, i32* %32
	%36 = icmp sle i32 %35, 100
	br i1 %36, label %for.block.NpXjpzvGVw, label %leave.gEunVREtRD

for.block.NpXjpzvGVw:
	%37 = load i32, i32* %32
	%38 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%39 = call i32 (i8*, ...) @printf(i8* %38, i32 %37)
	br label %for.declaration.zTMUOZYtBp

for.update.iUZFbFvIxL:
	%40 = load i32, i32* %32
	%41 = add i32 %40, 1
	store i32 %41, i32* %32
	br label %for.condition.pYlyOqrUIe

leave.IMLXxnJJum:
	br label %for.update.iUZFbFvIxL

for.declaration.zTMUOZYtBp:
	%42 = alloca i32
	store i32 0, i32* %42
	%43 = load i32, i32* %42
	%44 = icmp sle i32 %43, 100
	br i1 %44, label %for.block.RwsAeAClIF, label %leave.IMLXxnJJum

for.condition.PrtPwrbHyv:
	%45 = load i32, i32* %42
	%46 = icmp sle i32 %45, 100
	br i1 %46, label %for.block.RwsAeAClIF, label %leave.IMLXxnJJum

for.block.RwsAeAClIF:
	%47 = load i32, i32* %32
	%48 = load i32, i32* %42
	%49 = add i32 %47, %48
	%50 = icmp sle i32 %49, 10
	br i1 %50, label %if.then.FhwfrgIxXU, label %if.else.hxtmccBlkZ

for.update.cEAECvcnBt:
	%51 = load i32, i32* %42
	%52 = add i32 %51, 1
	store i32 %52, i32* %42
	br label %for.condition.PrtPwrbHyv

if.then.FhwfrgIxXU:
	%53 = load i32, i32* %32
	%54 = load i32, i32* %42
	%55 = add i32 %53, %54
	%56 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%57 = call i32 (i8*, ...) @printf(i8* %56, i32 %55)
	br label %lastLeave.vJuKqsNmzL

if.else.hxtmccBlkZ:
	br label %lastLeave.vJuKqsNmzL

lastLeave.vJuKqsNmzL:
	br label %for.update.cEAECvcnBt

if.then.xsvwdTRJXI:
	%58 = icmp ne i32 0, 0
	%59 = icmp ne i32 0, 0
	%60 = icmp ne i32 0, 0
	%61 = icmp ne i32 1, 0
	br i1 %58, label %if.then.SObDVweoBn, label %leave.KrQXNZcpwS

if.then.SObDVweoBn:
	%62 = getelementptr [6 x i8], [6 x i8]* @string.literal.XjbevbPhCY, i32 0, i32 0
	%63 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%64 = call i32 (i8*, ...) @printf(i8* %63, i8* %62)
	br label %lastLeave.zsmIAAWjHb

if.else.jRcqXUoiWD:
	%65 = getelementptr [6 x i8], [6 x i8]* @string.literal.kkPyWsgDzl, i32 0, i32 0
	%66 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%67 = call i32 (i8*, ...) @printf(i8* %66, i8* %65)
	br label %lastLeave.zsmIAAWjHb

elseif.then.muhnJQpwkA:
	%68 = getelementptr [6 x i8], [6 x i8]* @string.literal.LJtxwxJJct, i32 0, i32 0
	%69 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%70 = call i32 (i8*, ...) @printf(i8* %69, i8* %68)
	br label %lastLeave.zsmIAAWjHb

elseif.then.hjwjQYOUmn:
	%71 = getelementptr [6 x i8], [6 x i8]* @string.literal.lXtQaNiNkW, i32 0, i32 0
	%72 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%73 = call i32 (i8*, ...) @printf(i8* %72, i8* %71)
	br label %lastLeave.zsmIAAWjHb

elseif.then.ixovPXqHoC:
	%74 = getelementptr [3 x i8], [3 x i8]* @string.literal.hRDifGCSTU, i32 0, i32 0
	%75 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%76 = call i32 (i8*, ...) @printf(i8* %75, i8* %74)
	br label %lastLeave.zsmIAAWjHb

leave.KrQXNZcpwS:
	br i1 %59, label %elseif.then.muhnJQpwkA, label %leave.VehJJCkZNo

leave.VehJJCkZNo:
	br i1 %60, label %elseif.then.hjwjQYOUmn, label %leave.akLvuFgIrA

leave.akLvuFgIrA:
	br i1 %61, label %elseif.then.ixovPXqHoC, label %if.else.jRcqXUoiWD

lastLeave.zsmIAAWjHb:
	%77 = getelementptr [5 x i8], [5 x i8]* @string.literal.hbOaItKuZX, i32 0, i32 0
	%78 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%79 = call i32 (i8*, ...) @printf(i8* %78, i8* %77)
	br label %lastLeave.HELRZjeZMU

if.else.DpqraBusdP:
	br label %lastLeave.HELRZjeZMU

lastLeave.HELRZjeZMU:
	%80 = alloca [300 x [300 x i32]]
	br label %for.declaration.HgRYsROwIS

leave.HtNfSWfJdK:
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
	%91 = trunc i32 1 to i1
	%92 = alloca i1
	store i1 %91, i1* %92
	%93 = trunc i32 1 to i1
	%94 = alloca i1
	store i1 %93, i1* %94
	%95 = icmp ne i32 0, 0
	br i1 %95, label %if.then.DnRdYmuVeT, label %if.else.ZsPDfiZMIv

for.declaration.HgRYsROwIS:
	%96 = alloca i32
	store i32 0, i32* %96
	%97 = load i32, i32* %96
	%98 = icmp slt i32 %97, 300
	br i1 %98, label %for.block.YKSFTHqDdV, label %leave.HtNfSWfJdK

for.condition.DbERSuUnpt:
	%99 = load i32, i32* %96
	%100 = icmp slt i32 %99, 300
	br i1 %100, label %for.block.YKSFTHqDdV, label %leave.HtNfSWfJdK

for.block.YKSFTHqDdV:
	br label %for.declaration.pmSNZeDhNp

for.update.EpwrOHvJCj:
	%101 = load i32, i32* %96
	%102 = add i32 %101, 1
	store i32 %102, i32* %96
	br label %for.condition.DbERSuUnpt

leave.yxuMWhZDaW:
	br label %for.update.EpwrOHvJCj

for.declaration.pmSNZeDhNp:
	%103 = alloca i32
	store i32 0, i32* %103
	%104 = load i32, i32* %103
	%105 = icmp slt i32 %104, 300
	br i1 %105, label %for.block.OMSTsVNakH, label %leave.yxuMWhZDaW

for.condition.NJwbfhWpLK:
	%106 = load i32, i32* %103
	%107 = icmp slt i32 %106, 300
	br i1 %107, label %for.block.OMSTsVNakH, label %leave.yxuMWhZDaW

for.block.OMSTsVNakH:
	%108 = load i32, i32* %96
	%109 = icmp eq i32 %108, 200
	br i1 %109, label %if.then.HnbtiIUQsz, label %if.else.IRhJsIJliy

for.update.IVpYzLFvyj:
	%110 = load i32, i32* %103
	%111 = add i32 %110, 1
	store i32 %111, i32* %103
	br label %for.condition.NJwbfhWpLK

if.then.HnbtiIUQsz:
	%112 = load i32, i32* %96
	%113 = getelementptr [300 x [300 x i32]], [300 x [300 x i32]]* %80, i32 0, i32 %112
	%114 = load i32, i32* %103
	%115 = getelementptr [300 x i32], [300 x i32]* %113, i32 0, i32 %114
	store i32 1, i32* %115
	br label %lastLeave.QITtQzMdYN

if.else.IRhJsIJliy:
	%116 = load i32, i32* %96
	%117 = getelementptr [300 x [300 x i32]], [300 x [300 x i32]]* %80, i32 0, i32 %116
	%118 = load i32, i32* %103
	%119 = getelementptr [300 x i32], [300 x i32]* %117, i32 0, i32 %118
	%120 = load i32, i32* %96
	%121 = load i32, i32* %103
	%122 = add i32 %120, %121
	store i32 %122, i32* %119
	br label %lastLeave.QITtQzMdYN

lastLeave.QITtQzMdYN:
	br label %for.update.IVpYzLFvyj

if.then.DnRdYmuVeT:
	%123 = icmp ne i32 1, 0
	br i1 %123, label %if.then.MRHRugTNri, label %if.else.qsCkCjPdGd

if.then.MRHRugTNri:
	%124 = getelementptr [2 x i8], [2 x i8]* @string.literal.vLUAUgFKTN, i32 0, i32 0
	%125 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%126 = call i32 (i8*, ...) @printf(i8* %125, i8* %124)
	br label %lastLeave.qbjjDQTMPK

if.else.qsCkCjPdGd:
	%127 = getelementptr [2 x i8], [2 x i8]* @string.literal.StlfCYcfNH, i32 0, i32 0
	%128 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%129 = call i32 (i8*, ...) @printf(i8* %128, i8* %127)
	br label %lastLeave.qbjjDQTMPK

lastLeave.qbjjDQTMPK:
	br label %lastLeave.nlQpTJJBzz

if.else.ZsPDfiZMIv:
	%130 = getelementptr [2 x i8], [2 x i8]* @string.literal.AhkjoSVsJR, i32 0, i32 0
	%131 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%132 = call i32 (i8*, ...) @printf(i8* %131, i8* %130)
	br label %lastLeave.nlQpTJJBzz

lastLeave.nlQpTJJBzz:
	ret void
}
