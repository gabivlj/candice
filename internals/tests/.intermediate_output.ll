%aStruct-mbwyqfdiKU = type { i32, i64 }

@"%u " = global [4 x i8] c"%u \00"
@"%d " = global [4 x i8] c"%d \00"
@string.literal.UbNrXLpohl = global [21 x i8] c"this shouldn't print\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.NRaeSFEsiG = global [30 x i8] c"this should not print as well\00"
@string.literal.VvtJxjXnxE = global [25 x i8] c"this shouldnt print!!!!!\00"

define i32 @loadByPointer-mbwyqfdiKU(i32* %i-mbwyqfdiKU) {
loadByPointer-mbwyqfdiKU:
	%0 = alloca i32*
	store i32* %i-mbwyqfdiKU, i32** %0
	%1 = load i32*, i32** %0
	%2 = load i32, i32* %1
	ret i32 %2
}

define i32 @main() {
main:
	%0 = alloca i32
	store i32 u0xFFFFFF3C, i32* %0
	%1 = load i32, i32* %0
	%2 = add i32 %1, 100
	store i32 %2, i32* %0
	%3 = load i32, i32* %0
	%4 = load i32, i32* %0
	%5 = icmp eq i32 %3, %4
	br i1 %5, label %if.then.xDMYXogAGW, label %if.else.ZBVNQZludm

if.then.xDMYXogAGW:
	%6 = load i32, i32* %0
	%7 = getelementptr inbounds [4 x i8], [4 x i8]* @"%u ", i32 0, i32 0
	%8 = call i32 (i8*, ...) @printf(i8* %7, i32 %6)
	%9 = load i32, i32* %0
	%10 = getelementptr inbounds [4 x i8], [4 x i8]* @"%d ", i32 0, i32 0
	%11 = call i32 (i8*, ...) @printf(i8* %10, i32 %9)
	br label %lastLeave.FNnySxysen

if.else.ZBVNQZludm:
	br label %lastLeave.FNnySxysen

lastLeave.FNnySxysen:
	%12 = alloca i32
	store i32 3, i32* %12
	%13 = alloca i32*
	store i32* %12, i32** %13
	%14 = load i32*, i32** %13
	%15 = call i32 @loadByPointer-mbwyqfdiKU(i32* %14)
	%16 = alloca i32
	store i32 %15, i32* %16
	%17 = load i32, i32* %16
	%18 = icmp ne i32 %17, 3
	br i1 %18, label %if.then.XnsjeKkMiV, label %if.else.cyIKRaFBkY

if.then.XnsjeKkMiV:
	%19 = getelementptr inbounds [21 x i8], [21 x i8]* @string.literal.UbNrXLpohl, i32 0, i32 0
	%20 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%21 = call i32 (i8*, ...) @printf(i8* %20, i8* %19)
	br label %lastLeave.SEXQRsccMX

if.else.cyIKRaFBkY:
	br label %lastLeave.SEXQRsccMX

lastLeave.SEXQRsccMX:
	%22 = alloca %aStruct-mbwyqfdiKU
	%23 = getelementptr inbounds %aStruct-mbwyqfdiKU, %aStruct-mbwyqfdiKU* %22, i32 0, i32 0
	store i32 3, i32* %23
	%24 = alloca %aStruct-mbwyqfdiKU*
	store %aStruct-mbwyqfdiKU* %22, %aStruct-mbwyqfdiKU** %24
	%25 = load %aStruct-mbwyqfdiKU*, %aStruct-mbwyqfdiKU** %24
	%26 = bitcast %aStruct-mbwyqfdiKU* %25 to i32*
	%27 = alloca i32*
	store i32* %26, i32** %27
	%28 = load i32*, i32** %27
	%29 = getelementptr inbounds i32, i32* %28, i32 0
	%30 = load i32, i32* %29
	%31 = getelementptr inbounds %aStruct-mbwyqfdiKU, %aStruct-mbwyqfdiKU* %22, i32 0, i32 0
	%32 = load i32, i32* %31
	%33 = icmp ne i32 %30, %32
	br i1 %33, label %if.then.geszdQTewn, label %if.else.EdcOhsznXZ

if.then.geszdQTewn:
	%34 = getelementptr inbounds [30 x i8], [30 x i8]* @string.literal.NRaeSFEsiG, i32 0, i32 0
	%35 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%36 = call i32 (i8*, ...) @printf(i8* %35, i8* %34)
	br label %lastLeave.dLSmXHxzfi

if.else.EdcOhsznXZ:
	br label %lastLeave.dLSmXHxzfi

lastLeave.dLSmXHxzfi:
	%37 = load i32*, i32** %27
	%38 = bitcast i32* %37 to %aStruct-mbwyqfdiKU*
	%39 = alloca %aStruct-mbwyqfdiKU
	%40 = load %aStruct-mbwyqfdiKU, %aStruct-mbwyqfdiKU* %38
	store %aStruct-mbwyqfdiKU %40, %aStruct-mbwyqfdiKU* %39
	%41 = getelementptr inbounds %aStruct-mbwyqfdiKU, %aStruct-mbwyqfdiKU* %39, i32 0, i32 0
	%42 = load i32, i32* %41
	%43 = getelementptr inbounds %aStruct-mbwyqfdiKU, %aStruct-mbwyqfdiKU* %22, i32 0, i32 0
	%44 = load i32, i32* %43
	%45 = icmp ne i32 %42, %44
	br i1 %45, label %if.then.ogKexyydJk, label %if.else.PxquGByRLq

if.then.ogKexyydJk:
	%46 = getelementptr inbounds [25 x i8], [25 x i8]* @string.literal.VvtJxjXnxE, i32 0, i32 0
	%47 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%48 = call i32 (i8*, ...) @printf(i8* %47, i8* %46)
	br label %lastLeave.oZRnThNUGL

if.else.PxquGByRLq:
	br label %lastLeave.oZRnThNUGL

lastLeave.oZRnThNUGL:
	ret i32 0
}

declare ccc i32 @printf(i8* %0, ...)
