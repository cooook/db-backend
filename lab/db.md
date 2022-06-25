# 数据库lab
实验环境
MacOS Monterey Apple Silicon

## lab1
实验内容： 安装mysql server
运行如下内容
```bash 
brew install mysql 
```
安装之后查看mysql是否安装成功
```bash
➜  backend git:(master) ✗ mysql --version
mysql  Ver 8.0.28 for macos12.0 on arm64 (Homebrew)
```
如上所示安装mysql 8.0.28版本

## lab2
实验内容要求创建一个名为demo的数据库，并且创建customer
创建指令如下
```SQL
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
```
查看数据库是否创建成功
```SQL
mysql> use demo;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> desc customer;
+----------+---------------+------+-----+---------+-------+
| Field    | Type          | Null | Key | Default | Extra |
+----------+---------------+------+-----+---------+-------+
| customid | varchar(17)   | NO   | PRI | NULL    |       |
| name     | varchar(10)   | NO   |     | NULL    |       |
| sex      | varchar(2)    | NO   |     | NULL    |       |
| age      | int           | YES  |     | NULL    |       |
| xfg      | decimal(10,2) | YES  |     | 0.00    |       |
| address  | varchar(50)   | YES  |     | NULL    |       |
| memo     | varchar(100)  | YES  |     | NULL    |       |
+----------+---------------+------+-----+---------+-------+
7 rows in set (0.02 sec)
```
可以看到表和数据库已经被成功创建

## lab3 
### 任务1
建立student表，这里直接建立在demo数据库里
```SQL
create TABLE `student`(
`id` varchar(17) NOT NULL,
`name` varchar(10) NOT NULL,
`sex` varchar(2) ,
`age` int,
`score` numeric(6,2),
PRIMARY KEY (id)
) ENGINE=InnoDB;
```
### 任务2
向刚创建的student表中插入数据
指令如下
```SQL
insert into student(id, name, sex, age, score) values
('A0001','赵一','男',20,580.00),
('B0002','钱二','女',19,540.00),
('C0003','孙三','男',21,555.50),
('D0004','李四','男',22,480.00),
('E0005','周五','女',20,495.50),
('F0006','吴六','男',19,435.00);
```
### 任务3
练习查询语句，查找年龄大于等于20岁，成绩低于500分的记录。
指令如下
```SQL
mysql> select * from student where age >= 20 and score < 500;
+-------+--------+------+------+--------+
| id    | name   | sex  | age  | score  |
+-------+--------+------+------+--------+
| D0004 | 李四   | 男   |   22 | 480.00 |
| E0005 | 周五   | 女   |   20 | 495.50 |
+-------+--------+------+------+--------+
2 rows in set (0.01 sec)
```

## lab4 
### 任务1
创建aa表
```SQL
create TABLE `aa`(
    `Aa1`	Varchar(20),
    `Aa2`	Int,
    `Aa3`	Dec(10,2)
) ENGINE=InnoDB;
```
### 任务2
创建bb表
```SQL
create TABLE `bb`(
    `Bb1`	Varchar(30),
    `Bb2`	Int,
    `Bb3`	Dec(6,2)
) ENGINE=InnoDB;
```
### 任务3
删除创建的aa表
```SQL
drop table aa; 
```
### 任务4
修改表bb，添加一个字段Bb4，类型Varchar,长度20
修改前bb
```SQL
mysql> desc bb;
+-------+--------------+------+-----+---------+-------+
| Field | Type         | Null | Key | Default | Extra |
+-------+--------------+------+-----+---------+-------+
| Bb1   | varchar(30)  | YES  |     | NULL    |       |
| Bb2   | int          | YES  |     | NULL    |       |
| Bb3   | decimal(6,2) | YES  |     | NULL    |       |
+-------+--------------+------+-----+---------+-------+
3 rows in set (0.01 sec)
```
修改指令
```SQL
mysql> alter table bb add `Bb4` varchar(20);
Query OK, 0 rows affected (0.02 sec)
Records: 0  Duplicates: 0  Warnings: 0
```
查看修改后bb的结构
```SQL
mysql> desc bb;
+-------+--------------+------+-----+---------+-------+
| Field | Type         | Null | Key | Default | Extra |
+-------+--------------+------+-----+---------+-------+
| Bb1   | varchar(30)  | YES  |     | NULL    |       |
| Bb2   | int          | YES  |     | NULL    |       |
| Bb3   | decimal(6,2) | YES  |     | NULL    |       |
| Bb4   | varchar(20)  | YES  |     | NULL    |       |
+-------+--------------+------+-----+---------+-------+
4 rows in set (0.01 sec)
```
### 任务5
用Create View对表Bb的Bb1和Bb4建立一个视图Viewbb,字段名为Viewbb1和Viewbb2。
```SQL
create view Viewbb(Viewbb1, Viewbb2) as select Bb1, Bb4 from Bb;
```
展示如下
```SQL
mysql> desc Viewbb;
+---------+-------------+------+-----+---------+-------+
| Field   | Type        | Null | Key | Default | Extra |
+---------+-------------+------+-----+---------+-------+
| Viewbb1 | varchar(30) | YES  |     | NULL    |       |
| Viewbb2 | varchar(20) | YES  |     | NULL    |       |
+---------+-------------+------+-----+---------+-------+
2 rows in set (0.01 sec)
```
### 任务6
删除刚才建立的视图
```SQL
mysql> drop view Viewbb; 
Query OK, 0 rows affected (0.01 sec)
```
### 任务7
用Create Index对表Bb的Bb3字段建立一个升序索引，索引名Indexbb。
```SQL
Create Index Indexbb on bb(`Bb3` ASC);
```
### 任务8
删除索引
```SQL
drop Index Indexbb on bb;
```

