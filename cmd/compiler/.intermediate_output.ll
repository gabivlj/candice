%Something-lIorVMIlIe = type { i64 }
%Something-njIsArTVJq = type { i32 }
%Thing-nSzlUWcZYg = type { i32 }
%Array-JSWpaomwve = type { i32*, i32, i32 }

@string.literal.dqYySNTmFr = global [13 x i8] c"hello world!\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.hCHHKNGUVq = global [13 x i8] c"hello world!\00"
@"%d " = global [4 x i8] c"%d \00"

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc void @outsideFunction-lIorVMIlIe(%Something-lIorVMIlIe* %s-lIorVMIlIe) {
outsideFunction-lIorVMIlIe:
	%0 = alloca %Something-lIorVMIlIe*
	store %Something-lIorVMIlIe* %s-lIorVMIlIe, %Something-lIorVMIlIe** %0
	%1 = load %Something-lIorVMIlIe*, %Something-lIorVMIlIe** %0
	%2 = getelementptr %Something-lIorVMIlIe, %Something-lIorVMIlIe* %1, i32 0, i32 0
	%3 = load %Something-lIorVMIlIe*, %Something-lIorVMIlIe** %0
	%4 = getelementptr %Something-lIorVMIlIe, %Something-lIorVMIlIe* %3, i32 0, i32 0
	%5 = load i64, i64* %4
	%6 = sext i32 1 to i64
	%7 = add i64 %5, %6
	store i64 %7, i64* %2
	ret void
}

define ccc %Something-lIorVMIlIe @something-lIorVMIlIe() {
something-lIorVMIlIe:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.dqYySNTmFr, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-lIorVMIlIe
	%4 = sext i32 0 to i64
	%5 = getelementptr %Something-lIorVMIlIe, %Something-lIorVMIlIe* %3, i32 0, i32 0
	store i64 %4, i64* %5
	%6 = alloca %Something-lIorVMIlIe
	%7 = load %Something-lIorVMIlIe, %Something-lIorVMIlIe* %3
	store %Something-lIorVMIlIe %7, %Something-lIorVMIlIe* %6
	%8 = alloca %Something-lIorVMIlIe*
	store %Something-lIorVMIlIe* %6, %Something-lIorVMIlIe** %8
	%9 = load %Something-lIorVMIlIe*, %Something-lIorVMIlIe** %8
	call void @outsideFunction-lIorVMIlIe(%Something-lIorVMIlIe* %9)
	%10 = load %Something-lIorVMIlIe, %Something-lIorVMIlIe* %6
	ret %Something-lIorVMIlIe %10
}

define ccc void @outsideFunction-njIsArTVJq(%Something-njIsArTVJq* %s-njIsArTVJq) {
outsideFunction-njIsArTVJq:
	%0 = alloca %Something-njIsArTVJq*
	store %Something-njIsArTVJq* %s-njIsArTVJq, %Something-njIsArTVJq** %0
	%1 = load %Something-njIsArTVJq*, %Something-njIsArTVJq** %0
	%2 = getelementptr %Something-njIsArTVJq, %Something-njIsArTVJq* %1, i32 0, i32 0
	%3 = load %Something-njIsArTVJq*, %Something-njIsArTVJq** %0
	%4 = getelementptr %Something-njIsArTVJq, %Something-njIsArTVJq* %3, i32 0, i32 0
	%5 = load i32, i32* %4
	%6 = add i32 %5, 1
	store i32 %6, i32* %2
	ret void
}

define ccc %Something-njIsArTVJq @something-njIsArTVJq() {
something-njIsArTVJq:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.hCHHKNGUVq, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-njIsArTVJq
	%4 = getelementptr %Something-njIsArTVJq, %Something-njIsArTVJq* %3, i32 0, i32 0
	store i32 0, i32* %4
	%5 = alloca %Something-njIsArTVJq
	%6 = load %Something-njIsArTVJq, %Something-njIsArTVJq* %3
	store %Something-njIsArTVJq %6, %Something-njIsArTVJq* %5
	%7 = alloca %Something-njIsArTVJq*
	store %Something-njIsArTVJq* %5, %Something-njIsArTVJq** %7
	%8 = load %Something-njIsArTVJq*, %Something-njIsArTVJq** %7
	call void @outsideFunction-njIsArTVJq(%Something-njIsArTVJq* %8)
	%9 = load %Something-njIsArTVJq, %Something-njIsArTVJq* %5
	ret %Something-njIsArTVJq %9
}

