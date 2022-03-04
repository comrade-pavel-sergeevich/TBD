1.
SELECT n_group, COUNT(n_group)
FROM student
GROUP BY n_group
2.
SELECT n_group, AVG(score)
FROM student
GROUP BY n_group
3.
SELECT surname, (COUNT (DISTINCT surname))
FROM student
GROUP BY surname
4.
SELECT 2022-age as year, (COUNT (age))
FROM student
GROUP BY age
5.
SELECT n_group/1000 as course, (AVG (score))
FROM student
GROUP BY n_group/1000
6.
SELECT n_group, (AVG (score))
FROM student
WHERE n_group/1000=3
GROUP BY n_group
ORDER BY avg DESC LIMIT 1
7.
SELECT n_group, avg
FROM (
SELECT n_group, AVG(score)
FROM student
GROUP BY n_group
) as st
WHERE avg<=3.5
ORDER BY avg DESC
8.
SELECT n_group,COUNT(*),MAX(score),AVG(score),MIN(score)
FROM student
GROUP BY n_group;
9.
select *
from student
where score = (select max(score) from student WHERE n_group=2288)
10.
select s.*
from student s
INNER JOIN (
select n_group, MAX(score)
from student
group by n_group) as st 
on st.n_group=s.n_group and s.score=st.max