package projectprepare

// CopySlice 把切片src里的元素拷贝到dest里，返回成功拷贝的元素个数
func CopySlice[T any](dest, src []T) int {
	if len(dest) == 0 || len(src) == 0 {
		return 0
	}
	i, j := 0, 0
	for ; i < len(dest) && j < len(src); i, j = i+1, j+1 {
		dest[i] = src[j]
	}
	return i
}

// 这个是在讲可见性的，就是变量，结构体，函数都必须要大写
// 这样其他包可以访问到，否则是不行的
// var A int

// type User struct {
// 	Name string
// 	age  int
// }

// type Ifc interface{}
