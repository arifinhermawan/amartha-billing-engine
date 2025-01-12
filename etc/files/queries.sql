CREATE TABLE public."user" (
	id BIGSERIAL NOT NULL,
	"name" varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT user_pkey PRIMARY KEY (id)
);

CREATE TABLE public.loan (
	id BIGSERIAL NOT NULL,
	user_id BIGSERIAL NOT NULL,
	amount numeric(15, 2) NOT NULL,
	interest_rate numeric(5, 2) NOT NULL,
	is_active bool DEFAULT true NOT NULL,
	outstanding_balance numeric(15, 2) NOT NULL,
	start_date date NOT NULL,
	end_date date NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT loan_pkey PRIMARY KEY (id),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE CASCADE
);

CREATE INDEX loan_user_id_idx ON public.loan USING btree (user_id, is_active);

CREATE TABLE public.payment (
	id BIGSERIAL NOT NULL,
	loan_id BIGSERIAL NOT NULL,
	week_number int4 NOT NULL,
	amount numeric(15, 2) NOT NULL,
	due_date date NOT NULL,
	is_paid bool DEFAULT false,
	paid_date date NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT payment_pkey PRIMARY KEY (id),
    CONSTRAINT fk_loan FOREIGN KEY (loan_id) REFERENCES public.loan(id) ON DELETE CASCADE
);

CREATE INDEX payment_loan_id_idx ON public.payment USING btree (loan_id, is_paid, due_date);