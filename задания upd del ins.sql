1.
delete from student_hobby where student_id in (select id from student where date_birth is null);
delete  from student where date_birth is null;
2.
update student set date_birth='1-1-99' where date_birth is null
3.
delete from student where id=23
4.
update hobby set risk=risk-1 where id=(
select hobby_id
from student_hobby
where date_finish is null
group by hobby_id
order by count(*) desc limit 1)
5.
update student set score = score+0.01 where id in(
select distinct student_id from student_hobby where date_finish is null)
6.
delete from student_hobby where date_finish is not null
7.
insert into student_hobby (student_id, hobby_id, date_start)
values (4,5,'15-nov-2009')
8.
delete from student_hobby where id in
(select min(id) from student_hobby where ('('||hobby_id::varchar||','||student_id::varchar||',)') in(
select distinct (hobby_id, student_id, date_finish)::varchar as aa
from student_hobby
group by hobby_id, student_id, date_finish
order by aa) and date_finish is not null
group by student_id, hobby_id)
9.
UPDATE hobby SET name = CASE
   WHEN (name = 'Футбол') THEN 'Бальные танцы'
   WHEN (name = 'Баскетбол') THEN 'Вышивание крестиком'
END
WHERE name in('Футбол','Баскетбол')
10.
insert into hobby (name, risk) values('Учёба',0)
11.
with zan as((select distinct student_id, dd.id from student_hobby  cross join  (select id from hobby where name = 'Учёба') as dd where date_finish is null))
update student_hobby set hobby_id = (select id from hobby where name = 'Учёба')
where date_finish is null and student_id in (select student_id from zan) and student_id in(select student_id from student where score<3.2);
with zan as((select distinct student.id as stid, dd.id from student  cross join  (select id from hobby where name = 'Учёба') as dd where student.id not in (select student_id from student_hobby where date_finish is null)))
insert into student_hobby(student_id,hobby_id,date_start)
select stid, id, now() from zan;
12.
update student set n_group=n_group+1000 where (n_group/1000=4)
13.
delete from student_hobby where student_id=2;
delete from student where id=2
14.
update student set score =5 where id in
(select student_id from student_hobby where date_finish is null and (extract(days FROM(now()-date_start))>10))
15.
delete from student_hobby
where id in 
(select sh.id 
from student st right join student_hobby sh on st.id=sh.student_id 
where date_birth>date_start)