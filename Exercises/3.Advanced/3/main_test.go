package main

import (
	"fmt"
	"testing"
)

func Test_createList(t *testing.T) {
	tests := []struct {
		name    string
		sl      []any
		wantLen int
	}{
		{
			name:    "тест 1-2-3",
			sl:      []any{1, 2, 3},
			wantLen: 3,
		},
		{
			"тест 0",
			[]any{0},
			1,
		},
		{
			"тест nil",
			nil,
			0,
		},
		{
			"тест 27 чисел",
			[]any{-10, -8, -8, -5, -5, 7, 7, 8, 8, 8, 9, 9, 9, 10, 10, 11, 12, 15, 20, 29, 39, 49, 533, 555, 577, 599, 800},
			27,
		},
		{
			"тест 10 строк",
			[]any{"hex", "egn", "abc", "b123", "b3223", "c1", "c2", "drill", "ygr;", "xyz"},
			10,
		},
		{
			"тест 5 чисел с плавающей точкой",
			[]any{1.1, 2.2, 2.2, 3.3, 4.4},
			5,
		},
		{
			"тест с 7 рунами",
			[]any{'a', 'b', 'b', 'c', 'd', 'h', 'f'},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createList(tt.sl)
			if length(got) != tt.wantLen {
				t.Errorf("createList length = %d, want %d", length(got), tt.wantLen)
				return
			}

			if eq, reason := compareListWithSlice(got, tt.sl); !eq {
				t.Errorf("createList made incorrect list, reason %s", reason)
			}
		})
	}
}

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		name      string
		nodes     []any
		val       any
		wantCount int
	}{
		{"тест 1, 2, 2, 2, 3 тогда цифр 1 в кол-ве 1", []any{1, 2, 2, 2, 3, 3}, 1, 1},
		{"тест 1, 2, 2, 2, 3 тогда цифр 2 в кол-ве 1", []any{1, 2, 2, 2, 3, 3}, 2, 1},
		{"тест -10, -10, -2, -2, 2, 2, 3, 5, 6, 6, 6, 6, 7, 7 тогда цифр 2 в кол-ве 1", []any{-10, -10, -2, -2, 2, 2, 3, 5, 6, 6, 6, 6, 7, 7}, 2, 1},
		{"тест -10, -10, -2, -2, 2, 2, 3, 5, 6, 6, 6, 6, 7, 7 тогда цифр 7 в кол-ве 1", []any{-10, -10, -2, -2, 2, 2, 3, 5, 6, 6, 6, 6, 7, 7}, 7, 1},
		{"тест -10, -10, -2, -2, 2, 2, 3, 5, 6, 6, 6, 6, 7, 7 тогда цифр -10 в кол-ве 1", []any{-10, -10, -2, -2, 2, 2, 3, 5, 6, 6, 6, 6, 7, 7}, -10, 1},
		{"тест -10, -10, -2, -2, 2, 2, 3, 5, 6, 6, 6, 6, 7, 7 тогда цифр 0 в кол-ве 0", []any{-10, -10, -2, -2, 2, 2, 3, 5, 6, 6, 6, 6, 7, 7}, 0, 0},
		{"тест 0, 0, 0, 0, 0 тогда цифр 0 в кол-ве 1", []any{0, 0, 0, 0, 0}, 0, 1},
		{"тест 5, 5, 5, 5 тогда цифр 5 в кол-ве 1", []any{5, 5, 5, 5}, 5, 1},
		{"тест 1, 1, 1, 1, 1, 1, 1, 2, 3, 4, 4, 5, 5, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 9, 9 тогда цифр 5 в кол-ве 1", []any{1, 1, 1, 1, 1, 1, 1, 2, 3, 4, 4, 5, 5, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 9, 9}, 5, 1},
		{"nil", nil, 5, 0},
		{"тест 'a', 'b', 's', 'd', 'd', 'd', 'e,'g', 'o'", []any{'a', 'b', 's', 'd', 'd', 'd', 'e', 'g', 'o'}, 'd', 1},
		{`тест , "btmlt", "sewml", "dfefe", "drh", "dmb", "elgr", "gre", "orgr"`, []any{'a', 'b', 's', 'd', 'd', 'd', 'e', 'g', 'o'}, 'd', 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := createList(tt.nodes)
			got := deleteDuplicates(list)
			count := countVal(got, tt.val)
			if count != tt.wantCount {
				t.Errorf("после удаления дубликатов подсчет количества числа %d показал %d, а требуется %d", tt.val, count, tt.wantCount)
			}
		})
	}
}

