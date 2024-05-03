-- migrate:up
INSERT INTO
    roles (id, name, can_view, can_participate, can_create, is_free, is_admin)
VALUES
    (1, "Администратор", true, true, true, true, true);

INSERT INTO
    roles (id, name, can_view, can_participate, can_create, is_free, is_admin)
VALUES
    (2, "Судья", true, false, true, false, false);

INSERT INTO
    roles (id, name, can_view, can_participate, can_create, is_free, is_admin)
VALUES
    (3, "Неподтвержденный", false, false, false, false, false);

INSERT INTO
    roles (id, name, can_view, can_participate, can_create, is_free, is_admin)
VALUES
    (4, "Игрок", true, true, false, false, false);

INSERT INTO
    roles (id, name, can_view, can_participate, can_create, is_free, is_admin)
VALUES
    (5, "Спартаковец", true, true, false, true, false);

-- migrate:down
DELETE FROM
    roles
WHERE
    id in (1, 2, 3, 4, 5);

