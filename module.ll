@0 = global [6 x i8] c"%lld\0A\00"
@1 = global [4 x i8] c"%f\0A\00"
@2 = global [4 x i8] c"%c\0A\00"

declare i32 @puts(i8* %0)

declare i64 @printf(i8* %0, i64 %1)

define void @main() {
0:
	%1 = alloca [5 x i64]
	%2 = alloca i64
	store i64 10, i64* %2
	%3 = getelementptr [5 x i64], [5 x i64]* %1, i64 0, i64 0
	store i64 14, i64* %3
	%4 = getelementptr [5 x i64], [5 x i64]* %1, i64 0, i64 0
	%5 = load i64, i64* %4
	%6 = add i64 %5, 2
	store i64 %6, i64* %2
	%7 = load i64, i64* %2
	%8 = bitcast [6 x i8]* @0 to i8*
	%9 = call i64 @printf(i8* %8, i64 %7)
	%10 = load i64, i64* %2
	%11 = bitcast [6 x i8]* @0 to i8*
	%12 = call i64 @printf(i8* %11, i64 %10)
	%13 = load i64, i64* %2
	%14 = bitcast [6 x i8]* @0 to i8*
	%15 = call i64 @printf(i8* %14, i64 %13)
	ret void

16:
	ret void

17:
	br label %17
}
