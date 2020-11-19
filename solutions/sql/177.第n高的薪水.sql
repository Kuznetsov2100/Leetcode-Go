--
-- @lc app=leetcode.cn id=177 lang=mysql
--
-- [177] 第N高的薪水
--

-- @lc code=start
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  SET N := N-1;
  RETURN (
      # Write your MySQL query statement below.
    
    select ifnull((
    select distinct Salary
    from Employee
    order by Salary desc
    limit 1 offset N
), null)
  );
END
-- @lc code=end

