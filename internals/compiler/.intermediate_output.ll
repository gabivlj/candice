define i32 @main() {
_main:
	%0 = add i64 3, 3
	%1 = mul i64 3, 3
	%2 = add i64 %0, %1
	%3 = mul i64 %2, 5
	ret i32 0
}
