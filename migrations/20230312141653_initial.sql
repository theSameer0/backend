-- Thank you for giving goose a try!
-- 
-- This file was automatically created running goose init. If you're familiar with goose
-- feel free to remove/rename this file, write some SQL and goose up. Briefly,
-- 
-- Documentation can be found here: https://pressly.github.io/goose
--
-- A single goose .sql file holds both Up and Down migrations.
-- 
-- All goose .sql files are expected to have a -- +goose Up directive.
-- The -- +goose Down directive is optional, but recommended, and must come after the Up directive.
-- 
-- The -- +goose NO TRANSACTION directive may be added to the top of the file to run statements 
-- outside a transaction. Both Up and Down migrations within this file will be run without a transaction.
-- 
-- More complex statements that have semicolons within them must be annotated with 
-- the -- +goose StatementBegin and -- +goose StatementEnd directives to be properly recognized.
-- 
-- Use GitHub issues for reporting bugs and requesting features, enjoy!

-- +goose Up
CREATE TABLE movies (
    Id serial primary key ,
    Name text not null,
    Language text not null,
    Image text not null,
    HeadImage text not null,
    Tags text not null,
    Comment text not null
);
CREATE TABLE IF NOT EXISTS theatre (
    Id serial primary key ,
    Name text not null,
    Location text not null,
    Image text not null,
    City text not null,
    Screen int not null
);
CREATE TABLE IF NOT EXISTS shows (
    Id serial primary key ,
    Time text not null,
    Seats text,
    Date text not null,
    Screen int not null,
    MovieId int not null,
    TheatreId int not null,
    FOREIGN KEY (MovieId) REFERENCES movies(Id),
    FOREIGN KEY (TheatreId) REFERENCES theatre(Id)
);
CREATE TABLE IF NOT EXISTS ticket (
    Id text primary key ,
    Date text not null,
    Time text not null,
    Seats text not null,
    SeatCount int not null,
    Screen int not null,
    MovieId int not null,
    TheatreId int not null,
    ShowId int not null,
    FOREIGN KEY (MovieId) REFERENCES movies(Id),
    FOREIGN KEY (TheatreId) REFERENCES theatre(Id),
    FOREIGN KEY (ShowId) REFERENCES shows(Id)
);

-- +goose Down
DROP TABLE ticket;
DROP TABLE shows;
DROP TABLE theatre;
DROP TABLE movies;