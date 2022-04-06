# 5235.找出输掉零场或一场比赛的玩家

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/find-players-with-zero-or-one-losses/](https://leetcode-cn.com/problems/find-players-with-zero-or-one-losses/)

## 题目描述

<p>给你一个整数数组 <code>matches</code> 其中 <code>matches[i] = [winner<sub>i</sub>, loser<sub>i</sub>]</code> 表示在一场比赛中 <code>winner<sub>i</sub></code> 击败了 <code>loser<sub>i</sub></code> 。</p>

<p>返回一个长度为 2 的列表<em> </em><code>answer</code> ：</p>

<ul>
	<li><code>answer[0]</code> 是所有 <strong>没有</strong> 输掉任何比赛的玩家列表。</li>
	<li><code>answer[1]</code> 是所有恰好输掉 <strong>一场</strong> 比赛的玩家列表。</li>
</ul>

<p>两个列表中的值都应该按 <strong>递增</strong> 顺序返回。</p>

<p><strong>注意：</strong></p>

<ul>
	<li>只考虑那些参与 <strong>至少一场</strong> 比赛的玩家。</li>
	<li>生成的测试用例保证 <strong>不存在</strong> 两场比赛结果 <strong>相同</strong> 。</li>
</ul>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>matches = [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]
<strong>输出：</strong>[[1,2,10],[4,5,7,8]]
<strong>解释：</strong>
玩家 1、2 和 10 都没有输掉任何比赛。
玩家 4、5、7 和 8 每个都输掉一场比赛。
玩家 3、6 和 9 每个都输掉两场比赛。
因此，answer[0] = [1,2,10] 和 answer[1] = [4,5,7,8] 。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>matches = [[2,3],[1,3],[5,4],[6,4]]
<strong>输出：</strong>[[1,2,5,6],[]]
<strong>解释：</strong>
玩家 1、2、5 和 6 都没有输掉任何比赛。
玩家 3 和 4 每个都输掉两场比赛。
因此，answer[0] = [1,2,5,6] 和 answer[1] = [] 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= matches.length &lt;= 10<sup>5</sup></code></li>
	<li><code>matches[i].length == 2</code></li>
	<li><code>1 &lt;= winner<sub>i</sub>, loser<sub>i</sub> &lt;= 10<sup>5</sup></code></li>
	<li><code>winner<sub>i</sub> != loser<sub>i</sub></code></li>
	<li>所有 <code>matches[i]</code> <strong>互不相同</strong></li>
</ul>

## 标签

 - Graph 

## 代码

```java
     class Solution {
        public List<List<Integer>> findWinners(int[][] matches) {
            HashMap<Integer, Integer> winner = new HashMap<>();
            HashMap<Integer, Integer> loser = new HashMap<>();

            // int[][] result = new int[2][];
            for (int[] match : matches) {
                Integer winCnt = winner.getOrDefault(match[0], 0);
                Integer loseCnt = loser.getOrDefault(match[1], 0);
                winner.put(match[0], winCnt + 1);
                loser.put(match[1], loseCnt + 1);
            }
            ArrayList<List<Integer>> result = new ArrayList<>();
            // 计算结果
            ArrayList<Integer> ans0 = new ArrayList<>();
            for (Map.Entry<Integer, Integer> entry : winner.entrySet()) {
                Integer key = entry.getKey();
                Integer value = entry.getValue();
                if (loser.getOrDefault(key, 0).equals(0) )
                    ans0.add(key);
            }
            
            ans0.sort(Comparator.comparingInt(a -> a));
            result.add(ans0);
            ArrayList<Integer> ans1 = new ArrayList<>();
            for (Map.Entry<Integer, Integer> entry : loser.entrySet()) {
                Integer idx = entry.getKey();
                Integer times = entry.getValue();
                if (times == 1 )
                    ans1.add(idx);
            }
            ans1.sort(Comparator.comparingInt(a -> a));

            result.add(ans1);

            return result;


        }
    }
```