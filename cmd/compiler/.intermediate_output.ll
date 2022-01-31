%T-mFrmprpMDU = type i64
%Something-mFrmprpMDU = type { i64 }
%T-BzzWyAecWS = type i32
%Something-BzzWyAecWS = type { i32 }

@string.literal.dIgCbeDyXM = global [13 x i8] c"hello world!\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.uoTUKzfNsv = global [13 x i8] c"hello world!\00"
@"%d " = global [4 x i8] c"%d \00"

declare ccc i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc void @outsideFunction-mFrmprpMDU(%Something-mFrmprpMDU* %s-mFrmprpMDU) {
outsideFunction-mFrmprpMDU:
	%0 = alloca %Something-mFrmprpMDU*
	store %Something-mFrmprpMDU* %s-mFrmprpMDU, %Something-mFrmprpMDU** %0
	%1 = load %Something-mFrmprpMDU*, %Something-mFrmprpMDU** %0
	%2 = getelementptr %Something-mFrmprpMDU, %Something-mFrmprpMDU* %1, i32 0, i32 0
	%3 = load %Something-mFrmprpMDU*, %Something-mFrmprpMDU** %0
	%4 = getelementptr %Something-mFrmprpMDU, %Something-mFrmprpMDU* %3, i32 0, i32 0
	%5 = load i64, i64* %4
	%6 = sext i32 1 to i64
	%7 = add i64 %5, %6
	store i64 %7, i64* %2
	ret void
}

define ccc %Something-mFrmprpMDU @something-mFrmprpMDU() {
something-mFrmprpMDU:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.dIgCbeDyXM, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-mFrmprpMDU
	%4 = sext i32 0 to i64
	%5 = getelementptr %Something-mFrmprpMDU, %Something-mFrmprpMDU* %3, i32 0, i32 0
	store i64 %4, i64* %5
	%6 = alloca %Something-mFrmprpMDU
	%7 = load %Something-mFrmprpMDU, %Something-mFrmprpMDU* %3
	store %Something-mFrmprpMDU %7, %Something-mFrmprpMDU* %6
	%8 = alloca %Something-mFrmprpMDU*
	store %Something-mFrmprpMDU* %6, %Something-mFrmprpMDU** %8
	%9 = load %Something-mFrmprpMDU*, %Something-mFrmprpMDU** %8
	call void @outsideFunction-mFrmprpMDU(%Something-mFrmprpMDU* %9)
	%10 = load %Something-mFrmprpMDU, %Something-mFrmprpMDU* %6
	ret %Something-mFrmprpMDU %10
}

define ccc void @outsideFunction-BzzWyAecWS(%Something-BzzWyAecWS* %s-BzzWyAecWS) {
outsideFunction-BzzWyAecWS:
	%0 = alloca %Something-BzzWyAecWS*
	store %Something-BzzWyAecWS* %s-BzzWyAecWS, %Something-BzzWyAecWS** %0
	%1 = load %Something-BzzWyAecWS*, %Something-BzzWyAecWS** %0
	%2 = getelementptr %Something-BzzWyAecWS, %Something-BzzWyAecWS* %1, i32 0, i32 0
	%3 = load %Something-BzzWyAecWS*, %Something-BzzWyAecWS** %0
	%4 = getelementptr %Something-BzzWyAecWS, %Something-BzzWyAecWS* %3, i32 0, i32 0
	%5 = load i32, i32* %4
	%6 = add i32 %5, 1
	store i32 %6, i32* %2
	ret void
}

define ccc %Something-BzzWyAecWS @something-BzzWyAecWS() {
something-BzzWyAecWS:
	%0 = getelementptr [13 x i8], [13 x i8]* @string.literal.uoTUKzfNsv, i32 0, i32 0
	%1 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%2 = call i32 (i8*, ...) @printf(i8* %1, i8* %0)
	%3 = alloca %Something-BzzWyAecWS
	%4 = getelementptr %Something-BzzWyAecWS, %Something-BzzWyAecWS* %3, i32 0, i32 0
	store i32 0, i32* %4
	%5 = alloca %Something-BzzWyAecWS
	%6 = load %Something-BzzWyAecWS, %Something-BzzWyAecWS* %3
	store %Something-BzzWyAecWS %6, %Something-BzzWyAecWS* %5
	%7 = alloca %Something-BzzWyAecWS*
	store %Something-BzzWyAecWS* %5, %Something-BzzWyAecWS** %7
	%8 = load %Something-BzzWyAecWS*, %Something-BzzWyAecWS** %7
	call void @outsideFunction-BzzWyAecWS(%Something-BzzWyAecWS* %8)
	%9 = load %Something-BzzWyAecWS, %Something-BzzWyAecWS* %5
	ret %Something-BzzWyAecWS %9
}

define ccc %Something-mFrmprpMDU @something-DnpNFCDCHd() {
something-DnpNFCDCHd:
	%0 = alloca %Something-mFrmprpMDU
	%1 = sext i32 3 to i64
	%2 = getelementptr %Something-mFrmprpMDU, %Something-mFrmprpMDU* %0, i32 0, i32 0
	store i64 %1, i64* %2
	%3 = load %Something-mFrmprpMDU, %Something-mFrmprpMDU* %0
	ret %Something-mFrmprpMDU %3
}

define ccc void @main() {
main:
	%0 = call %Something-mFrmprpMDU @something-mFrmprpMDU()
	%1 = alloca %Something-mFrmprpMDU
	store %Something-mFrmprpMDU %0, %Something-mFrmprpMDU* %1
	%2 = call %Something-BzzWyAecWS @something-BzzWyAecWS()
	%3 = alloca %Something-BzzWyAecWS
	store %Something-BzzWyAecWS %2, %Something-BzzWyAecWS* %3
	%4 = call %Something-mFrmprpMDU @something-DnpNFCDCHd()
	%5 = alloca %Something-mFrmprpMDU
	store %Something-mFrmprpMDU %4, %Something-mFrmprpMDU* %5
	%6 = getelementptr %Something-BzzWyAecWS, %Something-BzzWyAecWS* %3, i32 0, i32 0
	%7 = load i32, i32* %6
	%8 = getelementptr %Something-mFrmprpMDU, %Something-mFrmprpMDU* %1, i32 0, i32 0
	%9 = load i64, i64* %8
	%10 = trunc i64 %9 to i32
	%11 = add i32 %7, %10
	%12 = getelementptr %Something-mFrmprpMDU, %Something-mFrmprpMDU* %5, i32 0, i32 0
	%13 = load i64, i64* %12
	%14 = trunc i64 %13 to i32
	%15 = add i32 %11, %14
	%16 = getelementptr [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%17 = call i32 (i8*, ...) @printf(i8* %16, i32 %15)
	ret void
}
