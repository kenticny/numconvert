package numconvert

import "testing"

type Operator struct {
	Number int64
	Radix  int64
	Expect string
}

func TestConvert(t *testing.T) {
	cases := []Operator{
		{
			Number: 987654321,
			Radix:  10,
			Expect: "987654321",
		}, {
			Number: 987654321,
			Radix:  12,
			Expect: "23691b269",
		}, {
			Number: 987654321,
			Radix:  16,
			Expect: "3ade68b1",
		}, {
			Number: 987654321,
			Radix:  25,
			Expect: "4139lml",
		}, {
			Number: 987654321,
			Radix:  36,
			Expect: "gc0uy9",
		}, {
			Number: 9876543219,
			Radix:  62,
			Expect: "aMoY4b",
		}, {
			Number: 765432198765,
			Radix:  62,
			Expect: "dtvdv3f",
		},
	}
	for _, c := range cases {
		result, err := Convert(c.Number, c.Radix)
		if err != nil {
			t.Errorf("%s -> %s failed: %s", c.Number, c.Radix, err)
		}
		if result != c.Expect {
			t.Errorf("expect is %s, actual is %s", c.Expect, result)
		}
	}
}
