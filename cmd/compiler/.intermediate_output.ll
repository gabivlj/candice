%Something-rdSwDByzlU = type { i64 }
%Something-BZzsvNVfUV = type { i32 }
%String-WTDITkWdPQ = type { i8*, i32 }
%Thing-eBEMnIDULZ = type { i32 }
%Array-lzOZQZvefr = type { i32*, i32, i32 }

@string.literal.avzBkIjQxZ = global [13 x i8] c"hello world!\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.tZYrFjDehO = global [13 x i8] c"hello world!\00"
@string.literal.pCgxfvhJVO = global [6 x i8] c"Hello\00"
@string.literal.HlDNpeMYvt = global [7 x i8] c"hellow\00"
@"%d " = global [4 x i8] c"%d \00"

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc void @outsideFunction-rdSwDByzlU(%Something-rdSwDByzlU* %s-rdSwDByzlU) {
outsideFunction-rdSwDByzlU:
	%0 = alloca %Something-rdSwDByzlU*
	store %Something-rdSwDByzlU* %s-rdSwDByzlU, %Something-rdSwDByzlU** %0
	%1 = load %Something-rdSwDByzlU*, %Something-rdSwDByzlU** %0
	%2 = getelementptr %Something-rdSwDByzlU, %Something-rdSwDByzlU* %1, i32 0, i32 0
	%3 = load %Something-rdSwDByzlU*, %Something-rdSwDByzlU** %0
	%4 = getelementptr %Something-rdSwDByzlU, %Something-rdSwDByzlU* %3, i32 0, i32 0
	%5 = load i64, i64* %4
	%6 = sext i32 1 to i64
	%7 = add i64 %5, %6
	store i64 %7, i64* %2
	ret void
}

define ccc %Something-rdSwDByzlU @something-rdSwDByzlU() {
something-rdSwDByzlU:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.avzBkIjQxZ, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-rdSwDByzlU
	%4 = sext i32 0 to i64
	%5 = getelementptr %Something-rdSwDByzlU, %Something-rdSwDByzlU* %3, i32 0, i32 0
	store i64 %4, i64* %5
	%6 = alloca %Something-rdSwDByzlU
	%7 = load %Something-rdSwDByzlU, %Something-rdSwDByzlU* %3
	store %Something-rdSwDByzlU %7, %Something-rdSwDByzlU* %6
	%8 = alloca %Something-rdSwDByzlU*
	store %Something-rdSwDByzlU* %6, %Something-rdSwDByzlU** %8
	%9 = load %Something-rdSwDByzlU*, %Something-rdSwDByzlU** %8
	call void @outsideFunction-rdSwDByzlU(%Something-rdSwDByzlU* %9)
	%10 = load %Something-rdSwDByzlU, %Something-rdSwDByzlU* %6
	ret %Something-rdSwDByzlU %10
}

define ccc void @outsideFunction-BZzsvNVfUV(%Something-BZzsvNVfUV* %s-BZzsvNVfUV) {
outsideFunction-BZzsvNVfUV:
	%0 = alloca %Something-BZzsvNVfUV*
	store %Something-BZzsvNVfUV* %s-BZzsvNVfUV, %Something-BZzsvNVfUV** %0
	%1 = load %Something-BZzsvNVfUV*, %Something-BZzsvNVfUV** %0
	%2 = getelementptr %Something-BZzsvNVfUV, %Something-BZzsvNVfUV* %1, i32 0, i32 0
	%3 = load %Something-BZzsvNVfUV*, %Something-BZzsvNVfUV** %0
	%4 = getelementptr %Something-BZzsvNVfUV, %Something-BZzsvNVfUV* %3, i32 0, i32 0
	%5 = load i32, i32* %4
	%6 = add i32 %5, 1
	store i32 %6, i32* %2
	ret void
}

define ccc %Something-BZzsvNVfUV @something-BZzsvNVfUV() {
something-BZzsvNVfUV:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.tZYrFjDehO, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-BZzsvNVfUV
	%4 = getelementptr %Something-BZzsvNVfUV, %Something-BZzsvNVfUV* %3, i32 0, i32 0
	store i32 0, i32* %4
	%5 = alloca %Something-BZzsvNVfUV
	%6 = load %Something-BZzsvNVfUV, %Something-BZzsvNVfUV* %3
	store %Something-BZzsvNVfUV %6, %Something-BZzsvNVfUV* %5
	%7 = alloca %Something-BZzsvNVfUV*
	store %Something-BZzsvNVfUV* %5, %Something-BZzsvNVfUV** %7
	%8 = load %Something-BZzsvNVfUV*, %Something-BZzsvNVfUV** %7
	call void @outsideFunction-BZzsvNVfUV(%Something-BZzsvNVfUV* %8)
	%9 = load %Something-BZzsvNVfUV, %Something-BZzsvNVfUV* %5
	ret %Something-BZzsvNVfUV %9
}

