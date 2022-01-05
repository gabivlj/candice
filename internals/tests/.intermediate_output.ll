@"%d " = global [3 x i8] c"%d "

define i32 @main() {
_main:
	%0 = alloca i32
	store i32 3, i32* %0
	%1 = alloca i32
	store i32 4, i32* %1
	%2 = icmp ne i32 0, 0
	%3 = load i32, i32* %0
	%4 = load i32, i32* %1
	%5 = icmp sgt i32 %3, %4
	%6 = load i32, i32* %0
	%7 = load i32, i32* %1
	%8 = icmp eq i32 %6, %7
	%9 = load i32, i32* %0
	%10 = load i32, i32* %1
	%11 = icmp sge i32 %9, %10
	%12 = load i32, i32* %0
	%13 = load i32, i32* %1
	%14 = icmp sle i32 %12, %13
	%15 = load i32, i32* %0
	%16 = icmp ne i32 %15, 0
	br i1 %2, label %if.then.mMThlunwKU, label %leave.hQxngVNOqq

if.then.mMThlunwKU:
	%17 = load i32, i32* %0
	%18 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%19 = call i32 (i8*, ...) @printf(i8* %18, i32 %17)
	br label %leave.OOHYwciEZk

if.else.vhNngJduLX:
	%20 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%21 = call i32 (i8*, ...) @printf(i8* %20, i32 1)
	br label %leave.OOHYwciEZk

elseif.then.OdraDksUTF:
	%22 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%23 = call i32 (i8*, ...) @printf(i8* %22, i32 100)
	br label %leave.OOHYwciEZk

elseif.then.KMUrSJAnPU:
	%24 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%25 = call i32 (i8*, ...) @printf(i8* %24, i32 200)
	br label %leave.OOHYwciEZk

elseif.then.mXGmoDoKrr:
	%26 = load i32, i32* %0
	%27 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%28 = call i32 (i8*, ...) @printf(i8* %27, i32 %26)
	br label %leave.OOHYwciEZk

elseif.then.OtsIRBpufa:
	%29 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%30 = call i32 (i8*, ...) @printf(i8* %29, i32 301)
	br label %leave.OOHYwciEZk

elseif.then.VpWENxcpaW:
	%31 = load i32, i32* %1
	%32 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%33 = call i32 (i8*, ...) @printf(i8* %32, i32 %31)
	br label %leave.OOHYwciEZk

leave.hQxngVNOqq:
	br i1 %5, label %elseif.then.OdraDksUTF, label %leave.imfLuYgzkp

leave.imfLuYgzkp:
	br i1 %8, label %elseif.then.KMUrSJAnPU, label %leave.DFcrhYrysw

leave.DFcrhYrysw:
	br i1 %11, label %elseif.then.mXGmoDoKrr, label %leave.FlfwfKdPdK

leave.FlfwfKdPdK:
	br i1 %14, label %elseif.then.OtsIRBpufa, label %leave.muYBDwCLac

leave.muYBDwCLac:
	br i1 %16, label %elseif.then.VpWENxcpaW, label %if.else.vhNngJduLX

leave.OOHYwciEZk:
	ret i32 0
}

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)
