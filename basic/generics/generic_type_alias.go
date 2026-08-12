package main

// 这是一个类型，不是结构体，可以拥有方法
type Set[T comparable] map[T]struct{}

// 构造函数
func NewSet[T comparable](n int) Set[T] {
	m := make(map[T]struct{}, n)
	return Set[T](m)
	// 将 map[T]struct{} 转换成 Set[T] 类型。
}

// 往Set里面添加元素
func (set Set[T]) Add(ele T) {
	set[ele] = struct{}{}
}

// 获取Set的长度
func (set Set[T]) Len() int {
	return len(set)
}

// 删除元素
func (set Set[T]) Remove(ele T) {
	delete(set, ele)
}

// 判断某个元素是否存在
func (set Set[T]) Exists(ele T) bool {
	_, exists := set[ele]
	return exists
}

// ForEach 遍历 Set，并对每个元素执行一次 handleElement。
// handleElement 的类型是 func(ele T)，表示 handleElement 是一个函数，并且这个函数接收一个 T 类型的元素。
func (set Set[T]) ForEach(handleElement func(ele T)) {
	for key := range set {
		handleElement(key) // 调用外面传进来的函数
	}
}
