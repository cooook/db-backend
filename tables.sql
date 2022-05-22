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