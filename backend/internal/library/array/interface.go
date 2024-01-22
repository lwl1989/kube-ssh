package array

type ICompare interface {
	Compare(ICompare) bool
	GetCompareValue() interface{}
}
