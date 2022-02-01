%Something-BGUYhRNvif = type { i64 }
%Something-gqWpLiJylC = type { i32 }
%Thing-yEABDKbPiS = type { i32 }

@string.literal.kSyYBTkLti = global [13 x i8] c"hello world!\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.euVwZElCfL = global [13 x i8] c"hello world!\00"
@"%d " = global [4 x i8] c"%d \00"

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc void @outsideFunction-BGUYhRNvif(%Something-BGUYhRNvif* %s-BGUYhRNvif) {
outsideFunction-BGUYhRNvif:
	%0 = alloca %Something-BGUYhRNvif*
	store %Something-BGUYhRNvif* %s-BGUYhRNvif, %Something-BGUYhRNvif** %0
	%1 = load %Something-BGUYhRNvif*, %Something-BGUYhRNvif** %0
	%2 = getelementptr %Something-BGUYhRNvif, %Something-BGUYhRNvif* %1, i32 0, i32 0
	%3 = load %Something-BGUYhRNvif*, %Something-BGUYhRNvif** %0
	%4 = getelementptr %Something-BGUYhRNvif, %Something-BGUYhRNvif* %3, i32 0, i32 0
	%5 = load i64, i64* %4
	%6 = sext i32 1 to i64
	%7 = add i64 %5, %6
	store i64 %7, i64* %2
	ret void
}

define ccc %Something-BGUYhRNvif @something-BGUYhRNvif() {
something-BGUYhRNvif:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.kSyYBTkLti, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-BGUYhRNvif
	%4 = sext i32 0 to i64
	%5 = getelementptr %Something-BGUYhRNvif, %Something-BGUYhRNvif* %3, i32 0, i32 0
	store i64 %4, i64* %5
	%6 = alloca %Something-BGUYhRNvif
	%7 = load %Something-BGUYhRNvif, %Something-BGUYhRNvif* %3
	store %Something-BGUYhRNvif %7, %Something-BGUYhRNvif* %6
	%8 = alloca %Something-BGUYhRNvif*
	store %Something-BGUYhRNvif* %6, %Something-BGUYhRNvif** %8
	%9 = load %Something-BGUYhRNvif*, %Something-BGUYhRNvif** %8
	call void @outsideFunction-BGUYhRNvif(%Something-BGUYhRNvif* %9)
	%10 = load %Something-BGUYhRNvif, %Something-BGUYhRNvif* %6
	ret %Something-BGUYhRNvif %10
}

define ccc void @outsideFunction-gqWpLiJylC(%Something-gqWpLiJylC* %s-gqWpLiJylC) {
outsideFunction-gqWpLiJylC:
	%0 = alloca %Something-gqWpLiJylC*
	store %Something-gqWpLiJylC* %s-gqWpLiJylC, %Something-gqWpLiJylC** %0
	%1 = load %Something-gqWpLiJylC*, %Something-gqWpLiJylC** %0
	%2 = getelementptr %Something-gqWpLiJylC, %Something-gqWpLiJylC* %1, i32 0, i32 0
	%3 = load %Something-gqWpLiJylC*, %Something-gqWpLiJylC** %0
	%4 = getelementptr %Something-gqWpLiJylC, %Something-gqWpLiJylC* %3, i32 0, i32 0
	%5 = load i32, i32* %4
	%6 = add i32 %5, 1
	store i32 %6, i32* %2
	ret void
}

define ccc %Something-gqWpLiJylC @something-gqWpLiJylC() {
something-gqWpLiJylC:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.euVwZElCfL, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-gqWpLiJylC
	%4 = getelementptr %Something-gqWpLiJylC, %Something-gqWpLiJylC* %3, i32 0, i32 0
	store i32 0, i32* %4
	%5 = alloca %Something-gqWpLiJylC
	%6 = load %Something-gqWpLiJylC, %Something-gqWpLiJylC* %3
	store %Something-gqWpLiJylC %6, %Something-gqWpLiJylC* %5
	%7 = alloca %Something-gqWpLiJylC*
	store %Something-gqWpLiJylC* %5, %Something-gqWpLiJylC** %7
	%8 = load %Something-gqWpLiJylC*, %Something-gqWpLiJylC** %7
	call void @outsideFunction-gqWpLiJylC(%Something-gqWpLiJylC* %8)
	%9 = load %Something-gqWpLiJylC, %Something-gqWpLiJylC* %5
	ret %Something-gqWpLiJylC %9
}

define ccc %Thing-yEABDKbPiS* @create-vfVpWSWhuY() {
create-vfVpWSWhuY:
	%0 = sext i32 1 to i64
	%1 = mul i64 %0, 4
	%2 = call i8* @malloc(i64 %1)
	%3 = bitcast i8* %2 to %Thing-yEABDKbPiS*
	%4 = alloca %Thing-yEABDKbPiS*
	store %Thing-yEABDKbPiS* %3, %Thing-yEABDKbPiS** %4
	%5 = load %Thing-yEABDKbPiS*, %Thing-yEABDKbPiS** %4
	ret %Thing-yEABDKbPiS* %5
}

define ccc %Something-BGUYhRNvif @something-yEABDKbPiS() {
something-yEABDKbPiS:
	%0 = alloca %Something-BGUYhRNvif
	%1 = sext i32 3 to i64
	%2 = getelementptr %Something-BGUYhRNvif, %Something-BGUYhRNvif* %0, i32 0, i32 0
	store i64 %1, i64* %2
	%3 = load %Something-BGUYhRNvif, %Something-BGUYhRNvif* %0
	ret %Something-BGUYhRNvif %3
}

define ccc void @main() {
main:
	%0 = call %Thing-yEABDKbPiS* @create-vfVpWSWhuY()
	%1 = alloca %Thing-yEABDKbPiS*
	store %Thing-yEABDKbPiS* %0, %Thing-yEABDKbPiS** %1
	%2 = load %Thing-yEABDKbPiS*, %Thing-yEABDKbPiS** %1
	%3 = getelementptr %Thing-yEABDKbPiS, %Thing-yEABDKbPiS* %2, i32 0, i32 0
	%4 = load i32, i32* %3
	%5 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%6 = call i32 (i8*, ...) @printf(i8* %5, i32 %4)
	%7 = call %Something-BGUYhRNvif @something-BGUYhRNvif()
	%8 = alloca %Something-BGUYhRNvif
	store %Something-BGUYhRNvif %7, %Something-BGUYhRNvif* %8
	%9 = call %Something-gqWpLiJylC @something-gqWpLiJylC()
	%10 = alloca %Something-gqWpLiJylC
	store %Something-gqWpLiJylC %9, %Something-gqWpLiJylC* %10
	%11 = call %Something-BGUYhRNvif @something-yEABDKbPiS()
	%12 = alloca %Something-BGUYhRNvif
	store %Something-BGUYhRNvif %11, %Something-BGUYhRNvif* %12
	%13 = getelementptr %Something-gqWpLiJylC, %Something-gqWpLiJylC* %10, i32 0, i32 0
	%14 = load i32, i32* %13
	%15 = getelementptr %Something-BGUYhRNvif, %Something-BGUYhRNvif* %8, i32 0, i32 0
	%16 = load i64, i64* %15
	%17 = trunc i64 %16 to i32
	%18 = add i32 %14, %17
	%19 = getelementptr %Something-BGUYhRNvif, %Something-BGUYhRNvif* %12, i32 0, i32 0
	%20 = load i64, i64* %19
	%21 = trunc i64 %20 to i32
	%22 = add i32 %18, %21
	%23 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%24 = call i32 (i8*, ...) @printf(i8* %23, i32 %22)
	ret void
}
