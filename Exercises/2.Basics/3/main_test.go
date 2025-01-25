package main

import "testing"

func Test_stringStat(t *testing.T) {
	tests := []struct {
		name string
		word string
		want string
	}{
		{
			"Старт",
			"Старт",
			`с - 1
т - 2
а - 1
р - 1`,
		},
		{
			"Финиш",
			"ФиНИшшш",
			`ф - 1
и - 2
н - 1
ш - 3`,
		},
		{
			"Булки с чаем",
			"Съешь ещё этих мягких французских булок, да выпей чаю",
			`с - 2
ъ - 1
е - 3
ш - 1
ь - 1
щ - 1
ё - 1
э - 1
т - 1
и - 3
х - 3
м - 1
я - 1
г - 1
к - 3
ф - 1
р - 1
а - 3
н - 1
ц - 1
у - 2
з - 1
б - 1
л - 1
о - 1
, - 1
д - 1
в - 1
ы - 1
п - 1
й - 1
ч - 1
ю - 1`,
		},
		{"FinishEN", "Finishh", `f - 1
i - 2
n - 1
s - 1
h - 2`},
		{
			"Empty", "", "",
		},
		{
			"FinishTR", "Bitiş", `b - 1
i - 2
t - 1
ş - 1`,
		},
		{
			"FinishJA", "終了", `終 - 1
了 - 1`,
		},
		{
			"FinishArabic", "إنهاء", `إ - 1
ن - 1
ه - 1
ا - 1
ء - 1`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringStat(tt.word); got != tt.want {
				t.Errorf("stringStat() = %v, want %v", got, tt.want)
			}
		})
	}
}
