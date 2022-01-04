%Point = type { i64, i64, %Point* }

@"%d " = global [3 x i8] c"%d "

define i32 @main() {
_main:
	%0 = alloca %Point
	%1 = sext i32 43 to i64
	%2 = getelementptr %Point, %Point* %0, i32 0, i32 0
	store i64 %1, i64* %2
	%3 = sext i32 55 to i64
	%4 = getelementptr %Point, %Point* %0, i32 0, i32 1
	store i64 %3, i64* %4
	%5 = sext i32 33 to i64
	%6 = mul i64 %5, 0
	%7 = call i8* @malloc(i64 %6)
	%8 = bitcast i8* %7 to %Point*
	%9 = alloca %Point*
	store %Point* %8, %Point** %9
	%10 = getelementptr %Point, %Point* %0, i32 0, i32 2
	%11 = load %Point*, %Point** %9
	store %Point* %11, %Point** %10
	%12 = getelementptr %Point, %Point* %0, i32 0, i32 0
	%13 = sext i32 3 to i64
	store i64 %13, i64* %12
	%14 = getelementptr %Point, %Point* %0, i32 0, i32 1
	%15 = sext i32 3 to i64
	store i64 %15, i64* %14
	%16 = getelementptr %Point, %Point* %0, i32 0, i32 1
	%17 = getelementptr %Point, %Point* %0, i32 0, i32 1
	%18 = load i64, i64* %17
	%19 = mul i64 %18, -1
	store i64 %19, i64* %16
	%20 = getelementptr %Point, %Point* %0, i32 0, i32 0
	%21 = alloca i64
	%22 = load i64, i64* %20
	store i64 %22, i64* %21
	%23 = getelementptr %Point, %Point* %0, i32 0, i32 1
	%24 = alloca i64*
	store i64* %23, i64** %24
	%25 = load i64*, i64** %24
	%26 = alloca i64
	%27 = load i64, i64* %25
	store i64 %27, i64* %26
	%28 = getelementptr %Point, %Point* %0, i32 0, i32 2
	%29 = load %Point*, %Point** %28
	%30 = getelementptr %Point, %Point* %29, i32 0
	%31 = alloca %Point
	%32 = load %Point, %Point* %30
	store %Point %32, %Point* %31
	%33 = getelementptr %Point, %Point* %0, i32 0, i32 2
	%34 = load %Point*, %Point** %33
	%35 = getelementptr %Point, %Point* %34, i32 0
	%36 = getelementptr %Point, %Point* %35, i32 0, i32 0
	%37 = mul i32 3, -1
	%38 = sext i32 %37 to i64
	store i64 %38, i64* %36
	%39 = sext i32 3 to i64
	%40 = mul i64 %39, 0
	%41 = call i8* @malloc(i64 %40)
	%42 = bitcast i8* %41 to %Point*
	%43 = alloca %Point*
	store %Point* %42, %Point** %43
	%44 = load %Point*, %Point** %43
	%45 = getelementptr %Point, %Point* %44, i32 0
	%46 = alloca %Point
	%47 = load %Point, %Point* %45
	store %Point %47, %Point* %46
	%48 = load %Point*, %Point** %43
	%49 = getelementptr %Point, %Point* %48, i32 0
	%50 = getelementptr %Point, %Point* %49, i32 0, i32 0
	%51 = mul i32 3, -1
	%52 = sext i32 %51 to i64
	store i64 %52, i64* %50
	%53 = load %Point*, %Point** %43
	%54 = getelementptr %Point, %Point* %53, i32 0
	%55 = getelementptr %Point, %Point* %54, i32 0, i32 0
	%56 = load i64, i64* %55
	%57 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%58 = call i32 (i8*, ...) @printf(i8* %57, i64 %56)
	%59 = getelementptr %Point, %Point* %46, i32 0, i32 0
	%60 = load i64, i64* %59
	%61 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%62 = call i32 (i8*, ...) @printf(i8* %61, i64 %60)
	%63 = load %Point*, %Point** %43
	%64 = getelementptr %Point, %Point* %63, i32 0
	%65 = alloca %Point*
	store %Point* %64, %Point** %65
	%66 = alloca %Point**
	store %Point** %65, %Point*** %66
	%67 = alloca %Point***
	store %Point*** %66, %Point**** %67
	%68 = alloca %Point****
	store %Point**** %67, %Point***** %68
	%69 = alloca %Point*****
	store %Point***** %68, %Point****** %69
	%70 = load %Point*****, %Point****** %69
	%71 = load %Point****, %Point***** %70
	%72 = load %Point***, %Point**** %71
	%73 = load %Point**, %Point*** %72
	%74 = alloca %Point*
	%75 = load %Point*, %Point** %73
	store %Point* %75, %Point** %74
	%76 = load %Point*, %Point** %74
	%77 = getelementptr %Point, %Point* %76, i32 0, i32 0
	%78 = alloca i64
	%79 = load i64, i64* %77
	store i64 %79, i64* %78
	%80 = load i64, i64* %78
	%81 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%82 = call i32 (i8*, ...) @printf(i8* %81, i64 %80)
	%83 = sext i32 1 to i64
	%84 = mul i64 %83, 0
	%85 = call i8* @malloc(i64 %84)
	%86 = bitcast i8* %85 to %Point*
	%87 = alloca %Point*
	store %Point* %86, %Point** %87
	%88 = load %Point*, %Point** %87
	%89 = alloca %Point
	%90 = sext i32 43 to i64
	%91 = getelementptr %Point, %Point* %89, i32 0, i32 0
	store i64 %90, i64* %91
	%92 = sext i32 55 to i64
	%93 = getelementptr %Point, %Point* %89, i32 0, i32 1
	store i64 %92, i64* %93
	%94 = sext i32 33 to i64
	%95 = mul i64 %94, 0
	%96 = call i8* @malloc(i64 %95)
	%97 = bitcast i8* %96 to %Point*
	%98 = alloca %Point*
	store %Point* %97, %Point** %98
	%99 = getelementptr %Point, %Point* %89, i32 0, i32 2
	%100 = load %Point*, %Point** %98
	store %Point* %100, %Point** %99
	%101 = load %Point, %Point* %89
	store %Point %101, %Point* %88
	%102 = load %Point*, %Point** %87
	%103 = getelementptr %Point, %Point* %102, i32 0, i32 0
	%104 = load i64, i64* %103
	%105 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%106 = call i32 (i8*, ...) @printf(i8* %105, i64 %104)
	ret i32 0
}

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)
