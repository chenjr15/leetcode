func groupAnagrams(strs []string) [][]string {
    // 考虑用字符哈希来逐个比较，
    // 逐个比较的过程中可以给他排序按长度排序，然后按哈希排序
    // 实在不行把每个字符串排序
    // -------------
    // 考虑使用字符串压缩方法，来做哈希表的key
    strHash:= map[[26]int][]string {}

    for _,str:=range strs{
        // 计算压缩哈希
        cntKey:=[26]int{}
        for i:=range str{
            cntKey[str[i]-'a'] ++
        }
        list:= strHash[cntKey]
        if list==nil{
            list=[]string{str}
        }else{
            list = append(list,str)
        }
        strHash[cntKey] = list
    }
    reslut := [][]string{}
    for _,str:= range strHash{
        reslut=append(reslut,str)
    }
    return reslut
}

