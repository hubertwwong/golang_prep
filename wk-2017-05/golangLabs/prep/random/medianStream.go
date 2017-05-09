package main

import "fmt"

type MedianFinder struct {
    nums []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
    fmt.Println(">ctor")
    var mf MedianFinder
    mf.nums = make([]int, 0)
    return mf
}

func (this *MedianFinder) AddNum(num int)  {
    fmt.Println(">AddNum", num)
    this.nums = append(this.nums, num)
    for i := len(this.nums)-1  ; i > 0 ; i-- {
        fmt.Println(i)
        if this.nums[i] < this.nums[i-1] {
            this.nums[i], this.nums[i-1] = this.nums[i-1], this.nums[i]
        }
    }
}

func (this *MedianFinder) FindMedian() float64 {
    fmt.Println(">findMedian")
    if len(this.nums) == 0 {
        return 0.0
    } else if len(this.nums) % 2 == 1 {
        return float64(this.nums[len(this.nums)/2])
    } else {
        fmt.Println(len(this.nums)/2)
        left := this.nums[len(this.nums)/2]
        right := this.nums[(len(this.nums)/2)-1]
        fmt.Println(left, right)
        return float64(float64(left + right) / 2.0)
    }
}

func main() {
  obj := Constructor()
  obj.AddNum(1)
  obj.AddNum(2)
  param2 := obj.FindMedian()
  fmt.Println(param2)
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */