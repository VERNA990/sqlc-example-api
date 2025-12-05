CREATE TABLE IF NOT EXISTS "users"( 
"user_id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36), 
"fullname" VARCHAR(50) NOT NULL, 
"username" VARCHAR(30) UNIQUE NOT NULL, 
"email" VARCHAR(100) UNIQUE NOT NULL, 
"created_at" TIMESTAMP DEFAULT now() 
); 

CREATE TABLE IF NOT EXISTS "dms"( 
"dm_id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36), 
"user1_id" VARCHAR(36) NOT NULL REFERENCES users(user_id), 
"user2_id" VARCHAR(36) NOT NULL REFERENCES users(user_id), 
"created_at" TIMESTAMP DEFAULT now(), 
UNIQUE (user1_id, user2_id) 
); 

CREATE TABLE IF NOT EXISTS "groups"( 
"gp_id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36), 
"gp_name" VARCHAR(50) NOT NULL, 
"created_by" VARCHAR(36) NOT NULL REFERENCES users(user_id),
"created_at" TIMESTAMP DEFAULT now()  
); 

CREATE TYPE role_type AS ENUM ('admin', 'member'); 

CREATE TABLE IF NOT EXISTS "gp_members"( 
"gp_members_id" VARCHAR(36) NOT NULL REFERENCES groups(gp_id) ON DELETE CASCADE, 
"gp_members_user_id" VARCHAR(36) NOT NULL REFERENCES users(user_id) ON DELETE CASCADE, 
"role" role_type DEFAULT 'member',
PRIMARY KEY (gp_members_id, gp_members_user_id) 
); 

CREATE TYPE visibility_enum AS ENUM ('public', 'restricted'); 

CREATE TABLE IF NOT EXISTS "threads"( 
"thread_id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36), 
"created_by" VARCHAR(36) REFERENCES users(user_id) ON DELETE SET NULL, 
"gp_id" VARCHAR(36) NOT NULL REFERENCES groups(gp_id) ON DELETE CASCADE, 
"thread_name" VARCHAR(50) NOT NULL, 
"visibility" visibility_enum DEFAULT 'public', 
"created_at" TIMESTAMP DEFAULT now() 
); 

CREATE TABLE IF NOT EXISTS "thread_members"( 
"thread_members_id" VARCHAR(36) NOT NULL REFERENCES threads(thread_id) ON DELETE CASCADE, 
"thread_members_user_id" VARCHAR(36) NOT NULL REFERENCES users(user_id) ON DELETE CASCADE, 
PRIMARY KEY (thread_members_id, thread_members_user_id) 
); 

CREATE TYPE chattype_enum AS ENUM ('dms', 'groups', 'threads'); 

CREATE TABLE IF NOT EXISTS "messages" ( 
"message_id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36), 
"sender"VARCHAR(36) REFERENCES users(user_id) ON DELETE SET NULL, 
"content" TEXT NOT NULL, 
"created_at" TIMESTAMP DEFAULT now(), 
"chat_id" VARCHAR(36) NOT NULL, 
"chat_type" chattype_enum NOT NULL 
);

ALTER TABLE IF EXISTS "messages"
ADD COLUMN edited_at TIMESTAMP NULL;