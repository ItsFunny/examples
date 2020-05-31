package utils

// 用于动态拼接or|and 的sql条件
// 如 假设有啊 a,b,c 的复合or 且like 条件查询,则最终会返回的是( A like a or b like or C like c )
// 若只是a 则返回的是 ( A like  a )
func CombineOrConditionSQL(orAnd string, judgeStr string, conditions []bool, colums []string, values ...interface{}) string {
	res := " ( "
	need := false
	l := len(conditions)
	for i := l - 1; i >= 0; i++ {
		if !conditions[i] {
			continue
		}

		// a | a,b | a,b,c
		if need {
			res += orAnd
		} else if i-1 >= 0 {
			need = true
		}

		s := values[i].(*interface{})
		strInterface := *s
		switch strInterface.(type) {
		case string:
			t := strInterface.(string)
			t += " % "
			*s = t
			res += colums[i] + " " + judgeStr + " ? "
		case int:
			// t := strconv.Itoa(strInterface.(int))
			t := strInterface.(int)
			t++
			*s = t
			res += colums[i] + " " + " = " + " ? "
		}
	}
	res += ")"
	return res
}

// 缺陷,如果通过values返回的话原先的nil值也会被传递进来
// values
func CombineOrConditionSQL2(orAnd string, judgeStr string, conditions []bool, colums []string, values *[]interface{}) (string, []interface{}) {
	res := " ( "
	need := false
	l := len(conditions)
	tValues := *values
	params := make([]interface{}, 0)
	for i := l - 1; i >= 0; i-- {
		if !conditions[i] {
			tValues = append(tValues[:i], tValues[i+1:])
			continue
		}

		// a | a,b | a,b,c
		if need {
			res += orAnd
		} else if i-1 >= 0 {
			need = true
		}

		s := tValues[i].(*interface{})
		strInterface := *s
		switch strInterface.(type) {
		case string:
			t := strInterface.(string)
			t += " % "
			*s = t
			params = append(params, t)
			res += colums[i] + " " + judgeStr + " ? "
		case int:
			// t := strconv.Itoa(strInterface.(int))
			t := strInterface.(int)
			t++
			*s = t
			params = append(params, t)
			res += colums[i] + " " + " = " + " ? "
		}
	}
	res += ")"
	return res, params
}

