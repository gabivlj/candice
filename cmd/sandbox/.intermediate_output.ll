%Point = type { i64, i64, %Point* }

@"%d " = global [3 x i8] c"%d "

define i32 @main() {
_main:
	%0 = alloca %Point
	%1 = alloca %Point
	%2 = getelementptr %Point, %Point* %1, i32 0, i32 0
	store i64 43, i64* %2
	%3 = getelementptr %Point, %Point* %1, i32 0, i32 1
	store i64 55, i64* %3
	%4 = mul i64 33, 0
	%5 = call i8* @malloc(i64 %4)
	%6 = bitcast i8* %5 to %Point*
	%7 = alloca %Point*
	store %Point* %6, %Point** %7
	%8 = getelementptr %Point, %Point* %1, i32 0, i32 2
	%9 = load %Point*, %Point** %7
	store %Point* %9, %Point** %8
	%10 = load %Point, %Point* %1
	store %Point %10, %Point* %0
	%11 = getelementptr %Point, %Point* %0, i32 0, i32 0
	store i64 3, i64* %11
	%12 = getelementptr %Point, %Point* %0, i32 0, i32 1
	store i64 3, i64* %12
	%13 = getelementptr %Point, %Point* %0, i32 0, i32 1
	%14 = getelementptr %Point, %Point* %0, i32 0, i32 1
	%15 = load i64, i64* %14
	%16 = mul i64 %15, -1
	store i64 %16, i64* %13
	%17 = alloca i64
	%18 = getelementptr %Point, %Point* %0, i32 0, i32 0
	%19 = load i64, i64* %18
	store i64 %19, i64* %17
	%20 = alloca i64*
	%21 = getelementptr %Point, %Point* %0, i32 0, i32 1
	%22 = alloca i64*
	store i64* %21, i64** %22
	%23 = load i64*, i64** %22
	store i64* %23, i64** %20
	%24 = alloca i64
	%25 = load i64*, i64** %20
	%26 = load i64, i64* %25
	store i64 %26, i64* %24
	%27 = alloca %Point
	%28 = getelementptr %Point, %Point* %0, i32 0, i32 2
	%29 = load %Point*, %Point** %28
	%30 = getelementptr %Point, %Point* %29, i64 0
	%31 = load %Point, %Point* %30
	store %Point %31, %Point* %27
	%32 = getelementptr %Point, %Point* %0, i32 0, i32 2
	%33 = load %Point*, %Point** %32
	%34 = getelementptr %Point, %Point* %33, i64 0
	%35 = getelementptr %Point, %Point* %34, i32 0, i32 0
	%36 = mul i64 3, -1
	store i64 %36, i64* %35
	%37 = alloca %Point*
	%38 = mul i64 3, 0
	%39 = call i8* @malloc(i64 %38)
	%40 = bitcast i8* %39 to %Point*
	%41 = alloca %Point*
	store %Point* %40, %Point** %41
	%42 = load %Point*, %Point** %41
	store %Point* %42, %Point** %37
	%43 = alloca %Point
	%44 = load %Point*, %Point** %37
	%45 = getelementptr %Point, %Point* %44, i64 0
	%46 = load %Point, %Point* %45
	store %Point %46, %Point* %43
	%47 = load %Point*, %Point** %37
	%48 = getelementptr %Point, %Point* %47, i64 0
	%49 = getelementptr %Point, %Point* %48, i32 0, i32 0
	%50 = mul i64 3, -1
	store i64 %50, i64* %49
	%51 = load %Point*, %Point** %37
	%52 = getelementptr %Point, %Point* %51, i64 0
	%53 = getelementptr %Point, %Point* %52, i32 0, i32 0
	%54 = load i64, i64* %53
	%55 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%56 = call i32 (i8*, ...) @printf(i8* %55, i64 %54)
	%57 = getelementptr %Point, %Point* %43, i32 0, i32 0
	%58 = load i64, i64* %57
	%59 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%60 = call i32 (i8*, ...) @printf(i8* %59, i64 %58)
	%61 = alloca %Point*****
	%62 = load %Point*, %Point** %37
	%63 = getelementptr %Point, %Point* %62, i64 0
	%64 = alloca %Point*
	store %Point* %63, %Point** %64
	%65 = alloca %Point**
	store %Point** %64, %Point*** %65
	%66 = alloca %Point***
	store %Point*** %65, %Point**** %66
	%67 = alloca %Point****
	store %Point**** %66, %Point***** %67
	%68 = alloca %Point*****
	store %Point***** %67, %Point****** %68
	%69 = load %Point*****, %Point****** %68
	store %Point***** %69, %Point****** %61
	%70 = alloca %Point*
	%71 = load %Point*****, %Point****** %61
	%72 = load %Point****, %Point***** %71
	%73 = load %Point***, %Point**** %72
	%74 = load %Point**, %Point*** %73
	%75 = load %Point*, %Point** %74
	store %Point* %75, %Point** %70
	%76 = alloca i64
	%77 = load %Point*, %Point** %70
	%78 = getelementptr %Point, %Point* %77, i32 0, i32 0
	%79 = load i64, i64* %78
	store i64 %79, i64* %76
	%80 = load i64, i64* %76
	%81 = getelementptr [3 x i8], [3 x i8]* @"%d ", i32 0, i32 0
	%82 = call i32 (i8*, ...) @printf(i8* %81, i64 %80)
	ret i32 0
}

declare i32 @printf(i8* %0, ...)

declare i8* @malloc(i64 %0)
