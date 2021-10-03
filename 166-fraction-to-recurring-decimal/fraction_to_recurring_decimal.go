package leetcode166

import "strings"

func fractionToDecimal(numerator int, denominator int) string {
	// 分数不能是无限不循环小数，循环部分记录余数和对应除数是否重复即可，余数和对应除数重复就开始重复了
	// 不能单独看余数或者除数，
	if numerator == 0 {
		return "0"
	}
	// ret = (isMinus)[integer](.[float])

	// 整数部分
	var integer int64 = 0
	// Quotient
	var q int64
	// Remainder
	var r int64 = int64(numerator)
	// 这样转会出问题
	// var r uint32 = uint32(numerator)
	// Denominator
	var d int64 = int64(denominator)

	// 存放小数部分
	var float []byte = make([]byte, 0)
	// 存放余数的位置的map
	var rMap map[Key]int = make(map[Key]int)
	isMinus := false
	if numerator < 0 {
		isMinus = !isMinus
		r = -r
	}
	if denominator < 0 {
		isMinus = !isMinus
		d = -d
	}

	integer = r / d
	r = r % d

	r *= 10
	var repeatIdx int = 10e5
	var ok bool = false
	var p int
	// 处理小数，寻找重复的点
	for i := 0; i < 10e4 && r != 0; i++ {
		q = r / d
		r = r % d
		key := Key{a: q, b: r}
		p, ok = rMap[key]
		if ok {
			repeatIdx = p
			break
		}
		rMap[key] = i
		float = append(float, byte(q))
		r *= 10
	}
	// fmt.Printf("(%v)[%v](.%v)\n",isMinus,integer,float)
	// 开始拼接 (isMinus)[integer](.[float])
	var builder strings.Builder
	if isMinus {
		builder.WriteByte('-')
	}
	if integer == 0 {
		builder.WriteByte('0')
	} else {
		var length int
		// 处理整数的除数
		buf := make([]uint8, 0)
		// 拼接整数部分
		for integer != 0 {
			buf = append(buf, uint8(integer%10))
			integer /= 10
		}
		length = len(buf)
		for length != 0 {
			builder.WriteByte(buf[length-1] + '0')
			buf = buf[:length-1]
			length = len(buf)
		}
	}
	if len(float) != 0 {
		// 写小数点
		builder.WriteByte('.')
	}
	i := 0
	length := len(float)
	// 没有重复部分
	for ; i < repeatIdx && i < length; i++ {
		builder.WriteByte(byte(float[i] + '0'))
	}
	if repeatIdx < 10e5 {
		builder.WriteByte('(')
		for i = repeatIdx; i < length; i++ {
			builder.WriteByte(byte(float[i] + '0'))
		}
		builder.WriteByte(')')

	}

	return builder.String()
}

type Key struct {
	a int64
	b int64
}
