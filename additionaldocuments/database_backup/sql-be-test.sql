--
-- PostgreSQL database cluster dump
--

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Drop databases (except postgres and template1)
--

DROP DATABASE fiber_starter;




--
-- Drop roles
--

DROP ROLE postgres;


--
-- Roles
--

CREATE ROLE postgres;
ALTER ROLE postgres WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS PASSWORD 'SCRAM-SHA-256$4096:7FuDCIh9wYLiiXLFke3Zsw==$kRp5s+ae+oCOu52N6ZopVJqYOLN62VmNGwsNtm6bIik=:BJwAXP6fFdSO4VYwp7/1JGEhGPCwyDUtXWUtWpE26mk=';

--
-- User Configurations
--








--
-- Databases
--

--
-- Database "template1" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

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

UPDATE pg_catalog.pg_database SET datistemplate = false WHERE datname = 'template1';
DROP DATABASE template1;
--
-- Name: template1; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE template1 WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE template1 OWNER TO postgres;

\connect template1

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

--
-- Name: DATABASE template1; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE template1 IS 'default template for new databases';


--
-- Name: template1; Type: DATABASE PROPERTIES; Schema: -; Owner: postgres
--

ALTER DATABASE template1 IS_TEMPLATE = true;


\connect template1

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

--
-- Name: DATABASE template1; Type: ACL; Schema: -; Owner: postgres
--

REVOKE CONNECT,TEMPORARY ON DATABASE template1 FROM PUBLIC;
GRANT CONNECT ON DATABASE template1 TO PUBLIC;


--
-- PostgreSQL database dump complete
--

--
-- Database "fiber_starter" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

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

--
-- Name: fiber_starter; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE fiber_starter WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE fiber_starter OWNER TO postgres;

\connect fiber_starter

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
-- Name: attendances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.attendances (
    attendance_id bigint NOT NULL,
    employee_id bigint,
    location_id bigint,
    absen_in timestamp with time zone,
    absen_out timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    created_by text,
    updated_by text
);


ALTER TABLE public.attendances OWNER TO postgres;

--
-- Name: attendances_attendance_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.attendances_attendance_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.attendances_attendance_id_seq OWNER TO postgres;

--
-- Name: attendances_attendance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.attendances_attendance_id_seq OWNED BY public.attendances.attendance_id;


--
-- Name: departments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.departments (
    department_id bigint NOT NULL,
    department_name text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    created_by text,
    updated_by text
);


ALTER TABLE public.departments OWNER TO postgres;

--
-- Name: departments_department_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.departments_department_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.departments_department_id_seq OWNER TO postgres;

--
-- Name: departments_department_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.departments_department_id_seq OWNED BY public.departments.department_id;


--
-- Name: employees; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.employees (
    employee_id bigint NOT NULL,
    employee_code text NOT NULL,
    employee_name text NOT NULL,
    department_id bigint NOT NULL,
    position_id bigint NOT NULL,
    superior bigint NOT NULL,
    password text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    created_by text,
    updated_by text
);


ALTER TABLE public.employees OWNER TO postgres;

--
-- Name: employees_employee_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.employees_employee_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.employees_employee_id_seq OWNER TO postgres;

--
-- Name: employees_employee_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.employees_employee_id_seq OWNED BY public.employees.employee_id;


--
-- Name: locations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.locations (
    location_id bigint NOT NULL,
    location_name text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    created_by text,
    updated_by text
);


ALTER TABLE public.locations OWNER TO postgres;

--
-- Name: locations_location_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.locations_location_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.locations_location_id_seq OWNER TO postgres;

--
-- Name: locations_location_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.locations_location_id_seq OWNED BY public.locations.location_id;


--
-- Name: positions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.positions (
    position_id bigint NOT NULL,
    department_id bigint NOT NULL,
    position_name text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    created_by text,
    updated_by text
);


ALTER TABLE public.positions OWNER TO postgres;

--
-- Name: positions_position_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.positions_position_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.positions_position_id_seq OWNER TO postgres;

--
-- Name: positions_position_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.positions_position_id_seq OWNED BY public.positions.position_id;


--
-- Name: attendances attendance_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.attendances ALTER COLUMN attendance_id SET DEFAULT nextval('public.attendances_attendance_id_seq'::regclass);


--
-- Name: departments department_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.departments ALTER COLUMN department_id SET DEFAULT nextval('public.departments_department_id_seq'::regclass);


--
-- Name: employees employee_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees ALTER COLUMN employee_id SET DEFAULT nextval('public.employees_employee_id_seq'::regclass);


--
-- Name: locations location_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.locations ALTER COLUMN location_id SET DEFAULT nextval('public.locations_location_id_seq'::regclass);


--
-- Name: positions position_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.positions ALTER COLUMN position_id SET DEFAULT nextval('public.positions_position_id_seq'::regclass);


