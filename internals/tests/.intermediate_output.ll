%Point = type { i64, i64, %Point* }

@"%d " = global [3 x i8] c"%d "

define i32 @main() {
_main:
	%0 = alloca %Point
	%1 = alloca %Point
	%2 = sext i32 43 to i64
	%3 = getelementptr %Point, %Point* %1, i32 0, i32 0
	store i64 %2, i64* %3
	%4 = sext i32 55 to i64
	%5 = getelementptr %Point, %Point* %1, i32 0, i32 1
	store i64 %4, i64* %5
	%6 = sext i32 33 to i64
	%7 = mul i64 %6, 0
	%8 = call i8* @malloc(i64 %7)
	%9 = bitcast i8* %8 to %Point*
	%10 = alloca %Point*
	store %Point* %9, %Point** %10
	%11 = getelementptr %Point, %Point* %1, i32 0, i32 2
	%12 = load %Point*, %Point** %10
	store %Point* %12, %Point** %11
	%13 = load %Point, %Point* %1
	store %Point %13, %Point* %0
	%14 = getelementptr %Point, %Point* %0, i32 0, i32 0
	%15 = sext i32 3 to i64
	store i64 %15, i64* %14
	%16 = getelementptr %Point, %Point* %0, i32 0, i32 1
	%17 = sext i32 3 to i64
	store i64 %17, i64* %16
	%18 = getelementptr %Point, %Point* %0, i32 0, i32 1
	%19 = getelementptr %Point, %Point* %0, i32 0, i32 1
	%20 = load i64, i64* %19
	%21 = mul i64 %20, -1
	store i64 %21, i64* %18
	%22 = alloca i64
	%23 = getelementptr %Point, %Point* %0, i32 0, i32 0
	%24 = load i64, i64* %23
	store i64 %24, i64* %22
	%25 = alloca i64*
	%26 = getelementptr %Point, %Point* %0, i32 0, i32 1
	%27 = alloca i64*
	store i64* %26, i64** %27
	%28 = load i64*, i64** %27
	store i64* %28, i64** %25
	%29 = alloca i64
	%30 = load i64*, i64** %25
	%31 = load i64, i64* %30
	store i64 %31, i64* %29
	%32 = alloca %Point
	%33 = getelementptr %Point, %Point* %0, i32 0, i32 2
	%34 = load %Point*, %Point** %33
	%35 = getelementptr %Point, %Point* %34, i32 0
	%36 = load %Point, %Point* %35
	store %Point %36, %Point* %32
	%37 = getelementptr %Point, %Point* %0, i32 0, i32 2
	%38 = load %Point*, %Point** %37
	%39 = getelementptr %Point, %Point* %38, i32 0
	%40 = getelementptr %Point, %Point* %39, i32 0, i32 0
	%41 = mul i32 3, -1
	%42 = sext i32 %41 to i64
	store i64 %42, i64* %40
	%43 = alloca %Point*
	%44 = sext i32 3 to i64
	%45 = mul i64 %44, 0
	%46 = call i8* @malloc(i64 %45)
	%47 = bitcast i8* %46 to %Point*
	%48 = alloca %Point*
	store %Point* %47, %Point** %48
	%49 = load %Point*, %Point** %48
	store %Point* %49, %Point** %43
	%50 = alloca %Point
	%51 = load %Point*, %Point** %43
	%52 = getelementptr %Point, %Point* %51, i32 0
	%53 = load %Point, %Point* %52
	store %Point %53, %Point* %50
	%54 = load %Point*, %Point** %43
	%55 = getelementptr %Point, %Point* %54, i32 0
	%56 = getelementptr %Point, %Point* %55, i32 0, i32 0
	%57 = mul i32 3, -1
	%58 = sext i32 %57 to i64
	store i64 %58, i64* %56
	%59 = load %Point*, %Point** %43
	%60 = getelementptr %Point, %Point* %59, i32 0
	%61 = getelementptr %Point, %Point* %60, i32 0, i32 0
	%62 = load i64, i64* %61
	%63 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%64 = call i32 (i8*, ...) @printf(i8* %63, i64 %62)
	%65 = getelementptr %Point, %Point* %50, i32 0, i32 0
	%66 = load i64, i64* %65
	%67 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%68 = call i32 (i8*, ...) @printf(i8* %67, i64 %66)
	%69 = alloca %Point*****
	%70 = load %Point*, %Point** %43
	%71 = getelementptr %Point, %Point* %70, i32 0
	%72 = alloca %Point*
	store %Point* %71, %Point** %72
	%73 = alloca %Point**
	store %Point** %72, %Point*** %73
	%74 = alloca %Point***
	store %Point*** %73, %Point**** %74
	%75 = alloca %Point****
	store %Point**** %74, %Point***** %75
	%76 = alloca %Point*****
	store %Point***** %75, %Point****** %76
	%77 = load %Point*****, %Point****** %76
	store %Point***** %77, %Point****** %69
	%78 = alloca %Point*
	%79 = load %Point*****, %Point****** %69
	%80 = load %Point****, %Point***** %79
	%81 = load %Point***, %Point**** %80
	%82 = load %Point**, %Point*** %81
	%83 = load %Point*, %Point** %82
	store %Point* %83, %Point** %78
	%84 = alloca i64
	%85 = load %Point*, %Point** %78
	%86 = getelementptr %Point, %Point* %85, i32 0, i32 0
	%87 = load i64, i64* %86
	store i64 %87, i64* %84
	%88 = load i64, i64* %84
	%89 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%90 = call i32 (i8*, ...) @printf(i8* %89, i64 %88)
	ret i32 0
}

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)
