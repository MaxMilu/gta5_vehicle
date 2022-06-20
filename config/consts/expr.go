package consts

const (
	Expr1 = `\[\d+\]`       // [1234]
	Expr2 = `\[.+\]`        // [....]
	Expr3 = `[\[\]]`        // [ 或者 ]
	Expr4 = `(\\+r)|(\\+n)` // \r（1个或多个\） 或者 \n（1个或多个\）
)