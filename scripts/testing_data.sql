INSERT INTO course (id, created_at, updated_at, deleted_at, name) VALUES 
('course1', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Introduction to Programming'),
('course2', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Database Management'),
('course3', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Web Development');
INSERT INTO "group" (id, created_at, updated_at, deleted_at, name, max_size, course_id) VALUES 
('group1', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Programming Group A', 10, 'course1'),
('group2', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Database Group B', 8, 'course2'),
('group3', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Web Development Group', 12, 'course3');
INSERT INTO user (id, created_at, updated_at, deleted_at, name, email) VALUES 
('user1', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'John Doe', 'john@example.com'),
('user2', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Jane Smith', 'jane@example.com'),
('user3', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Alice Johnson', 'alice@example.com');
INSERT INTO user_tag (id, created_at, updated_at, deleted_at, value, user_id) VALUES 
('usertag1', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Beginner', 'user1'),
('usertag2', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Intermediate', 'user2'),
('usertag3', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Advanced', 'user3');
INSERT INTO rating (id, created_at, updated_at, deleted_at, text_content, star_rating, user_id) VALUES 
('rating1', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Great course!', 5, 'user1'),
('rating2', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Excellent content', 5, 'user2'),
('rating3', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Very informative', 4, 'user3');
INSERT INTO rating_tag (id, created_at, updated_at, deleted_at, value, rating_id) VALUES 
('tag1', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Beginner-friendly', 'rating1'),
('tag2', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Challenging', 'rating2'),
('tag3', '2024-03-03 10:00:00', '2024-03-03 10:00:00', NULL, 'Engaging', 'rating3');
INSERT INTO user_group (group_id, user_id) VALUES 
('group1', 'user1'),
('group2', 'user2'),
('group3', 'user3');

