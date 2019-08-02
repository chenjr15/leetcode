func merge(nums1 []int, m int, nums2 []int, n int)  {
    // 遇事不决，从后往前
    done := m+n -1
    p1 := m-1
    p2 := n-1
    for p1>=0&&p2>=0{
        if nums1[p1]>=nums2[p2]{
            nums1[done] = nums1[p1]
            done--
            p1--
        }else{
            nums1[done] = nums2[p2]
            done--
            p2--
        }
    }
    for p1>=0 {
        nums1[done] = nums1[p1]
        done--
        p1--
    }
    for p2>=0{
        nums1[done] = nums2[p2]
        done--
        p2--
    }
    
}
