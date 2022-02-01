%Something-dKSMNaoynp = type { i64 }
%Something-aAxoqycuYF = type { i32 }
%String-nNVuGbPJIJ = type { i8*, i32 }
%Thing-YDYIkPdlTs = type { i32 }
%Array-gyhqVuTbky = type { i32*, i32, i32 }

@string.literal.OwYAQjrBNr = global [13 x i8] c"hello world!\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.EhkiILfWGM = global [13 x i8] c"hello world!\00"
@string.literal.MSkoWhcrNQ = global [6 x i8] c"Hello\00"
@string.literal.JerkcKUMwT = global [7 x i8] c"hellow\00"
@"%d " = global [4 x i8] c"%d \00"

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc void @outsideFunction-dKSMNaoynp(%Something-dKSMNaoynp* %s-dKSMNaoynp) {
outsideFunction-dKSMNaoynp:
	%0 = alloca %Something-dKSMNaoynp*
	store %Something-dKSMNaoynp* %s-dKSMNaoynp, %Something-dKSMNaoynp** %0
	%1 = load %Something-dKSMNaoynp*, %Something-dKSMNaoynp** %0
	%2 = getelementptr %Something-dKSMNaoynp, %Something-dKSMNaoynp* %1, i32 0, i32 0
	%3 = load %Something-dKSMNaoynp*, %Something-dKSMNaoynp** %0
	%4 = getelementptr %Something-dKSMNaoynp, %Something-dKSMNaoynp* %3, i32 0, i32 0
	%5 = load i64, i64* %4
	%6 = sext i32 1 to i64
	%7 = add i64 %5, %6
	store i64 %7, i64* %2
	ret void
}

define ccc %Something-dKSMNaoynp @something-dKSMNaoynp() {
something-dKSMNaoynp:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.OwYAQjrBNr, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-dKSMNaoynp
	%4 = sext i32 0 to i64
	%5 = getelementptr %Something-dKSMNaoynp, %Something-dKSMNaoynp* %3, i32 0, i32 0
	store i64 %4, i64* %5
	%6 = alloca %Something-dKSMNaoynp
	%7 = load %Something-dKSMNaoynp, %Something-dKSMNaoynp* %3
	store %Something-dKSMNaoynp %7, %Something-dKSMNaoynp* %6
	%8 = alloca %Something-dKSMNaoynp*
	store %Something-dKSMNaoynp* %6, %Something-dKSMNaoynp** %8
	%9 = load %Something-dKSMNaoynp*, %Something-dKSMNaoynp** %8
	call void @outsideFunction-dKSMNaoynp(%Something-dKSMNaoynp* %9)
	%10 = load %Something-dKSMNaoynp, %Something-dKSMNaoynp* %6
	ret %Something-dKSMNaoynp %10
}

define ccc void @outsideFunction-aAxoqycuYF(%Something-aAxoqycuYF* %s-aAxoqycuYF) {
outsideFunction-aAxoqycuYF:
	%0 = alloca %Something-aAxoqycuYF*
	store %Something-aAxoqycuYF* %s-aAxoqycuYF, %Something-aAxoqycuYF** %0
	%1 = load %Something-aAxoqycuYF*, %Something-aAxoqycuYF** %0
	%2 = getelementptr %Something-aAxoqycuYF, %Something-aAxoqycuYF* %1, i32 0, i32 0
	%3 = load %Something-aAxoqycuYF*, %Something-aAxoqycuYF** %0
	%4 = getelementptr %Something-aAxoqycuYF, %Something-aAxoqycuYF* %3, i32 0, i32 0
	%5 = load i32, i32* %4
	%6 = add i32 %5, 1
	store i32 %6, i32* %2
	ret void
}

define ccc %Something-aAxoqycuYF @something-aAxoqycuYF() {
something-aAxoqycuYF:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.EhkiILfWGM, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-aAxoqycuYF
	%4 = getelementptr %Something-aAxoqycuYF, %Something-aAxoqycuYF* %3, i32 0, i32 0
	store i32 0, i32* %4
	%5 = alloca %Something-aAxoqycuYF
	%6 = load %Something-aAxoqycuYF, %Something-aAxoqycuYF* %3
	store %Something-aAxoqycuYF %6, %Something-aAxoqycuYF* %5
	%7 = alloca %Something-aAxoqycuYF*
	store %Something-aAxoqycuYF* %5, %Something-aAxoqycuYF** %7
	%8 = load %Something-aAxoqycuYF*, %Something-aAxoqycuYF** %7
	call void @outsideFunction-aAxoqycuYF(%Something-aAxoqycuYF* %8)
	%9 = load %Something-aAxoqycuYF, %Something-aAxoqycuYF* %5
	ret %Something-aAxoqycuYF %9
}

