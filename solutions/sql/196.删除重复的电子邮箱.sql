--
-- @lc app=leetcode.cn id=196 lang=mysql
--
-- [196] 删除重复的电子邮箱
--

-- @lc code=start
# Write your MySQL query statement below
delete P1
from Person P1, Person P2
where P1.Email = P2.Email and P1.Id > P2.Id
-- @lc code=end

