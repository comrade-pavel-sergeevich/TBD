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
WHERE sh.finished_at IS NULL
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