define ccc %String-nNVuGbPJIJ @New-nNVuGbPJIJ(i8* %s-nNVuGbPJIJ) {
New-nNVuGbPJIJ:
	%0 = alloca i8*
	store i8* %s-nNVuGbPJIJ, i8** %0
	%1 = alloca %String-nNVuGbPJIJ
	%2 = getelementptr %String-nNVuGbPJIJ, %String-nNVuGbPJIJ* %1, i32 0, i32 0
	%3 = load i8*, i8** %0
	store i8* %3, i8** %2
	%4 = getelementptr %String-nNVuGbPJIJ, %String-nNVuGbPJIJ* %1, i32 0, i32 1
	store i32 0, i32* %4
	%5 = load %String-nNVuGbPJIJ, %String-nNVuGbPJIJ* %1
	ret %String-nNVuGbPJIJ %5
}

define ccc %String-nNVuGbPJIJ @CreateString-yXirOcarPd() {
CreateString-yXirOcarPd:
	%0 = getelementptr [6 x i8], [6 x i8]* @string.literal.MSkoWhcrNQ, i32 0, i32 0
	%1 = call %String-nNVuGbPJIJ @New-nNVuGbPJIJ(i8* %0)
	ret %String-nNVuGbPJIJ %1
}

define ccc %String-nNVuGbPJIJ @UseString-yXirOcarPd(%String-nNVuGbPJIJ* %z-yXirOcarPd) {
UseString-yXirOcarPd:
	%0 = alloca %String-nNVuGbPJIJ*
	store %String-nNVuGbPJIJ* %z-yXirOcarPd, %String-nNVuGbPJIJ** %0
	%1 = load %String-nNVuGbPJIJ*, %String-nNVuGbPJIJ** %0
	%2 = load %String-nNVuGbPJIJ, %String-nNVuGbPJIJ* %1
	ret %String-nNVuGbPJIJ %2
}

define ccc %Thing-YDYIkPdlTs* @create-DEDyATuyxQ() {
create-DEDyATuyxQ:
	%0 = sext i32 1 to i64
	%1 = mul i64 %0, 4
	%2 = call i8* @malloc(i64 %1)
	%3 = bitcast i8* %2 to %Thing-YDYIkPdlTs*
	%4 = alloca %Thing-YDYIkPdlTs*
	store %Thing-YDYIkPdlTs* %3, %Thing-YDYIkPdlTs** %4
	%5 = load %Thing-YDYIkPdlTs*, %Thing-YDYIkPdlTs** %4
	ret %Thing-YDYIkPdlTs* %5
}

declare ccc void @free(i8* %0)

define ccc %Array-gyhqVuTbky @New-gyhqVuTbky() {
New-gyhqVuTbky:
	%0 = alloca %Array-gyhqVuTbky
	%1 = sext i32 10 to i64
	%2 = mul i64 %1, 4
	%3 = call i8* @malloc(i64 %2)
	%4 = bitcast i8* %3 to i32*
	%5 = alloca i32*
	store i32* %4, i32** %5
	%6 = getelementptr %Array-gyhqVuTbky, %Array-gyhqVuTbky* %0, i32 0, i32 0
	%7 = load i32*, i32** %5
	store i32* %7, i32** %6
	%8 = getelementptr %Array-gyhqVuTbky, %Array-gyhqVuTbky* %0, i32 0, i32 1
	store i32 0, i32* %8
	%9 = getelementptr %Array-gyhqVuTbky, %Array-gyhqVuTbky* %0, i32 0, i32 2
	store i32 10, i32* %9
	%10 = load %Array-gyhqVuTbky, %Array-gyhqVuTbky* %0
	ret %Array-gyhqVuTbky %10
}

define ccc void @Free-gyhqVuTbky(%Array-gyhqVuTbky* %a-gyhqVuTbky) {
Free-gyhqVuTbky:
	%0 = alloca %Array-gyhqVuTbky*
	store %Array-gyhqVuTbky* %a-gyhqVuTbky, %Array-gyhqVuTbky** %0
	%1 = load %Array-gyhqVuTbky*, %Array-gyhqVuTbky** %0
	%2 = getelementptr %Array-gyhqVuTbky, %Array-gyhqVuTbky* %1, i32 0, i32 0
	%3 = load i32*, i32** %2
	%4 = bitcast i32* %3 to i8*
	call void @free(i8* %4)
	ret void
}

