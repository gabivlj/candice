%HelloWorlds = type { i32 }
%FILE = type {}
%Point = type { i32, i32 }
%Points = type { [300 x %Point] }
%Player = type { %Point, i32, i8* }

@"%d " = global [4 x i8] c"%d \00"
@string.literal.hLZnytlvZn = global [8 x i8] c"./aFile\00"
@string.literal.mTJKFXKNLs = global [3 x i8] c"a+\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.PSCPoacBRw = global [7 x i8] c"hello!\00"
@string.literal.lSPoDZTQvT = global [12 x i8] c"Hello world\00"

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
	%7 = getelementptr [8 x i8], [8 x i8]* @string.literal.hLZnytlvZn, i32 0, i32 0
	%8 = getelementptr [3 x i8], [3 x i8]* @string.literal.mTJKFXKNLs, i32 0, i32 0
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
	br label %for.declaration.xIJSXAMedh

leave.vqpeXNriJL:
	%18 = getelementptr [7 x i8], [7 x i8]* @string.literal.PSCPoacBRw, i32 0, i32 0
	%19 = sext i32 1 to i64
	%20 = sext i32 6 to i64
	%21 = load %FILE*, %FILE** %10
	call void @fwrite(i8* %18, i64 %19, i64 %20, %FILE* %21)
	%22 = load i8*, i8** %17
	call void @free(i8* %22)
	%23 = load %FILE*, %FILE** %10
	call void @fclose(%FILE* %23)
	ret void

for.declaration.xIJSXAMedh:
	%24 = load i8*, i8** %17
	%25 = load %FILE*, %FILE** %10
	%26 = call i8* @fgets(i8* %24, i32 10, %FILE* %25)
	%27 = ptrtoint i8* %26 to i32
	%28 = icmp ne i32 %27, 0
	br i1 %28, label %for.block.ZZfBYDDGHv, label %leave.vqpeXNriJL

for.condition.EWUrpubMjg:
	%29 = load i8*, i8** %17
	%30 = load %FILE*, %FILE** %10
	%31 = call i8* @fgets(i8* %29, i32 10, %FILE* %30)
	%32 = ptrtoint i8* %31 to i32
	%33 = icmp ne i32 %32, 0
	br i1 %33, label %for.block.ZZfBYDDGHv, label %leave.vqpeXNriJL

for.block.ZZfBYDDGHv:
	%34 = load i8*, i8** %17
	%35 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%36 = call i32 (i8*, ...) @printf(i8* %35, i8* %34)
	br label %for.update.FSXxLKMauJ

for.update.FSXxLKMauJ:
	br label %for.condition.EWUrpubMjg
}

define ccc void @main() {
main:
	%0 = alloca %Player
	%1 = alloca %Point
	%2 = getelementptr %Point, %Point* %1, i32 0, i32 0
	store i32 3, i32* %2
	%3 = getelementptr %Point, %Point* %1, i32 0, i32 1
	store i32 3, i32* %3
	%4 = getelementptr %Player, %Player* %0, i32 0, i32 0
	%5 = load %Point, %Point* %1
	store %Point %5, %Point* %4
	%6 = getelementptr %Player, %Player* %0, i32 0, i32 1
	store i32 35, i32* %6
	%7 = getelementptr [12 x i8], [12 x i8]* @string.literal.lSPoDZTQvT, i32 0, i32 0
	%8 = getelementptr %Player, %Player* %0, i32 0, i32 2
	store i8* %7, i8** %8
	%9 = getelementptr %Player, %Player* %0, i32 0, i32 2
	%10 = load i8*, i8** %9
	%11 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%12 = call i32 (i8*, ...) @printf(i8* %11, i8* %10)
	%13 = getelementptr %Player, %Player* %0, i32 0, i32 0
	%14 = getelementptr %Point, %Point* %13, i32 0, i32 0
	%15 = getelementptr %Player, %Player* %0, i32 0, i32 0
	%16 = getelementptr %Point, %Point* %15, i32 0, i32 1
	%17 = load i32, i32* %16
	%18 = add i32 %17, 5
	store i32 %18, i32* %14
	%19 = getelementptr %Player, %Player* %0, i32 0, i32 0
	%20 = getelementptr %Point, %Point* %19, i32 0, i32 0
	%21 = load i32, i32* %20
	%22 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%23 = call i32 (i8*, ...) @printf(i8* %22, i32 %21)
	ret void
}
