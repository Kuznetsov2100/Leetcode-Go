--
-- @lc app=leetcode.cn id=183 lang=mysql
--
-- [183] 从不订购的客户
--

-- @lc code=start
# Write your MySQL query statement below
select Name as Customers
from Customers left join Orders
on Customers.Id = Orders.CustomerId
where Orders.CustomerId is null
-- @lc code=end