--
-- Data for Name: attendances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.attendances (attendance_id, employee_id, location_id, absen_in, absen_out, created_at, updated_at, deleted_at, created_by, updated_by) FROM stdin;
2	20	2	2024-06-20 14:38:59.15372+00	0001-01-01 00:00:00+00	2024-06-20 14:38:59.269359+00	2024-06-20 14:38:59.269359+00	\N	22100002	
1	19	2	2024-06-20 14:36:09.869322+00	2024-06-20 14:40:12.240999+00	2024-06-20 14:36:09.982945+00	2024-06-20 14:40:12.354338+00	2024-06-20 14:42:27.125109+00	22100002	22100002
3	24	1	2024-06-20 17:27:46.680592+00	0001-01-01 00:00:00+00	2024-06-20 17:27:46.783103+00	2024-06-20 17:27:46.783103+00	\N	22100002	
4	24	2	2024-06-20 17:30:54.88801+00	0001-01-01 00:00:00+00	2024-06-20 17:30:55.003398+00	2024-06-20 17:30:55.003398+00	\N	22100002	
5	24	3	2024-06-20 17:30:59.681457+00	0001-01-01 00:00:00+00	2024-06-20 17:30:59.784821+00	2024-06-20 17:30:59.784821+00	\N	22100002	
6	24	4	2024-06-20 17:31:03.459907+00	0001-01-01 00:00:00+00	2024-06-20 17:31:03.562815+00	2024-06-20 17:31:03.562815+00	\N	22100002	
7	24	5	2024-06-20 17:31:07.652053+00	0001-01-01 00:00:00+00	2024-06-20 17:31:07.76631+00	2024-06-20 17:31:07.76631+00	\N	22100002	
\.


--
-- Data for Name: departments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.departments (department_id, department_name, created_at, updated_at, deleted_at, created_by, updated_by) FROM stdin;
1	Front Office	2024-06-20 15:30:07.811861+00	2024-06-20 15:30:07.811861+00	\N	22100002	
2	Back Office	2024-06-20 15:30:21.147622+00	2024-06-20 15:30:21.147622+00	\N	22100002	
3	Middle Office	2024-06-20 15:30:31.63168+00	2024-06-20 15:30:31.63168+00	\N	22100002	
4	Upper Office	2024-06-20 15:30:39.936096+00	2024-06-20 15:30:39.936096+00	\N	22100002	
5	In Office	2024-06-20 15:30:48.281616+00	2024-06-20 15:32:11.620911+00	2024-06-20 15:32:53.784906+00	22100002	22100002
\.


--
-- Data for Name: employees; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.employees (employee_id, employee_code, employee_name, department_id, position_id, superior, password, created_at, updated_at, deleted_at, created_by, updated_by) FROM stdin;
3			0	0	0		2024-06-20 09:57:57.222041+00	2024-06-20 09:57:57.222041+00	\N		
4			0	0	0		2024-06-20 10:01:06.244075+00	2024-06-20 10:01:06.244075+00	\N		
5			0	0	0		2024-06-20 10:03:49.745107+00	2024-06-20 10:03:49.745107+00	\N		
6			0	0	0		2024-06-20 10:03:51.693521+00	2024-06-20 10:03:51.693521+00	\N		
7			0	0	0		2024-06-20 10:15:19.084759+00	2024-06-20 10:15:19.084759+00	\N		
8			0	0	0		2024-06-20 10:16:35.302139+00	2024-06-20 10:16:35.302139+00	\N		
9			0	0	0		2024-06-20 10:17:18.621208+00	2024-06-20 10:17:18.621208+00	\N		
10			0	0	0		2024-06-20 10:18:16.0096+00	2024-06-20 10:18:16.0096+00	\N		
12	22100002	Mamat	13	11	15	ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f	2024-06-20 10:25:41.527533+00	2024-06-20 10:25:41.527533+00	\N		
13	22100003	Mamat	13	11	15	ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f	2024-06-20 12:40:00.674563+00	2024-06-20 12:40:00.674563+00	\N		
14	22100004	Memet	13	11	15	ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f	2024-06-20 12:55:14.581017+00	2024-06-20 12:55:14.581017+00	\N	22100002	
1			0	0	0		2024-06-20 09:48:14.168713+00	2024-06-20 09:48:14.168713+00	2024-06-20 12:58:22.826023+00		
2			0	0	0		2024-06-20 09:50:50.391203+00	2024-06-20 09:50:50.391203+00	2024-06-20 12:58:52.067885+00		
11	22100003	Mamat	13	11	15	341cf9022a3e6469098ed7a7d056071f73087d283407c3b108c6e0fd8996ef3c	2024-06-20 10:20:44.751741+00	2024-06-20 13:00:57.717014+00	\N		22100002
15	22100004	Memet	13	11	15	ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f	2024-06-20 13:28:20.771767+00	2024-06-20 13:28:20.771767+00	\N	22100002	
16	22100004	Memet	13	11	15	ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f	2024-06-20 13:30:45.266191+00	2024-06-20 13:30:45.266191+00	\N	22100002	
17	22100004	Memet	13	11	15	ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f	2024-06-20 13:37:05.51094+00	2024-06-20 13:37:05.51094+00	\N	22100002	
18	22100004	Memet	13	11	15	ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f	2024-06-20 13:49:03.029785+00	2024-06-20 13:49:03.029785+00	\N	22100002	
19	22100004	Memet	13	11	15	ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f	2024-06-20 13:50:06.51855+00	2024-06-20 13:50:06.51855+00	\N	22100002	
20	22100004	Memet	13	11	15	ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f	2024-06-20 13:56:32.343023+00	2024-06-20 13:56:32.343023+00	\N	22100002	
21	22100004	Memet	13	11	15	ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f	2024-06-20 13:57:54.972787+00	2024-06-20 13:57:54.972787+00	\N	22100002	
22	24060022	Memet	13	11	15	ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f	2024-06-20 14:00:21.605362+00	2024-06-20 14:00:21.605362+00	\N	22100002	
23	24060023	Memet	13	11	15	ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f	2024-06-20 14:03:13.064931+00	2024-06-20 14:03:13.064931+00	\N	22100002	
24	24060024	Ahmad	2	2	2	ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f	2024-06-20 17:26:56.282298+00	2024-06-20 17:26:56.282298+00	\N	22100002	
\.


--
-- Data for Name: locations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.locations (location_id, location_name, created_at, updated_at, deleted_at, created_by, updated_by) FROM stdin;
1	Head Office	2024-06-20 15:36:25.646237+00	2024-06-20 15:36:25.646237+00	\N	22100002	
2	KC. Jakarta	2024-06-20 15:36:40.579676+00	2024-06-20 15:36:40.579676+00	\N	22100002	
3	KC. Bandung	2024-06-20 15:36:50.547736+00	2024-06-20 15:36:50.547736+00	\N	22100002	
6	KC. Bali	2024-06-20 15:37:13.455584+00	2024-06-20 15:37:13.455584+00	\N	22100002	
5	KC. Medan	2024-06-20 15:37:06.023991+00	2024-06-20 15:38:04.500998+00	\N	22100002	22100002
4	KC. Semarang	2024-06-20 15:36:58.226767+00	2024-06-20 15:36:58.226767+00	2024-06-20 15:38:34.040546+00	22100002	
\.


--
-- Data for Name: positions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.positions (position_id, department_id, position_name, created_at, updated_at, deleted_at, created_by, updated_by) FROM stdin;
1	2	IT	2024-06-20 15:42:28.118313+00	2024-06-20 15:42:28.118313+00	\N	22100002	
2	2	Analyst	2024-06-20 15:42:37.512677+00	2024-06-20 15:42:37.512677+00	\N	22100002	
3	1	Teller	2024-06-20 15:42:46.881906+00	2024-06-20 15:42:46.881906+00	\N	22100002	
6	4	CTO	2024-06-20 15:43:28.553812+00	2024-06-20 15:43:28.553812+00	\N	22100002	
5	4	Kang Marah-Marah	2024-06-20 15:43:16.431963+00	2024-06-20 15:45:26.619971+00	\N	22100002	22100002
4	1	HR	2024-06-20 15:42:53.357213+00	2024-06-20 15:42:53.357213+00	2024-06-20 15:46:12.736686+00	22100002	
\.


--
-- Name: attendances_attendance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.attendances_attendance_id_seq', 7, true);


--
-- Name: departments_department_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.departments_department_id_seq', 5, true);


--
-- Name: employees_employee_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.employees_employee_id_seq', 24, true);


--
-- Name: locations_location_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.locations_location_id_seq', 6, true);


--
-- Name: positions_position_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.positions_position_id_seq', 6, true);


--
-- Name: attendances attendances_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.attendances
    ADD CONSTRAINT attendances_pkey PRIMARY KEY (attendance_id);


--
-- Name: departments departments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.departments
    ADD CONSTRAINT departments_pkey PRIMARY KEY (department_id);


--
-- Name: employees employees_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_pkey PRIMARY KEY (employee_id);


--
-- Name: locations locations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.locations
    ADD CONSTRAINT locations_pkey PRIMARY KEY (location_id);


--
-- Name: positions positions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.positions
    ADD CONSTRAINT positions_pkey PRIMARY KEY (position_id);


--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

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

DROP DATABASE postgres;
--
-- Name: postgres; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE postgres OWNER TO postgres;

\connect postgres

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

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database cluster dump complete
--

