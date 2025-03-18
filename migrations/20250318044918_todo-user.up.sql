CREATE TABLE public.users (
	user_id uuid DEFAULT gen_random_uuid() NULL,
	username varchar NOT NULL,
	email varchar NOT NULL,
	password varchar NOT NULL,
	created_at timestamp DEFAULT NOW() NULL,
	updated_at timestamp NULL,
	CONSTRAINT users_pk PRIMARY KEY (user_id),
	CONSTRAINT users_unique UNIQUE (username)
);
