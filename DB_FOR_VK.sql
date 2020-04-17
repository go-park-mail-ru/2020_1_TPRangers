DROP TABLE IF EXISTS Photos CASCADE;
DROP TABLE IF EXISTS Users CASCADE;
DROP TABLE IF EXISTS Friends;
DROP TABLE IF EXISTS Posts CASCADE;
DROP TABLE IF EXISTS Feeds;
DROP TABLE IF EXISTS UsersPosts;
DROP TABLE IF EXISTS Albums;
DROP TABLE IF EXISTS UsersPostsLikes;
DROP TABLE IF EXISTS UsersPhotosLikes;
DROP TABLE IF EXISTS AlbumsPhotos;
DROP TABLE IF EXISTS PhotosFromAlbums;
DROP TABLE IF EXISTS PostsAuthor;
DROP TABLE IF EXISTS Chats CASCADE;
DROP TABLE IF EXISTS ChatsUsers;
DROP TABLE IF EXISTS Messages;



CREATE EXTENSION IF NOT EXISTS CITEXT;
CREATE EXTENSION IF NOT EXISTS BYTEA;

CREATE TABLE Photos
(
    photo_id           SERIAL PRIMARY KEY,
    url                TEXT,
    photos_likes_count INT
);

CREATE TABLE Users
(
    u_id      SERIAL PRIMARY KEY,
    login     CITEXT COLLATE "C" NOT NULL UNIQUE,
    phone     VARCHAR(20),
    mail      CITEXT COLLATE "C" NOT NULL UNIQUE,
    name      TEXT,
    surname   TEXT,
    password  BYTEA,
    photo_id  INT DEFAULT 1 REFERENCES Photos,
    birthdate VARCHAR(20)
);

CREATE TABLE Friends
(
    u_id INT NOT NULL REFERENCES Users,
    f_id INT
);

CREATE UNIQUE INDEX userfriend_idx ON Friends (u_id, f_id);

CREATE TABLE Posts
(
    post_id           SERIAL PRIMARY KEY,
    txt_data          TEXT,
    photo_id          INT,
    posts_likes_count INT,
    creation_date     TIMESTAMP,
    attachments       TEXT
);

CREATE TABLE PostsAuthor
(
    post_id INT NOT NULL REFERENCES Posts ON DELETE CASCADE,
    u_id    INT NOT NULL REFERENCES Users
);

CREATE TABLE Feeds
(
    u_id    INT NOT NULL REFERENCES Users,
    post_id INT NOT NULL REFERENCES Posts
);

CREATE TABLE UsersPosts
(
    u_id       INT NOT NULL REFERENCES Users,
    post_owner INT NOT NULL REFERENCES Users,
    post_id    INT NOT NULL REFERENCES Posts
);

CREATE TABLE Albums
(
    album_id SERIAL PRIMARY KEY,
    name     TEXT,
    u_id     INT
);

CREATE TABLE PhotosFromAlbums
(
    photo_id  INT NOT NULL REFERENCES Photos,
    photo_url TEXT,
    album_id  INT
);

CREATE TABLE UsersPostsLikes
(
    postlike_id BIGSERIAL PRIMARY KEY,
    u_id        INT NOT NULL REFERENCES Users,
    post_id     INT NOT NULL REFERENCES Posts
);

CREATE UNIQUE INDEX userpostlike_idx ON UsersPostsLikes (u_id, post_id);

CREATE TABLE UsersPhotosLikes
(
    photolike_id BIGSERIAL PRIMARY KEY,
    u_id         INT NOT NULL REFERENCES Users,
    photo_id     INT NOT NULL REFERENCES Photos
);

CREATE UNIQUE INDEX userphotolike_idx ON UsersPhotosLikes (u_id, photo_id);

CREATE TABLE Chats
(
    ch_id BIGSERIAL PRIMARY KEY,
    photo_id  INT DEFAULT 2 REFERENCES Photos,
    name  TEXT  DEFAULT ''
);

CREATE TABLE ChatsUsers
(
    ch_id BIGINT NOT NULL REFERENCES Chats,
    u_id  INT    NOT NULL REFERENCES Users
);

CREATE UNIQUE INDEX chatuser_idx ON ChatsUsers (u_id, ch_id);

CREATE TABLE Messages
(
    msg_id BIGINT PRIMARY KEY,
    ch_id  BIGINT NOT NULL REFERENCES Chats,
    u_id   INT    NOT NULL REFERENCES Users,
    txt    TEXT
);



INSERT INTO photos (url, photos_likes_count)
VALUES ('https://social-hub.ru/uploads/img/default.png', 0);

INSERT INTO photos (url, photos_likes_count)
VALUES ('https://social-hub.ru/uploads/img/default_chat.png', 0);







