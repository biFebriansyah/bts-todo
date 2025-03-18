CREATE TABLE public.todos (
	todo_id uuid DEFAULT gen_random_uuid() NULL,
	card_id uuid NOT NULL,
	todo_name varchar NOT NULL,
	todo_status boolean DEFAULT FALSE NULL,
	created_at timestamp DEFAULT NOW() NULL,
	updated_at timestamp NULL,
	CONSTRAINT todos_pk PRIMARY KEY (todo_id),
	CONSTRAINT todos_fk1 FOREIGN KEY (card_id) REFERENCES public.cards(card_id) ON DELETE CASCADE
);
