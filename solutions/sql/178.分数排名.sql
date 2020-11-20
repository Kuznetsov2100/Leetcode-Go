--
-- @lc app=leetcode.cn id=178 lang=mysql
--
-- [178] 分数排名
--

-- @lc code=start
# Write your MySQL query statement below

select Score, DENSE_RANK() over(order by Score desc) as `Rank`
from Scores
order by `Rank`
-- @lc code=end

