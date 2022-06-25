create TABLE `customer`(
    `customid`	varchar(17) NOT NULL, 
    `name`	varchar(10) NOT NULL,
    `sex`	varchar(2)  NOT NULL,
    `age`	int(4)              ,
    `xfg`	dec(10,2)   default 0,
    `address`	varchar(50) ,
    `memo`	varchar(100),
    PRIMARY KEY (customid)
) ENGINE=InnoDB;

create TABLE `student`(
`id` varchar(17) NOT NULL,
`name` varchar(10) NOT NULL,
`sex` varchar(2) ,
`age` int,
`score` numeric(6,2),
PRIMARY KEY (id)
) ENGINE=InnoDB;

insert into student(id, name, sex, age, score) values
('A0001','赵一','男',20,580.00),
('B0002','钱二','女',19,540.00),
('C0003','孙三','男',21,555.50),
('D0004','李四','男',22,480.00),
('E0005','周五','女',20,495.50),
('F0006','吴六','男',19,435.00);

create TABLE `aa`(
    `Aa1`	Varchar(20),
    `Aa2`	Int,
    `Aa3`	Dec(10,2)
) ENGINE=InnoDB;

create TABLE `bb`(
    `Bb1`	Varchar(30),
    `Bb2`	Int,
    `Bb3`	Dec(6,2)
) ENGINE=InnoDB;

drop TABLE aa; 
Alter table bb add  `Bb4` Varchar(20); 
create View Viewbb(Viewbb1, Viewbb2) as select Bb1, Bb4 from bb;
drop view Viewbb;
create Index Indexbb on bb (`Bb3` ASC);
drop Index Indexbb on bb;

create TABLE `cc` (
    `Cc1`	Varchar(20),
    `Cc2`	Int	,
    `Cc3`	Dec	(10,2), 
    `Cc4`	Varchar(60)
) ENGINE=InnoDB;

insert into `cc` (Cc1, Cc2, Cc3, Cc4) values
('赵一', 20, 580.00, '重邮宿舍12-3-5'),
('钱二', 19, 540.00, '南福苑5-2-9'),
('孙三', 21, 555.50, '学生新区21-5-15'),
('李四', 22, 480.00, '重邮宿舍8-2-22'),
('周五', 20, 495.50, '学生新区23-4-8'),
('吴六', 19, 435.00, '南福苑2-5-12');
update cc set Cc3=Cc3+5 where Cc2<=20;

delete from cc where Cc2>=20 and Cc3>=500;

create TABLE `Student`(
    `ID`	Varchar(20),
    `Name`	Varchar(10),
    `Age`	Int	,
    `Department`	Varchar(30)
)ENGINE=InnoDB;

create TABLE `Course`( 
    `CourseID`	Varchar(15),
    `CourseName`	Varchar(30),
    `CourseBefore`	Varchar(15)
) ENGINE=InnoDB;

create TABLE `Choose`(
    `ID`	Varchar(20),
    `CourseID`	Varchar(30),
    `Score`	Dec(5,2)
) ENGINE=InnoDB;

insert into `Student` (`ID`, `Name`, `Age`, `Department`) values
('00001','张三',20,'计算机系'),
('00002','李四',19,'计算机系'),
('00003','王五',21,'计算机系');

insert into `Course` (`CourseID`, `CourseName`, `CourseBefore`) values
('C1','计算机引论', NULL),
('C2','PASCAL语言','C1'),
('C3','数据结构', 'C2');

insert into `Choose` (`ID`, `CourseID`, `Score`) values
('00001','C1',95),
('00001','C2',80),
('00001','C3',84),
('00002','C1',80),
('00002','C2',85),
('00003','C1',78),
('00003','C3',70);

select ID, Name from Student where Department<=>"计算机系";

select Student.ID, Student.Name, Course.CourseName, Choose.Score
from Student, Choose, Course
where Student.ID = Choose.ID and Course.CourseID = Choose.CourseID
order by Student.ID;

select ID, Score from `Choose` where CourseID='C1' and ID in
(select ID from `Choose` where CourseID='C1' and Score<
(select Choose.Score from Choose, Student where Student.Name='张三' and Choose.CourseID='C1' and Choose.ID = Student.ID
));

select Student.ID from Student where ID in(select ID from Choose where Choose.CourseID = 'C3') 
and ID in(select ID from Choose where Choose.CourseID='C2');

create user 'DCL' identified by 'DCL';