define ccc void @Push-gyhqVuTbky(%Array-gyhqVuTbky* %a-gyhqVuTbky, i32 %element-gyhqVuTbky) {
Push-gyhqVuTbky:
	%0 = alloca %Array-gyhqVuTbky*
	store %Array-gyhqVuTbky* %a-gyhqVuTbky, %Array-gyhqVuTbky** %0
	%1 = alloca i32
	store i32 %element-gyhqVuTbky, i32* %1
	%2 = load %Array-gyhqVuTbky*, %Array-gyhqVuTbky** %0
	%3 = getelementptr %Array-gyhqVuTbky, %Array-gyhqVuTbky* %2, i32 0, i32 1
	%4 = load i32, i32* %3
	%5 = load %Array-gyhqVuTbky*, %Array-gyhqVuTbky** %0
	%6 = getelementptr %Array-gyhqVuTbky, %Array-gyhqVuTbky* %5, i32 0, i32 2
	%7 = load i32, i32* %6
	%8 = icmp sge i32 %4, %7
	br i1 %8, label %if.then.MruYDtAnKg, label %if.else.ycqDvrVFJq

if.then.MruYDtAnKg:
	br label %lastLeave.xhVxxqhfQW

if.else.ycqDvrVFJq:
	br label %lastLeave.xhVxxqhfQW

lastLeave.xhVxxqhfQW:
	%9 = load %Array-gyhqVuTbky*, %Array-gyhqVuTbky** %0
	%10 = getelementptr %Array-gyhqVuTbky, %Array-gyhqVuTbky* %9, i32 0, i32 0
	%11 = load %Array-gyhqVuTbky*, %Array-gyhqVuTbky** %0
	%12 = getelementptr %Array-gyhqVuTbky, %Array-gyhqVuTbky* %11, i32 0, i32 1
	%13 = load i32, i32* %12
	%14 = load i32*, i32** %10
	%15 = getelementptr i32, i32* %14, i32 %13
	%16 = load i32, i32* %1
	store i32 %16, i32* %15
	%17 = load %Array-gyhqVuTbky*, %Array-gyhqVuTbky** %0
	%18 = getelementptr %Array-gyhqVuTbky, %Array-gyhqVuTbky* %17, i32 0, i32 1
	%19 = load %Array-gyhqVuTbky*, %Array-gyhqVuTbky** %0
	%20 = getelementptr %Array-gyhqVuTbky, %Array-gyhqVuTbky* %19, i32 0, i32 1
	%21 = load i32, i32* %20
	%22 = add i32 %21, 1
	store i32 %22, i32* %18
	ret void
}

define ccc %Something-dKSMNaoynp @something-YDYIkPdlTs() {
something-YDYIkPdlTs:
	%0 = alloca %Something-dKSMNaoynp
	%1 = sext i32 3 to i64
	%2 = getelementptr %Something-dKSMNaoynp, %Something-dKSMNaoynp* %0, i32 0, i32 0
	store i64 %1, i64* %2
	%3 = load %Something-dKSMNaoynp, %Something-dKSMNaoynp* %0
	ret %Something-dKSMNaoynp %3
}

define ccc i32 @call-JedcQKKIra(i32 (i32)* %f-JedcQKKIra) {
call-JedcQKKIra:
	%0 = alloca i32 (i32)*
	store i32 (i32)* %f-JedcQKKIra, i32 (i32)** %0
	%1 = load i32 (i32)*, i32 (i32)** %0
	%2 = call i32 %1(i32 1)
	ret i32 %2
}

define ccc i32 @wow-YDYIkPdlTs(i32 %i-YDYIkPdlTs) {
wow-YDYIkPdlTs:
	%0 = alloca i32
	store i32 %i-YDYIkPdlTs, i32* %0
	%1 = load i32, i32* %0
	ret i32 %1
}

