package life

import (
	"strconv"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestHouse(t *testing.T) {
	tl := []struct {
		data    []int64
		wantOld int64
		wantNew int64
	}{
		{
			data:    []int64{1000, 1500, 1500},
			wantOld: 4000 * 20,
			wantNew: (1000*3 + 1500*2 + 1500*1) * 9 / 10, // 整百才能这样算
		},
	}
	for i, s := range tl {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			fund := NewHosingFund(WithData(s.data))
			oldH := fund.HouseOld() // 老算法可以贷款的额度
			if oldH != s.wantOld {
				t.Errorf("data is (%+v) wantOld (%d) got (%d)", s.data, s.wantOld, oldH)
			}
			newH := fund.HouseNew() // 新算法可以贷款的额度
			if newH != s.wantNew {
				t.Errorf("data is (%+v) wantNew (%d) got (%d)", s.data, s.wantOld, newH)
			}
		})
	}
}

func TestInterestMonth(t *testing.T) {
	tl := []struct {
		hosingFundAmount int64 // 公积金贷款额度
		loan             int64 // 需要贷款总额
		want             int64 // 每月需要还款金额
	}{
		{
			hosingFundAmount: 40 * 10000,
			loan:             100 * 10000,
			want:             4486,
		},
		{
			hosingFundAmount: 20 * 10000,
			loan:             100 * 10000,
			want:             4558,
		},
	}
	for i, s := range tl {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			fund := NewHosingFund(WithHosingFundAmount(s.hosingFundAmount), WithLoan(s.loan))
			interestMonth := fund.InterestMonth()
			assert.Equal(t, int64(interestMonth), s.want)
		})
	}
}
