CREATE TYPE UserRole AS ENUM ('engineer', 'worker', 'master', 'manager', 'admin');
CREATE TYPE TaskPriority AS ENUM ('low', 'middle', 'hot_task', 'hot_fix');
CREATE TYPE TaskFrequency AS ENUM ('one_time', 'daily', 'weekly', 'monthly','quarterly');
CREATE TYPE TaskStatus AS ENUM ('todo', 'inProgress', 'failed', 'completed', 'verified');