--
-- @lc app=leetcode.cn id=180 lang=mysql
--
-- [180] 连续出现的数字
--

-- @lc code=start
# Write your MySQL query statement below
select distinct L1.Num as ConsecutiveNums
from Logs L1, Logs L2, Logs L3
where L1.Id = L2.Id-1 and L2.Id = L3.Id-1 and
L1.Num = L2.Num and L2.Num = L3.Num
-- @lc code=end

