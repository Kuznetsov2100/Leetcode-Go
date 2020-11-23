--
-- @lc app=leetcode.cn id=262 lang=mysql
--
-- [262] 行程和用户
--

-- @lc code=start
# Write your MySQL query statement below

select T.Request_at as `Day`, ROUND(SUM(IF(T.Status = 'completed', 0, 1)) / COUNT(*), 2) as `Cancellation Rate`
from Trips as T  join Users as U1 on (T.Client_Id = U1.Users_Id and U1.Banned = 'No')
                 join Users as U2 on (T.Driver_Id = U2.Users_Id and U2.Banned = 'No')
where T.Request_at between '2013-10-01' and '2013-10-03'
group by T.Request_at
-- @lc code=end

