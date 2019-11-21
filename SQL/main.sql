##175. 组合两个表
##links https://leetcode-cn.com/problems/combine-two-tables/

Create table Person (PersonId int, FirstName varchar(255), LastName varchar(255))
Create table Address (AddressId int, PersonId int, City varchar(255), State varchar(255))
Truncate table Person
insert into Person (PersonId, LastName, FirstName) values ('1', 'Wang', 'Allen')
Truncate table Address
insert into Address (AddressId, PersonId, City, State) values ('1', '2', 'New York City', 'New York')

##answer:
select FirstName, LastName, City, State from Person
left join Address on Person.PersonId = Address.PersonId


##176. 第二高的薪水
## links https://leetcode-cn.com/problems/second-highest-salary/
## 编写一个 SQL 查询, 获取 Employee 表中第二高的薪水(Salary)

Create table If Not Exists Employee (Id int, Salary int)
Truncate table Employee
insert into Employee (Id, Salary) values ('1', '100')
insert into Employee (Id, Salary) values ('2', '200')
insert into Employee (Id, Salary) values ('3', '300')

##answer:
##1
select IFNULL(
    (select Distinct Salary as SecondHighestSalary
        from Employee order by Salary Desc limit 1 offset 1), NULL)
        as SecondHighestSalary
##2
select
    (select Distinct Salary as SecondHighestSalary
        from Employee order by Salary Desc limit 1 offset 1)
        as SecondHighestSalary



##178. 分数排名
## links https://leetcode-cn.com/problems/rank-scores/
## 编写一个 SQL 查询来实现分数排名.如果两个分数相同,则两个分数排名(Rank)相同.请注意,平分后的下一个名次应该是下一个连续的整数值.换句话说,名次之间不应该有"间隔".

select count(distinct b.Score) from Scores b where b.Scores > X ## 计算出当前数据的排名
select a.Scores, (
    select count(dintinct b.Score) from Scores b where b.Scores > a.Scores
) as Rank
from Scores a order by a.Score Desc


##180. 连续出现的数字
## @links https://leetcode-cn.com/problems/consecutive-numbers/
## 编写一个 SQL 查询, 查找所有至少连续出现三次的数字.


Create table If Not Exists Logs (Id int, Num int)
Truncate table Logs
insert into Logs (Id, Num) values ('1', '1')
insert into Logs (Id, Num) values ('2', '1')
insert into Logs (Id, Num) values ('3', '1')
insert into Logs (Id, Num) values ('4', '2')
insert into Logs (Id, Num) values ('5', '1')
insert into Logs (Id, Num) values ('6', '2')
insert into Logs (Id, Num) values ('7', '2')

##answer
select distinct l1.num as ConsecutiveNums
from Logs l1, Logs l2, Logs l3
where
    l1.Id = l2.Id - 1
and l1.Id = l3.Id - 2
and l1.Num = l2.Num
and l1.Num = l3.Num



##181. 超过经理收入的员工
## @links https://leetcode-cn.com/problems/employees-earning-more-than-their-managers/submissions/
## Employee 表包含所有员工, 他们的经理也属于员工. 每个员工都有一个 Id, 此外还有一列对应员工的经理的 Id


select e1.Name as Employee from Employee as e1
join Employee as e2 on
e1.ManagerId = e2.Id
and a.Salary > b.Salary

##184. 部门工资最高的员工
## @links https://leetcode-cn.com/problems/department-highest-salary/


select Department.Name as Department, Employee.Name Employee, Salary
from Employee inner join Department
on Employee.DepartmentId = Department.Id
where (DepartmentId, Salary) in ( ##算出各部门的最高薪水, 多个where条件要用于自查询, 需要用括号包裹
    select DepartmentId, max(Salary) as Salary from Employee
    group by DepartmentId
)




##185. 部门工资前三高的所有员工
##links https://leetcode-cn.com/problems/department-top-three-salaries/solution/bu-men-gong-zi-qian-san-gao-de-yuan-gong-by-leetco/
select d.Name AS 'Department', e1.Name AS 'Employee', e1.Salary
    from Employee e1
        JOIN Department d ON e1.DepartmentId = d.Id
where 3 > (
    select count(distinct e2.Salary) from Employee e2
        where e2.Salary > e1.Salary
        and e1.DepartmentId = e2.DepartmentId
)