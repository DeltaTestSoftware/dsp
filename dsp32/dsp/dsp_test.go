package dsp

import (
	"math"
	"testing"

	"github.com/gonutz/check"
)

func TestCopyReturnsANewArray(t *testing.T) {
	a := []float32{1, 2, 3}
	b := Copy(a)
	a[1] = 0
	check.Eq(t, a, []float32{1, 0, 3})
	check.Eq(t, b, []float32{1, 2, 3})
}

func TestCopyingEmptySliceReturnsEmptySlice(t *testing.T) {
	check.Eq(t, len(Copy(nil)), 0)
}

func TestMinMaxReturnsIndicesAndValuesOfFirstExtremes(t *testing.T) {
	a := []float32{3, 2, 1, 3, 1, 4, 4}
	minIndex, minValue, maxIndex, maxValue := MinMax(a)
	check.Eq(t, minIndex, 2)
	check.Eq(t, minValue, 1)
	check.Eq(t, maxIndex, 5)
	check.Eq(t, maxValue, 4)
}

func TestMinMaxOfEmptyArrayIsNegativeAndInfinite(t *testing.T) {
	minIndex, minValue, maxIndex, maxValue := MinMax(nil)
	check.Eq(t, minIndex, -1)
	check.Eq(t, math.IsInf(float64(minValue), 1), true)
	check.Eq(t, maxIndex, -1)
	check.Eq(t, math.IsInf(float64(maxValue), -1), true)
}

func TestMinMaxIndicesAndValuesCanBeComputedOnTheirOwn(t *testing.T) {
	a := []float32{5, 0, 5, 9, 5}
	check.Eq(t, MinIndex(a), 1)
	check.Eq(t, MinValue(a), 0)
	check.Eq(t, MaxIndex(a), 3)
	check.Eq(t, MaxValue(a), 9)
}

func TestAverageFilter(t *testing.T) {
	check.Eq(t, AverageFilter([]float32{2, 4, 6, 8}, 2), []float32{3, 5, 7})
	check.Eq(t, AverageFilter([]float32{1, 2, 3, 4, 5}, 3), []float32{2, 3, 4})
	check.Eq(t, AverageFilter([]float32{0, 2, 4, 6, 8, 10}, 4), []float32{3, 5, 7})
}

func TestAverageFilterLeavesAtLeastOneElement(t *testing.T) {
	check.Eq(t, AverageFilter([]float32{1, 2, 3}, 3), []float32{2})
	check.Eq(t, AverageFilter([]float32{1, 2, 3}, 4), []float32{2})
	check.Eq(t, AverageFilter([]float32{1, 2, 3}, 999), []float32{2})
}

func TestAverageFilterOverEmptyInputReturnsEmptyOutput(t *testing.T) {
	check.Eq(t, AverageFilter(nil, 0), nil)
	check.Eq(t, AverageFilter(nil, 1), nil)
	check.Eq(t, AverageFilter(nil, 2), nil)
}

func TestAverageFilterOfWidthOneOrLessReturnsCopyOfInput(t *testing.T) {
	for width := 1; width >= -2; width-- {
		a := []float32{1, 2, 3}
		avg := AverageFilter(a, width)
		a[1] = 0
		check.Eq(t, avg, []float32{1, 2, 3})
		check.Eq(t, a, []float32{1, 0, 3})
	}
}

func TestMedianFilter(t *testing.T) {
	check.Eq(t, MedianFilter([]float32{2, 1, 30, 50, 44}, 3), []float32{2, 30, 44})
	check.Eq(t, MedianFilter([]float32{1, 3, 2}, 2), []float32{3, 3})
	check.Eq(t, MedianFilter([]float32{1, 3, 2}, 1), []float32{1, 3, 2})
}

func TestMedianFilterLeavesAtLeastOneElement(t *testing.T) {
	check.Eq(t, MedianFilter([]float32{1, 2, 3}, 3), []float32{2})
	check.Eq(t, MedianFilter([]float32{1, 2, 3}, 4), []float32{2})
	check.Eq(t, MedianFilter([]float32{1, 2, 3}, 999), []float32{2})
}

func TestMedianFilterOverEmptyInputReturnsEmptyOutput(t *testing.T) {
	check.Eq(t, MedianFilter(nil, 0), nil)
	check.Eq(t, MedianFilter(nil, 1), nil)
	check.Eq(t, MedianFilter(nil, 2), nil)
}

func TestMedianFilterOfWidthOneOrLessReturnsCopyOfInput(t *testing.T) {
	for width := 1; width >= -2; width-- {
		a := []float32{1, 2, 3}
		avg := MedianFilter(a, width)
		a[1] = 0
		check.Eq(t, avg, []float32{1, 2, 3})
		check.Eq(t, a, []float32{1, 0, 3})
	}
}

