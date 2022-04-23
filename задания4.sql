1.
SELECT st.name, st.surname, h.name
FROM student st
INNER JOIN student_hobby sh ON st.id = sh.student_id
INNER JOIN hobby h ON h.id = sh.hobby_id
2.
SELECT (NOW ()-sh.started_at) as zanimaetsia, *
FROM student st
INNER JOIN student_hobby sh ON st.id = sh.student_id
INNER JOIN hobby h ON h.id = sh.hobby_id
WHERE sh.finished_at IS NULL
ORDER BY sh.started_at
LIMIT 1
3.
SELECT st.name, st.surname, st.id, st.age
FROM student st LEFT JOIN (
SELECT student_id, SUM(risk)
FROM student_hobby sh
JOIN hobby h on sh.hobby_id = h.id
GROUP BY student_id) as nt 
ON st.id=nt.student_id
WHERE st.score>= (select AVG(score) 
				  FROM student) and nt.sum>9
4.
SELECT st.surname, st.name, st.id, age, h.name hobby, 1/30.*extract(days from (sh.finished_at-sh.started_at)) as zanimalsia
FROM student st
INNER JOIN student_hobby sh ON st.id = sh.student_id
INNER JOIN hobby h ON h.id = sh.hobby_id
WHERE sh.finished_at IS NOT NULL
5.
SELECT st.surname, st.name, st.id, st.age
FROM student st Inner JOIN (
SELECT student_id, (count(sh.started_at)-count(sh.finished_at)) as cnt
FROM student_hobby sh
JOIN hobby h on sh.hobby_id = h.id
GROUP BY student_id) as nt 
ON st.id=nt.student_id
WHERE st.age>19 and nt.cnt>1
6.
SELECT st.n_group, AVG(st.score)
FROM student st INNER JOIN student_hobby sh on st.id=sh.student_id
WHERE sh.finished_at IS NULL and sh.started_at IS NOT NULL
GROUP BY st.n_group
7.
SELECT h.name as hobby, h.risk, 1/30.*extract( days from (NOW ()-sh.started_at)) as zanimaetsia, st.id
FROM student st
INNER JOIN student_hobby sh ON st.id = sh.student_id
INNER JOIN hobby h ON h.id = sh.hobby_id
WHERE sh.finished_at IS NULL and st.id=5
ORDER BY sh.started_at
LIMIT 1
8.
SELECT h.name as hobby
FROM hobby h
INNER JOIN  student_hobby sh on h.id=sh.hobby_id
INNER JOIN(
SELECT st.id
FROM student st
WHERE st.score=(select MAX(score) from student)) as nt on nt.id=sh.student_id
WHERE sh.finished_at IS NULL
9.
SELECT h.name as hobby
FROM hobby h
INNER JOIN
student_hobby sh on h.id = sh.hobby_id
INNER JOIN (SELECT st.id
FROM student st
WHERE ABS(st.score-3)<0.5 and n_group/1000=2)  st on st.id=sh.student_id
WHERE sh.finished_at IS NULL
10.
SELECT course_act.course
FROM
(SELECT course, COUNT(DISTINCT sh.student_id )as zanim
FROM student_hobby sh
INNER JOIN (
SELECT DISTINCT n_group/1000 as course, st.id
FROM student st) as nt on nt.id = sh.student_id
WHERE finished_at IS NULL
GROUP BY course) as course_act
RIGHT JOIN
(SELECT n_group/1000 as course, COUNT(*) from student st GROUP BY n_group/1000) as course_all 
on course_act.course = course_all.course
WHERE course_act.zanim*1./course_all.count>0.5
11.
SELECT n_group
FROM student
GROUP BY n_group
HAVING 1.* COUNT(CASE
    WHEN score>=4 THEN 1
    ELSE NULL
END)/COUNT(*)>=0.6
12.
SELECT n_group/1000 as course, COUNT(distinct h.name)
FROM student as st
RIGHT JOIN student_hobby as sh on sh.student_id=st.id
Left JOIN hobby as h on h.id=sh.hobby_id
GROUP BY n_group/1000
13.
SELECT st.id, st.surname, st.name, age, n_group/1000 as course
FROM student as st
LEFT JOIN student_hobby as sh on sh.student_id=st.id
Left JOIN hobby as h on h.id=sh.hobby_id
WHERE sh.hobby_id is NULL and score=5
ORDER BY course, age DESC
14.
CREATE OR REPLACE VIEW z14 AS
SELECT distinct st.id, st.name, st.surname,st.address,st.score,st.n_group,st.age FROM
student st RIGHT JOIN student_hobby sh on st.id=sh.student_id
WHERE finished_at is null and  1./365*extract(days from (clock_timestamp ( )-sh.started_at))>5
15.
SELECT h.name, COUNT(distinct (sh.student_id, sh.hobby_id))-count(*)+count(started_at) as count --> так хитро, потому что студент может заниматься одним хобби в двух кружках сразу, и надо убирать хобби, у которых нет started_at, а то они посчитаются
FROM hobby h LEFT JOIN student_hobby sh on h.id=sh.hobby_id
WHERE sh.finished_at is null
GROUP BY h.name
16.
SELECT id FROM(
SELECT h.id, COUNT(distinct (sh.student_id, sh.hobby_id))-count(*)+count(started_at) as count --> так хитро, потому что студент может заниматься одним хобби в двух кружках сразу, и надо убирать хобби, у которых нет started_at, а то они посчитаются
FROM hobby h LEFT JOIN student_hobby sh on h.id=sh.hobby_id
WHERE sh.finished_at is null
GROUP BY h.id
ORDER BY count desc limit 1) as foo
17.
SELECT * FROM student st RIGHT JOIN (
Select student_id from student_hobby sh
RIGHT JOIN(
SELECT h.id 
FROM hobby h LEFT JOIN student_hobby sh on h.id=sh.hobby_id
WHERE sh.finished_at is null
GROUP BY h.id
ORDER BY COUNT(distinct (sh.student_id, sh.hobby_id))-count(*)+count(started_at) desc limit 1) as besth on besth.id=sh.hobby_id
WHERE finished_at is null) as stid on st.id = stid.student_id
18.
SELECT id FROM hobby ORDER BY risk desc limit 3
19.
SELECT distinct st.id, st.name, st.surname,st.address,st.score,st.n_group,st.age, started_at FROM
student st RIGHT JOIN student_hobby sh on st.id=sh.student_id
WHERE finished_at is null
ORDER BY started_at limit 10
20.
SELECT distinct n_group from
(SELECT distinct st.id, st.name, st.surname,st.address,st.score,st.n_group,st.age, started_at FROM
student st RIGHT JOIN student_hobby sh on st.id=sh.student_id
WHERE finished_at is null
ORDER BY started_at limit 10) as t1
21.
CREATE OR REPLACE VIEW z21 AS
SELECT id, name, surname
FROM student
ORDER BY score desc
22.
CREATE OR REPLACE VIEW z22 AS
with tab1 as (SELECT course, h.name, COUNT(*) as cnt
FROM(
SELECT distinct course, student_id, hobby_id
FROM (SELECT st.id, st.n_group/1000 as course FROM student st) as st 
RIGHT JOIN student_hobby sh on st.id=sh.student_id
WHERE finished_at is null) as t1 LEFT JOIN hobby h on h.id=t1.hobby_id
GROUP BY course, name
ORDER BY cnt desc)

