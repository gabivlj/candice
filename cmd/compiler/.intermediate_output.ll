%Punto = type { i32, i32 }

@"%d " = global [4 x i8] c"%d \00"

define i32 @main() {
_main:
	%0 = alloca i32
	store i32 3, i32* %0
	%1 = alloca [30 x i32]
	br label %for.declaration.cMOQPTKxij

leave.pbqcqllEPl:
	%2 = getelementptr [30 x i32], [30 x i32]* %1, i32 0, i32 29
	%3 = load i32, i32* %2
	%4 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%5 = call i32 (i8*, ...) @printf(i8* %4, i32 %3)
	%6 = alloca [30 x %Punto]
	%7 = call %Punto @CrearPunto(i32 3, i32 4)
	%8 = alloca %Punto
	store %Punto %7, %Punto* %8
	%9 = load %Punto, %Punto* %8
	%10 = getelementptr [30 x %Punto], [30 x %Punto]* %6, i32 0, i32 0
	store %Punto %9, %Punto* %10
	%11 = getelementptr [30 x %Punto], [30 x %Punto]* %6, i32 0, i32 0
	%12 = getelementptr %Punto, %Punto* %11, i32 0, i32 0
	%13 = load i32, i32* %12
	%14 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%15 = call i32 (i8*, ...) @printf(i8* %14, i32 %13)
	ret i32 0

for.declaration.cMOQPTKxij:
	%16 = alloca i32
	store i32 0, i32* %16
	%17 = load i32, i32* %16
	%18 = icmp slt i32 %17, 30
	br i1 %18, label %for.block.kIHiuVcxTX, label %leave.pbqcqllEPl

for.condition.cfTpDUVPDF:
	%19 = load i32, i32* %16
	%20 = icmp slt i32 %19, 30
	br i1 %20, label %for.block.kIHiuVcxTX, label %leave.pbqcqllEPl

for.block.kIHiuVcxTX:
	%21 = load i32, i32* %16
	%22 = getelementptr [30 x i32], [30 x i32]* %1, i32 0, i32 %21
	%23 = load i32, i32* %16
	store i32 %23, i32* %22
	br label %for.update.REneHEygLc

for.update.REneHEygLc:
	%24 = load i32, i32* %16
	%25 = add i32 %24, 1
	store i32 %25, i32* %16
	br label %for.condition.cfTpDUVPDF
}

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc %Punto @CrearPunto(i32 %x, i32 %y) {
CrearPunto:
	%0 = alloca i32
	store i32 %x, i32* %0
	%1 = alloca i32
	store i32 %y, i32* %1
	%2 = alloca %Punto
	%3 = getelementptr %Punto, %Punto* %2, i32 0, i32 0
	%4 = load i32, i32* %0
	store i32 %4, i32* %3
	%5 = getelementptr %Punto, %Punto* %2, i32 0, i32 1
	%6 = load i32, i32* %1
	store i32 %6, i32* %5
	%7 = load %Punto, %Punto* %2
	ret %Punto %7
}