define ccc %Thing-nSzlUWcZYg* @create-AdQwByBNQt() {
create-AdQwByBNQt:
	%0 = sext i32 1 to i64
	%1 = mul i64 %0, 4
	%2 = call i8* @malloc(i64 %1)
	%3 = bitcast i8* %2 to %Thing-nSzlUWcZYg*
	%4 = alloca %Thing-nSzlUWcZYg*
	store %Thing-nSzlUWcZYg* %3, %Thing-nSzlUWcZYg** %4
	%5 = load %Thing-nSzlUWcZYg*, %Thing-nSzlUWcZYg** %4
	ret %Thing-nSzlUWcZYg* %5
}

declare ccc void @free(i8* %0)

define ccc %Array-JSWpaomwve @New-JSWpaomwve() {
New-JSWpaomwve:
	%0 = alloca %Array-JSWpaomwve
	%1 = sext i32 10 to i64
	%2 = mul i64 %1, 4
	%3 = call i8* @malloc(i64 %2)
	%4 = bitcast i8* %3 to i32*
	%5 = alloca i32*
	store i32* %4, i32** %5
	%6 = getelementptr %Array-JSWpaomwve, %Array-JSWpaomwve* %0, i32 0, i32 0
	%7 = load i32*, i32** %5
	store i32* %7, i32** %6
	%8 = getelementptr %Array-JSWpaomwve, %Array-JSWpaomwve* %0, i32 0, i32 1
	store i32 0, i32* %8
	%9 = getelementptr %Array-JSWpaomwve, %Array-JSWpaomwve* %0, i32 0, i32 2
	store i32 10, i32* %9
	%10 = load %Array-JSWpaomwve, %Array-JSWpaomwve* %0
	ret %Array-JSWpaomwve %10
}

define ccc void @Free-JSWpaomwve(%Array-JSWpaomwve* %a-JSWpaomwve) {
Free-JSWpaomwve:
	%0 = alloca %Array-JSWpaomwve*
	store %Array-JSWpaomwve* %a-JSWpaomwve, %Array-JSWpaomwve** %0
	%1 = load %Array-JSWpaomwve*, %Array-JSWpaomwve** %0
	%2 = getelementptr %Array-JSWpaomwve, %Array-JSWpaomwve* %1, i32 0, i32 0
	%3 = load i32*, i32** %2
	%4 = bitcast i32* %3 to i8*
	call void @free(i8* %4)
	ret void
}

define ccc void @Push-JSWpaomwve(%Array-JSWpaomwve* %a-JSWpaomwve, i32 %element-JSWpaomwve) {
Push-JSWpaomwve:
	%0 = alloca %Array-JSWpaomwve*
	store %Array-JSWpaomwve* %a-JSWpaomwve, %Array-JSWpaomwve** %0
	%1 = alloca i32
	store i32 %element-JSWpaomwve, i32* %1
	%2 = load %Array-JSWpaomwve*, %Array-JSWpaomwve** %0
	%3 = getelementptr %Array-JSWpaomwve, %Array-JSWpaomwve* %2, i32 0, i32 1
	%4 = load i32, i32* %3
	%5 = load %Array-JSWpaomwve*, %Array-JSWpaomwve** %0
	%6 = getelementptr %Array-JSWpaomwve, %Array-JSWpaomwve* %5, i32 0, i32 2
	%7 = load i32, i32* %6
	%8 = icmp sge i32 %4, %7
	br i1 %8, label %if.then.nklNMPPZHY, label %if.else.ydfreDMniN

if.then.nklNMPPZHY:
	br label %lastLeave.LsYKhmsluE

if.else.ydfreDMniN:
	br label %lastLeave.LsYKhmsluE

lastLeave.LsYKhmsluE:
	%9 = load %Array-JSWpaomwve*, %Array-JSWpaomwve** %0
	%10 = getelementptr %Array-JSWpaomwve, %Array-JSWpaomwve* %9, i32 0, i32 0
	%11 = load %Array-JSWpaomwve*, %Array-JSWpaomwve** %0
	%12 = getelementptr %Array-JSWpaomwve, %Array-JSWpaomwve* %11, i32 0, i32 1
	%13 = load i32, i32* %12
	%14 = load i32*, i32** %10
	%15 = getelementptr i32, i32* %14, i32 %13
	%16 = load i32, i32* %1
	store i32 %16, i32* %15
	%17 = load %Array-JSWpaomwve*, %Array-JSWpaomwve** %0
	%18 = getelementptr %Array-JSWpaomwve, %Array-JSWpaomwve* %17, i32 0, i32 1
	%19 = load %Array-JSWpaomwve*, %Array-JSWpaomwve** %0
	%20 = getelementptr %Array-JSWpaomwve, %Array-JSWpaomwve* %19, i32 0, i32 1
	%21 = load i32, i32* %20
	%22 = add i32 %21, 1
	store i32 %22, i32* %18
	ret void
}

