package one

import "testing"

func TestAdd(t *testing.T) {
	var a, b = 10, 20
	sum := Add(a, b)
	if sum != 30 {
		t.Fatalf("测试失败，预期得到%d，实际得到%d\n", 30, sum)
	}
	t.Log("测试成功")

	//输出日志，使用fmt.Sprintln的返回值作为日志
	//t.Log()
	//使用fmt.Sprintf的返回值作为日志
	//t.Logf()

	//宣布测试失败，但继续执行
	//t.Fail()
	//宣布测试失败，且立刻终止
	//t.FailNow()

	//t.Log()+t.Fail()
	//t.Error()
	//to.Logf()+t.Fail()
	//t.Errorf()

	//t.Log()+t.FailNow()
	//t.Fatal()
	//t.Logf()+t.FailNow()
	//t.Fatalf()
}

func TestAddMany(t *testing.T) {
	rets := make([][3]int, 0)
	rets = append(rets, [3]int{5, 3, 8}, [3]int{3, 4, 7}, [3]int{100, 200, 300})
	for _, v := range rets {
		sum := Add(v[0], v[1])
		if sum != v[2] {
			t.Fatalf("测试失败，预期得到%d，实际得到%d\n", v[2], sum)
		}
	}
	t.Log("测试成功")
}

func TestGetSum(t *testing.T) {
	n := 10
	sum := GetSum(n)
	if sum != 55 {
		t.Fatalf("测试失败，预期得到%d，实际得到%d\n", 55, sum)
	}
	t.Log("测试成功")
}

func TestGetSumRecursively(t *testing.T) {
	n := 10
	sum := GetSum(n)
	if sum != 55 {
		t.Fatalf("测试失败，预期得到%d，实际得到%d\n", 55, sum)
	}
	t.Log("测试成功")
}