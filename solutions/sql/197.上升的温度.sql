--
-- @lc app=leetcode.cn id=197 lang=mysql
--
-- [197] 上升的温度
--

-- @lc code=start
# Write your MySQL query statement below
select W1.id
from Weather W1, Weather W2
where datediff(W1.recordDate, W2.recordDate) = 1 and
    W1.Temperature > W2.Temperature
-- @lc code=end

