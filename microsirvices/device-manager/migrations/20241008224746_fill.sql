-- +goose Up
-- +goose StatementBegin
insert into device_type (id, name) values (1, 'heating_system');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from device_type where id=1;
-- +goose StatementEnd
