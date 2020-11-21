--
-- @lc app=leetcode.cn id=185 lang=mysql
--
-- [185] 部门工资前三高的所有员工
--

-- @lc code=start
# Write your MySQL query statement below
select B.Name as Department, A.Name as Employee, Salary
from
(select *, dense_rank() over (partition by DepartmentId order by Salary desc) as ranking
from Employee) as A, Department as B
where A.DepartmentId = B.Id and ranking <= 3
-- @lc code=end

