type NumMatrix struct {
    Msum [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return NumMatrix{}
    }
    // msum[i][j] 表示每个 [0,0] 到 [i,j]位置(包含)的和
    msum:=make([][]int,len(matrix))
    for row:=range matrix{
        msum[row] = make([]int,len(matrix[0]))
        msum[row][0] = matrix[row][0]
       
        for col:=1 ;col<len(matrix[0]);col++{
            msum[row][col] = matrix[row][col]+msum[row][col-1]
            
        }
        
    }
    for col:=range matrix[0]{
        for row:= range matrix{
            if row>0{
                msum[row][col] += msum[row-1][col]
            }
        }
    }

    return NumMatrix{Msum:msum}

}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    // (0,0)-(row2,col2)代表的矩形 减去 (0,0)-(row1-1,col1-1)代表的矩形 的面积 再减去 上面和左边 分别多出来的面积就是答案
    // 
    sum:= this.Msum[row2][col2]
    // 减去 左边的小矩形
    if col1 != 0 {
        sum -= this.Msum[row2][col1-1]
    }
    // 减去 上面的小矩形
    if row1 !=0{
        sum -= this.Msum[row1-1][col2]
    }
    // (0,0)-(row1-1,col1-1) 因为和左边上面的都有重合，可能被减去两次，因此加回来
    if row1 !=0 && col1 !=0 {
        sum += this.Msum[row1-1][col1-1]
    }
    return sum

}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */