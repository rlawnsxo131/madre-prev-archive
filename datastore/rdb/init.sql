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
  email character varying(255) NOT NULL,
  origin_name character varying(50) DEFAULT NULL,
  username character varying(50) NOT NULL,
  photo_url character varying(255) DEFAULT NULL,
  created_at timestamp with time zone DEFAULT now() NOT NULL,
  updated_at timestamp with time zone DEFAULT now() NOT NULL,
  PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS user_ix_email ON public.user USING btree (email);
CREATE UNIQUE INDEX IF NOT EXISTS user_ix_username ON public.user USING btree (username);

-- ALTER TABLE public.user OWNER TO madre;

--
-- social_account
--

CREATE TABLE IF NOT EXISTS public.social_account (
  id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
  user_id uuid NOT NULL,
  social_id character varying(255) NOT NULL,
  provider character varying(10) NOT NULL DEFAULT 'GOOGLE',
  created_at timestamp with time zone DEFAULT now() NOT NULL,
  updated_at timestamp with time zone DEFAULT now() NOT NULL,
  PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS social_account_ix_user_id ON public.social_account USING btree (user_id);
CREATE UNIQUE INDEX IF NOT EXISTS social_account_ix_social_id_provider ON public.social_account USING btree (social_id, provider);

-- ALTER TABLE public.social_account OWNER TO madre;

--
-- data
--
CREATE TABLE IF NOT EXISTS public.data (
  id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
  user_id uuid NOT NULL,
  file_url character varying(255)  NOT NULL,
  title character varying(255)  NOT NULL,
  description character varying(255) DEFAULT NULL,
  is_public boolean NOT NULL DEFAULT false,
  created_at timestamp with time zone DEFAULT now() NOT NULL,
  updated_at timestamp with time zone DEFAULT now() NOT NULL,
  PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS data_ix_user_id ON public.data USING btree (user_id);
CREATE INDEX IF NOT EXISTS data_ix_created_at ON public.data USING btree(created_at);

-- ALTER TABLE public.data OWNER TO madre;



