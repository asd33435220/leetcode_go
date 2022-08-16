package main

import "fmt"

func main() {
	//OrderedStream os= new OrderedStream(5);
	//os.insert(3, "ccccc"); // 插入 (3, "ccccc")，返回 []
	//os.insert(1, "aaaaa"); // 插入 (1, "aaaaa")，返回 ["aaaaa"]
	//os.insert(2, "bbbbb"); // 插入 (2, "bbbbb")，返回 ["bbbbb", "ccccc"]
	//os.insert(5, "eeeee"); // 插入 (5, "eeeee")，返回 []
	//os.insert(4, "ddddd"); // 插入 (4, "ddddd")，返回 ["ddddd", "eeeee"]
	//
	os := Constructor(5)
	fmt.Println(os.Insert(3, "ccccc"))
	fmt.Println(os.Insert(1, "aaaaa"))
	fmt.Println(os.Insert(2, "bbbbb"))
	fmt.Println(os.Insert(5, "eeeee"))
	fmt.Println(os.Insert(4, "ddddd"))
}

type OrderedStream struct {
	kv  []string
	ptr int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		kv:  make([]string, n),
		ptr: 1,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) (result []string) {
	this.kv[idKey-1] = value
	for i := this.ptr - 1; i < len(this.kv); i++ {
		if this.kv[i] == "" {
			return result
		} else {
			result = append(result, this.kv[i])
		}
		this.ptr++
	}
	return result
}