SELECT tab1.course, name
from tab1 RIGHT JOIN (select course, max(cnt) as max from tab1 group by course) as tab2 on tab1.course=tab2.course and tab1.cnt=tab2.max
23.
CREATE OR REPLACE VIEW z23 AS
with tab1 as
(
SELECT distinct hobby_id, count(*)
FROM student st right join student_hobby sh on st.id = sh.student_id
WHERE n_group/1000=2 and finished_at is null
group by hobby_id
ORDER BY count desc),
tab2 as(
SELECT name, risk
FROM hobby h RIGHT JOIN tab1 on h.id = tab1.hobby_id
WHERE count = (select count from tab1 limit 1)
ORDER BY risk desc)

SELECT name
FROM tab2
WHERE risk = (SELECT risk from tab2 limit 1)
24.
CREATE OR REPLACE VIEW z24 AS
SELECT n_group/1000 as course, count(*) as vsego, SUM(floor(score)::bigint/5) as otli4niki
FROM student
GROUP BY n_group/1000
25.
CREATE OR REPLACE VIEW z25 AS
SELECT name FROM(
SELECT h.name, COUNT(distinct (sh.student_id, sh.hobby_id))-count(*)+count(started_at) as count --> так хитро, потому что студент может заниматься одним хобби в двух кружках сразу, и надо убирать хобби, у которых нет started_at, а то они посчитаются
FROM hobby h LEFT JOIN student_hobby sh on h.id=sh.hobby_id
WHERE sh.finished_at is null
GROUP BY h.name
ORDER BY count desc limit 1) as foo
26.
CREATE VIEW obnovl as
SELECT * FROM hobby
--Пример запроса к ней:
INSERT INTO obnovl (name, risk)
VALUES ('плавание', 3)
27.
SELECT left(name,1), max(score), avg(score), min(score)
from student
group by left(name,1)
HAVING max(score)>3.6
28.
SELECT n_group/1000 as course, surname, max(score),min(score) 
from student
group by n_group/1000, surname
29.
SELECT age, COUNT(distinct hobby_id)
FROM student st RIGHT JOIN student_hobby sh on st.id = sh.student_id
GROUP BY age
30.
SELECT left(tab1.name,1), min(risk), max(risk)
FROM (student st RIGHT JOIN student_hobby sh on st.id = sh.student_id) as tab1 left join hobby h on h.id=tab1.student_id
GROUP BY left(tab1.name,1)
31.
--начиная с этого момента я, наконец, смирился с тем, что нужно создать поле "дата рождения", "возраста" не хватит
SELECT EXTRACT(MONTH from birth) as month, avg(score)
from (student st right join student_hobby sh on sh.student_id=st.id) as tab1 left join hobby h on h.id=tab1.hobby_id
WHERE finished_at is null and h.name='football'
GROUP BY month
32.
SELECT distinct 'Имя: '||name||', фамилия: '||surname||', группа: '||n_group
from student st right  join student_hobby sh on st.id=sh.student_id
33.
SELECT case position('ов' in surname)::varchar
when '0' then 'не найдено'
else position('ов' in surname)::varchar
end
from student
34.
SELECT case 
when (length(surname)>10) then surname
else rpad(surname,10,'#')
end
from student
35.
SELECT rtrim(fam,'#') from(
SELECT case 
when (length(surname)>10) then surname
else rpad(surname,10,'#')
end as fam
from student) as tab1
36.
SELECT extract(days FROM date_trunc('month', '4-1-2018'::date) + interval '1 month - 1 day');
37.
SELECT current_date+7-((cast(extract(dow from current_date) as int))+1)%8;
38.
Select extract(century from current_date) as century, to_char(current_date, 'IW') as weeknumber, extract(doy from current_date) as daynumber
39.
SELECT st.name, st.surname, h.name, case
when (sh.finished_at is null) then 'Закончил'
else 'Занимается' end
FROM student st right join student_hobby sh on st.id=sh.student_id left join hobby h on h.id = sh.hobby_id
40.
SELECT n_group, count(case round(score) when 2 then 1 end) as "2", count(case round(score) when 3 then 1 end) as "3", count(case round(score) when 4 then 1 end) as "4", count(case round(score) when 5 then 1 end) as "5"
from student
group by n_group
