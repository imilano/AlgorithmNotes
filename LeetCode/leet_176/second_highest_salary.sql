# Write your MySQL query statement below
SELECT max(Salary) AS SecondHighestSalary FROM Employee WHERE Salary < (SELECT max(Salary) FROM Employee);

# LIMIT N-1来选出Nth个salary
SELECT
    (SELECT DISTINCT
         Salary
     FROM
         Employee
     ORDER BY Salary DESC
        LIMIT 1 OFFSET 1) AS SecondHighestSalary
;