# 剑指 Offer II 037.小行星碰撞

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/XagZNi/](https://leetcode-cn.com/problems/XagZNi/)

## 题目描述

<p>给定一个整数数组 <code>asteroids</code>，表示在同一行的小行星。</p>

<p>对于数组中的每一个元素，其绝对值表示小行星的大小，正负表示小行星的移动方向（正表示向右移动，负表示向左移动）。每一颗小行星以相同的速度移动。</p>

<p>找出碰撞后剩下的所有小行星。碰撞规则：两个行星相互碰撞，较小的行星会爆炸。如果两颗行星大小相同，则两颗行星都会爆炸。两颗移动方向相同的行星，永远不会发生碰撞。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>asteroids = [5,10,-5]
<strong>输出：</strong>[5,10]
<b>解释：</b>10 和 -5 碰撞后只剩下 10 。 5 和 10 永远不会发生碰撞。</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>asteroids = [8,-8]
<strong>输出：</strong>[]
<b>解释：</b>8 和 -8 碰撞后，两者都发生爆炸。</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>asteroids = [10,2,-5]
<strong>输出：</strong>[10]
<b>解释：</b>2 和 -5 发生碰撞后剩下 -5 。10 和 -5 发生碰撞后剩下 10 。</pre>

<p><strong>示例 4：</strong></p>

<pre>
<strong>输入：</strong>asteroids = [-2,-1,1,2]
<strong>输出：</strong>[-2,-1,1,2]
<b>解释</b><strong>：</strong>-2 和 -1 向左移动，而 1 和 2 向右移动。 由于移动方向相同的行星不会发生碰撞，所以最终没有行星发生碰撞。 </pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= asteroids.length&nbsp;&lt;= 10<sup>4</sup></code></li>
	<li><code>-1000 &lt;= asteroids[i] &lt;= 1000</code></li>
	<li><code>asteroids[i] != 0</code></li>
</ul>

<p>&nbsp;</p>

<p><meta charset="UTF-8" />注意：本题与主站 735&nbsp;题相同：&nbsp;<a href="https://leetcode-cn.com/problems/asteroid-collision/">https://leetcode-cn.com/problems/asteroid-collision/</a></p>

## 标签

 - Stack 
 - Array 

## 代码

```golang
func asteroidCollision(asteroids []int) []int {
    // 1. 符号相同的不会相遇
    // 2. 相遇后绝对值大的留下，相等的同时消失
    // 考虑用双链表模拟，先构建双链表，然后从做到右找第一个异号的数字开始碰
    // 碰了以后将消失的删掉，然后留下的在和左边的比符号，如果左边为空则和右边比符号，重复上述步骤，直到没有异号
    // 注意最左边的负号将没得碰，因此直接跳掉
    // -- 这个过程可以直接用一个栈来简化
    stack:= make([]int,0,len(asteroids))
    for _,star := range asteroids{
        if len(stack) == 0 || stack[len(stack)-1] < 0||star >0{
            // 左边是空或者全是符号的,或者当前行星是往右的，直接进栈
            stack=append(stack,star)
        }else if star*stack[len(stack)-1] <0{
            // fmt.Println("开始撞",stack,star)
            // 就一直撞直到：1. 左边没了，2.右边的撞没 3.左边的是往左走的
            for len(stack) != 0 && star!=0 && stack[len(stack)-1]>0{
                // 异号，且star<0 ,开始撞
                if -star == stack[len(stack)-1] {
                    // 同归于尽
                    stack=stack[:len(stack)-1]
                    // 消除star
                    star = 0 
                }else if -star < stack[len(stack)-1]{
                    // 右边消失
                    star = 0 
                    break
                }else{
                    // 左边消失，需要将左边的提上来再比
                    stack=stack[:len(stack)-1]
                }
            }
            // fmt.Println("结果",stack,star)
            if star != 0{
                stack=append(stack,star)
            }

        }
    }
    return stack

}
```