func singleNumber(nums []int) []int {
    sum:= 0
    x1 := 0
    x2 := 0
    // 根据答案的思路
    // 将所有的元素异或一次以后，因为对同一个数异或两次能够消除这个数，所以最后异或的结果已经是x1^x2 即只剩下x1和x2
    for _,v :=range nums{
        sum^=v
    }
    // 接下去是想办法获得x1 和x2
    // 首先x1 不等于x2， 因此x1 和x2 异或必定不等于0，即sum不为0
    // 用lowbit求出sum最低位的1的位置为k，那么这个最低位的1。那么在x1 和x2中对应的第k位肯定相等。
    // 为此我们将该位为0 和为1数分开，x1 和x2 肯定分别在两个类中。同时，其他元素因为都有两个，这两个一定在同一个类中
    // 因此只要对这两个类分别异或即可得到对应的元素
    lowbit:= sum & -sum
    for _,v:=range nums{
        if (v&lowbit) == 0{
            x1 ^= v
        }else{
            x2 ^= v
        }
    }
    return []int{x1,x2}

}