## lab5 
### 任务1
创建cc表
```SQL
create TABLE `cc` (
    `Cc1`	Varchar(20),
    `Cc2`	Int	,
    `Cc3`	Dec	(10,2), 
    `Cc4`	Varchar(60)
) ENGINE=InnoDB;
```
### 任务2
向cc插入数据
```SQL 
insert into `cc` (Cc1, Cc2, Cc3, Cc4) values
('赵一', 20, 580.00, '重邮宿舍12-3-5'),
('钱二', 19, 540.00, '南福苑5-2-9'),
('孙三', 21, 555.50, '学生新区21-5-15'),
('李四', 22, 480.00, '重邮宿舍8-2-22'),
('周五', 20, 495.50, '学生新区23-4-8'),
('吴六', 19, 435.00, '南福苑2-5-12');
```
### 任务3
UPDATE数据
更新前
```SQL
mysql> select * from cc;
+--------+------+--------+---------------------+
| Cc1    | Cc2  | Cc3    | Cc4                 |
+--------+------+--------+---------------------+
| 赵一   |   20 | 580.00 | 重邮宿舍12-3-5      |
| 钱二   |   19 | 540.00 | 南福苑5-2-9         |
| 孙三   |   21 | 555.50 | 学生新区21-5-15     |
| 李四   |   22 | 480.00 | 重邮宿舍8-2-22      |
| 周五   |   20 | 495.50 | 学生新区23-4-8      |
| 吴六   |   19 | 435.00 | 南福苑2-5-12        |
+--------+------+--------+---------------------+
6 rows in set (0.00 sec)
```
```SQL
update cc set cc3 = cc3 + 5 where cc2 <= 20;
```
更新后
```SQL
mysql> update cc set cc3 = cc3 + 5 where cc2 <= 20;
Query OK, 4 rows affected (0.00 sec)
Rows matched: 4  Changed: 4  Warnings: 0

mysql> select * from cc;
+--------+------+--------+---------------------+
| Cc1    | Cc2  | Cc3    | Cc4                 |
+--------+------+--------+---------------------+
| 赵一   |   20 | 585.00 | 重邮宿舍12-3-5      |
| 钱二   |   19 | 545.00 | 南福苑5-2-9         |
| 孙三   |   21 | 555.50 | 学生新区21-5-15     |
| 李四   |   22 | 480.00 | 重邮宿舍8-2-22      |
| 周五   |   20 | 500.50 | 学生新区23-4-8      |
| 吴六   |   19 | 440.00 | 南福苑2-5-12        |
+--------+------+--------+---------------------+
6 rows in set (0.00 sec)
```
### 任务4
删除指定的条目
```SQL
mysql> delete from cc where cc2 >= 20 and cc3 >= 500;
Query OK, 3 rows affected (0.01 sec)

mysql> select * from cc;
+--------+------+--------+--------------------+
| Cc1    | Cc2  | Cc3    | Cc4                |
+--------+------+--------+--------------------+
| 钱二   |   19 | 545.00 | 南福苑5-2-9        |
| 李四   |   22 | 480.00 | 重邮宿舍8-2-22     |
| 吴六   |   19 | 440.00 | 南福苑2-5-12       |
+--------+------+--------+--------------------+
3 rows in set (0.00 sec)
```
## lab6
### 任务1
新建表Student
```SQL
create TABLE `Student`(
    `ID`	Varchar(20),
    `Name`	Varchar(10),
    `Age`	Int	,
    `Department`	Varchar(30)
)ENGINE=InnoDB;
```
### 任务2
新建Course表
```SQL
create TABLE `Course`( 
    `CourseID`	Varchar(15),
    `CourseName`	Varchar(30),
    `CourseBefore`	Varchar(15)
) ENGINE=InnoDB;
```
### 任务3
新建Choose表
```SQL
create TABLE `Choose`(
    `ID`	Varchar(20),
    `CourseID`	Varchar(30),
    `Score`	Dec(5,2)
) ENGINE=InnoDB;
```
### 任务4
插入Student表
```SQL
insert into `Student` (`ID`, `Name`, `Age`, `Department`) values
('00001','张三',20,'计算机系'),
('00002','李四',19,'计算机系'),
('00003','王五',21,'计算机系');
```
### 任务5
插入Course表
```SQL
insert into `Course` (`CourseID`, `CourseName`, `CourseBefore`) values
('C1','计算机引论', NULL),
('C2','PASCAL语言','C1'),
('C3','数据结构', 'C2');
```
### 任务6
插入Choose表
```SQL
insert into `Choose` (`ID`, `CourseID`, `Score`) values
('00001','C1',95),
('00001','C2',80),
('00001','C3',84),
('00002','C1',80),
('00002','C2',85),
('00003','C1',78),
('00003','C3',70);
```
### 任务7
用SELECT语句求计算机系学生的学号和姓名。
```SQL
mysql> select ID, Name from student where Department = '计算机系';
+-------+--------+
| ID    | Name   |
+-------+--------+
| 00001 | 张三   |
| 00002 | 李四   |
| 00003 | 王五   |
+-------+--------+
3 rows in set (0.01 sec)
```
### 任务8
用SELECT语句求学生的学号、姓名、选的课程名及成绩。
```SQL
mysql> select Student.ID, Student.Name, Course.CourseName, Choose.Score from Choose, Student, Course where Choose.ID = Student.ID and Course.CourseID = Choose.CourseID;
+-------+--------+-----------------+-------+
| ID    | Name   | CourseName      | Score |
+-------+--------+-----------------+-------+
| 00003 | 王五    | 计算机引论        | 78.00 |
| 00002 | 李四    | 计算机引论        | 80.00 |
| 00001 | 张三    | 计算机引论        | 95.00 |
| 00002 | 李四    | PASCAL语言       | 85.00 |
| 00001 | 张三    | PASCAL语言       | 80.00 |
| 00003 | 王五    | 数据结构          | 70.00 |
| 00001 | 张三    | 数据结构          | 84.00 |
+-------+--------+-----------------+-------+
7 rows in set (0.01 sec)
```
### 任务9
用SELECT语句求C1课程的成绩低于张三的学生的学号和成绩。
```SQL
mysql> select ID, Score from `Choose` where CourseID='C1' and ID in
    -> (select ID from `Choose` where CourseID='C1' and Score<
    -> (select Choose.Score from Choose, Student where Student.Name='张三' and Choose.CourseID='C1' and Choose.ID = Student.ID
    -> ));
+-------+-------+
| ID    | Score |
+-------+-------+
| 00002 | 80.00 |
| 00003 | 78.00 |
+-------+-------+
2 rows in set (0.00 sec)
```

