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
    Id text primary key,
    Name text,
    Language text,
    Image text,
    HeadImage text,
    Tags text,
    Comment text
);
CREATE TABLE IF NOT EXISTS theatre (
    Id text primary key,
    Name text,
    Location text,
    Image text,
    City text,
    Screen int
);
CREATE TABLE IF NOT EXISTS shows (
    Id text primary key,
    Time text,
    Seats text,
    Date text,
    Screen int,
    MovieId text,
    TheatreId text,
    FOREIGN KEY (MovieId) REFERENCES movies(Id),
    FOREIGN KEY (TheatreId) REFERENCES theatre(Id)
);
CREATE TABLE IF NOT EXISTS ticket (
    Id text primary key,
    Date text,
    Time text,
    Seats text,
    SeatCount int,
    Screen int,
    MovieId text,
    TheatreId text,
    ShowId text,
    FOREIGN KEY (MovieId) REFERENCES movies(Id),
    FOREIGN KEY (TheatreId) REFERENCES theatre(Id),
    FOREIGN KEY (ShowId) REFERENCES shows(Id)
);

-- +goose Down
DROP TABLE ticket;
DROP TABLE shows;
DROP TABLE theatre;
DROP TABLE movies;