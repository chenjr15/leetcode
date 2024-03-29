# 2217.找到指定长度的回文数

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/find-palindrome-with-fixed-length/](https://leetcode-cn.com/problems/find-palindrome-with-fixed-length/)

## 题目描述

<p>给你一个整数数组&nbsp;<code>queries</code>&nbsp;和一个 <strong>正</strong>&nbsp;整数&nbsp;<code>intLength</code>&nbsp;，请你返回一个数组&nbsp;<code>answer</code>&nbsp;，其中&nbsp;<code>answer[i]</code> 是长度为&nbsp;<code>intLength</code>&nbsp;的&nbsp;<strong>正回文数</strong> 中第<em>&nbsp;</em><code>queries[i]</code>&nbsp;小的数字，如果不存在这样的回文数，则为 <code>-1</code>&nbsp;。</p>

<p><strong>回文数</strong> 指的是从前往后和从后往前读一模一样的数字。回文数不能有前导 0 。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<b>输入：</b>queries = [1,2,3,4,5,90], intLength = 3
<b>输出：</b>[101,111,121,131,141,999]
<strong>解释：</strong>
长度为 3 的最小回文数依次是：
101, 111, 121, 131, 141, 151, 161, 171, 181, 191, 202, ...
第 90 个长度为 3 的回文数是 999 。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<b>输入：</b>queries = [2,4,6], intLength = 4
<b>输出：</b>[1111,1331,1551]
<strong>解释：</strong>
长度为 4 的前 6 个回文数是：
1001, 1111, 1221, 1331, 1441 和 1551 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= queries.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>1 &lt;= queries[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= intLength&nbsp;&lt;= 15</code></li>
</ul>

## 标签

 - Array 
 - Math 

## 代码

```java
class Solution {
    public long[] kthPalindrome(int[] queries, int intLength) {
        long[] answer = new long[queries.length];
        // 中间的位数， 取值范围是0-9 共十位，
        long centerLen = (intLength - 1) / 2 ;
        // 个数最多为，头尾的取值范围是1-9 共9位，乘上 中间的位数，每位有10个选项
        long maxCnt = (long) (9 * Math.pow(10, centerLen));
        byte [] answerParts  = new byte[intLength];

        for (int i = 0, queriesLength = queries.length; i < queriesLength; i++) {
            int q = queries[i];
            // 计算每个查询
            if (q > maxCnt) {
                answer[i] = -1;
            }else{
                q--;
                // 先计算一半, 计算右边
                int right = intLength /2 ;
                while(right<intLength-1){
                    answerParts[right] = (byte) (q%10);
                    q/=10;
                    right++;
                }
                answerParts[intLength-1]= (byte) (q+1);

                //
                for (int l =0,r=intLength-1;l<r;l++,r--){
                    answerParts[l] = answerParts[r];
                }
                // 转换成long
                answer[i]= 0;
                for (byte part : answerParts) {
                    answer[i]*=10;
                    answer[i]+=part;
                }
                // System.out.println("q = "+(queries[i])+" "+Arrays.toString(answerParts)+answer[i]);
            }
        }
        return answer;
    }
}
```