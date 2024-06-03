@0 = global [6 x i8] c"%lld\0A\00"
@1 = global [4 x i8] c"%f\0A\00"
@2 = global [4 x i8] c"%c\0A\00"

declare i32 @puts(i8* %0)

declare i64 @printf(i8* %0, i64 %1)

define void @main() {
0:
	%1 = alloca [4 x i64]
	%2 = alloca i64
	store i64 10, i64* %2
	%3 = alloca i64
	store i64 54, i64* %3
	%4 = alloca i64
	store i64 418, i64* %4
	%5 = getelementptr [4 x i64], [4 x i64]* %1, i64 0, i64 0
	store i64 19, i64* %5
	%6 = getelementptr [4 x i64], [4 x i64]* %1, i64 0, i64 1
	store i64 15, i64* %6
	%7 = add i64 0, 1
	%8 = getelementptr [4 x i64], [4 x i64]* %1, i64 0, i64 %7
	%9 = load i64, i64* %8
	%10 = load i64, i64* %4
	%11 = add i64 %10, %9
	%12 = bitcast [6 x i8]* @0 to i8*
	%13 = call i64 @printf(i8* %12, i64 %11)
	%14 = add i64 0, 1
	%15 = getelementptr [4 x i64], [4 x i64]* %1, i64 0, i64 %14
	%16 = load i64, i64* %15
	%17 = load i64, i64* %4
	%18 = add i64 %17, %16
	br label %19

19:
	ret void

20:
	br label %20
}
