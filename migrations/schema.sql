--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.16
-- Dumped by pg_dump version 9.5.16

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: admin_privs; Type: TABLE; Schema: public; Owner: buffalo
--

CREATE TABLE public.admin_privs (
    id uuid NOT NULL,
    administrator_id uuid NOT NULL,
    privilege_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.admin_privs OWNER TO buffalo;

--
-- Name: administrators; Type: TABLE; Schema: public; Owner: buffalo
--

CREATE TABLE public.administrators (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.administrators OWNER TO buffalo;

--
-- Name: privileges; Type: TABLE; Schema: public; Owner: buffalo
--

CREATE TABLE public.privileges (
    id uuid NOT NULL,
    label text NOT NULL,
    description text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.privileges OWNER TO buffalo;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: buffalo
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO buffalo;

--
-- Name: admin_privs_pkey; Type: CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.admin_privs
    ADD CONSTRAINT admin_privs_pkey PRIMARY KEY (id);


--
-- Name: administrators_pkey; Type: CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.administrators
    ADD CONSTRAINT administrators_pkey PRIMARY KEY (id);


--
-- Name: privileges_pkey; Type: CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.privileges
    ADD CONSTRAINT privileges_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: buffalo
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: admin_privs_administrator_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.admin_privs
    ADD CONSTRAINT admin_privs_administrator_id_fkey FOREIGN KEY (administrator_id) REFERENCES public.administrators(id) ON DELETE CASCADE;


--
-- Name: admin_privs_privilege_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: buffalo
--

ALTER TABLE ONLY public.admin_privs
    ADD CONSTRAINT admin_privs_privilege_id_fkey FOREIGN KEY (privilege_id) REFERENCES public.privileges(id) ON DELETE CASCADE;


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

