%S-nrkpRhSffM = type { i32 }

@string.literal.pMVDnhvzCF = global [5 x i8] c"badS\00"
@"%s " = global [4 x i8] c"%s \00"
@string.literal.LilBmTJzfz = global [6 x i8] c"badSS\00"

define void @s-nrkpRhSffM(%S-nrkpRhSffM %c-nrkpRhSffM) {
s-nrkpRhSffM:
	%0 = alloca %S-nrkpRhSffM
	store %S-nrkpRhSffM %c-nrkpRhSffM, %S-nrkpRhSffM* %0
	%1 = getelementptr inbounds %S-nrkpRhSffM, %S-nrkpRhSffM* %0, i32 0, i32 0
	%2 = load i32, i32* %1
	%3 = icmp ne i32 %2, 4
	br i1 %3, label %if.then.sYZPAPrUZv, label %if.else.XhYGxPxGTI

if.then.sYZPAPrUZv:
	%4 = getelementptr inbounds [5 x i8], [5 x i8]* @string.literal.pMVDnhvzCF, i32 0, i32 0
	%5 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%6 = call i32 (i8*, ...) @printf(i8* %5, i8* %4)
	br label %lastLeave.amrfGBWifp

if.else.XhYGxPxGTI:
	br label %lastLeave.amrfGBWifp

lastLeave.amrfGBWifp:
	ret void
}

define void @ss-nrkpRhSffM(%S-nrkpRhSffM* %c-nrkpRhSffM) {
ss-nrkpRhSffM:
	%0 = alloca %S-nrkpRhSffM*
	store %S-nrkpRhSffM* %c-nrkpRhSffM, %S-nrkpRhSffM** %0
	%1 = load %S-nrkpRhSffM*, %S-nrkpRhSffM** %0
	%2 = getelementptr inbounds %S-nrkpRhSffM, %S-nrkpRhSffM* %1, i32 0, i32 0
	%3 = load i32, i32* %2
	%4 = icmp ne i32 %3, 3
	br i1 %4, label %if.then.FZHDRGTrrC, label %if.else.wlFMagWzCo

if.then.FZHDRGTrrC:
	%5 = getelementptr inbounds [6 x i8], [6 x i8]* @string.literal.LilBmTJzfz, i32 0, i32 0
	%6 = getelementptr inbounds [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%7 = call i32 (i8*, ...) @printf(i8* %6, i8* %5)
	br label %lastLeave.ZHbKijcsaf

if.else.wlFMagWzCo:
	br label %lastLeave.ZHbKijcsaf

lastLeave.ZHbKijcsaf:
	ret void
}

define i32 @main() {
main:
	%0 = alloca %S-nrkpRhSffM
	%1 = getelementptr inbounds %S-nrkpRhSffM, %S-nrkpRhSffM* %0, i32 0, i32 0
	store i32 4, i32* %1
	%2 = load %S-nrkpRhSffM, %S-nrkpRhSffM* %0
	call void @s-nrkpRhSffM(%S-nrkpRhSffM %2)
	%3 = alloca %S-nrkpRhSffM
	%4 = getelementptr inbounds %S-nrkpRhSffM, %S-nrkpRhSffM* %3, i32 0, i32 0
	store i32 3, i32* %4
	%5 = alloca %S-nrkpRhSffM*
	store %S-nrkpRhSffM* %3, %S-nrkpRhSffM** %5
	%6 = load %S-nrkpRhSffM*, %S-nrkpRhSffM** %5
	call void @ss-nrkpRhSffM(%S-nrkpRhSffM* %6)
	%7 = alloca %S-nrkpRhSffM
	%8 = getelementptr inbounds %S-nrkpRhSffM, %S-nrkpRhSffM* %7, i32 0, i32 0
	store i32 4, i32* %8
	%9 = load %S-nrkpRhSffM, %S-nrkpRhSffM* %7
	call void @s-nrkpRhSffM(%S-nrkpRhSffM %9)
	ret i32 0
}

declare ccc i32 @printf(i8* %0, ...)
