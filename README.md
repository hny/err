## QuckStart:

```golang

package service

import (
	"strconv"

	Err "github.com/hny/err"
)


func TestErrUse(t *testing.T) {
	t.Log(
		`Use("9527"): `, use("9527"),
	)
	t.Log(
		`Use(""): `, use(""),
	)
	t.Log(
		`Use("fce180"): `, use("fce180"),
	)
}
func use(userId string) Err.Err {
	if userId == "" {
		return Err.New().Err(
			"userId is empty",
		).Msg(
			Err.PARAMS_INVALIDED_ZN,
		).Code(
			Err.PARAMS_INVALIDED_CODE,
		)
	}

	err := marashal(userId)
	if err != nil {
		return err.Failed(
			"marashal",
		).Info(
			"userId", userId,
		).Code(
			Err.PARAMS_INVALIDED_CODE,
		).Msg(
			Err.PARAMS_INVALIDED_ZN,
		)
	}

	return nil
}
func marashal(str string) Err.Err {
	_, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return Err.New().Err(
			"strconv.ParseUint", err,
		)
		// or
		// return Err.New(
		// 	"strconv.ParseUint", err,
		// )
	}

	return nil
}

```

### 输出:
``` bash
// $ go test -timeout 30s -run ^Test$ ……/service -v
// === RUN   Test
//     ……/help_test.go:11: Use("9527"):  <nil>
//     ……/help_test.go:14: Use(""):  {"err": 'userId is empty;',"code": 10000,"msg": "参数有误，请检查"}
//     ……/help_test.go:17: Use("fce180"):  {"err": 'marashal failed; strconv.ParseUint err: strconv.ParseUint: parsing "fce180": invalid syntax; userId: [fce180];',"code": 10000,"msg": "参数有误，请检查"}
// --- PASS: Test (0.00s)
// PASS
// ok      ……/service      0.202s
```