CREATE TABLE public.users(
    id serial4 not null,
    "name" varchar null,
    "username" varchar null,
    "email" varchar null,
    CONSTRAINT users_pk PRIMARY KEY (id)
)
