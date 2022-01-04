@"%d " = global [3 x i8] c"%d "

define i32 @main() {
_main:
	%0 = alloca i32
	store i32 3, i32* %0
	%1 = alloca i32
	store i32 4, i32* %1
	%2 = load i32, i32* %0
	%3 = load i32, i32* %1
	%4 = icmp eq i32 %2, %3
	br i1 %4, label %if.then.HpgxJhmNiR, label %if.else.TynCpAKyAR

if.then.HpgxJhmNiR:
	%5 = load i32, i32* %0
	%6 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%7 = call i32 (i8*, ...) @printf(i8* %6, i32 %5)
	br label %leave.gpUpAwAuwv

if.else.TynCpAKyAR:
	br label %leave.gpUpAwAuwv

leave.gpUpAwAuwv:
	%8 = load i32, i32* %0
	%9 = load i32, i32* %1
	%10 = icmp ne i32 %8, %9
	br i1 %10, label %if.then.mllsfuaxfl, label %if.else.xspwFMLNVm

if.then.mllsfuaxfl:
	%11 = load i32, i32* %1
	%12 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%13 = call i32 (i8*, ...) @printf(i8* %12, i32 %11)
	br label %leave.oQBIlJiTwQ

if.else.xspwFMLNVm:
	br label %leave.oQBIlJiTwQ

leave.oQBIlJiTwQ:
	%14 = load i32, i32* %0
	%15 = load i32, i32* %1
	%16 = icmp sle i32 %14, %15
	br i1 %16, label %if.then.rJXVcHKiqJ, label %if.else.YPztxeVWtD

if.then.rJXVcHKiqJ:
	%17 = load i32, i32* %1
	%18 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%19 = call i32 (i8*, ...) @printf(i8* %18, i32 %17)
	br label %leave.yjnSuhFKFo

if.else.YPztxeVWtD:
	br label %leave.yjnSuhFKFo

leave.yjnSuhFKFo:
	%20 = load i32, i32* %0
	%21 = load i32, i32* %1
	%22 = icmp sge i32 %20, %21
	br i1 %22, label %if.then.GfBSKQzTox, label %if.else.bDMhYcIFqW

if.then.GfBSKQzTox:
	%23 = load i32, i32* %0
	%24 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%25 = call i32 (i8*, ...) @printf(i8* %24, i32 %23)
	br label %leave.CQdNAWXfvy

if.else.bDMhYcIFqW:
	br label %leave.CQdNAWXfvy

leave.CQdNAWXfvy:
	%26 = load i32, i32* %0
	%27 = load i32, i32* %1
	%28 = icmp eq i32 %26, %27
	%29 = icmp eq i1 %28, 0
	br i1 %29, label %if.then.UtEVtPKLjK, label %if.else.niOoByAetg

if.then.UtEVtPKLjK:
	%30 = load i32, i32* %1
	%31 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%32 = call i32 (i8*, ...) @printf(i8* %31, i32 %30)
	br label %leave.dmomXeJvvf

if.else.niOoByAetg:
	br label %leave.dmomXeJvvf

leave.dmomXeJvvf:
	%33 = load i32, i32* %0
	%34 = icmp ne i32 %33, 0
	%35 = icmp eq i1 %34, 0
	br i1 %35, label %if.then.IqdSprxaIa, label %if.else.xnOeuhggwN

if.then.IqdSprxaIa:
	%36 = load i32, i32* %1
	%37 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%38 = call i32 (i8*, ...) @printf(i8* %37, i32 %36)
	br label %leave.WpHkeAtIDm

if.else.xnOeuhggwN:
	br label %leave.WpHkeAtIDm

leave.WpHkeAtIDm:
	%39 = load i32, i32* %0
	%40 = icmp ne i32 %39, 0
	br i1 %40, label %if.then.mkMHfhOcJw, label %if.else.qZGBFWGgiJ

if.then.mkMHfhOcJw:
	%41 = load i32, i32* %1
	%42 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%43 = call i32 (i8*, ...) @printf(i8* %42, i32 %41)
	br label %leave.nzIviZeTog

if.else.qZGBFWGgiJ:
	br label %leave.nzIviZeTog

leave.nzIviZeTog:
	ret i32 0
}

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)
