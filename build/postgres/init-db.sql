create table Users(
    id  SERIAL NOT NULL PRIMARY KEY,
    login varchar NOT NULL,
    created_at timestamp default (current_timestamp),
    updated_at timestamp default (current_timestamp)
);

create table Slugs(
    name varchar NOT NULL PRIMARY KEY,
    created_at timestamp default (current_timestamp),
    updated_at timestamp default (current_timestamp)
);

create table Users_With_Slugs(
    id SERIAL NOT NULL PRIMARY KEY,
    user_id int not null references Users on delete cascade ,
    slug_name varchar not null references Slugs on delete cascade ,
    is_valid bool default (True),
    created_at timestamp default (current_timestamp),
    updated_at timestamp default (current_timestamp)
);

CREATE FUNCTION tr_updated_at() RETURNS trigger AS $tr_updated_at$
    BEGIN
        NEW.updated_at := current_timestamp;
        RETURN NEW;
    END;
$tr_updated_at$ LANGUAGE plpgsql;

CREATE TRIGGER tr_updated_at_employ BEFORE UPDATE ON Users
    FOR EACH ROW EXECUTE PROCEDURE tr_updated_at();

CREATE TRIGGER tr_updated_at_employ BEFORE UPDATE ON Slugs
    FOR EACH ROW EXECUTE PROCEDURE tr_updated_at();

CREATE TRIGGER tr_updated_at_employ BEFORE UPDATE ON Users_With_Slugs
    FOR EACH ROW EXECUTE PROCEDURE tr_updated_at();

insert into users (login) values ('user1'), ('user2'), ('user3'), ('user4'), ('user5');