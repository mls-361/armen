-- Database: armen

-- DROP TABLE public.workflows;
-- DROP TYPE public.workflow_status;

-- Type: public.workflow_status

CREATE TYPE public.workflow_status AS ENUM
    ('running', 'succeeded', 'failed');

-- Table: public.workflows

CREATE TABLE public.workflows
(
    id uuid NOT NULL,
    description character varying(100) COLLATE pg_catalog."default" NOT NULL,
    origin character varying(50) COLLATE pg_catalog."default" NOT NULL,
    priority integer NOT NULL,
    first_step character varying(30) COLLATE pg_catalog."default" NOT NULL,
    all_steps jsonb NOT NULL,
    external_reference character varying(50) COLLATE pg_catalog."default",
    emails character varying(50) COLLATE pg_catalog."default",
    data jsonb,
    created_at timestamp(3) with time zone NOT NULL,
    status workflow_status NOT NULL,
    finished_at timestamp(3) with time zone,
    CONSTRAINT workflows_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

GRANT DELETE, INSERT, SELECT, UPDATE ON TABLE public.workflows TO armen;

COMMENT ON TABLE public.workflows
    IS 'La table de gestion des workflows.';

COMMENT ON COLUMN public.workflows.id
    IS 'L''identifiant du workflow.';

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

