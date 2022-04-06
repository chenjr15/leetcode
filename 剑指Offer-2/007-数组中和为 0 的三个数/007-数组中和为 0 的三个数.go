func threeSum(nums []int) [][]int {
    triples := make([][]int,0)
    length := len(nums)
    if length <3 {
        return triples
    }
    // 排序
    sort.Ints(nums)
    var idxB,idxC,target int
    for idxA,a :=range nums{
        // 如果a 大于0 则直接break，因为后面的所有都是正数再加也不会等于0了
        if a>0{
            break
        }
        // 对第一个数去重
        if idxA > 0 && nums[idxA]== nums[idxA-1]{
            continue
        }
        target = -a
        // 双指针搜索TwoSum 同时去重
        idxB=idxA+1
        idxC=length-1
        for idxB<idxC{
            
            if (nums[idxB] + nums[idxC]) == target{
                triples = append(triples,[]int{a,nums[idxB],nums[idxC]})
                // 数组有序且要求答案不重复，因此两个指针一定要前后移动
                idxB++
                idxC--
                for idxB<idxC && nums[idxB] == nums[idxB-1]{
                    idxB++
                }
                for idxB<idxC && nums[idxC] == nums[idxC+1]{
                    idxC--
                }
                
            }else if (nums[idxB] + nums[idxC]) > target{
                idxC--
            }else{
                idxB++
            }
        }
        

    }
    return triples

}