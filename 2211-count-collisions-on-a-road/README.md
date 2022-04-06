# 2211.统计道路上的碰撞次数

难度：`Medium`

 链接：[https://leetcode-cn.com/problems/count-collisions-on-a-road/](https://leetcode-cn.com/problems/count-collisions-on-a-road/)

## 题目描述

<p>在一条无限长的公路上有 <code>n</code> 辆汽车正在行驶。汽车按从左到右的顺序按从 <code>0</code> 到 <code>n - 1</code> 编号，每辆车都在一个 <strong>独特的</strong> 位置。</p>

<p>给你一个下标从 <strong>0</strong> 开始的字符串 <code>directions</code> ，长度为 <code>n</code> 。<code>directions[i]</code> 可以是 <code>'L'</code>、<code>'R'</code> 或 <code>'S'</code> 分别表示第 <code>i</code> 辆车是向 <strong>左</strong> 、向 <strong>右</strong> 或者 <strong>停留</strong> 在当前位置。每辆车移动时 <strong>速度相同</strong> 。</p>

<p>碰撞次数可以按下述方式计算：</p>

<ul>
	<li>当两辆移动方向&nbsp;<strong>相反</strong>&nbsp;的车相撞时，碰撞次数加 <code>2</code> 。</li>
	<li>当一辆移动的车和一辆静止的车相撞时，碰撞次数加 <code>1</code> 。</li>
</ul>

<p>碰撞发生后，涉及的车辆将无法继续移动并停留在碰撞位置。除此之外，汽车不能改变它们的状态或移动方向。</p>

<p>返回在这条道路上发生的 <strong>碰撞总次数</strong> 。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>directions = "RLRSLL"
<strong>输出：</strong>5
<strong>解释：</strong>
将会在道路上发生的碰撞列出如下：
- 车 0 和车 1 会互相碰撞。由于它们按相反方向移动，碰撞数量变为 0 + 2 = 2 。
- 车 2 和车 3 会互相碰撞。由于 3 是静止的，碰撞数量变为 2 + 1 = 3 。
- 车 3 和车 4 会互相碰撞。由于 3 是静止的，碰撞数量变为 3 + 1 = 4 。
- 车 4 和车 5 会互相碰撞。在车 4 和车 3 碰撞之后，车 4 会待在碰撞位置，接着和车 5 碰撞。碰撞数量变为 4 + 1 = 5 。
因此，将会在道路上发生的碰撞总次数是 5 。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>directions = "LLRR"
<strong>输出：</strong>0
<strong>解释：</strong>
不存在会发生碰撞的车辆。因此，将会在道路上发生的碰撞总次数是 0 。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= directions.length &lt;= 10<sup>5</sup></code></li>
	<li><code>directions[i]</code> 的值为 <code>'L'</code>、<code>'R'</code> 或 <code>'S'</code></li>
</ul>

## 标签

 - Stack 
 - String 

## 代码

```golang
func countCollisions(directions string) int {
    // stack := []int{}
    cnt :=0
    haveStay := 0
    haveGoRight := 0
    for _,dir:=range directions{

        if dir == 'L'{
            // 向左，判断一下栈是不是有，有就会撞
            if haveGoRight>0 {
                // 碰撞一次后停止第一次是+2，左边向右的也开始碰撞, cnt= 2 + haveGoRight-1 
                cnt+=haveGoRight+1
                haveGoRight = 0
                haveStay = 1
            }else if haveStay>0 {
                cnt++
            }
        }else if dir == 'R'{
            
            haveGoRight++
            
        }else if dir == 'S'{
            if haveGoRight>0 {
                // 碰撞一次后停止，左边向右的也开始碰撞, cnt= 1 + haveGoRight-1 
                cnt+=haveGoRight
                haveGoRight = 0
            }
            haveStay = 1
        }
    }
    return cnt

}
```