--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3 (Debian 15.3-1.pgdg110+1)
-- Dumped by pg_dump version 15.3 (Debian 15.3-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    name character varying(255) NOT NULL,
    family character varying(255) NOT NULL,
    id integer NOT NULL,
    age integer NOT NULL,
    sex character varying(255) NOT NULL,
    createdat time without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (name, family, id, age, sex, createdat) FROM stdin;
AQADANIAL	BOZORG	1111	21	MALE	20:10:01.340512
HAJMEHDI	AKBAR	1112	20	MALE	20:10:21.374697
HAJHOSSEIN	RAEIS	1113	22	MALE	20:10:36.554669
YAZDAN	PANDA	1114	22	MALE	20:10:59.089122
SEYYEDE_MANA	ABBASZADEH_QOMI	1115	21	FEMALE	20:11:26.810975
MICHELE	OBAMA	1116	72	FEMALE	20:11:47.763411
A	TEXAS	1117	55	FEMALE	20:12:06.835642
L	RHOADS	1118	32	FEMALE	20:12:21.210096
DONALD	TRUMP	1119	65	MALE	20:12:38.933047
JOE	BIDEN	1120	73	MALE	20:12:53.903547
RYAN	GOSSLING	1121	36	MALE	20:13:09.626096
EMMA	STONES	1122	33	MALE	20:13:28.016564
KATY	PERRY	1123	46	FEMALE	20:14:17.522923
BILLIE	EILISH	1124	19	FEMALE	20:14:49.238055
CR	RONALDO	1125	43	MALE	20:16:19.857988
KARIM	BENZEMA	1135	48	MALE	20:16:47.81234
DASHALI	MANSOORIAN	1235	53	MALE	20:17:22.670049
\.


--
-- Name: users users_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_id_key UNIQUE (id);


--
-- PostgreSQL database dump complete
--

