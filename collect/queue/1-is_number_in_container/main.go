package main

/*
已知数字k在容器内，则数宇k*2+1和k*3+1也在容器内。编写
一段程序计算已知数字k的情况下，另一个数字是否在容器内
*/

func isNumberInContainer(k, target int) bool {
	queue := []int{k}

	for len(queue) != 0 {
		temp := queue[0]
		queue = queue[1:]

		if temp == target {
			return true
		}
		if temp*2+1 == target || temp*3+1 == target {
			return true
		}
		if target > 0 {
			if temp*2+1 < target {
				queue = append(queue, temp*2+1)
			}
			if temp*3+1 < target {
				queue = append(queue, temp*3+1)
			}
		}
		if target < 0 {
			if temp*2+1 > target {
				queue = append(queue, temp*2+1)
			}
			if temp*3+1 > target {
				queue = append(queue, temp*3+1)
			}
		}
	}
	return false
}