define ccc %String-WTDITkWdPQ @New-WTDITkWdPQ(i8* %s-WTDITkWdPQ) {
New-WTDITkWdPQ:
	%0 = alloca i8*
	store i8* %s-WTDITkWdPQ, i8** %0
	%1 = alloca i32
	store i32 0, i32* %1
	%2 = trunc i32 0 to i8
	%3 = alloca i8
	store i8 %2, i8* %3
	br label %for.declaration.oHfLRRMwDu

leave.PwXSmMoMNS:
	%4 = alloca %String-WTDITkWdPQ
	%5 = getelementptr %String-WTDITkWdPQ, %String-WTDITkWdPQ* %4, i32 0, i32 0
	%6 = load i8*, i8** %0
	store i8* %6, i8** %5
	%7 = getelementptr %String-WTDITkWdPQ, %String-WTDITkWdPQ* %4, i32 0, i32 1
	%8 = load i32, i32* %1
	store i32 %8, i32* %7
	%9 = load %String-WTDITkWdPQ, %String-WTDITkWdPQ* %4
	ret %String-WTDITkWdPQ %9

for.declaration.oHfLRRMwDu:
	%10 = load i32, i32* %1
	%11 = load i8*, i8** %0
	%12 = getelementptr i8, i8* %11, i32 %10
	%13 = load i8, i8* %12
	%14 = load i8, i8* %3
	%15 = icmp ne i8 %13, %14
	br i1 %15, label %for.block.bmRJyNPQaj, label %leave.PwXSmMoMNS

for.condition.JOtYrMbyUh:
	%16 = load i32, i32* %1
	%17 = load i8*, i8** %0
	%18 = getelementptr i8, i8* %17, i32 %16
	%19 = load i8, i8* %18
	%20 = load i8, i8* %3
	%21 = icmp ne i8 %19, %20
	br i1 %21, label %for.block.bmRJyNPQaj, label %leave.PwXSmMoMNS

for.block.bmRJyNPQaj:
	%22 = load i32, i32* %1
	%23 = add i32 %22, 1
	store i32 %23, i32* %1
	br label %for.update.kgCMNzzBvR

for.update.kgCMNzzBvR:
	br label %for.condition.JOtYrMbyUh
}

define ccc %String-WTDITkWdPQ @CreateString-LTrxvBadFL() {
CreateString-LTrxvBadFL:
	%0 = getelementptr [6 x i8], [6 x i8]* @string.literal.pCgxfvhJVO, i32 0, i32 0
	%1 = call %String-WTDITkWdPQ @New-WTDITkWdPQ(i8* %0)
	ret %String-WTDITkWdPQ %1
}

define ccc %String-WTDITkWdPQ @UseString-LTrxvBadFL(%String-WTDITkWdPQ* %z-LTrxvBadFL) {
UseString-LTrxvBadFL:
	%0 = alloca %String-WTDITkWdPQ*
	store %String-WTDITkWdPQ* %z-LTrxvBadFL, %String-WTDITkWdPQ** %0
	%1 = load %String-WTDITkWdPQ*, %String-WTDITkWdPQ** %0
	%2 = load %String-WTDITkWdPQ, %String-WTDITkWdPQ* %1
	ret %String-WTDITkWdPQ %2
}

define ccc %Thing-eBEMnIDULZ* @create-hZUVCrYJYG() {
create-hZUVCrYJYG:
	%0 = sext i32 1 to i64
	%1 = mul i64 %0, 4
	%2 = call i8* @malloc(i64 %1)
	%3 = bitcast i8* %2 to %Thing-eBEMnIDULZ*
	%4 = alloca %Thing-eBEMnIDULZ*
	store %Thing-eBEMnIDULZ* %3, %Thing-eBEMnIDULZ** %4
	%5 = load %Thing-eBEMnIDULZ*, %Thing-eBEMnIDULZ** %4
	ret %Thing-eBEMnIDULZ* %5
}

declare ccc void @free(i8* %0)

define ccc %Array-lzOZQZvefr @New-lzOZQZvefr() {
New-lzOZQZvefr:
	%0 = alloca %Array-lzOZQZvefr
	%1 = sext i32 10 to i64
	%2 = mul i64 %1, 4
	%3 = call i8* @malloc(i64 %2)
	%4 = bitcast i8* %3 to i32*
	%5 = alloca i32*
	store i32* %4, i32** %5
	%6 = getelementptr %Array-lzOZQZvefr, %Array-lzOZQZvefr* %0, i32 0, i32 0
	%7 = load i32*, i32** %5
	store i32* %7, i32** %6
	%8 = getelementptr %Array-lzOZQZvefr, %Array-lzOZQZvefr* %0, i32 0, i32 1
	store i32 0, i32* %8
	%9 = getelementptr %Array-lzOZQZvefr, %Array-lzOZQZvefr* %0, i32 0, i32 2
	store i32 10, i32* %9
	%10 = load %Array-lzOZQZvefr, %Array-lzOZQZvefr* %0
	ret %Array-lzOZQZvefr %10
}

