type MovingAverage struct {
    size int
    cnt int
    avg float64
    head int
    nums []int
    
}


/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
    return MovingAverage{size:size,nums:make([]int,0,size)}
}


func (this *MovingAverage) Next(val int) float64 {
   
    if this.cnt<this.size{
         if this.cnt==0{
            this.avg = float64(val)
        }else{
            this.avg =( this.avg * float64(this.cnt) +float64(val))/float64(this.cnt+1)
        }
        this.nums = append(this.nums,val)
       
    }else{
        this.avg =this.avg - float64(this.nums[this.cnt%this.size])/float64(this.size) + float64(val)/float64(this.size)
        this.nums[this.cnt%this.size] = val

    }
    this.cnt++
  
    return this.avg
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */