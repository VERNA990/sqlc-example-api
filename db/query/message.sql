/* adduser,
 create group,
  add group members,
   start thread,
    add thread members,
     get dms messages(dm messages for a particular dm chat) ,
      get group messages(messages for a particular group excluding thread messages),
       get thread messages(oder by group and thread name.show all messages for a particular thread,
        the name of the thread and the group to which it belongs),
         get user groups(groups user belongs to), get user threads(threads user is part of),
          get list of groups and users who created the groups,
           get list of threads and users who started the threads,
            get number of thread belonging to a group,
             get list of thread belonging to a group,
              create a message,
               list all groups */ 

-- name: AddUser :one 
INSERT INTO users (fullname, username, email) 
VALUES ($1, $2, $3) RETURNING *; 

-- name: CreateGroup :one 
INSERT INTO groups (gp_name, created_by) 
VALUES ($1, $2) 
RETURNING *; 

-- name: AddGroupMember :one 
INSERT INTO gp_members (gp_members_id, gp_members_user_id, role) 
VALUES ($1, $2, COALESCE(NULLIF($3, '')::role_type, 'member'::role_type))
RETURNING *;

-- name: StartThread :one 
INSERT INTO threads (created_by, gp_id, thread_name, visibility) 
VALUES ($1, $2, $3, COALESCE(NULLIF($4, '')::visibility_enum, 'public'::visibility_enum)) 
RETURNING *; 

-- name: AddThreadMember :one 
INSERT INTO thread_members (thread_members_id, thread_members_user_id) 
VALUES ($1, $2)
RETURNING *; 

-- name: StartDm :one 
INSERT INTO dms (user1_id, user2_id) 
VALUES (LEAST($1, $2), GREATEST($1, $2))
RETURNING *; 

-- name: GetDmMessages :many 
SELECT m.chat_id, m.chat_type, 
(SELECT fullname FROM users 
WHERE user_id = m.sender 
) AS sender_name,

m.content As message, TO_CHAR(m.created_at, 'HH12:MI AM') As time_sent, DATE(m.created_at) AS day_sent
FROM messages m 
WHERE m.chat_type = 'dms' AND m.chat_id = $1 
ORDER BY m.created_at DESC; 

-- name: GetGroupMessages :many 
SELECT m.chat_id, m.chat_type, 
(SELECT fullname FROM users 
WHERE user_id = m.sender 
) AS sender_name, 

m.content AS message, TO_CHAR(m.created_at, 'HH12:MI AM') As time_sent, DATE(m.created_at) AS day_sent
FROM messages m
WHERE m.chat_type = 'groups' AND m.chat_id = $1 
ORDER BY m.created_at DESC; 

-- name: GetThreadMessages :many 
SELECT m.chat_id, m.chat_type,
(SELECT gp_name FROM groups 
WHERE gp_id = (SELECT gp_id FROM threads 
WHERE thread_id = m.chat_id) 
) AS group_name, 

(SELECT thread_name FROM threads 
WHERE thread_id = m.chat_id 
) AS thread_name, 

(SELECT fullname FROM users 
WHERE user_id = m.sender 
) AS sender_name, 

content As message, TO_CHAR(created_at, 'HH12:MI AM') As time_sent, DATE(created_at) AS day_sent
FROM messages m 
WHERE m.chat_type = 'threads' AND m.chat_id = $1 
ORDER BY m.created_at DESC; 

-- name: GetUserGroups :many 
SELECT gp_name FROM groups 
WHERE gp_id IN (SELECT gp_members_id FROM gp_members 
                WHERE gp_members_user_id = $1); 

-- name: GetUserThreads :many 
SELECT thread_name FROM threads 
WHERE thread_id IN (SELECT thread_members_id FROM thread_members 
                    WHERE thread_members_user_id = $1); 

-- name: GetGroupsUserCreated :many 
SELECT g.gp_name As group_name, 
(SELECT fullname FROM users 
WHERE user_id = g.created_by 
) As creators_name 

FROM groups g 
WHERE g.created_by = $1 
ORDER BY creators_name; 

-- name: GetThreadsUserStarted :many 
SELECT 
(SELECT fullname FROM users 
WHERE user_id = t.created_by 
) As creators_name,

t.thread_name FROM threads t 
WHERE t.created_by = $1 
ORDER BY creators_name; 

-- name: GetGroupThreads :many 
SELECT 
(SELECT gp_name FROM groups 
WHERE gp_id = t.gp_id 
) As group_name, 

t.thread_name FROM threads t 
WHERE t.gp_id = $1 
ORDER BY group_name, t.created_at; 

-- name: GetListOfGroups :many 
SELECT gp_id, gp_name FROM groups 
ORDER BY created_at DESC; 

-- name: CreateMessage :one 
INSERT INTO messages (sender, content, chat_id, chat_type) 
VALUES ($1, $2, $3, $4) 
RETURNING *; 

-- name: GetMessageByID :one 
SELECT m.message_id, m.chat_type AS message_from,
(SELECT fullname FROM users 
WHERE user_id = m.sender 
) AS sender_name, 

m.content As message, TO_CHAR(created_at, 'HH12:MI AM') As time_sent, DATE(created_at) AS day_sent
FROM messages m 
WHERE m.message_id = $1;