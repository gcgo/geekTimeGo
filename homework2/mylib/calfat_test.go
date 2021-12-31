package gobmi

import "testing"

func TestBMI(t *testing.T) {
	inH, inW := 1.0, 1.0
	expectedOut := 1.0
	actualOut, err := BMI(inW, inH)
	if err != nil {
		t.Fatalf("something wrong..:%v", err)
	}
	if expectedOut != actualOut {
		t.Fatalf("expect:%f,but:%f", expectedOut, actualOut)
	}
}

func TestTizhi(t *testing.T) {
	bmi, age, sex := 20.0, 30.0, 0.0
	expectedOut := 25.5
	actualOut := Tizhi(bmi, age, sex)

	if expectedOut != actualOut {
		t.Fatalf("expect:%f,but:%f", expectedOut, actualOut)
	}
}

func TestGiveAdvice(t *testing.T) {
	tizhi, age, sex := 20.0, 30.0, 0.0
	expectedOut1 := "偏瘦"
	expectedOut2 := "多吃"
	actualOut1, actualOut2, err := GiveAdvice(age, tizhi, sex)
	if err != nil {
		t.Fatalf("something wrong..:%v", err)
	}
	if expectedOut1 != actualOut1 {
		t.Fatalf("expect:%s,but:%s", expectedOut1, actualOut1)
	}
	if expectedOut2 != actualOut2 {
		t.Fatalf("expect:%s,but:%s", expectedOut2, actualOut2)
	}
}
