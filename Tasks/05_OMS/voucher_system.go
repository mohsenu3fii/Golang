package main

import "fmt"

func NewPercentageVoucher(code string, percent int) (*PercentageVoucher, error) {
	// - Validate percent (1–100).
	if percent > 0 && percent <= 100 {
		return &PercentageVoucher{code: code, percent: percent}, nil
	}
	return nil, fmt.Errorf("The code or percentage was not entered correctly")
}

func (v *PercentageVoucher) Apply(total int) int {
	// - Handle nil interface correctly.
	if v == nil {
		return total
	}
	// - Prevent applying voucher twice.

	result := total - (total * v.percent / 100)

	// - Do not allow negative result totals.
	if result < 0 {
		return total
	}

	return result
}
func (v *PercentageVoucher) Code() string {
	// - Handle nil interface correctly.
	if v == nil {
		return ""
	}
	return v.code
}

func NewFixedAmountVoucher(code string, amount int) (*FixedAmountVoucher, error) {
	// - Validate fixed amount (>0).
	if amount > 0 {
		return &FixedAmountVoucher{code: code, amount: amount}, nil
	}
	return &FixedAmountVoucher{}, fmt.Errorf("The code or percentage was not entered correctly")
}

func (v *FixedAmountVoucher) Apply(total int) int {

	// - Handle nil interface correctly.
	if v == nil {
		return total
	}
	result := total - v.amount

	// - Do not allow negative result totals.
	if result < 0 {
		return total
	}
	return result
}
func (v *FixedAmountVoucher) Code() string {
	// - Handle nil interface correctly.
	if v == nil {
		return ""
	}
	return v.code
}
