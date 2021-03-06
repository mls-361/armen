-- Database: armen

-- DROP TABLE public.history;

-- Table: public.history

CREATE TABLE public.history
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( CYCLE INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    created_at timestamp(3) with time zone NOT NULL,
    job uuid,
    workflow uuid,
    type character varying(50) COLLATE pg_catalog."default" NOT NULL,
    status character varying(10) COLLATE pg_catalog."default" NOT NULL,
    data jsonb NOT NULL,
    CONSTRAINT history_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

GRANT DELETE, INSERT, SELECT ON TABLE public.history TO armen;

COMMENT ON TABLE public.history
    IS 'L''historique de la vie des jobs et des workflows.';

COMMENT ON COLUMN public.history.id
    IS 'Un identifiant unique.';

COMMENT ON COLUMN public.history.created_at
    IS 'La date et l''heure de création de l''enregistrement.';

COMMENT ON COLUMN public.history.job
    IS 'L''éventuel identifiant du job.';

COMMENT ON COLUMN public.history.workflow
    IS 'L''éventuel identifiant du workflow.';

COMMENT ON COLUMN public.history.type
    IS 'Le type du job ou du workflow';

COMMENT ON COLUMN public.history.status
    IS 'Le statut du job ou du workflow';

COMMENT ON COLUMN public.history.data
    IS 'Le job ou le workflow.';

