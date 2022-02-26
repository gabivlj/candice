@string.literal.otVosLuoFi = global [4 x i8] c"bad\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.OyPuSvcQzR = global [5 x i8] c"bad1\00"
@string.literal.UhBvpKjKQL = global [5 x i8] c"bad2\00"
@string.literal.srigOTALDx = global [5 x i8] c"bad3\00"

declare ccc void @free(i8* %0, ...)

declare ccc i8* @realloc(i8* %0, i64 %1, ...)

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define i32 @main() {
main:
	%0 = alloca float
	store float 1.0, float* %0
	%1 = alloca float
	store float 2.0, float* %1
	%2 = load float, float* %0
	%3 = load float, float* %1
	%4 = fcmp oeq float %2, %3
	br i1 %4, label %if.then.FBcklYofsS, label %if.else.EgdRKjaLjw

if.then.FBcklYofsS:
	%5 = getelementptr [4 x i8], [4 x i8]* @string.literal.otVosLuoFi, i32 0, i32 0
	%6 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%7 = call i32 (i8*, ...) @printf(i8* %6, i8* %5)
	br label %lastLeave.kZjxUmFPVU

if.else.EgdRKjaLjw:
	br label %lastLeave.kZjxUmFPVU

lastLeave.kZjxUmFPVU:
	%8 = load float, float* %0
	%9 = load float, float* %1
	%10 = fcmp ogt float %8, %9
	br i1 %10, label %if.then.ufoGJMbCwr, label %if.else.jPJNOGccMT

if.then.ufoGJMbCwr:
	%11 = getelementptr [5 x i8], [5 x i8]* @string.literal.OyPuSvcQzR, i32 0, i32 0
	%12 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%13 = call i32 (i8*, ...) @printf(i8* %12, i8* %11)
	br label %lastLeave.TBHjJcSxcL

if.else.jPJNOGccMT:
	br label %lastLeave.TBHjJcSxcL

lastLeave.TBHjJcSxcL:
	%14 = load float, float* %0
	%15 = load float, float* %1
	%16 = fcmp oge float %14, %15
	br i1 %16, label %if.then.uQaxMSdpXe, label %if.else.zbKjmmzAkr

if.then.uQaxMSdpXe:
	%17 = getelementptr [5 x i8], [5 x i8]* @string.literal.UhBvpKjKQL, i32 0, i32 0
	%18 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%19 = call i32 (i8*, ...) @printf(i8* %18, i8* %17)
	br label %lastLeave.TYdfDqXfui

if.else.zbKjmmzAkr:
	br label %lastLeave.TYdfDqXfui

lastLeave.TYdfDqXfui:
	store float 1.0, float* %1
	%20 = load float, float* %0
	%21 = load float, float* %1
	%22 = fcmp one float %20, %21
	br i1 %22, label %if.then.dHHjEfnGfn, label %if.else.tQNmexwHJD

if.then.dHHjEfnGfn:
	%23 = getelementptr [5 x i8], [5 x i8]* @string.literal.srigOTALDx, i32 0, i32 0
	%24 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%25 = call i32 (i8*, ...) @printf(i8* %24, i8* %23)
	br label %lastLeave.IZfGMlghhj

if.else.tQNmexwHJD:
	br label %lastLeave.IZfGMlghhj

lastLeave.IZfGMlghhj:
	ret i32 0
}
