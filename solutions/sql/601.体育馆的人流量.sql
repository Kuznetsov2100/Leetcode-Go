--
-- @lc app=leetcode.cn id=601 lang=mysql
--
-- [601] 体育馆的人流量
--

-- @lc code=start
# Write your MySQL query statement below
select id, visit_date, people
from
(
    select id, visit_date, 
    lag(people, 2) over (order by id) as lag2,
    lag(people, 1) over (order by id) as lag1,
    people,
    lead(people, 1) over (order by id) as lead1,
    lead(people, 2) over (order by id) as lead2
    from stadium
) as T
where (T.lag2 >= 100 and T.lag1 >= 100 and T.people >= 100) or
    (T.lag1 >= 100 and T.people >= 100 and T.lead1 >= 100) or
    (T.people >= 100 and T.lead1 >= 100 and T.lead2 >= 100)
-- @lc code=end