### 任务10
用SELECT语句求选了C2课程并且也选了C3课程的学生的学号。
```SQL
mysql> select Student.ID from Student where ID in(select ID from Choose where Choose.CourseID = 'C3') 
    -> and ID in(select ID from Choose where Choose.CourseID='C2');
+-------+
| ID    |
+-------+
| 00001 |
+-------+
1 row in set (0.00 sec)
```

## lab7
### 任务1
```SQL
create user 'DCL'@'localhost' identified by 'DCL';
```
### 任务2
```SQL
grant all privileges on *.* to DCL@'localhost'
```
### 任务3
```bash
➜  backend git:(master) ✗ mysql -uDCL -pDCL
mysql: [Warning] Using a password on the command line interface can be insecure.
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 45
Server version: 8.0.28 Homebrew

Copyright (c) 2000, 2022, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> 
```
### 任务4
```SQL
mysql> create Database DCLDemo;
Query OK, 1 row affected (0.00 sec)
mysql> create table Abc( `A1` varchar(20), `B2` Dec(4, 2), `C3` int)ENGINE=InnoDB;
Query OK, 0 rows affected (0.03 sec)
```
### 任务5
```SQL
mysql> insert into Abc (A1, B2, C3) values  ('DCL测试', 90.5, 30);
Query OK, 1 row affected (0.00 sec)

mysql> select * from Abc;
+-----------+-------+------+
| A1        | B2    | C3   |
+-----------+-------+------+
| DCL测试   | 90.50 |   30 |
+-----------+-------+------+
1 row in set (0.00 sec)
```
### 任务6
```bash

➜  backend git:(master) ✗ mysql -u root -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 47
Server version: 8.0.28 Homebrew

Copyright (c) 2000, 2022, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> 
```
```SQL
mysql> revoke all on *.* from DCL@'localhost';
Query OK, 0 rows affected (0.01 sec)
```
## lab8
### 任务1
```SQL
mysql> create table Exam(
    ->     `id` varchar(17),
    ->     `name` varchar(10),
    ->     `sex`varchar(2),
    ->     `age`integer,
    ->     `score`numeric(6,2),
    ->     `address`varchar(50),
    ->     `memo`varchar(100)
    -> )ENGINE=InnoDB;
Query OK, 0 rows affected (0.02 sec)
```
### 任务2
```SQL
mysql> insert into Exam (id, name, sex, age, score, address, memo) values
    -> ('A0001','赵一','男',20,580.00,'重邮宿舍12-3-5',    '学习委员'),
    -> ('B0002','钱二','女',19,540.00,'南福苑5-2-9',       '班长'),
    -> ('C0003','孙三','男',21,555.50,'学生新区21-5-15','优秀共青团员'),
    -> ('D0004','李四','男',22,480.00,'重邮宿舍8-2-22','暂无相关信息'),
    -> ('E0005','周五','女',20,495.50,'学生新区23-4-8','暂无相关信息'),
    -> ('F0006','吴六','男',19,435.00,'南福苑2-5-12', '暂无相关信息');
Query OK, 6 rows affected (0.00 sec)
Records: 6  Duplicates: 0  Warnings: 0

mysql> select * from Exam;
+-------+--------+------+------+--------+---------------------+--------------------+
| id    | name   | sex  | age  | score  | address             | memo               |
+-------+--------+------+------+--------+---------------------+--------------------+
| A0001 | 赵一   | 男   |   20 | 580.00 | 重邮宿舍12-3-5      | 学习委员           |
| B0002 | 钱二   | 女   |   19 | 540.00 | 南福苑5-2-9         | 班长               |
| C0003 | 孙三   | 男   |   21 | 555.50 | 学生新区21-5-15     | 优秀共青团员       |
| D0004 | 李四   | 男   |   22 | 480.00 | 重邮宿舍8-2-22      | 暂无相关信息       |
| E0005 | 周五   | 女   |   20 | 495.50 | 学生新区23-4-8      | 暂无相关信息       |
| F0006 | 吴六   | 男   |   19 | 435.00 | 南福苑2-5-12        | 暂无相关信息       |
+-------+--------+------+------+--------+---------------------+--------------------+
6 rows in set (0.00 sec)
```
### 任务3
```SQL
create index IndexScore on Exam(Score ASC);
```
### 任务4
```SQL
create view ViewExam(ViewExam1, ViewExam2) as select Name, Address from Exam;
```
### 任务6
求Exam中男性同学平均分
```SQL
mysql> select AVG(Score) from Exam where sex = '男';
+------------+
| AVG(Score) |
+------------+
| 512.625000 |
+------------+
1 row in set (0.01 sec)
```
## 课程设计
### 简介
使用golang语言完成了选课系统的后端
基本功能
- 展示所有用户信息
- 修改用户属性
- 增加用户
- 查看课程信息
- 增加课程
- 学生选课
- 为学生所选课程打分
- 展示学生所选课程

