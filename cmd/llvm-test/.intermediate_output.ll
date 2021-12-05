@seed = global i32 12
@printinteger = global [11 x i8] c"Result: %d\0A"

declare i32 @printf(i8* %0, ...)

declare void @printCool()

declare i32 @abs(i32 %x)

define i32 @rand() {
0:
	%1 = load i32, i32* @seed
	%2 = mul i32 %1, 22695477
	%3 = add i32 %2, 1
	store i32 %3, i32* @seed
	%4 = call i32 @abs(i32 %3)
	ret i32 %4
}

define i32 @main() {
0:
	%1 = call i32 @rand()
	%2 = getelementptr [11 x i8], [11 x i8]* @printinteger, i32 0, i32 0
	%3 = call i8* @malloc(i64 16)
	%4 = bitcast i8* %3 to i32*
	%5 = alloca i32*
	store i32* %4, i32** %5
	%6 = load i32*, i32** %5
	store i32 32, i32* %6
	%7 = getelementptr i32, i32* %6, i32 0
	%8 = load i32, i32* %7
	%9 = call i32 (i8*, ...) @printf(i8* %2, i32 %8)
	%10 = call i32 (i8*, ...) @printf(i8* %2, i32 %1)
	call void @printCool()
	%11 = alloca { i32, i32 }
	%12 = getelementptr { i32, i32 }, { i32, i32 }* %11, i32 0, i32 0
	store i32 43, i32* %12
	%13 = getelementptr { i32, i32 }, { i32, i32 }* %11, i32 0, i32 1
	store i32 3223, i32* %13
	%14 = load i32, i32* %12
	%15 = call i32 (i8*, ...) @printf(i8* %2, i32 %14)
	%16 = load i32, i32* %13
	%17 = call i32 (i8*, ...) @printf(i8* %2, i32 %16)
	ret i32 0
}

declare i8* @malloc(i64 %size)
