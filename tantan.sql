-- Table: public.users

-- DROP TABLE public.users;

CREATE TABLE public.users
(
    id integer NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    name text COLLATE pg_catalog."default" NOT NULL DEFAULT ''::text,
    gender smallint NOT NULL DEFAULT 0,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_gender_check CHECK (gender >= '-127'::integer AND gender <= 128)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.users
    OWNER to lixiaodong;


-- Table: public.relation

-- DROP TABLE public.relation;

CREATE TABLE public.relation
(
    id integer NOT NULL DEFAULT nextval('relation_id_seq'::regclass),
    id1 bigint NOT NULL DEFAULT 0,
    id2 bigint NOT NULL DEFAULT 0,
    state smallint NOT NULL DEFAULT 0,
    CONSTRAINT relation_pkey PRIMARY KEY (id),
    CONSTRAINT relation_id1_id2_idx UNIQUE (id1, id2),
    CONSTRAINT relation_id_key UNIQUE (id),
    CONSTRAINT relation_id1_fkey FOREIGN KEY (id1)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT relation_id2_fkey FOREIGN KEY (id2)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT relation_state_check CHECK (state >= '-127'::integer AND state <= 128)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.relation
    OWNER to lixiaodong;

-- Index: relation_id_state_idx

-- DROP INDEX public.relation_id_state_idx;

CREATE INDEX relation_id_state_idx
    ON public.relation USING btree
    (id1, state)
    TABLESPACE pg_default;