func TestAverage(t *testing.T) {
	check.Eq(t, Average(nil), 0)
	check.Eq(t, Average([]float32{8}), 8)
	check.Eq(t, Average([]float32{1, 2}), 1.5)
	check.Eq(t, Average([]float32{1, 2, 3}), 2)
}

func TestNegation(t *testing.T) {
	check.Eq(t, Negative([]float32{1, -2, 3}), []float32{-1, 2, -3})
}

func TestDerivative(t *testing.T) {
	check.Eq(t, Derivative([]float32{}), []float32{})
	check.Eq(t, Derivative([]float32{1}), []float32{0})
	check.Eq(t, Derivative([]float32{1, 3}), []float32{2})
	check.Eq(t, Derivative([]float32{1, 3, 4}), []float32{2, 1})
	check.Eq(t, Derivative([]float32{1, 3, 4, 2}), []float32{2, 1, -2})
}

func TestNthDerivative(t *testing.T) {
	check.Eq(t, NthDerivative([]float32{1, 3, 4, 2}, -1), []float32{1, 3, 4, 2})
	check.Eq(t, NthDerivative([]float32{1, 3, 4, 2}, 0), []float32{1, 3, 4, 2})
	check.Eq(t, NthDerivative([]float32{1, 3, 4, 2}, 1), []float32{2, 1, -2})
	check.Eq(t, NthDerivative([]float32{1, 3, 4, 2}, 2), []float32{-1, -3})
	check.Eq(t, NthDerivative([]float32{1, 3, 4, 2}, 3), []float32{-2})
	check.Eq(t, NthDerivative([]float32{1, 3, 4, 2}, 4), []float32{0})
	check.Eq(t, NthDerivative([]float32{1, 3, 4, 2}, 5), []float32{0})
}

func TestAddUsesTheLowestCommonElementCount(t *testing.T) {
	check.Eq(t, Add(), nil)
	check.Eq(t, Add(nil, nil), nil)
	check.Eq(t, Add([]float32{1, 2, 3}), []float32{1, 2, 3})
	check.Eq(t, Add([]float32{1, 2, 3}, []float32{4, 7, 9}), []float32{5, 9, 12})
	check.Eq(t, Add([]float32{1, 2}, []float32{4, 7, 9}), []float32{5, 9})
	check.Eq(t, Add([]float32{1, 2, 3}, []float32{4, 7}), []float32{5, 9})
	check.Eq(t, Add([]float32{1}, []float32{2}, []float32{3}), []float32{6})
}

func TestSubUsesTheLowestCommonElementCount(t *testing.T) {
	check.Eq(t, Sub(), nil)
	check.Eq(t, Sub(nil, nil), nil)
	check.Eq(t, Sub([]float32{5, 2, 1}), []float32{5, 2, 1})
	check.Eq(t, Sub([]float32{9, 8, 3}, []float32{4, 7, 9}), []float32{5, 1, -6})
	check.Eq(t, Sub([]float32{9, 8}, []float32{4, 7, 9}), []float32{5, 1})
	check.Eq(t, Sub([]float32{9, 8, 3}, []float32{4, 7}), []float32{5, 1})
	check.Eq(t, Sub([]float32{5}, []float32{1}, []float32{2}), []float32{2})
}

func TestAddOffsetAddsValueToAll(t *testing.T) {
	check.Eq(t, AddOffset(nil, 1), nil)
	check.Eq(t, AddOffset([]float32{2}, 1), []float32{3})
	check.Eq(t, AddOffset([]float32{2, 3, 4}, -1), []float32{1, 2, 3})
}

func TestEveryNthTakesEveryNthElement(t *testing.T) {
	check.Eq(t, EveryNth(nil, 3), nil)
	check.Eq(t, EveryNth([]float32{1}, 3), []float32{1})
	check.Eq(t, EveryNth([]float32{1, 2}, 3), []float32{1})
	check.Eq(t, EveryNth([]float32{1, 2, 3}, 3), []float32{1})
	check.Eq(t, EveryNth([]float32{1, 2, 3, 4}, 3), []float32{1, 4})
	check.Eq(t, EveryNth([]float32{1, 2, 3, 4, 5}, 3), []float32{1, 4})
	check.Eq(t, EveryNth([]float32{1, 2, 3, 4, 5, 6}, 3), []float32{1, 4})
	check.Eq(t, EveryNth([]float32{1, 2, 3, 4, 5, 6, 7}, 3), []float32{1, 4, 7})

	check.Eq(t, EveryNth([]float32{1, 2, 3}, 1), []float32{1, 2, 3})

	check.Eq(t, EveryNth([]float32{1, 2, 3}, 0), nil)
	check.Eq(t, EveryNth([]float32{1, 2, 3}, -1), nil)
}

