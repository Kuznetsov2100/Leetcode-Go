--
-- @lc app=leetcode.cn id=626 lang=mysql
--
-- [626] 换座位
--

-- @lc code=start
# Write your MySQL query statement below
select
    (case
        when mod(id, 2) != 0 and id != counts then id+1
        when mod(id, 2) != 0 and id = counts then id
        else id-1
    end) as id, student
from seat, (select count(*) as counts from seat) as seat_count
order by id asc



-- select if(id%2=0,
--          id-1,
--          if(id=(select count(distinct id) from seat),
--          id,
--          id+1)) as id, student
-- from seat
-- order by id
-- @lc code=end

