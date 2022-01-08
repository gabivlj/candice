%HelloWorlds = type { i32 }
%FILE = type {}

@"%d " = global [4 x i8] c"%d \00"
@string.literal.rZkULSNkFz = global [6 x i8] c"./pog\00"
@string.literal.EvZUlgXtRr = global [2 x i8] c"w\00"
@string.literal.qYoxJWcVOg = global [6 x i8] c"./pog\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.fTheKPmDHN = global [7 x i8] c"hello!\00"

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

declare ccc %HelloWorlds* @helloWorldsFunction()

declare ccc %FILE* @fopen(i8* %0, i8* %1)

declare ccc void @fwrite(i8* %0, i64 %1, i64 %2, %FILE* %3)

declare ccc void @fclose(%FILE* %0)

define ccc void @main() {
main:
	%0 = call %HelloWorlds* @helloWorldsFunction()
	%1 = alloca %HelloWorlds*
	store %HelloWorlds* %0, %HelloWorlds** %1
	%2 = load %HelloWorlds*, %HelloWorlds** %1
	%3 = getelementptr %HelloWorlds, %HelloWorlds* %2, i32 0, i32 0
	%4 = load i32, i32* %3
	%5 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%6 = call i32 (i8*, ...) @printf(i8* %5, i32 %4)
	%7 = getelementptr [6 x i8], [6 x i8]* @string.literal.rZkULSNkFz, i32 0, i32 0
	%8 = getelementptr [2 x i8], [2 x i8]* @string.literal.EvZUlgXtRr, i32 0, i32 0
	%9 = call %FILE* @fopen(i8* %7, i8* %8)
	%10 = alloca %FILE*
	store %FILE* %9, %FILE** %10
	%11 = getelementptr [6 x i8], [6 x i8]* @string.literal.qYoxJWcVOg, i32 0, i32 0
	%12 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%13 = call i32 (i8*, ...) @printf(i8* %12, i8* %11)
	%14 = getelementptr [7 x i8], [7 x i8]* @string.literal.fTheKPmDHN, i32 0, i32 0
	%15 = sext i32 1 to i64
	%16 = sext i32 6 to i64
	%17 = load %FILE*, %FILE** %10
	call void @fwrite(i8* %14, i64 %15, i64 %16, %FILE* %17)
	%18 = load %FILE*, %FILE** %10
	call void @fclose(%FILE* %18)
	ret void
}
