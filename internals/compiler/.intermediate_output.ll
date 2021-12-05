define i32 @main() {
_main:
	%0 = add i64 3, 3
	%1 = sub i64 332, 1
	%2 = sdiv i64 %1, 3
	%3 = add i64 %0, %2
	%4 = mul i64 %3, 5
	ret i32 0
}
