-- BEGIN

-- Database: armen

/*
CREATE DATABASE armen
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;
*/

-- Table: public.locks
------------------------------------------------------------------------------------------------------------------------

CREATE TABLE public.locks
(
    name character varying(20) COLLATE pg_catalog."default" NOT NULL,
    owner uuid,
    expiry timestamp(3) with time zone,
    CONSTRAINT locks_pkey PRIMARY KEY (name)
)

TABLESPACE pg_default;

ALTER TABLE public.locks
    OWNER to postgres;

GRANT UPDATE, SELECT ON TABLE public.locks TO armen;

GRANT ALL ON TABLE public.locks TO postgres;

COMMENT ON TABLE public.locks
    IS 'Cette table permet la gestion de verrous.';

COMMENT ON COLUMN public.locks.name
    IS 'Le nom du verrou.';

COMMENT ON COLUMN public.locks.owner
    IS 'Le propriétaire du verrou.';

COMMENT ON COLUMN public.locks.expiry
    IS 'La date d''expiration du verrou.';

INSERT INTO public.locks VALUES ('leader', NULL, NULL);

-- Table: public.history
------------------------------------------------------------------------------------------------------------------------

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

ALTER TABLE public.history
    OWNER to postgres;

GRANT DELETE, INSERT, SELECT ON TABLE public.history TO armen;

GRANT ALL ON TABLE public.history TO postgres;

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

-- Type: workflow_status
------------------------------------------------------------------------------------------------------------------------

CREATE TYPE public.workflow_status AS ENUM
    ('running', 'succeeded', 'failed');

ALTER TYPE public.workflow_status
    OWNER TO postgres;

-- Table: public.workflows
------------------------------------------------------------------------------------------------------------------------

