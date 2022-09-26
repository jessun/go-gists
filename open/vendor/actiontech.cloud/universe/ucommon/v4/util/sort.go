package util

import "reflect"

type Slice []SliceItem

func ToSlice(slice interface{}) Slice {
	s := reflect.ValueOf(slice)
	ret := make(Slice, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

type SliceItem interface{}

type CompareFunc func(before, after SliceItem) bool

type PagingSorter struct {
	slice                        Slice       // 原始数组
	sortSlice                    Slice       // 因为在分页等过程中会减少被操作数组中的条目数, 所以添加一个中间态数组保证原数组不受影响
	sort                         Sort        // 排序器
	pager                        Pager       // 分页器
	ifBeforeLessOrEqualThanAfter CompareFunc // 比较两个条目大小用
}

// 可以使用ToSlice(slice interface{})函数,此函数可以自动将任意数组转换为Slice类型, 文档: http://10.186.18.11/confluence/pages/viewpage.action?pageId=32059663
func NewPagingSorter(slice Slice, ifBeforeLessOrEqualThanAfter CompareFunc) *PagingSorter {
	sortSlice := make(Slice, len(slice))
	for i := range slice {
		sortSlice[i] = slice[i]
	}
	return &PagingSorter{
		slice:                        slice,
		sortSlice:                    sortSlice,
		sort:                         &QuickSort{},
		pager:                        &DefaultPager{},
		ifBeforeLessOrEqualThanAfter: ifBeforeLessOrEqualThanAfter,
	}
}

func (p *PagingSorter) SetSorter(sort Sort) {
	p.sort = sort
}

func (p *PagingSorter) SetPager(pager Pager) {
	p.pager = pager
}

// 分页方法
func (p *PagingSorter) Paging(pageIndex, pageSize, defaultPageSize uint32) *PagingSorter {
	p.sortSlice = p.pager.Paging(p.sortSlice, pageIndex, pageSize, defaultPageSize)
	return p
}

/*
	排序方法:
	默认选型: 快速排序
	快速排序原理: 找到一个数pivot将数组'平均'分成两份,使其中一组均大于另一组, 此时pivot的位置是正确的位置, 然后再对这两个数组做相同操作
	默认原因: 数据量越大快速排序速度优势越大, 数据量小时用哪个排序速度都可以忽略
	实际案例: 随机数排序 ↓ ↓ ↓
		n=10
		冒泡排序耗时：2.4080276489257812e-05
		选择排序耗时：1.9311904907226562e-05
		插入排序耗时：1.5020370483398438e-05
		希尔排序耗时：1.5974044799804688e-05
		归并排序耗时：2.8848648071289062e-05
		快速排序耗时：1.9073486328125e-05

		n=100
		冒泡排序耗时：0.000782012939453125
		选择排序耗时：0.0004570484161376953
		插入排序耗时：0.00039076805114746094
		希尔排序耗时：0.00018095970153808594
		归并排序耗时：0.0003409385681152344
		快速排序耗时：0.00017905235290527344

		n=1000
		冒泡排序耗时：0.08327889442443848
		选择排序耗时：0.03776884078979492
		插入排序耗时：0.04986977577209473
		希尔排序耗时：0.0034036636352539062
		归并排序耗时：0.005920886993408203
		快速排序耗时：0.0021750926971435547

		n=10000
		冒泡排序耗时：8.781844854354858
		选择排序耗时：3.438148021697998
		插入排序耗时：4.186453819274902
		希尔排序耗时：0.05663800239562988
		归并排序耗时：0.06386470794677734
		快速排序耗时：0.02335190773010254

		n=100000
		冒泡排序耗时：900.5480690002441
		选择排序耗时：879.1669909954071
		插入排序耗时：428.66180515289307
		希尔排序耗时：0.967015266418457
		归并排序耗时：1.4872560501098633
		快速排序耗时：0.3050980567932129
—————————————————————————————————————————————————————————————
		基本有序数组排序 ↓ ↓ ↓
		n=10
		冒泡排序耗时：2.288818359375e-05
		选择排序耗时：1.9788742065429688e-05
		插入排序耗时：1.3113021850585938e-05
		希尔排序耗时：1.5974044799804688e-05
		归并排序耗时：2.9087066650390625e-05
		快速排序耗时：1.811981201171875e-05

		n=100
		冒泡排序耗时：0.0004851818084716797
		选择排序耗时：0.0004131793975830078
		插入排序耗时：0.00013065338134765625
		希尔排序耗时：0.00015997886657714844
		归并排序耗时：0.00032019615173339844
		快速排序耗时：0.00015974044799804688

		n=1000
		冒泡排序耗时：0.05040717124938965
		选择排序耗时：0.03394508361816406
		插入排序耗时：0.009570121765136719
		希尔排序耗时：0.0029370784759521484
		归并排序耗时：0.005821943283081055
		快速排序耗时：0.0022530555725097656

		n=10000
		冒泡排序耗时：5.24026083946228
		选择排序耗时：3.340329885482788
		插入排序耗时：0.8101489543914795
		希尔排序耗时：0.04622912406921387
		归并排序耗时：0.05988883972167969
		快速排序耗时：0.023930788040161133
——————————————————————————————————————————————————————————
*/
func (p *PagingSorter) Sorting(isAsc bool) *PagingSorter {
	p.sortSlice = p.sort.Sorting(p.sortSlice, isAsc, p.ifBeforeLessOrEqualThanAfter)
	return p
}

func (p *PagingSorter) Result() Slice {
	return p.sortSlice
}

type Pager interface {
	Paging(nums Slice, pageIndex, pageSize, defaultPageSize uint32) Slice
}

type DefaultPager struct {
}

func (d *DefaultPager) Paging(nums Slice, pageIndex, pageSize, defaultPageSize uint32) Slice {
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	if pageIndex == 0 {
		pageIndex = 1
	}

	// 如果起始点比最后一条还要靠后就显示空值
	if ((pageIndex - 1) * pageSize) >= uint32(len(nums)) {
		return Slice{}
	}

	// 开始点是 当前页的上一页 * 每页条目数,即不显示的条目数
	start := (pageIndex - 1) * pageSize
	// 结束点是 当前页的最后一条
	end := pageSize * pageIndex
	//如果结束点比最后一条多就显示到最后一条即可
	if end > uint32(len(nums)) {
		end = uint32(len(nums))
	}
	// 收集当前页内容
	temp := Slice{}
	for i := start; i < end; i++ {
		temp = append(temp, nums[i])
	}
	nums = temp
	return nums
}

type Sort interface {
	Sorting(nums Slice, isAsc bool, ifBeforeLessOrEqualThanAfter CompareFunc) Slice
}

type QuickSort struct {
}

func (q *QuickSort) Sorting(nums Slice, isAsc bool, ifBeforeLessOrEqualThanAfter CompareFunc) Slice {
	q.recursionSort(nums, 0, len(nums)-1, ifBeforeLessOrEqualThanAfter)
	if !isAsc {
		ss := make(Slice, len(nums))
		for i, num := range nums {
			ss[len(nums)-i-1] = num
		}
		nums = ss
	}
	return nums
}

func (q *QuickSort) recursionSort(nums Slice, left int, right int, ifBeforeLessOrEqualThanAfter CompareFunc) {
	if left < right {
		pivot := q.partition(nums, left, right, ifBeforeLessOrEqualThanAfter)
		q.recursionSort(nums, left, pivot-1, ifBeforeLessOrEqualThanAfter)
		q.recursionSort(nums, pivot+1, right, ifBeforeLessOrEqualThanAfter)
	}
}

func (q *QuickSort) partition(nums Slice, left int, right int, ifBeforeLessOrEqualThanAfter CompareFunc) int {
	for left < right {
		for left < right && ifBeforeLessOrEqualThanAfter(nums[left], nums[right]) {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}

		for left < right && ifBeforeLessOrEqualThanAfter(nums[left], nums[right]) {
			left++
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}
	return left
}
