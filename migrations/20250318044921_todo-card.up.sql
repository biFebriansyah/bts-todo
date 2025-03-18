CREATE TABLE public.cards (
	card_id uuid DEFAULT gen_random_uuid() NULL,
	user_id uuid NOT NULL,
	card_name varchar NOT NULL,
	created_at timestamp DEFAULT NOW() NULL,
	updated_at timestamp NULL,
	CONSTRAINT cards_pk PRIMARY KEY (card_id),
	CONSTRAINT cards_fk1 FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE
);