CREATE TABLE public.workflows
(
    id uuid NOT NULL,
    type character varying(20) COLLATE pg_catalog."default" NOT NULL,
    description character varying(100) COLLATE pg_catalog."default" NOT NULL,
    origin character varying(50) COLLATE pg_catalog."default" NOT NULL,
    priority smallint NOT NULL,
    first_step character varying(30) COLLATE pg_catalog."default" NOT NULL,
    all_steps jsonb NOT NULL,
    external_reference character varying(50) COLLATE pg_catalog."default",
    emails character varying(50) COLLATE pg_catalog."default",
    data jsonb,
    created_at timestamp(3) with time zone NOT NULL,
    status public.workflow_status NOT NULL,
    finished_at timestamp(3) with time zone,
    CONSTRAINT workflows_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.workflows
    OWNER to postgres;

GRANT DELETE, INSERT, SELECT, UPDATE ON TABLE public.workflows TO armen;

GRANT ALL ON TABLE public.workflows TO postgres;

COMMENT ON TABLE public.workflows
    IS 'La table de gestion des workflows.';

COMMENT ON COLUMN public.workflows.id
    IS 'L''identifiant du workflow.';

COMMENT ON COLUMN public.workflows.type
    IS 'Le type de ce workflow.';

COMMENT ON COLUMN public.workflows.description
    IS 'La description du workflow.';

COMMENT ON COLUMN public.workflows.origin
    IS 'L''origine du workflow.';

COMMENT ON COLUMN public.workflows.priority
    IS 'La priorité du workflow.
C''est une valeur entre 0 (la priorité la plus basse) et 100 (la priorité la plus haute).';

COMMENT ON COLUMN public.workflows.first_step
    IS 'La première étape du workflow à exécuter.';

COMMENT ON COLUMN public.workflows.all_steps
    IS 'L''ensemble des étapes du workflow.';

COMMENT ON COLUMN public.workflows.external_reference
    IS 'Une éventuelle référence externe.';

COMMENT ON COLUMN public.workflows.emails
    IS 'Liste d''adresses email à informer sur différents évènements de la vie du workflow.';

COMMENT ON COLUMN public.workflows.data
    IS 'Les données de départ du workflow.';

COMMENT ON COLUMN public.workflows.created_at
    IS 'La date et l''heure à laquelle le workflow a été créé.';

COMMENT ON COLUMN public.workflows.status
    IS 'Le statut du workflow.';

COMMENT ON COLUMN public.workflows.finished_at
    IS 'La date et l''heure à laquelle l''exécution du workflow s''est terminée.';

-- Type: job_status
------------------------------------------------------------------------------------------------------------------------

CREATE TYPE public.job_status AS ENUM
    ('todo', 'running', 'pending', 'succeeded', 'failed');

ALTER TYPE public.job_status
    OWNER TO postgres;

-- Table: public.jobs
------------------------------------------------------------------------------------------------------------------------

CREATE TABLE public.jobs
(
    id uuid NOT NULL,
    name character varying(30) COLLATE pg_catalog."default" NOT NULL,
    namespace character varying(10) COLLATE pg_catalog."default" NOT NULL,
    type character varying(50) COLLATE pg_catalog."default" NOT NULL,
    origin character varying(50) COLLATE pg_catalog."default" NOT NULL,
    priority smallint NOT NULL,
    key character varying(50) COLLATE pg_catalog."default",
    workflow uuid,
    workflow_failed boolean,
    emails character varying(50) COLLATE pg_catalog."default",
    config jsonb,
    private jsonb,
    public jsonb,
    created_at timestamp(3) with time zone NOT NULL,
    status public.job_status NOT NULL,
    error character varying COLLATE pg_catalog."default",
    attempts smallint NOT NULL,
    finished_at timestamp(3) with time zone,
    run_after timestamp(3) with time zone NOT NULL,
    result character varying(20) COLLATE pg_catalog."default",
    next_step character varying(30) COLLATE pg_catalog."default",
    weight smallint NOT NULL,
    time_reference timestamp(3) with time zone NOT NULL,
    CONSTRAINT jobs_pkey PRIMARY KEY (id),
    CONSTRAINT jobs_fkey FOREIGN KEY (workflow)
        REFERENCES public.workflows (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
)

TABLESPACE pg_default;

ALTER TABLE public.jobs
    OWNER to postgres;

GRANT DELETE, INSERT, SELECT, UPDATE ON TABLE public.jobs TO armen;

GRANT ALL ON TABLE public.jobs TO postgres;

COMMENT ON TABLE public.jobs
    IS 'La table de gestion des jobs.';

COMMENT ON COLUMN public.jobs.id
    IS 'L''identifiant du job.';

COMMENT ON COLUMN public.jobs.name
    IS 'Le nom du job.
Ce nom est utilisé dans les workflows pour identifier le job.
Il doit donc être unique pour un workflow.';

COMMENT ON COLUMN public.jobs.namespace
    IS 'L''espace de nom qui implémente ce job.';

COMMENT ON COLUMN public.jobs.type
    IS 'Le type de ce job.
Ce type est unique par espace de nom.';

COMMENT ON COLUMN public.jobs.origin
    IS 'L''origine du job.
Elle correspond à l''origine du workflow ou à l''entité qui a créée ce job.';

COMMENT ON COLUMN public.jobs.priority
    IS 'La priorité du job.
C''est une valeur entre 0 (la priorité la plus basse) et 100 (la priorité la plus haute).';

COMMENT ON COLUMN public.jobs.key
    IS 'Clé d''unicité du job.
Il ne peut y avoir qu''un seul job actif (todo, running, pending) ayant cette clé.
Cela permet d''éviter de créer plusieurs fois le même job.';

COMMENT ON COLUMN public.jobs.workflow
    IS 'L''éventuel identifiant du workflow auquel appartient ce job.';

COMMENT ON COLUMN public.jobs.workflow_failed
    IS 'Est-ce que le workflow auquel appartient ce job est en erreur ?';

COMMENT ON COLUMN public.jobs.emails
    IS 'Liste d''adresses email à informer sur différents évènements de la vie du job.';

COMMENT ON COLUMN public.jobs.config
    IS 'Configuration nécessaire à l''exécution du job.
Ce sont des données statiques.';

COMMENT ON COLUMN public.jobs.private
    IS 'Ce champ permet de partager des données entre les différentes sessions d''exécution de ce job.
Ce sont des données dynamiques.';

COMMENT ON COLUMN public.jobs.public
    IS 'Ce champ permet de partager des données entre les jobs.
Ce sont des données dynamiques.';

COMMENT ON COLUMN public.jobs.created_at
    IS 'La date et l''heure à laquelle le job a été créé.';

COMMENT ON COLUMN public.jobs.status
    IS 'Le statut du job.';

COMMENT ON COLUMN public.jobs.error
    IS 'L''éventuel dernier message d''erreur résultant de l''exécution du job.';

COMMENT ON COLUMN public.jobs.attempts
    IS 'Le nombre de tentatives effectuées pour tenter d''exécuter le job.';

COMMENT ON COLUMN public.jobs.finished_at
    IS 'La date et l''heure à laquelle l''exécution du job s''est terminée.';

COMMENT ON COLUMN public.jobs.run_after
    IS 'La date et l''heure à partir de laquelle le job peut être sélectionné pour être exécuté.';

COMMENT ON COLUMN public.jobs.result
    IS 'Un éventuel résultat d''exécution du job.
Ce résultat permettant de déterminer le prochain job à exécuter.';

COMMENT ON COLUMN public.jobs.next_step
    IS 'L''éventuel nom du prochain job (step) à exécuter.';

COMMENT ON COLUMN public.jobs.weight
    IS 'Ce champ sert uniquement à l''algorithme de sélection du prochain job à exécuter.';

COMMENT ON COLUMN public.jobs.time_reference
    IS 'La date et l''heure de référence utilisées par l''algorithme de sélection du prochain job à exécuter.
Quand le job appartient à un workflow, il s''agit de la date de création du workflow sinon de la date de création du job.';

-- END