define ccc void @main() {
main:
	%0 = getelementptr [7 x i8], [7 x i8]* @string.literal.JerkcKUMwT, i32 0, i32 0
	%1 = call %String-nNVuGbPJIJ @New-nNVuGbPJIJ(i8* %0)
	%2 = alloca %String-nNVuGbPJIJ
	store %String-nNVuGbPJIJ %1, %String-nNVuGbPJIJ* %2
	%3 = alloca %String-nNVuGbPJIJ*
	store %String-nNVuGbPJIJ* %2, %String-nNVuGbPJIJ** %3
	%4 = load %String-nNVuGbPJIJ*, %String-nNVuGbPJIJ** %3
	%5 = call %String-nNVuGbPJIJ @UseString-yXirOcarPd(%String-nNVuGbPJIJ* %4)
	%6 = alloca %String-nNVuGbPJIJ
	store %String-nNVuGbPJIJ %5, %String-nNVuGbPJIJ* %6
	%7 = call %Array-gyhqVuTbky @New-gyhqVuTbky()
	%8 = alloca %Array-gyhqVuTbky
	store %Array-gyhqVuTbky %7, %Array-gyhqVuTbky* %8
	%9 = getelementptr %Array-gyhqVuTbky, %Array-gyhqVuTbky* %8, i32 0, i32 0
	%10 = load i32*, i32** %9
	%11 = getelementptr i32, i32* %10, i32 0
	%12 = load i32, i32* %11
	%13 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%14 = call i32 (i8*, ...) @printf(i8* %13, i32 %12)
	%15 = alloca %Array-gyhqVuTbky*
	store %Array-gyhqVuTbky* %8, %Array-gyhqVuTbky** %15
	%16 = load %Array-gyhqVuTbky*, %Array-gyhqVuTbky** %15
	call void @Push-gyhqVuTbky(%Array-gyhqVuTbky* %16, i32 3)
	%17 = getelementptr %Array-gyhqVuTbky, %Array-gyhqVuTbky* %8, i32 0, i32 0
	%18 = load i32*, i32** %17
	%19 = getelementptr i32, i32* %18, i32 0
	%20 = load i32, i32* %19
	%21 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%22 = call i32 (i8*, ...) @printf(i8* %21, i32 %20)
	%23 = call i32 @call-JedcQKKIra(i32 (i32)* @wow-YDYIkPdlTs)
	%24 = alloca i32
	store i32 %23, i32* %24
	%25 = load i32, i32* %24
	%26 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%27 = call i32 (i8*, ...) @printf(i8* %26, i32 %25)
	%28 = call %Thing-YDYIkPdlTs* @create-DEDyATuyxQ()
	%29 = alloca %Thing-YDYIkPdlTs*
	store %Thing-YDYIkPdlTs* %28, %Thing-YDYIkPdlTs** %29
	%30 = load %Thing-YDYIkPdlTs*, %Thing-YDYIkPdlTs** %29
	%31 = getelementptr %Thing-YDYIkPdlTs, %Thing-YDYIkPdlTs* %30, i32 0, i32 0
	%32 = load i32, i32* %31
	%33 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%34 = call i32 (i8*, ...) @printf(i8* %33, i32 %32)
	%35 = call %Something-dKSMNaoynp @something-dKSMNaoynp()
	%36 = alloca %Something-dKSMNaoynp
	store %Something-dKSMNaoynp %35, %Something-dKSMNaoynp* %36
	%37 = call %Something-aAxoqycuYF @something-aAxoqycuYF()
	%38 = alloca %Something-aAxoqycuYF
	store %Something-aAxoqycuYF %37, %Something-aAxoqycuYF* %38
	%39 = call %Something-dKSMNaoynp @something-YDYIkPdlTs()
	%40 = alloca %Something-dKSMNaoynp
	store %Something-dKSMNaoynp %39, %Something-dKSMNaoynp* %40
	%41 = getelementptr %Something-aAxoqycuYF, %Something-aAxoqycuYF* %38, i32 0, i32 0
	%42 = load i32, i32* %41
	%43 = getelementptr %Something-dKSMNaoynp, %Something-dKSMNaoynp* %36, i32 0, i32 0
	%44 = load i64, i64* %43
	%45 = trunc i64 %44 to i32
	%46 = add i32 %42, %45
	%47 = getelementptr %Something-dKSMNaoynp, %Something-dKSMNaoynp* %40, i32 0, i32 0
	%48 = load i64, i64* %47
	%49 = trunc i64 %48 to i32
	%50 = add i32 %46, %49
	%51 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%52 = call i32 (i8*, ...) @printf(i8* %51, i32 %50)
	ret void
}
