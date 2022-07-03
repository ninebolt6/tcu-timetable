-- syllabus.Lectures definition

CREATE TABLE `Lectures` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '講義ID',
  `name` varchar(96) NOT NULL COMMENT '講義名',
  `place_id` int DEFAULT NULL COMMENT '教室',
  PRIMARY KEY (`id`),
  KEY `Lectures_FK` (`place_id`),
  CONSTRAINT `Lectures_FK` FOREIGN KEY (`place_id`) REFERENCES `Classrooms` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='講義情報';


-- syllabus.Classrooms definition

CREATE TABLE `Classrooms` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '教室ID',
  `campus_id` int NOT NULL COMMENT 'キャンパス',
  `name` varchar(100) NOT NULL COMMENT '教室名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='教室情報';


-- syllabus.Timetables definition

CREATE TABLE `Timetables` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '時間割ID',
  `lecture_id` int NOT NULL COMMENT '講義ID',
  `classroom_id` int DEFAULT NULL COMMENT '教室ID',
  `day_of_week` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '開講曜日',
  `period` int NOT NULL COMMENT '開講時限',
  `memo` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'メモ',
  PRIMARY KEY (`id`),
  KEY `Timetables_FK` (`lecture_id`),
  KEY `Timetables_FK_1` (`classroom_id`),
  CONSTRAINT `Timetables_FK` FOREIGN KEY (`lecture_id`) REFERENCES `Lectures` (`id`),
  CONSTRAINT `Timetables_FK_1` FOREIGN KEY (`classroom_id`) REFERENCES `Classrooms` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='時間割情報';