define ccc void @Free-lzOZQZvefr(%Array-lzOZQZvefr* %a-lzOZQZvefr) {
Free-lzOZQZvefr:
	%0 = alloca %Array-lzOZQZvefr*
	store %Array-lzOZQZvefr* %a-lzOZQZvefr, %Array-lzOZQZvefr** %0
	%1 = load %Array-lzOZQZvefr*, %Array-lzOZQZvefr** %0
	%2 = getelementptr %Array-lzOZQZvefr, %Array-lzOZQZvefr* %1, i32 0, i32 0
	%3 = load i32*, i32** %2
	%4 = bitcast i32* %3 to i8*
	call void @free(i8* %4)
	ret void
}

define ccc void @Push-lzOZQZvefr(%Array-lzOZQZvefr* %a-lzOZQZvefr, i32 %element-lzOZQZvefr) {
Push-lzOZQZvefr:
	%0 = alloca %Array-lzOZQZvefr*
	store %Array-lzOZQZvefr* %a-lzOZQZvefr, %Array-lzOZQZvefr** %0
	%1 = alloca i32
	store i32 %element-lzOZQZvefr, i32* %1
	%2 = load %Array-lzOZQZvefr*, %Array-lzOZQZvefr** %0
	%3 = getelementptr %Array-lzOZQZvefr, %Array-lzOZQZvefr* %2, i32 0, i32 1
	%4 = load i32, i32* %3
	%5 = load %Array-lzOZQZvefr*, %Array-lzOZQZvefr** %0
	%6 = getelementptr %Array-lzOZQZvefr, %Array-lzOZQZvefr* %5, i32 0, i32 2
	%7 = load i32, i32* %6
	%8 = icmp sge i32 %4, %7
	br i1 %8, label %if.then.rxjByVLrtw, label %if.else.wTNrazGMvF

if.then.rxjByVLrtw:
	br label %lastLeave.DQfNnHZyKR

if.else.wTNrazGMvF:
	br label %lastLeave.DQfNnHZyKR

lastLeave.DQfNnHZyKR:
	%9 = load %Array-lzOZQZvefr*, %Array-lzOZQZvefr** %0
	%10 = getelementptr %Array-lzOZQZvefr, %Array-lzOZQZvefr* %9, i32 0, i32 0
	%11 = load %Array-lzOZQZvefr*, %Array-lzOZQZvefr** %0
	%12 = getelementptr %Array-lzOZQZvefr, %Array-lzOZQZvefr* %11, i32 0, i32 1
	%13 = load i32, i32* %12
	%14 = load i32*, i32** %10
	%15 = getelementptr i32, i32* %14, i32 %13
	%16 = load i32, i32* %1
	store i32 %16, i32* %15
	%17 = load %Array-lzOZQZvefr*, %Array-lzOZQZvefr** %0
	%18 = getelementptr %Array-lzOZQZvefr, %Array-lzOZQZvefr* %17, i32 0, i32 1
	%19 = load %Array-lzOZQZvefr*, %Array-lzOZQZvefr** %0
	%20 = getelementptr %Array-lzOZQZvefr, %Array-lzOZQZvefr* %19, i32 0, i32 1
	%21 = load i32, i32* %20
	%22 = add i32 %21, 1
	store i32 %22, i32* %18
	ret void
}

define ccc %Something-rdSwDByzlU @something-eBEMnIDULZ() {
something-eBEMnIDULZ:
	%0 = alloca %Something-rdSwDByzlU
	%1 = sext i32 3 to i64
	%2 = getelementptr %Something-rdSwDByzlU, %Something-rdSwDByzlU* %0, i32 0, i32 0
	store i64 %1, i64* %2
	%3 = load %Something-rdSwDByzlU, %Something-rdSwDByzlU* %0
	ret %Something-rdSwDByzlU %3
}

define ccc i32 @call-nuhiSiYgRJ(i32 (i32)* %f-nuhiSiYgRJ) {
call-nuhiSiYgRJ:
	%0 = alloca i32 (i32)*
	store i32 (i32)* %f-nuhiSiYgRJ, i32 (i32)** %0
	%1 = load i32 (i32)*, i32 (i32)** %0
	%2 = call i32 %1(i32 1)
	ret i32 %2
}

define ccc i32 @wow-eBEMnIDULZ(i32 %i-eBEMnIDULZ) {
wow-eBEMnIDULZ:
	%0 = alloca i32
	store i32 %i-eBEMnIDULZ, i32* %0
	%1 = load i32, i32* %0
	ret i32 %1
}

