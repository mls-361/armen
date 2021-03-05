-- Database: armen

-- DROP TABLE public.locks;

-- Table: public.locks

CREATE TABLE public.locks
(
    name character varying(20) COLLATE pg_catalog."default" NOT NULL,
    owner uuid,
    expiry timestamp(3) with time zone,
    CONSTRAINT locks_pkey PRIMARY KEY (name)
)

TABLESPACE pg_default;

GRANT UPDATE, SELECT ON TABLE public.locks TO armen;

COMMENT ON TABLE public.locks
    IS 'Cette table permet la gestion de verrous.';

COMMENT ON COLUMN public.locks.name
    IS 'Le nom du verrou.';

COMMENT ON COLUMN public.locks.owner
    IS 'Le propriétaire du verrou.';

COMMENT ON COLUMN public.locks.expiry
    IS 'La date d''expiration du verrou.';

INSERT INTO public.locks(name, owner, expiry) VALUES ('leader', NULL, NULL);

