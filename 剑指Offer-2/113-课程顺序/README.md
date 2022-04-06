# 剑指 Offer II 113.课程顺序

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/QA2IGt/](https://leetcode-cn.com/problems/QA2IGt/)

## 题目描述

<p>现在总共有 <code>numCourses</code>&nbsp;门课需要选，记为&nbsp;<code>0</code>&nbsp;到&nbsp;<code>numCourses-1</code>。</p>

<p>给定一个数组&nbsp;<code>prerequisites</code> ，它的每一个元素&nbsp;<code>prerequisites[i]</code>&nbsp;表示两门课程之间的先修顺序。&nbsp;例如&nbsp;<code>prerequisites[i] = [a<sub>i</sub>, b<sub>i</sub>]</code>&nbsp;表示想要学习课程 <code>a<sub>i</sub></code>&nbsp;，需要先完成课程 <code>b<sub>i</sub></code>&nbsp;。</p>

<p>请根据给出的总课程数 &nbsp;<code>numCourses</code> 和表示先修顺序的&nbsp;<code>prerequisites</code>&nbsp;得出一个可行的修课序列。</p>

<p>可能会有多个正确的顺序，只要任意返回一种就可以了。如果不可能完成所有课程，返回一个空数组。</p>

<p>&nbsp;</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre>
<strong>输入:</strong> numCourses = 2, prerequisites = [[1,0]] 
<strong>输出: </strong><code>[0,1]</code>
<strong>解释:</strong>&nbsp;总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 <code>[0,1] 。</code></pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre>
<strong>输入:</strong> numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
<strong>输出: </strong><code>[0,1,2,3] or [0,2,1,3]</code>
<strong>解释:</strong>&nbsp;总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
&nbsp;因此，一个正确的课程顺序是&nbsp;<code>[0,1,2,3]</code> 。另一个正确的排序是&nbsp;<code>[0,2,1,3]</code> 。
</pre>

<p><strong>示例 3:</strong></p>

<pre>
<strong>输入:</strong> numCourses = 1, prerequisites = [] 
<strong>输出: </strong><code>[0]</code>
<strong>解释:</strong>&nbsp;总共 1 门课，直接修第一门课就可。</pre>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>

<ul>
	<li><code>1 &lt;= numCourses &lt;= 2000</code></li>
	<li><code>0 &lt;= prerequisites.length &lt;= numCourses * (numCourses - 1)</code></li>
	<li><code>prerequisites[i].length == 2</code></li>
	<li><code>0 &lt;= ai, bi &lt; numCourses</code></li>
	<li><code>ai != bi</code></li>
	<li><code>prerequisites</code>&nbsp;中不存在重复元素</li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 210&nbsp;题相同：<a href="https://leetcode-cn.com/problems/course-schedule-ii/">https://leetcode-cn.com/problems/course-schedule-ii/</a></p>

## 标签

 - Depth-First Search 
 - Breadth-First Search 
 - Graph 
 - Topological Sort 

## 代码

```java
class Solution {
    public int[] findOrder(int numCourses, int[][] prerequisites) {
        // 找入度为0 的节点
        ArrayList<Integer>[] afterThis = new ArrayList[numCourses];
        int[] beforeThis = new int[numCourses];

        // for(int i = 0;i<numCourses;++i){
        //    beforeThis[i] =  new ArrayList<>();
        // }
        for(int[] req: prerequisites){
            int a = req[0];
            int b = req[1];
            // b before a
            if( afterThis[b] ==null){
                 afterThis[b]=  new ArrayList<>();
            }
            afterThis[b].add(a);
            beforeThis[a]++;
            // System.out.println(Arrays.toString(afterThis)+""+Arrays.toString(beforeThis));
        }
        int[] result = new int[numCourses];
        for(int i = 0;i<numCourses;++i){
            // 找入度为0的点
            boolean found=false;
            for(int j= 0;j<numCourses;++j){
                if (beforeThis[j] == 0  ){
                    found = true;
                    result[i] = j;
                    // 把该点指向的所有边删掉，再把该节点删掉
                    if (afterThis[j]!=null){
                        for(int depThis : afterThis[j]){
                            beforeThis[depThis]--;
                        }
                    }
                  
                    beforeThis[j] = -1;
                    break;
                }
            }
            if (!found){
                // 有环
                return new int[0];
            }
        }
        return result;

    }
}
```