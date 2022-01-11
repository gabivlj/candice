%Person = type { i8*, i8, i1, %Person*, [10 x i32] }

@string.literal.ignGMdYsqD = global [6 x i8] c"James\00"
@string.literal.GjgLKuYzKS = global [6 x i8] c"Pedro\00"
@"%s " = global [4 x i8] c"%s \00"

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)

define ccc %Person* @DefaultPerson(i1 %isCool) {
DefaultPerson:
	%0 = alloca i1
	store i1 %isCool, i1* %0
	%1 = sext i32 1 to i64
	%2 = mul i64 %1, 0
	%3 = call i8* @malloc(i64 %2)
	%4 = bitcast i8* %3 to %Person*
	%5 = alloca %Person*
	store %Person* %4, %Person** %5
	%6 = load i1, i1* %0
	br i1 %6, label %if.then.bGvBplhbbB, label %if.else.TSAoCnZlJE

if.then.bGvBplhbbB:
	%7 = load %Person*, %Person** %5
	%8 = getelementptr %Person, %Person* %7, i32 0, i32 1
	%9 = trunc i32 21 to i8
	store i8 %9, i8* %8
	%10 = load %Person*, %Person** %5
	%11 = getelementptr %Person, %Person* %10, i32 0, i32 0
	%12 = getelementptr [6 x i8], [6 x i8]* @string.literal.ignGMdYsqD, i32 0, i32 0
	store i8* %12, i8** %11
	br label %leave.DsduvsuWDq

if.else.TSAoCnZlJE:
	%13 = load %Person*, %Person** %5
	%14 = getelementptr %Person, %Person* %13, i32 0, i32 1
	%15 = trunc i32 66 to i8
	store i8 %15, i8* %14
	%16 = load %Person*, %Person** %5
	%17 = getelementptr %Person, %Person* %16, i32 0, i32 0
	%18 = getelementptr [6 x i8], [6 x i8]* @string.literal.GjgLKuYzKS, i32 0, i32 0
	store i8* %18, i8** %17
	br label %leave.DsduvsuWDq

leave.DsduvsuWDq:
	%19 = load %Person*, %Person** %5
	%20 = getelementptr %Person, %Person* %19, i32 0, i32 2
	%21 = load i1, i1* %0
	store i1 %21, i1* %20
	%22 = load %Person*, %Person** %5
	%23 = getelementptr %Person, %Person* %22, i32 0, i32 3
	%24 = inttoptr i32 0 to %Person*
	store %Person* %24, %Person** %23
	%25 = load %Person*, %Person** %5
	%26 = getelementptr %Person, %Person* %25, i32 0, i32 4
	%27 = alloca [10 x i32]
	%28 = getelementptr [10 x i32], [10 x i32]* %27, i32 0, i32 0
	store i32 1, i32* %28
	%29 = getelementptr [10 x i32], [10 x i32]* %27, i32 0, i32 1
	store i32 1, i32* %29
	%30 = getelementptr [10 x i32], [10 x i32]* %27, i32 0, i32 2
	store i32 1, i32* %30
	%31 = getelementptr [10 x i32], [10 x i32]* %27, i32 0, i32 3
	store i32 1, i32* %31
	%32 = load [10 x i32], [10 x i32]* %27
	store [10 x i32] %32, [10 x i32]* %26
	%33 = load %Person*, %Person** %5
	ret %Person* %33
}

declare ccc void @free(i8* %0)

define ccc void @main() {
main:
	%0 = trunc i32 1 to i1
	%1 = call %Person* @DefaultPerson(i1 %0)
	%2 = alloca %Person*
	store %Person* %1, %Person** %2
	%3 = load %Person*, %Person** %2
	%4 = getelementptr %Person, %Person* %3, i32 0, i32 0
	%5 = load i8*, i8** %4
	%6 = getelementptr [4 x i8], [4 x i8]* @"%s ", i32 0, i32 0
	%7 = call i32 (i8*, ...) @printf(i8* %6, i8* %5)
	%8 = sext i32 300 to i64
	%9 = mul i64 %8, 0
	%10 = call i8* @malloc(i64 %9)
	%11 = bitcast i8* %10 to %Person*
	%12 = alloca %Person*
	store %Person* %11, %Person** %12
	br label %for.declaration.pDqBqTHMOx

leave.ZyKZwzDlUe:
	%13 = load %Person*, %Person** %12
	%14 = bitcast %Person* %13 to i8*
	call void @free(i8* %14)
	ret void

for.declaration.pDqBqTHMOx:
	%15 = alloca i32
	store i32 0, i32* %15
	%16 = load i32, i32* %15
	%17 = icmp slt i32 %16, 300
	br i1 %17, label %for.block.ocfQpxTRcQ, label %leave.ZyKZwzDlUe

for.condition.gmhQqBcINQ:
	%18 = load i32, i32* %15
	%19 = icmp slt i32 %18, 300
	br i1 %19, label %for.block.ocfQpxTRcQ, label %leave.ZyKZwzDlUe

for.block.ocfQpxTRcQ:
	%20 = load i32, i32* %15
	%21 = load %Person*, %Person** %12
	%22 = getelementptr %Person, %Person* %21, i32 %20
	%23 = trunc i32 1 to i1
	%24 = call %Person* @DefaultPerson(i1 %23)
	%25 = load %Person, %Person* %24
	store %Person %25, %Person* %22
	br label %for.update.rmFURuiSWF

for.update.rmFURuiSWF:
	%26 = load i32, i32* %15
	%27 = add i32 %26, 1
	store i32 %27, i32* %15
	br label %for.condition.gmhQqBcINQ
}
