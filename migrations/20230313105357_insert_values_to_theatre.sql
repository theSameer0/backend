-- +goose Up
-- +goose StatementBegin
INSERT INTO theatre (Name,Location,Image,City,Screen) VALUES 
('Urvashi Cinema Hall','Shivaji Nagar','https://ik.imagekit.io/2h0gcydui/images/UrvashiCinemaHall.png','Bangalore',4),
('Inox','1mg Mall','https://ik.imagekit.io/2h0gcydui/images/Theatre.png','Hyderabad',4),
('Cinepolis','Herbal','https://ik.imagekit.io/2h0gcydui/images/Theatre.png','Chennai',4),
('Matrix','English','https://ik.imagekit.io/2h0gcydui/images/Theatre.png','Bangalore',4),
('Cinema Hall','SLN Mall','https://ik.imagekit.io/2h0gcydui/images/Theatre.png','Hyderabad',4),
('Cinema Hall','SLN Mall','https://ik.imagekit.io/2h0gcydui/images/Theatre.png','Chennai',4),
('Cinema Hall','SLN Mall','https://ik.imagekit.io/2h0gcydui/images/Theatre.png','Bangalore',4),
('Cinema Hall','DLF Mall','https://ik.imagekit.io/2h0gcydui/images/Theatre.png','Hyderabad',4);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM theatre WHERE 1=1;
-- +goose StatementEnd