func TestRepeatMakesArrayOfSameValues(t *testing.T) {
	check.Eq(t, Repeat(1.5, -1), nil)
	check.Eq(t, Repeat(1.5, 0), nil)
	check.Eq(t, Repeat(1.5, 1), []float32{1.5})
	check.Eq(t, Repeat(1.5, 2), []float32{1.5, 1.5})
	check.Eq(t, Repeat(1.5, 3), []float32{1.5, 1.5, 1.5})
}

func TestReverseReturnsValuesInFlippedOrder(t *testing.T) {
	check.Eq(t, Reverse(nil), nil)
	check.Eq(t, Reverse([]float32{1}), []float32{1})
	check.Eq(t, Reverse([]float32{1, 2}), []float32{2, 1})
	check.Eq(t, Reverse([]float32{1, 2, 3}), []float32{3, 2, 1})
}

func TestScaleMultipliesEveryElementWithGivenFactor(t *testing.T) {
	check.Eq(t, Scale(nil, 2), nil)
	check.Eq(t, Scale([]float32{1}, 2), []float32{2})
	check.Eq(t, Scale([]float32{1, 2}, 2), []float32{2, 4})
	check.Eq(t, Scale([]float32{1, 2, 3}, 2), []float32{2, 4, 6})
}

func TestAbsReturnsAbsoluteValues(t *testing.T) {
	check.Eq(t, Abs(nil), nil)
	check.Eq(t, Abs([]float32{1}), []float32{1})
	check.Eq(t, Abs([]float32{-1}), []float32{1})
	check.Eq(t, Abs([]float32{1, -2, 3, -4}), []float32{1, 2, 3, 4})
}

func TestAbsValueReturnsAbsoluteValueOfSingleInput(t *testing.T) {
	check.Eq(t, AbsValue(1), 1)
	check.Eq(t, AbsValue(-1), 1)
	nan := float32(math.NaN())
	posInf := float32(math.Inf(1))
	negInf := float32(math.Inf(-1))
	check.Eq(t, AbsValue(nan), nan)
	check.Eq(t, AbsValue(posInf), posInf)
	check.Eq(t, AbsValue(negInf), posInf)
}

func TestRangeEnumeratesIntegersAsFloats(t *testing.T) {
	check.Eq(t, Range(0, 0), []float32{0.0})
	check.Eq(t, Range(0, 1), []float32{0.0, 1.0})
	check.Eq(t, Range(10, 8), []float32{10.0, 9.0, 8.0})
	check.Eq(t, Range(-2, 3), []float32{-2, -1, 0, 1, 2, 3})
	check.Eq(t, Range(3, -2), []float32{3, 2, 1, 0, -1, -2})
}

func TestMulUsesTheLowestCommonElementCount(t *testing.T) {
	check.Eq(t, Mul(), nil)
	check.Eq(t, Mul(nil, nil), nil)
	check.Eq(t, Mul([]float32{2, 3, 4}), []float32{2, 3, 4})
	check.Eq(t, Mul([]float32{2, 3, 4}, []float32{5, 6, 7}), []float32{2 * 5, 3 * 6, 4 * 7})
	check.Eq(t, Mul([]float32{2, 3}, []float32{4, 5, 6}), []float32{2 * 4, 3 * 5})
	check.Eq(t, Mul([]float32{2, 3, 4}, []float32{5, 6}), []float32{2 * 5, 3 * 6})
	check.Eq(t, Mul([]float32{2}, []float32{3}, []float32{4}), []float32{2 * 3 * 4})
}

func TestReciprocal(t *testing.T) {
	check.Eq(t, Reciprocal(nil), nil)
	check.Eq(t, Reciprocal([]float32{999}), []float32{1.0 / 999})
	check.Eq(t, Reciprocal([]float32{1, 2, 3}), []float32{1.0 / 1, 1.0 / 2, 1.0 / 3})
	var zero float32
	check.Eq(t, Reciprocal([]float32{0}), []float32{1.0 / zero})
}

func TestSafeReciprocal(t *testing.T) {
	check.Eq(t, SafeReciprocal(nil, 0), nil)
	check.Eq(t, SafeReciprocal([]float32{999}, 0), []float32{1.0 / 999})
	check.Eq(t, SafeReciprocal([]float32{1, 2, 3}, 0), []float32{1.0 / 1, 1.0 / 2, 1.0 / 3})
	check.Eq(t, SafeReciprocal([]float32{0}, 123), []float32{123})
}
