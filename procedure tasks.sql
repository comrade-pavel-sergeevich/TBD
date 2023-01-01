1.
DO $$ 
BEGIN
raise notice 'Heil';
END $$;

2.
DO $$ 
DECLARE
now TIMESTAMP WITH TIME ZONE;
curs2 CURSOR FOR Select now() AT TIME ZONE 'EUROPE/MOSCOW' as timenow;
BEGIN
SET TIMEZONE='EUROPE/MOSCOW';
OPEN curs2;
FETCH curs2 into now;
raise notice '%',now;
END $$;

3.
DO $$ 
DECLARE
ch1 double precision;
ch2 double precision;
BEGIN
ch1=228;ch2=1337;
ch1=ch1-77;
ch2=(ch2*(ch1+9850));
raise notice '%',ch1+ch2;
END $$;

4.
DO $$ 
DECLARE
ch1 integer;
BEGIN
ch1= random()*10;
if ch1=2 THEN raise notice 'Неуд.'; else
if ch1=3 THEN raise notice 'Уд.'; else
if ch1=4 THEN raise notice 'Хор.'; else
if ch1=5 THEN raise notice 'Отл.'; else
raise notice 'Ошибка: неправильная оценка (%)',ch1; end if; end if; end if; end if;
CASE ch1 
WHEN 2 THEN raise notice 'Неуд.';
WHEN 3 THEN raise notice 'Уд.';
WHEN 4 THEN raise notice 'Хор.';
WHEN 5 THEN raise notice 'Отл.';
ELSE raise notice 'Ошибка: неправильная оценка (%)',ch1;
END CASE;
END $$;

5.
DO $$
declare
   n integer:= 20;
BEGIN

for i in 20..30 loop
	raise notice '%^2 = : %', i,i^2;
end loop;

loop
	raise notice '%^2 = : %', n,n^2;
	n=n+1;
	exit when n=31;
end loop;

n=20;

WHILE n<=30 loop
	raise notice '%^2 = : %', n,n^2;
	n=n+1;
end loop;

END $$;

6.
create or replace function col(ch1 integer)
   returns integer 
   language plpgsql
  as
$$
DECLARE
i integer;
BEGIN
i=0;
while ch1<>1 loop
i=i+1;
 if ch1%2 then ch1=ch1*3+1; else ch1=ch1/2; end if; 
 end loop;
RETURN i;
end;
$$;

create or replace procedure coln(ch1 integer)
   language plpgsql
  as
$$
DECLARE
i integer;
BEGIN
i=0;
raise notice '%',ch1;
while ch1<>1 loop
i=i+1;
 if ch1%2 then ch1=ch1*3+1; else ch1=ch1/2; end if; 
 raise notice '%',ch1;
 end loop;
end;
$$;

DO $$
DECLARE
ch1 integer;
i integer;
BEGIN
ch1= random()*1000+1;
select col(ch1) into i;
CALL coln(ch1);
raise notice 'Число шагов составило % шт.',i;
END $$

7.
create or replace function luk(ch1 integer)
   returns integer 
   language plpgsql
  as
$$
DECLARE
i integer;
aa integer;
bb integer;
BEGIN
if ch1=1 then return 2; end if;
if ch1=2 then return 1; end if;
i=2;
aa=2;
bb=1;
while i<>ch1 loop
i=i+1;
 bb = aa + bb;
 aa = bb - aa;
 end loop;
RETURN bb;
end;
$$;

create or replace procedure lukn(ch1 integer)
   language plpgsql
  as
$$

DECLARE
i integer;
aa integer;
bb integer;
BEGIN
if ch1=1 then
	 raise notice '2';
	 return;
else 
raise notice '2';
end if;

if ch1=2 then
	 raise notice '1';
	 return;
else  raise notice '1';
end if;

i=2;
aa=2;
bb=1;
while i<>ch1 loop
i=i+1;
 bb = aa + bb;
 aa = bb - aa;
 raise notice '%',bb;
 end loop;
end;
$$;

DO $$
DECLARE
ch1 integer;
i integer;
BEGIN
ch1= random()*20+1;
select luk(ch1) into i;
CALL lukn(ch1);
raise notice '%-е число %.',ch1,i;
END $$

