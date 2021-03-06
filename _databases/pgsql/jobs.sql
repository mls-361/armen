-- Database: armen

-- DROP TABLE public.jobs;
-- DROP TYPE public.job_status;

-- Type: public.job_status

CREATE TYPE public.job_status AS ENUM
    ('todo', 'running', 'pending', 'succeeded', 'failed');

-- Table: public.jobs

CREATE TABLE public.jobs
(
    id uuid NOT NULL,
    name character varying(30) COLLATE pg_catalog."default" NOT NULL,
    namespace character varying(10) COLLATE pg_catalog."default" NOT NULL,
    type character varying(50) COLLATE pg_catalog."default" NOT NULL,
    origin character varying(50) COLLATE pg_catalog."default" NOT NULL,
    priority integer NOT NULL,
    key character varying(50) COLLATE pg_catalog."default",
    workflow uuid,
    workflow_failed boolean,
    emails character varying(50) COLLATE pg_catalog."default",
    config jsonb,
    private jsonb,
    public jsonb,
    created_at timestamp(3) with time zone NOT NULL,
    status job_status NOT NULL,
    error character varying COLLATE pg_catalog."default",
    attempts integer NOT NULL,
    finished_at timestamp(3) with time zone,
    run_after timestamp(3) with time zone NOT NULL,
    result character varying(20) COLLATE pg_catalog."default",
    next_step character varying(30) COLLATE pg_catalog."default",
    weight integer NOT NULL,
    time_reference timestamp(3) with time zone NOT NULL,
    CONSTRAINT jobs_pkey PRIMARY KEY (id),
    CONSTRAINT jobs_fkey FOREIGN KEY (workflow)
        REFERENCES public.workflows (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
)

TABLESPACE pg_default;

GRANT DELETE, INSERT, SELECT, UPDATE ON TABLE public.jobs TO armen;

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

