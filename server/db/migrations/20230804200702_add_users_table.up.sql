create table "users" (
    "id" bigserial PRIMARY KEY,
    "username" varchar not null,
    "email" varchar not null,
    "password" varchar not null
)