8.
create or replace function rodvgod(god integer, out i integer)
   returns integer 
   language plpgsql
  as
$$
DECLARE
curs2 CURSOR (year integer) for select count(*) from people where extract(years from birth_date)=year;
BEGIN

open curs2 (year:=god);
fetch curs2 into i;
RETURN;
end;
$$;


DO $$
DECLARE
god integer;
i integer;
BEGIN
god = 1985+random()*20;
Select rodvgod(god) into i;
raise notice '% человек родились в % году.',i,god;
END $$

9.
create or replace function tsvetdivanov(glaz varchar, out i integer)
   returns integer 
   language plpgsql
  as
$$
DECLARE
curs2 CURSOR (eyess varchar) for select count(*) from people where eyes=eyess;
BEGIN

open curs2 (eyess:=glaz);
fetch curs2 into i;
RETURN;
end;
$$;


DO $$
DECLARE
glaza varchar[];
glazid integer;
i integer;
BEGIN
glaza[0]='blue';
glaza[1]='brown';
glazid = random()*2;
Select tsvetdivanov(glaza[glazid]) into i;
raise notice '% человек с % глазами.',i,glaza[glazid];
END $$

10.
create or replace function youngest(out i integer)
   returns integer 
   language plpgsql
  as
$$
DECLARE
curs2 CURSOR for select id from people order by birth_date desc limit 1;
BEGIN

open curs2;
fetch curs2 into i;
RETURN;
end;
$$;


DO $$
DECLARE
i integer;
BEGIN

Select youngest() into i;
raise notice '% id yougest человека.',i;
END $$

11.
create or replace function zhirik(imt double precision)
   returns table(
   				user_id integer,
	   			user_name varchar,
	   			surname varchar
   ) 
   language plpgsql
  as
$$
BEGIN

RETURN query
select people.id user_id,name user_name,people.surname surname from people where weight/(growth*growth)*10000>imt;
end;
$$;

SELECT * FROM (Select (15+random()*15) imt) ddd, zhirik(ddd.imt);

12.
BEGIN;

CREATE TABLE relationships(
	id int PRIMARY KEY,
    person1 INT NOT NULL REFERENCES people(id),
   	person2 INT NOT NULL REFERENCES people(id),
    whoissecond VARCHAR(256) NOT NULL
);

INSERT INTO relationships
VALUES (,3, 4, 'wife'), (,4, 3, 'husband'), (,2, 6, 'brother'), (,6, 2, 'brother');

COMMIT;

13.
create or replace procedure addloh(
	name varchar,
	surname varchar,
	birth_date DATE,
	growth real,
	weight real,
	eyes varchar,
	hair varchar,
	relations relationships[])
   language plpgsql
  as
$$
DECLARE
 tmp relationships;
BEGIN
INSERT INTO people(name, surname,birth_date,growth,weight,eyes,hair) VALUES(name, surname,birth_date,growth,weight,eyes,hair);
FOR tmp IN SELECT * FROM unnest(relations) LOOP
        INSERT INTO relationships(person1, person2, whoissecond)
        VALUES ((select last_value from people_id_seq), tmp.person2, tmp.whoissecond),
		(tmp.person2,(select last_value from people_id_seq),  ('person 1 is '||tmp.whoissecond||' for person 2'));
    END LOOP;
end;
$$;
CALL  addloh('Sergay','Dorogov', '12-05-1959', 172, 92,'gray','black',
    ARRAY['(,,6,"brother-in-law")', '(,,4,"sister")']::relationships[]);

14.
BEGIN;
ALTER TABLE people
ADD COLUMN data_proverki_otk TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP;
COMMIT;

15.
CREATE OR REPLACE PROCEDURE zhirniyidlinniy(person_id INT, rost REAL, nagiev REAL)
	language plpgsql
AS
$$
BEGIN
SET TIMEZONE='EUROPE/MOSCOW';
    UPDATE people
    SET growth = rost, weight = nagiev, data_proverki_otk = CURRENT_TIMESTAMP
    WHERE id = person_id;
	
end; 
$$;

CALL zhirniyidlinniy(4, 193, 88);