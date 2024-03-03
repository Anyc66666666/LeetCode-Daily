package main

/*
游戏内有需求要把图片的像素数据沿横轴做垂直翻转，一个像素内的数据由rgba4个数字组成
1111 2222 3333 4444 =>(2)=>  2222 1111 4444 3333
*/

func verticalFlip(k int, pixelData []int) []int {
	splits := len(pixelData) / 4
	var ans []int

	for i := 0; i < splits; i += k {
		tmp := pixelData[i*4 : (i+k)*4]
		reverse(tmp)
		ans = append(ans, tmp...)
	}
	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
