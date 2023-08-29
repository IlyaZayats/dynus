create table Users(
    id integer NOT NULL PRIMARY KEY,
    created_at timestamp default (current_timestamp),
    updated_at timestamp default (current_timestamp)
);

create table Slugs(
    id SERIAL NOT NULL PRIMARY KEY,
    name varchar NOT NULL unique,
    chance float default (0.0),
    created_at timestamp default (current_timestamp),
    updated_at timestamp default (current_timestamp)
);

create table Link(
    id SERIAL NOT NULL PRIMARY KEY,
    user_id integer not null references Users(id) on delete cascade ,
    slug_id integer not null references Slugs(id) on delete cascade ,
    ttl interval default (NULL),
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

CREATE TRIGGER tr_updated_at_employ BEFORE UPDATE ON Link
    FOR EACH ROW EXECUTE PROCEDURE tr_updated_at();

insert into users (id) select * from generate_series(1, 100)