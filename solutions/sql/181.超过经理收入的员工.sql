--
-- @lc app=leetcode.cn id=181 lang=mysql
--
-- [181] 超过经理收入的员工
--

-- @lc code=start
# Write your MySQL query statement below
# 自连接
select a.Name as Employee
from Employee as a, Employee as b
where a.ManagerId = b.Id and a.Salary > b.Salary
-- @lc code=end