func Test_insertAt(t *testing.T) {
	tests := []struct {
		name  string
		nodes []any
		index int
		val   any
	}{
		{
			"Вставка элемента 3 по индексу 2",
			[]any{1, 2, 3, 4, 5},
			2,
			3,
		},
		{
			"Вставка элемента 'a' по индексу 4",
			[]any{'a', 'b'},
			4,
			'a',
		}, {
			`Вставка элемента "abc" по индексу 0`,
			[]any{},
			0,
			"abc",
		}, {
			"Вставка элемента -1.1 по индексу -1",
			[]any{1.1, 2.2, 3.3},
			-1,
			-1.1,
		},
	}

	for _, tt := range tests {
		list := createList(tt.nodes)
		initLength := length(list)
		got := insertAt(list, tt.index, tt.val)
		gotLength := length(got)
		found := getElement(got, tt.index)
		if initLength != gotLength && found != tt.val {
			t.Errorf("После вставки элемета, элемент, находящийся по индексу: %d не соответствует желаемому: %v. Найден: %v", tt.index, tt.val, found)
			return
		}
		if initLength == gotLength && tt.index < initLength && tt.index >= 0 {
			t.Errorf("Данный индекс был меньше длины листа, но элемент не был вставлен. Индекс: %d, длина: %d, необходимый элемент %v:", tt.index, length(list), tt.val)

			return
		}

		if initLength != gotLength && (tt.index < 0 || tt.index > initLength) {
			t.Errorf("Индекс вне границ, но элемент был вставлен. Индекс: %d, элемент: %v", tt.index, tt.val)
			return
		}

	}
}

// helpers

// lastNode returns pointer to last item.
func lastNode[T comparable](head *ListNode[T]) *ListNode[T] {
	var last *ListNode[T]
	for node := head; node != nil; node = node.Next {
		if node.Next == nil {
			last = node
		}
	}
	return last
}

// countVal returns count of the item with Val=val.
func countVal[T comparable](head *ListNode[T], val T) int {
	res := 0
	for node := head; node != nil; node = node.Next {
		if node.Val == val {
			res++
		}
	}
	return res
}

// length returns count of the items.
func length[T comparable](head *ListNode[T]) int {
	res := 0
	for node := head; node != nil; node = node.Next {
		res++
	}
	return res
}

// compareListWithSlice returns flag of equal values in listNode and slice []ini
// and if not equal returns difference.
func compareListWithSlice[T comparable](head *ListNode[T], sl []T) (bool, string) {
	lenSL := len(sl)
	lenHead := length(head)
	if lenSL != lenHead {
		return false, fmt.Sprintf("heal length=%v <> slice length=%v", lenHead, lenHead)
	}
	i := 0
	for node := head; node != nil; node = node.Next {
		if node.Val != sl[i] {
			return false, fmt.Sprintf("heal[%v]=%v <> sl[%v]=%v", i, node.Val, i, sl[i])
		}
		i++
	}
	return true, ""
}

// getElement returns an element of index i in the listnode or default value for T type.
func getElement[T comparable](head *ListNode[T], i int) T {
	searchInd := 0
	for node := head; node != nil; node = node.Next {
		if i == searchInd {
			return node.Val
		}
		searchInd++
	}
	return *new(T)
}
