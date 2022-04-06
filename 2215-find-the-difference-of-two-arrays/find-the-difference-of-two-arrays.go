func findDifference(nums1 []int, nums2 []int) [][]int {
    ans:=[][]int{
        []int{},
        []int{},
    }
    hash1 := make(map[int]struct{})    
    hash2 := make(map[int]struct{})    
    for _,n:=range nums1{
        hash1[n] = struct{}{}
    }
    for _,n:=range nums2{
        if _,ok:=hash1[n];!ok{
            ans[1] = append(ans[1],n)
            // 去重
            hash1[n] = struct{}{}

        }
         hash2[n] = struct{}{}
    }
    for _,n:=range nums1{
        if _,ok:=hash2[n];!ok{
            ans[0] = append(ans[0],n)
            // 去重
            hash2[n] = struct{}{}
        }
    }
    return ans

}