功能特点
- 为接口增加JWT鉴权
- 使用go可支持高并发
### 资源设计图
![](./截屏2022-06-25%2017.57.46.png)
### 数据库表设置
```SQL
create TABLE `courses` (
    `course_id` varchar(32),
    `teacher_id` INT NOT NULL,
    `teacher` varchar(32),
    `name` varchar(32) NOT NULL, 
    `max_num` INT NOT NULL,
    `num` INT default 0, 
    `score` INT NOT NULL,
    PRIMARY KEY (course_id, teacher_id) 
) ENGINE=InnoDB; 

create TABLE `users` (
    `id` int PRIMARY KEY,
    `username` varchar(128) NOT NULL,
    `password` varchar(128) NOT NULL,
    `college` varchar(128) NOT NULL,
    `profession` varchar(128),
    `type` int default 0
) ENGINE=InnoDB;

create TABLE `course_select` (
    `student_id` int NOT NULL, 
    `course_id` varchar(32) NOT NULL, 
    `teacher_id` int NOT NULL, 
    `score` int , 
    PRIMARY KEY (student_id, course_id)
) ENGINE=InnoDB;
```
users表保存所有的用户，用type来判断用户的类型，username，password保存用户的登录信息来提供登录功能
courses表保存所有的课程
course_select表保存每一条选课记录
### 具体实现
代码已上传到[github](https://github.com/cooook/db-backend)
后端采用了go-gin框架，使用中间件对api进行权限控制和鉴权，鉴权使用go-jwt框架
**api设置**
```go
func Register_api(r *gin.Engine) {
	r.POST("/auth", auth.AuthHandler)

	api_group := r.Group("/v1").Use(auth.JWTAuth())

	api_group.GET("/users", auth.IsTypeMiddleWare(auth.Admin_Type, true), handler.User_get_all)
	api_group.POST("/users", auth.IsTypeMiddleWare(auth.Admin_Type, true), handler.User_post_handler)     
	api_group.PUT("/users", auth.IsTypeMiddleWare(auth.Admin_Type, true), handler.User_put_handler)      
	api_group.POST("/courses", auth.IsTypeMiddleWare(auth.Admin_Type, true), handler.Course_post_handler)

	api_group.GET("/users/:user_id", handler.User_get_handler)
	api_group.GET("/courses", handler.Course_get_all)
	api_group.GET("/courses/:course_id", handler.Course_get_handler)

	api_group.GET("/course_table", handler.Select_getby_student_handler)
	api_group.PUT("/course_table", handler.Update_Score_handler)
	api_group.POST("/course_table", handler.Add_course_select_entry)
}
```

**JWT鉴权**
```go

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{"error": "No token."})
			c.Abort()
			return
		}

		payload, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		})

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		claims, ok := payload.Claims.(jwt.MapClaims)
		if !(ok && payload.Valid) {
			c.JSON(http.StatusOK, gin.H{"error": "cannot convert claim to mapClaim"})
			c.Abort()
			return
		}
		c.Set("user_id", int(claims["user_id"].(float64)))
		user_id, _ := c.Get("user_id")
		log.Print("user_id=", user_id)
	}
}
```
**用户类型中间件**
```go
func IsTypeMiddleWare(Type User_Type, Is bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id, ok := c.Get("user_id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{"error": fmt.Errorf("no user_id in middleware").Error()})
			c.Abort()
			return
		}

		var user []handler.User
		if err := handler.Db.Select(&user, "select type from users where id = ?", user_id); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if len(user) == 0 {
			c.JSON(http.StatusOK, gin.H{"error": fmt.Errorf("no such user:%d", user_id).Error()})
			c.Abort()
			return
		}

		result := (user[0].Type != int(Type))
		if !Is {
			result = !result
		}

		if result {
			c.JSON(http.StatusOK, gin.H{"error": fmt.Errorf("except type:%v, user type:%v", Type, user[0].Type).Error()})
			c.Abort()
			return
		}
	}
}
```
数据库部分使用了Db模块，代码中含有较多的鉴权判断，在这里不做展示，可以查看附件中的source code/github上的代码
### 运行示例
对api的测试使用了postman
**代码运行**
![](截屏2022-06-25%2018.29.01.png)
**获得token**
![](%E6%88%AA%E5%B1%8F2022-06-25%2018.30.38.png)
以161920319身份尝试创建用户
![](截屏2022-06-25%2018.35.01.png)
此时161920319type为teacher
在数据库中手动修改161920319type为Admin
```SQL
mysql> update users set type=2 where id=161920319;
Query OK, 1 row affected (0.01 sec)
Rows matched: 1  Changed: 1  Warnings: 0
```
再次尝试添加
![](截屏2022-06-25%2018.36.21.png)
在数据库中查询，发现对应用户
```SQL
mysql> select * from users;
+-----------+--------------+--------------+-----------------+--------------+------+
| id        | username     | password     | college         | profession   | type |
+-----------+--------------+--------------+-----------------+--------------+------+
| 161910126 | 赵安         | ekko         | 摆烂学院        | 摆烂学       |    3 |
| 161910127 | 啊吧啊吧     | asd          | 长空学院        | 信息安全     |    3 |
| 161920319 | 张一帆       | password     | 计算机学院      | 信息安全     |    2 |
| 161920320 | 曹玉林       | password     | 计算机学院      | 信息安全     |    3 |
| 161920321 | 陆天舆       | asd          | 长空学院        | 信息安全     |    3 |
| 161920501 | 颜宇明       | TestPassword | 长空学院        | 信息安全     |    3 |
+-----------+--------------+--------------+-----------------+--------------+------+
6 rows in set (0.00 sec)
```
**修改用户密码**
![](截屏2022-06-25%2018.40.39.png)
数据库中查询
```SQL
mysql> select * from users;
+-----------+--------------+--------------+-----------------+--------------+------+
| id        | username     | password     | college         | profession   | type |
+-----------+--------------+--------------+-----------------+--------------+------+
| 161910126 | 赵安         | ekko         | 摆烂学院        | 摆烂学       |    3 |
| 161910127 | 啊吧啊吧     | asd          | 长空学院        | 信息安全     |    3 |
| 161920319 | 张一帆       | password123  | 计算机学院      | 信息安全     |    2 |
| 161920320 | 曹玉林       | password     | 计算机学院      | 信息安全     |    3 |
| 161920321 | 陆天舆       | asd          | 长空学院        | 信息安全     |    3 |
| 161920501 | 颜宇明       | TestPassword | 长空学院        | 信息安全     |    3 |
+-----------+--------------+--------------+-----------------+--------------+------+
6 rows in set (0.00 sec)
```
**查询开课**
![](截屏2022-06-25%2018.42.24.png)
**增加课程**
![](%E6%88%AA%E5%B1%8F2022-06-25%2018.54.20.png)
数据库中也有了对应的课程··
```SQL
mysql> select * from courses;
+-----------+------------+---------+------------------+---------+------+-------+
| course_id | teacher_id | teacher | name             | max_num | num  | score |
+-----------+------------+---------+------------------+---------+------+-------+
| 000.1     |          0 | 1       | apex爆杀教程     |       1 |    0 |     0 |
| 1         |          0 | 123     | apex爆杀教程     |       1 |    1 |     0 |
| 1         |          1 | 233     | Test             |       2 |    0 |     2 |
+-----------+------------+---------+------------------+---------+------+-------+
3 rows in set (0.00 sec)
```
**学生选课**
![](截屏2022-06-25%2018.44.14.png)
```SQL
mysql> select * from course_select;
+------------+-----------+------------+-------+
| student_id | course_id | teacher_id | score |
+------------+-----------+------------+-------+
|  161920319 | 1         |          0 |     0 |
+------------+-----------+------------+-------+
1 row in set (0.00 sec)
mysql> select * from courses;
+-----------+------------+---------+------------------+---------+------+-------+
| course_id | teacher_id | teacher | name             | max_num | num  | score |
+-----------+------------+---------+------------------+---------+------+-------+
| 1         |          0 | 123     | apex爆杀教程     |       1 |    1 |     0 |
| 1         |          1 | 233     | Test             |       2 |    0 |     2 |
+-----------+------------+---------+------------------+---------+------+-------+
2 rows in set (0.00 sec)
```
courses中num也对应的发生了变化
**教师赋分**
![](截屏2022-06-25%2018.45.52.png)
当teacher_id为1的老师尝试为选择同样course但是选择teacher_id为0老师教授的课时会抛出错误。
登录teacher_id为0的用户并且更换token再次尝试
![](截屏2022-06-25%2018.50.29.png)
数据库中查询
```SQL
mysql> select * from course_select;
+------------+-----------+------------+-------+
| student_id | course_id | teacher_id | score |
+------------+-----------+------------+-------+
|  161920319 | 1         |          0 |    90 |
+------------+-----------+------------+-------+
1 row in set (0.00 sec)
```
### 实验心得
本次实验让我从0开始接触了go，并且选择go作为后端语言，从安全角度考虑设计数据库后端，学习了go语言jwt鉴权的使用，gin框架中间件的使用，实践了课上学到的mysql数据库知识。