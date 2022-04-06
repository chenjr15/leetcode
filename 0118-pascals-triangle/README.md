# 118.杨辉三角

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/pascals-triangle/](https://leetcode-cn.com/problems/pascals-triangle/)

## 题目描述

<p>给定一个非负整数 <em><code>numRows</code>，</em>生成「杨辉三角」的前 <em><code>numRows</code> </em>行。</p>

<p><small>在「杨辉三角」中，每个数是它左上方和右上方的数的和。</small></p>

<p><img alt="" src="https://pic.leetcode-cn.com/1626927345-DZmfxB-PascalTriangleAnimated2.gif" /></p>

<p> </p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> numRows = 5
<strong>输出:</strong> [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> numRows = 1
<strong>输出:</strong> [[1]]
</pre>

<p> </p>

<p><strong>提示:</strong></p>

<ul>
	<li><code>1 <= numRows <= 30</code></li>
</ul>

## 标签

 - Array 
 - Dynamic Programming 

## 代码

```java
class Solution {
    public List<List<Integer>> generate(int numRows) {
        List<List<Integer>> result = new ArrayList<List<Integer>>();
        ArrayList<Integer> list=  new ArrayList<Integer>();
        list.add(1);
        result.add(list);
        for(int row = 1;row < numRows;++row){
            int colMax = row+1;
            list=  new ArrayList<Integer>();
            list.add(1);
            for (int col = 1 ;col<colMax;col++){
                // System.out.println("row:"+row+" col:"+col);
                int val = result.get(row-1).get(col-1);
                if (col < row){
    
                    val += result.get(row-1).get(col);
                }
                list.add(val);
            }
            result.add(list);
        }
        return result;    
    }

}
```