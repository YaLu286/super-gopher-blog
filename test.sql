-- create table articles(
--   ID bigint primary key not null,
--   Title varchar not null,
--   Text text not null, 
--   PostDate timestamp not null
-- );

-- SELECT id, title, text, postdate FROM articles WHERE id BETWEEN 0 AND 3

-- INSERT INTO articles VALUES (1, 'Lorem Ipsum', 'But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the system, and expound the actual teachings of the great explorer of the truth, the master-builder of human happiness. No one rejects, dislikes, or avoids pleasure itself, because it is pleasure, but because those who do not know how to pursue pleasure rationally encounter consequences that are extremely painful. Nor again is there anyone who loves or pursues or desires to obtain pain of itself, because it is pain, but because occasionally circumstances occur in which toil and pain can procure him some great pleasure. To take a trivial example, which of us ever undertakes laborious physical exercise, except to obtain some advantage from it? But who has any right to find fault with a man who chooses to enjoy a pleasure that has no annoying consequences, or one who avoids a pain that produces no resultant pleasure?', now());
-- INSERT INTO articles VALUES (2, 'Muspi Merol', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', now());
-- INSERT INTO articles VALUES (3, 'Lorem Ipsum', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', now());

-- DROP table users;

-- create table users(
--     login varchar not null,
--     password varchar not null
-- );

-- INSERT INTO users VALUES('super-gopher', '$2a$14$1uTQ/.kg8FAei6fJWTOGMOzcokrs9Ho4djNDciyhz4/pMsOCDwVBm')

-- SELECT * FROM articles;

select * from users;

-- DELETE FROM articles WHERE id = 5;
