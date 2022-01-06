@"%d " = global [3 x i8] c"%d "

define i32 @main() {
_main:
	%0 = sext i32 10 to i64
	%1 = mul i64 %0, 4
	%2 = call i8* @malloc(i64 %1)
	%3 = bitcast i8* %2 to i32*
	%4 = alloca i32*
	store i32* %3, i32** %4
	%5 = load i32*, i32** %4
	call void @fillArray(i32* %5, i32 10, i32 1)
	%6 = load i32*, i32** %4
	call void @printArray(i32* %6, i32 10)
	ret i32 0
}

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc void @fillArray(i32* %array, i32 %len, i32 %value) {
fillArray:
	%0 = alloca i32*
	store i32* %array, i32** %0
	%1 = alloca i32
	store i32 %len, i32* %1
	%2 = alloca i32
	store i32 %value, i32* %2
	br label %for.declaration.MMOhOkJaqv

leave.KcjkKLLHiq:
	ret void

for.declaration.MMOhOkJaqv:
	%3 = alloca i32
	store i32 0, i32* %3
	%4 = load i32, i32* %3
	%5 = load i32, i32* %1
	%6 = icmp slt i32 %4, %5
	br i1 %6, label %for.block.oiZVPXfqQq, label %leave.KcjkKLLHiq

for.condition.ltHoNLzVDy:
	%7 = load i32, i32* %3
	%8 = load i32, i32* %1
	%9 = icmp slt i32 %7, %8
	br i1 %9, label %for.block.oiZVPXfqQq, label %leave.KcjkKLLHiq

for.block.oiZVPXfqQq:
	%10 = load i32, i32* %3
	%11 = load i32*, i32** %0
	%12 = getelementptr i32, i32* %11, i32 %10
	%13 = load i32, i32* %2
	%14 = load i32, i32* %3
	%15 = add i32 %13, %14
	store i32 %15, i32* %12
	br label %for.update.otLjYdYJrA

for.update.otLjYdYJrA:
	%16 = load i32, i32* %3
	%17 = add i32 %16, 1
	store i32 %17, i32* %3
	br label %for.condition.ltHoNLzVDy
}

define ccc void @printArray(i32* %array, i32 %len) {
printArray:
	%0 = alloca i32*
	store i32* %array, i32** %0
	%1 = alloca i32
	store i32 %len, i32* %1
	br label %for.declaration.qzADMXDoUB

leave.woBGVpPGfN:
	ret void

for.declaration.qzADMXDoUB:
	%2 = alloca i32
	store i32 0, i32* %2
	%3 = load i32, i32* %2
	%4 = load i32, i32* %1
	%5 = icmp slt i32 %3, %4
	br i1 %5, label %for.block.ijieWEgJfL, label %leave.woBGVpPGfN

for.condition.DLjnbvhXpr:
	%6 = load i32, i32* %2
	%7 = load i32, i32* %1
	%8 = icmp slt i32 %6, %7
	br i1 %8, label %for.block.ijieWEgJfL, label %leave.woBGVpPGfN

for.block.ijieWEgJfL:
	%9 = load i32, i32* %2
	%10 = load i32*, i32** %0
	%11 = getelementptr i32, i32* %10, i32 %9
	%12 = load i32, i32* %11
	%13 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%14 = call i32 (i8*, ...) @printf(i8* %13, i32 %12)
	br label %for.update.OSmbKTveXK

for.update.OSmbKTveXK:
	%15 = load i32, i32* %2
	%16 = add i32 %15, 1
	store i32 %16, i32* %2
	br label %for.condition.DLjnbvhXpr
}
