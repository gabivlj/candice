%Some-OSpkGXunBk = type {}

@string.literal.LNLkpUsMpT = global [28 x i8] c"We can't handle this error!\00"
@"%s " = global [4 x i8] c"%s \00"

define { %Some-OSpkGXunBk*, i32 }* @CreateSome-OSpkGXunBk() {
CreateSome-OSpkGXunBk:
	%0 = inttoptr i32 0 to %Some-OSpkGXunBk*
	%1 = alloca { %Some-OSpkGXunBk*, i32 }
	%2 = getelementptr { %Some-OSpkGXunBk*, i32 }, { %Some-OSpkGXunBk*, i32 }* %1, i32 0, i32 0
	store %Some-OSpkGXunBk* %0, %Some-OSpkGXunBk** %2
	%3 = getelementptr { %Some-OSpkGXunBk*, i32 }, { %Some-OSpkGXunBk*, i32 }* %1, i32 0, i32 1
	store i32 1, i32* %3
	ret { %Some-OSpkGXunBk*, i32 }* %1
}

define i32 @main() {
main:
	%0 = call { %Some-OSpkGXunBk*, i32 }* @CreateSome-OSpkGXunBk()
	%1 = getelementptr { %Some-OSpkGXunBk*, i32 }, { %Some-OSpkGXunBk*, i32 }* %0, i32 0, i32 0
	%2 = getelementptr { %Some-OSpkGXunBk*, i32 }, { %Some-OSpkGXunBk*, i32 }* %0, i32 0, i32 1
	%3 = load i32, i32* %2
	switch i32 %3, label %default.faZtmdAIPc [
		i32 1, label %case-0-uWyzxjfdzk
	]

default.faZtmdAIPc:
	%4 = getelementptr inbounds [28 x i8], [28 x i8]* @string.literal.LNLkpUsMpT, i32 0, i32 0
	%5 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%6 = call i32 (i8*, ...) @printf(i8* %5, i8* %4)
	br label %leaveSwitchBlock.eXwxzzPYLU

case-0-uWyzxjfdzk:
	br label %leaveSwitchBlock.eXwxzzPYLU

leaveSwitchBlock.eXwxzzPYLU:
	%7 = call { %Some-OSpkGXunBk*, i32 }* @CreateSome-OSpkGXunBk()
	%8 = getelementptr { %Some-OSpkGXunBk*, i32 }, { %Some-OSpkGXunBk*, i32 }* %7, i32 0, i32 0
	%9 = getelementptr { %Some-OSpkGXunBk*, i32 }, { %Some-OSpkGXunBk*, i32 }* %7, i32 0, i32 1
	ret i32 0
}

declare ccc i32 @printf(i8* %0, ...)
