# 599.两个列表的最小索引总和

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists/](https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists/)

## 题目描述

<p>假设 Andy 和 Doris 想在晚餐时选择一家餐厅，并且他们都有一个表示最喜爱餐厅的列表，每个餐厅的名字用字符串表示。</p>

<p>你需要帮助他们用<strong>最少的索引和</strong>找出他们<strong>共同喜爱的餐厅</strong>。 如果答案不止一个，则输出所有答案并且不考虑顺序。 你可以假设答案总是存在。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入: </strong>list1 = ["Shogun", "Tapioca Express", "Burger King", "KFC"]，list2 = ["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
<strong>输出:</strong> ["Shogun"]
<strong>解释:</strong> 他们唯一共同喜爱的餐厅是“Shogun”。
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong>list1 = ["Shogun", "Tapioca Express", "Burger King", "KFC"]，list2 = ["KFC", "Shogun", "Burger King"]
<strong>输出:</strong> ["Shogun"]
<strong>解释:</strong> 他们共同喜爱且具有最小索引和的餐厅是“Shogun”，它有最小的索引和1(0+1)。
</pre>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>

<ul>
	<li><code>1 &lt;= list1.length, list2.length &lt;= 1000</code></li>
	<li><code>1 &lt;= list1[i].length, list2[i].length &lt;= 30</code>&nbsp;</li>
	<li><code>list1[i]</code> 和 <code>list2[i]</code> 由空格<meta charset="UTF-8" />&nbsp;<code>' '</code>&nbsp;和英文字母组成。</li>
	<li><code>list1</code> 的所有字符串都是 <strong>唯一</strong> 的。</li>
	<li><code>list2</code> 中的所有字符串都是 <strong>唯一</strong> 的。</li>
</ul>

## 标签

 - Array 
 - Hash Table 
 - String 

## 代码

```java
class Solution {
    public String[] findRestaurant(String[] list1, String[] list2) {
        if (list1 == null || list2==null ){
            return new String[0];
        }
        HashMap<String,Integer> map1 = new HashMap<>();
        ArrayList<String> common = new ArrayList();;
        int minSum = list2.length+list1.length;
        for(int i =0;i<list1.length;++i){
            map1.put(list1[i],i);
        }
        String s;
        for(int i =0;i<list2.length;++i){
           s = list2[i];
           if (s==""){
               continue;
           }
           Integer sum = map1.get(s);
           if(sum == null){
               continue;
           }
           sum += i;
           if(sum == minSum){
               common.add(s);
           }else if (sum < minSum){
               common.clear();
               common.add(s);
               minSum = sum;
           }
        //    System.out.println(common);
        }
        String[] arr = new String[common.size()];
        common.toArray(arr);
        return arr;
    }
}
```