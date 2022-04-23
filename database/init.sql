SELECT 'CREATE DATABASE madre'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'madre');

--
-- Create user
-- 
-- CREATE USER madre WITH ENCRYPTED PASSWORD '1234';
-- GRANT ALL PRIVILEGES ON DATABASE madre to madre;

-- \c madre;

-- PostgreSQL database dump
-- SET statement_timeout = 0;
-- SET lock_timeout = 0;
-- SET idle_in_transaction_session_timeout = 0;
-- SET client_encoding = 'UTF8';
-- SET standard_conforming_strings = on;
-- SELECT pg_catalog.set_config('search_path', '', false);
-- SET check_function_bodies = false;
-- SET xmloption = content;
-- SET client_min_messages = warning;
-- SET row_security = off;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: 
--
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--
COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';

--
-- user
--
CREATE TABLE IF NOT EXISTS public.user(
  id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
  email varchar(255) NOT NULL,
  origin_name varchar(16),
  display_name varchar(48) NOT NULL,
  photo_url varchar(255),
  created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
  updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
  PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS user_ix_email ON public.user USING btree (email);

-- ALTER TABLE public.user OWNER TO madre;

--
-- social_account
--
-- DROP TYPE IF EXISTS social_account_provider;
-- CREATE TYPE social_account_provider AS ENUM ('GOOGLE');

CREATE TABLE IF NOT EXISTS public.social_account (
  id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
  user_id uuid NOT NULL,
  provider social_account_provider NOT NULL DEFAULT 'GOOGLE',
  social_id varchar(255) NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
  updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
  PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS social_account_ix_provider_social_id ON public.social_account USING btree (provider, social_id);
CREATE INDEX IF NOT EXISTS social_account_ix_user_id ON public.social_account USING btree (user_id);

-- ALTER TABLE public.social_account OWNER TO madre;

--
-- data
--
CREATE TABLE IF NOT EXISTS public.data (
  id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
  user_id uuid NOT NULL,
  file_url varchar(255)  NOT NULL,
  title varchar(255)  NOT NULL,
  description varchar(255),
  is_public boolean NOT NULL DEFAULT false,
  created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
  updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
  PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS data_ix_user_id ON public.data USING btree (user_id);

-- ALTER TABLE public.data OWNER TO madre;



