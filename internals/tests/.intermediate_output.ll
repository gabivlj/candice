@"%d " = global [4 x i8] c"%d \00"
@string.literal.KGJeGJfAZy = global [15 x i8] c"wow this works\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.KzBHaOSpqk = global [16 x i8] c"shouldn't print\00"
@string.literal.znfXCeSLdV = global [16 x i8] c"shouldn't print\00"

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc void @main() {
main:
	br label %for.declaration.adxrIfGqla

leave.jmLrShKhsc:
	%0 = icmp sge i32 6, 5
	br i1 %0, label %if.then.gLbuyLLofO, label %if.else.IlIcsAOQBc

for.declaration.adxrIfGqla:
	%1 = alloca i32
	store i32 0, i32* %1
	%2 = load i32, i32* %1
	%3 = icmp sle i32 %2, 100
	br i1 %3, label %for.block.pdNZmYXwhb, label %leave.jmLrShKhsc

for.condition.bkROhXlMuv:
	%4 = load i32, i32* %1
	%5 = icmp sle i32 %4, 100
	br i1 %5, label %for.block.pdNZmYXwhb, label %leave.jmLrShKhsc

for.block.pdNZmYXwhb:
	%6 = load i32, i32* %1
	%7 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%8 = call i32 (i8*, ...) @printf(i8* %7, i32 %6)
	br label %for.declaration.ZGptEbBGZE

for.update.tDERwOzxMw:
	%9 = load i32, i32* %1
	%10 = add i32 %9, 1
	store i32 %10, i32* %1
	br label %for.condition.bkROhXlMuv

leave.OympTuxkTD:
	br label %for.update.tDERwOzxMw

for.declaration.ZGptEbBGZE:
	%11 = alloca i32
	store i32 0, i32* %11
	%12 = load i32, i32* %11
	%13 = icmp sle i32 %12, 100
	br i1 %13, label %for.block.FAfjIgveei, label %leave.OympTuxkTD

for.condition.ulViVmmSeO:
	%14 = load i32, i32* %11
	%15 = icmp sle i32 %14, 100
	br i1 %15, label %for.block.FAfjIgveei, label %leave.OympTuxkTD

for.block.FAfjIgveei:
	%16 = load i32, i32* %1
	%17 = load i32, i32* %11
	%18 = add i32 %16, %17
	%19 = icmp sle i32 %18, 10
	br i1 %19, label %if.then.pQuYroODxO, label %if.else.bTxHXGOZWF

for.update.JvinWIlSXW:
	%20 = load i32, i32* %11
	%21 = add i32 %20, 1
	store i32 %21, i32* %11
	br label %for.condition.ulViVmmSeO

if.then.pQuYroODxO:
	%22 = load i32, i32* %1
	%23 = load i32, i32* %11
	%24 = add i32 %22, %23
	%25 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%26 = call i32 (i8*, ...) @printf(i8* %25, i32 %24)
	br label %lastLeave.BWGHAqjUTN

if.else.bTxHXGOZWF:
	br label %lastLeave.BWGHAqjUTN

lastLeave.BWGHAqjUTN:
	br label %for.update.JvinWIlSXW

if.then.gLbuyLLofO:
	%27 = icmp sge i32 6, 4
	br i1 %27, label %if.then.YUaaFuToFS, label %if.else.zwmlnLwxSu

if.then.YUaaFuToFS:
	%28 = icmp sge i32 6, 3
	br i1 %28, label %if.then.SIRXAjDcrA, label %if.else.TLfjIXelKn

if.then.SIRXAjDcrA:
	%29 = getelementptr [15 x i8], [15 x i8]* @string.literal.KGJeGJfAZy, i32 0, i32 0
	%30 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%31 = call i32 (i8*, ...) @printf(i8* %30, i8* %29)
	br label %lastLeave.WmdzRXMKBB

if.else.TLfjIXelKn:
	br label %lastLeave.WmdzRXMKBB

lastLeave.WmdzRXMKBB:
	br label %lastLeave.tnMwbAbjJs

if.else.zwmlnLwxSu:
	br label %lastLeave.tnMwbAbjJs

lastLeave.tnMwbAbjJs:
	br label %lastLeave.nXtPKgnZrb

if.else.IlIcsAOQBc:
	br label %lastLeave.nXtPKgnZrb

lastLeave.nXtPKgnZrb:
	br label %for.declaration.BYVhJIgtAT

leave.LyxgTLENsq:
	ret void

for.declaration.BYVhJIgtAT:
	%32 = icmp ne i32 1, 0
	br i1 %32, label %for.block.oiXllzZtJA, label %leave.LyxgTLENsq

for.condition.wNQFvLLzFp:
	%33 = icmp ne i32 1, 0
	br i1 %33, label %for.block.oiXllzZtJA, label %leave.LyxgTLENsq

for.block.oiXllzZtJA:
	%34 = icmp ne i32 1, 0
	br i1 %34, label %if.then.hHKovNZoGX, label %if.else.eaQqwskQge

for.update.EbUlufmAUr:
	br label %for.condition.wNQFvLLzFp

if.then.hHKovNZoGX:
	br label %for.declaration.CAxMaPefzc

leave.fkzBfijthr:
	br label %leave.LyxgTLENsq

for.declaration.CAxMaPefzc:
	%35 = icmp ne i32 1, 0
	br i1 %35, label %for.block.EmSWHeQSnQ, label %leave.fkzBfijthr

for.condition.LApDcHMklb:
	%36 = icmp ne i32 1, 0
	br i1 %36, label %for.block.EmSWHeQSnQ, label %leave.fkzBfijthr

for.block.EmSWHeQSnQ:
	%37 = icmp ne i32 2, 0
	br i1 %37, label %if.then.vFHpcddUfj, label %if.else.cItIMBEmeP

for.update.skZeoBgaDi:
	br label %for.condition.LApDcHMklb

if.then.vFHpcddUfj:
	%38 = icmp ne i32 2, 0
	br i1 %38, label %if.then.OyqFAgYCAz, label %if.else.AIEpwjmJdT

if.then.OyqFAgYCAz:
	%39 = icmp ne i32 2, 0
	br i1 %39, label %if.then.eYwnlAcxnq, label %if.else.LOCJvfkUJI

if.then.eYwnlAcxnq:
	%40 = icmp ne i32 2, 0
	br i1 %40, label %if.then.sKgJHlEOui, label %if.else.EvutrGXFfw

if.then.sKgJHlEOui:
	br label %leave.fkzBfijthr

if.else.EvutrGXFfw:
	br label %lastLeave.WMqKTkVQrT

lastLeave.WMqKTkVQrT:
	br label %lastLeave.gLNLLrasin

if.else.LOCJvfkUJI:
	br label %lastLeave.gLNLLrasin

lastLeave.gLNLLrasin:
	br label %lastLeave.wYsEoHAsts

if.else.AIEpwjmJdT:
	br label %lastLeave.wYsEoHAsts

lastLeave.wYsEoHAsts:
	br label %lastLeave.yNQDRcSNZZ

if.else.cItIMBEmeP:
	br label %lastLeave.yNQDRcSNZZ

lastLeave.yNQDRcSNZZ:
	%41 = getelementptr [16 x i8], [16 x i8]* @string.literal.KzBHaOSpqk, i32 0, i32 0
	%42 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%43 = call i32 (i8*, ...) @printf(i8* %42, i8* %41)
	br label %for.update.skZeoBgaDi

if.else.eaQqwskQge:
	br label %lastLeave.KljEqvamGg

lastLeave.KljEqvamGg:
	%44 = getelementptr [16 x i8], [16 x i8]* @string.literal.znfXCeSLdV, i32 0, i32 0
	%45 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%46 = call i32 (i8*, ...) @printf(i8* %45, i8* %44)
	br label %for.update.EbUlufmAUr
}