define ccc %Something-lIorVMIlIe @something-nSzlUWcZYg() {
something-nSzlUWcZYg:
	%0 = alloca %Something-lIorVMIlIe
	%1 = sext i32 3 to i64
	%2 = getelementptr %Something-lIorVMIlIe, %Something-lIorVMIlIe* %0, i32 0, i32 0
	store i64 %1, i64* %2
	%3 = load %Something-lIorVMIlIe, %Something-lIorVMIlIe* %0
	ret %Something-lIorVMIlIe %3
}

define ccc i32 @call-IldxvUtDXz(i32 (i32)* %f-IldxvUtDXz) {
call-IldxvUtDXz:
	%0 = alloca i32 (i32)*
	store i32 (i32)* %f-IldxvUtDXz, i32 (i32)** %0
	%1 = load i32 (i32)*, i32 (i32)** %0
	%2 = call i32 %1(i32 1)
	ret i32 %2
}

define ccc i32 @wow-nSzlUWcZYg(i32 %i-nSzlUWcZYg) {
wow-nSzlUWcZYg:
	%0 = alloca i32
	store i32 %i-nSzlUWcZYg, i32* %0
	%1 = load i32, i32* %0
	ret i32 %1
}

define ccc void @main() {
main:
	%0 = call %Array-JSWpaomwve @New-JSWpaomwve()
	%1 = alloca %Array-JSWpaomwve
	store %Array-JSWpaomwve %0, %Array-JSWpaomwve* %1
	%2 = getelementptr %Array-JSWpaomwve, %Array-JSWpaomwve* %1, i32 0, i32 0
	%3 = load i32*, i32** %2
	%4 = getelementptr i32, i32* %3, i32 0
	%5 = load i32, i32* %4
	%6 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%7 = call i32 (i8*, ...) @printf(i8* %6, i32 %5)
	%8 = alloca %Array-JSWpaomwve*
	store %Array-JSWpaomwve* %1, %Array-JSWpaomwve** %8
	%9 = load %Array-JSWpaomwve*, %Array-JSWpaomwve** %8
	call void @Push-JSWpaomwve(%Array-JSWpaomwve* %9, i32 3)
	%10 = getelementptr %Array-JSWpaomwve, %Array-JSWpaomwve* %1, i32 0, i32 0
	%11 = load i32*, i32** %10
	%12 = getelementptr i32, i32* %11, i32 0
	%13 = load i32, i32* %12
	%14 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%15 = call i32 (i8*, ...) @printf(i8* %14, i32 %13)
	%16 = call i32 @call-IldxvUtDXz(i32 (i32)* @wow-nSzlUWcZYg)
	%17 = alloca i32
	store i32 %16, i32* %17
	%18 = load i32, i32* %17
	%19 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%20 = call i32 (i8*, ...) @printf(i8* %19, i32 %18)
	%21 = call %Thing-nSzlUWcZYg* @create-AdQwByBNQt()
	%22 = alloca %Thing-nSzlUWcZYg*
	store %Thing-nSzlUWcZYg* %21, %Thing-nSzlUWcZYg** %22
	%23 = load %Thing-nSzlUWcZYg*, %Thing-nSzlUWcZYg** %22
	%24 = getelementptr %Thing-nSzlUWcZYg, %Thing-nSzlUWcZYg* %23, i32 0, i32 0
	%25 = load i32, i32* %24
	%26 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%27 = call i32 (i8*, ...) @printf(i8* %26, i32 %25)
	%28 = call %Something-lIorVMIlIe @something-lIorVMIlIe()
	%29 = alloca %Something-lIorVMIlIe
	store %Something-lIorVMIlIe %28, %Something-lIorVMIlIe* %29
	%30 = call %Something-njIsArTVJq @something-njIsArTVJq()
	%31 = alloca %Something-njIsArTVJq
	store %Something-njIsArTVJq %30, %Something-njIsArTVJq* %31
	%32 = call %Something-lIorVMIlIe @something-nSzlUWcZYg()
	%33 = alloca %Something-lIorVMIlIe
	store %Something-lIorVMIlIe %32, %Something-lIorVMIlIe* %33
	%34 = getelementptr %Something-njIsArTVJq, %Something-njIsArTVJq* %31, i32 0, i32 0
	%35 = load i32, i32* %34
	%36 = getelementptr %Something-lIorVMIlIe, %Something-lIorVMIlIe* %29, i32 0, i32 0
	%37 = load i64, i64* %36
	%38 = trunc i64 %37 to i32
	%39 = add i32 %35, %38
	%40 = getelementptr %Something-lIorVMIlIe, %Something-lIorVMIlIe* %33, i32 0, i32 0
	%41 = load i64, i64* %40
	%42 = trunc i64 %41 to i32
	%43 = add i32 %39, %42
	%44 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%45 = call i32 (i8*, ...) @printf(i8* %44, i32 %43)
	ret void
}
