-- +goose Up
-- +goose StatementBegin

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from ticket where 1=1;
-- +goose StatementEnd