define ccc void @main() {
main:
	%0 = getelementptr [7 x i8], [7 x i8]* @string.literal.HlDNpeMYvt, i32 0, i32 0
	%1 = call %String-WTDITkWdPQ @New-WTDITkWdPQ(i8* %0)
	%2 = alloca %String-WTDITkWdPQ
	store %String-WTDITkWdPQ %1, %String-WTDITkWdPQ* %2
	%3 = alloca %String-WTDITkWdPQ*
	store %String-WTDITkWdPQ* %2, %String-WTDITkWdPQ** %3
	%4 = load %String-WTDITkWdPQ*, %String-WTDITkWdPQ** %3
	%5 = call %String-WTDITkWdPQ @UseString-LTrxvBadFL(%String-WTDITkWdPQ* %4)
	%6 = alloca %String-WTDITkWdPQ
	store %String-WTDITkWdPQ %5, %String-WTDITkWdPQ* %6
	%7 = call %Array-lzOZQZvefr @New-lzOZQZvefr()
	%8 = alloca %Array-lzOZQZvefr
	store %Array-lzOZQZvefr %7, %Array-lzOZQZvefr* %8
	%9 = getelementptr %Array-lzOZQZvefr, %Array-lzOZQZvefr* %8, i32 0, i32 0
	%10 = load i32*, i32** %9
	%11 = getelementptr i32, i32* %10, i32 0
	%12 = load i32, i32* %11
	%13 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%14 = call i32 (i8*, ...) @printf(i8* %13, i32 %12)
	%15 = alloca %Array-lzOZQZvefr*
	store %Array-lzOZQZvefr* %8, %Array-lzOZQZvefr** %15
	%16 = load %Array-lzOZQZvefr*, %Array-lzOZQZvefr** %15
	call void @Push-lzOZQZvefr(%Array-lzOZQZvefr* %16, i32 3)
	%17 = getelementptr %Array-lzOZQZvefr, %Array-lzOZQZvefr* %8, i32 0, i32 0
	%18 = load i32*, i32** %17
	%19 = getelementptr i32, i32* %18, i32 0
	%20 = load i32, i32* %19
	%21 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%22 = call i32 (i8*, ...) @printf(i8* %21, i32 %20)
	%23 = call i32 @call-nuhiSiYgRJ(i32 (i32)* @wow-eBEMnIDULZ)
	%24 = alloca i32
	store i32 %23, i32* %24
	%25 = load i32, i32* %24
	%26 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%27 = call i32 (i8*, ...) @printf(i8* %26, i32 %25)
	%28 = call %Thing-eBEMnIDULZ* @create-hZUVCrYJYG()
	%29 = alloca %Thing-eBEMnIDULZ*
	store %Thing-eBEMnIDULZ* %28, %Thing-eBEMnIDULZ** %29
	%30 = load %Thing-eBEMnIDULZ*, %Thing-eBEMnIDULZ** %29
	%31 = getelementptr %Thing-eBEMnIDULZ, %Thing-eBEMnIDULZ* %30, i32 0, i32 0
	%32 = load i32, i32* %31
	%33 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%34 = call i32 (i8*, ...) @printf(i8* %33, i32 %32)
	%35 = call %Something-rdSwDByzlU @something-rdSwDByzlU()
	%36 = alloca %Something-rdSwDByzlU
	store %Something-rdSwDByzlU %35, %Something-rdSwDByzlU* %36
	%37 = call %Something-BZzsvNVfUV @something-BZzsvNVfUV()
	%38 = alloca %Something-BZzsvNVfUV
	store %Something-BZzsvNVfUV %37, %Something-BZzsvNVfUV* %38
	%39 = call %Something-rdSwDByzlU @something-eBEMnIDULZ()
	%40 = alloca %Something-rdSwDByzlU
	store %Something-rdSwDByzlU %39, %Something-rdSwDByzlU* %40
	%41 = getelementptr %Something-BZzsvNVfUV, %Something-BZzsvNVfUV* %38, i32 0, i32 0
	%42 = load i32, i32* %41
	%43 = getelementptr %Something-rdSwDByzlU, %Something-rdSwDByzlU* %36, i32 0, i32 0
	%44 = load i64, i64* %43
	%45 = trunc i64 %44 to i32
	%46 = add i32 %42, %45
	%47 = getelementptr %Something-rdSwDByzlU, %Something-rdSwDByzlU* %40, i32 0, i32 0
	%48 = load i64, i64* %47
	%49 = trunc i64 %48 to i32
	%50 = add i32 %46, %49
	%51 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%52 = call i32 (i8*, ...) @printf(i8* %51, i32 %50)
	ret void
}
