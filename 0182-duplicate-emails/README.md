# 182.查找重复的电子邮箱

难度：`Easy`

 链接：[https://leetcode-cn.com/problems/duplicate-emails/](https://leetcode-cn.com/problems/duplicate-emails/)

## 题目描述

<p>编写一个 SQL 查询，查找&nbsp;<code>Person</code> 表中所有重复的电子邮箱。</p>

<p><strong>示例：</strong></p>

<pre>+----+---------+
| Id | Email   |
+----+---------+
| 1  | a@b.com |
| 2  | c@d.com |
| 3  | a@b.com |
+----+---------+
</pre>

<p>根据以上输入，你的查询应返回以下结果：</p>

<pre>+---------+
| Email   |
+---------+
| a@b.com |
+---------+
</pre>

<p><strong>说明：</strong>所有电子邮箱都是小写字母。</p>

## 标签

 - Database 

## 代码

```mysql
select distinct p1.Email  from Person p1,Person p2 where p1.Email=p2.Email and p1.id!=p2.id
```