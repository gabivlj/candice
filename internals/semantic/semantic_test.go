package semantic

import (
	"testing"

	"github.com/gabivlj/candice/internals/lexer"
	"github.com/gabivlj/candice/internals/parser"
	"github.com/gabivlj/candice/pkg/a"
)

func TestSemantic_Analyze(t *testing.T) {
	tests := []struct {
		program    string
		shouldBeOk bool
	}{
		{
			`variable : i64 = @cast(i64, 32); variable : i32 = 32`,
			true,
		},
		{
			`variable : i32 = @cast(i64, 32);`,
			false,
		},
		{
			`variable : i64 = 32`,
			false,
		},
		{
			`variable : i64 = @cast(i64, 32 + 32 * 32 * 32 * 44 / 33 / 44 / 444 - 333 + 333 * 333 + 44)`,
			true,
		},
		{
			`variable : bool = !!!*&44
					 variable2 := variable as i32 + variable as i32`,
			true,
		},
		{
			`variable : i32 = *!!!*&44`,
			false,
		},
		{
			`struct Point { x i32 y i32 point *Point }`,
			true,
		},
		{
			`struct Point { x i32 y i32 point Point }`,
			false,
		},
		{
			`func aFunction(i i32, j i32) i32 {
	if 1 {
		return 3
	}
	return;
}`,
			false,
		},
		{
			`func aFunction(i i32, j i32) i32 {
	if 1 {
		return i + j;
	} else { return 3; }
	
}`,
			true,
		},
		{
			`func aFunction(i i32, j i32) i32 {
	if 1 {
		return i + j;
	}
}`,
			false,
		},
		{
			`func aFunction(i i32, j i32) {
	if 1 {
		return;
	}
}`,
			true,
		},
		{
			`func aFunction(i i32, j i32) {}`,
			true,
		},
		{
			`for 1 < 100 { }`,
			true,
		},
		{
			`for i := 100; i < 100; { }`,
			true,
		},
		{
			`for i := 100; i < 100; i = i + 1 { }`,
			true,
		},
		{
			`func callMe() i32 { return callMe() }
					called := callMe() + 3 + 10 * 100 / 300;
					func callback(c func() i32) { c(); } callback(callMe)`,
			true,
		},
		{
			`func callMe() i32 { return callMe() }
					called := callMe() + 3 + 10 * 100 / 300;
					func callback(c func() i64) { c(); } callback(callMe)`,
			false,
		},
		{
			`arr := [4]i32{1, 2, 3, 4}`,
			true,
		},
		{
			`arr := [3]i32{1, 2, 3, 4}`,
			false,
		},
		{
			`struct RecursivePoint { x i32 y i32 point *RecursivePoint } point := @RecursivePoint { x: 3, y: 4, point: @alloc(RecursivePoint, 3) }`,
			true,
		},
		{
			`point := @RecursivePoint { x: 3, y: 4, point: @alloc(RecursivePoint, 3) }`,
			false,
		},
		{
			`struct RecursivePoint { x i32 y i32 point *RecursivePoint } struct Point { x i32 y i32 point *RecursivePoint }point := @RecursivePoint { x: 3, y: 4, point: @alloc(Point, 3) }`,
			false,
		},
		{
			`thing := @alloc(RecursivePoint, 3)`,
			false,
		},
		{
			`string :*i8 = "hello world"`,
			true,
		},
		{
			`arr :[4]i32 = [4]i32{1, 2, 3, 4}
					for i := 0; i < 4; i = i + 1 { arr[i] = 1 }
					if arr[0] == 0 {
						arr[1] = 3
					}`,
			true,
		},
		{
			`extern func print(*void);
					print(@cast(*void, @alloc(i32, 1)))`,
			true,
		},
		{
			`struct Point { p Point }`,
			false,
		},
		{
			`switch 1 { case 1 {} case 2 {} default {} }`,
			true,
		},
		{
			`switch 1 as i64 { case 1 {} case 2 {} default {} }`,
			false,
		},
		{
			`func returnsI32() i32 {
				switch 1 {
					case 1 {
						return 1;
					}

					case 2 {
						return 2;
					}

					default {
						return 0;
					}
				}		
			}`,
			true,
		},
		{
			`func returnsI32() i32 {
				switch 1 {
					case 1 {
						return 1;
					}

					case 2 {
						return 2;
					}

					default {
						// return 0;
					}
				}		
			}`,
			false,
		},
		{
			`func returnsI32() i32 {
				switch 1 {
					case 1 {
						return 1;
					}

					case 2 {
						return 2;
					}

					default {
						// return 0;
					}
				}
				return 0;		
			}`,
			true,
		},
		{
			`func returnsI32() i32 {
				switch 1 {
					case 1 {
						return 1;
					}

					case 2 {
						// return 2;
					}

					default {
						return 0;
					}
				}				
			}`,
			false,
		},
		{
			`func returnsI32() i32 {
				switch 1 {
					case 1 {
						// return 1;
					}

					case 2 {
						return 2;
					}

					default {
						return 0;
					}
				}				
			}`,
			false,
		},
		// This still doesn't work...
		// {
		// 	`struct C { p Point } struct Point { p C }`,
		// 	false,
		// },
	}

	for _, test := range tests {
		semantic := New()
		p := parser.New(lexer.New(test.program))
		program := p.Parse()
		a.Assert(len(p.Errors) == 0, p.Errors)
		semantic.Analyze(program)
		if test.shouldBeOk && len(semantic.Errors) != 0 {
			t.Fatal(test, semantic.Errors)
		} else if !test.shouldBeOk && len(semantic.Errors) == 0 {
			t.Fatal(test, "shouldn't be ok but we got 0 Errors...")
		}
	}
}
