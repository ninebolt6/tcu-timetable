-- syllabus.Lectures definition

CREATE TABLE `Lectures` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '講義ID',
  `name` varchar(96) NOT NULL COMMENT '講義名',
  `teacher` varchar(32) DEFAULT NULL COMMENT '教員名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='講義情報';


-- syllabus.Timetables definition

CREATE TABLE `Timetables` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '時間割ID',
  `lecture_id` int NOT NULL COMMENT '講義ID',
  `classroom` varchar(64) DEFAULT NULL COMMENT '教室',
  `day_of_week` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '開講曜日',
  `period` int NOT NULL COMMENT '開講時限',
  `memo` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'メモ',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='時間割情報';
