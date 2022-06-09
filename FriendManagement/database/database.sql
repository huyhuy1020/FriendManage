--Define user to allow users use email
CREATE TABLE user_profile(
    email VARCHAR(100) NOT NULL,
    CONSTRAINT pk_user_profile PRIMARY KEY(email),
);

--Define Friend to contact between 2 emails
CREATE TABLE friend(
    id int NOT NULL GENERATED ALWAYS AS IDENTITY,
    the_first_user VARCHAR(100) NULL,
    the_second_user VARCHAR(100) NULL,
    CONSTRAINT pk_friend PRIMARY KEY(id),
    CONSTRAINT fk_first_user FOREIGN KEY(the_first_user) REFERENCES user_profile(email),
    CONSTRAINT fk_second_user FOREIGN KEY(the_second_user) REFERENCES user_profile(email),
);

--Define Block_User to block a people has been pointed by requestor
CREATE TABLE block_user (
    id int NOT NULL GENERATED ALWAYS AS IDENTITY,
    requestor VARCHAR(100) NULL,
    target VARCHAR(100) NULL,
    CONSTRAINT PK_Block PRIMARY KEY(id),
    CONSTRAINT FK_Request FOREIGN KEY(requestor) REFERENCES user_profile(email),
    CONSTRAINT FK_Blocked FOREIGN KEY(target) REFERENCES user_profile(email));

--Define Subcription to update a subcribe from email
CREATE TABLE subcription(
    id int NOT NULL GENERATED ALWAYS AS IDENTITY,
    sender VARCHAR(100) NULL,
    reciever VARCHAR(100) NULL,
    CONSTRAINT PK_Subcription PRIMARY KEY(id),
    CONSTRAINT FK_sender FOREIGN KEY(sender) REFERENCES user_profile(email),
    CONSTRAINT FK_reciever FOREIGN KEY(reciever) REFERENCES user_profile(email));
