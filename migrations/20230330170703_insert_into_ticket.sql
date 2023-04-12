-- +goose Up
-- +goose StatementBegin

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from ticket where Id != '';
-- +goose StatementEnd
