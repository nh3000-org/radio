--
-- PostgreSQL database dump
--

-- Dumped from database version 17.1 (Ubuntu 17.1-1.pgdg22.04+1)
-- Dumped by pg_dump version 17.1 (Ubuntu 17.1-1.pgdg22.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    rowid integer NOT NULL,
    id character varying(32),
    description text NOT NULL
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: categories_rowid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_rowid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_rowid_seq OWNER TO postgres;

--
-- Name: categories_rowid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_rowid_seq OWNED BY public.categories.rowid;


--
-- Name: days; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.days (
    rowid integer NOT NULL,
    id character(3),
    description text NOT NULL,
    dayofweek integer
);


ALTER TABLE public.days OWNER TO postgres;

--
-- Name: days_rowid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.days_rowid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.days_rowid_seq OWNER TO postgres;

--
-- Name: days_rowid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.days_rowid_seq OWNED BY public.days.rowid;


--
-- Name: hours; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.hours (
    rowid integer NOT NULL,
    id character(2),
    description text NOT NULL
);


ALTER TABLE public.hours OWNER TO postgres;

--
-- Name: hours_rowid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.hours_rowid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.hours_rowid_seq OWNER TO postgres;

--
-- Name: hours_rowid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.hours_rowid_seq OWNED BY public.hours.rowid;


--
-- Name: inventory; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.inventory (
    rowid integer NOT NULL,
    category character varying(32) NOT NULL,
    artist text NOT NULL,
    song text NOT NULL,
    album text,
    songlength integer,
    rndorder text,
    expireson timestamp without time zone,
    lastplayed timestamp without time zone,
    dateadded timestamp without time zone,
    spinstoday integer,
    spinsweek integer,
    spinstotal integer,
    sourcelink text
);


ALTER TABLE public.inventory OWNER TO postgres;

--
-- Name: inventory_rowid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.inventory_rowid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.inventory_rowid_seq OWNER TO postgres;

--
-- Name: inventory_rowid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.inventory_rowid_seq OWNED BY public.inventory.rowid;


--
-- Name: schedule; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schedule (
    rowid integer NOT NULL,
    days character varying(3),
    hours character(2),
    "position" character(2),
    categories character varying(32),
    spinstoplay integer
);


ALTER TABLE public.schedule OWNER TO postgres;

--
-- Name: schedule_rowid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.schedule_rowid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.schedule_rowid_seq OWNER TO postgres;

--
-- Name: schedule_rowid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.schedule_rowid_seq OWNED BY public.schedule.rowid;


--
-- Name: categories rowid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN rowid SET DEFAULT nextval('public.categories_rowid_seq'::regclass);


--
-- Name: days rowid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.days ALTER COLUMN rowid SET DEFAULT nextval('public.days_rowid_seq'::regclass);


--
-- Name: hours rowid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.hours ALTER COLUMN rowid SET DEFAULT nextval('public.hours_rowid_seq'::regclass);


--
-- Name: inventory rowid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inventory ALTER COLUMN rowid SET DEFAULT nextval('public.inventory_rowid_seq'::regclass);


--
-- Name: schedule rowid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedule ALTER COLUMN rowid SET DEFAULT nextval('public.schedule_rowid_seq'::regclass);


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (rowid, id, description) FROM stdin;
1	STATIONID	Station ID
2	IMAGINGID	Imaging ID
3	NEXT	Play Next
4	LIVE	Live
5	ADDSTOH	Advertising Top Of Hour
6	ADDSDRIVETIME	Advertising Drive Time
7	ADDS	Advertising
8	TOP40	Top 40 MUSIC
9	ROOTS	Roots MUSIC
10	MUSIC	Music Library
11	FILLTOTOH	Fill To TOH Schedule
\.


--
-- Data for Name: days; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.days (rowid, id, description, dayofweek) FROM stdin;
1	MON	Monday	1
2	TUE	Tuesday	2
3	WED	Wednesday	3
4	THU	Thursday	4
5	FRI	Friday	5
6	SAT	Saturday	6
7	SUN	Sunday	7
\.


--
-- Data for Name: hours; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.hours (rowid, id, description) FROM stdin;
1	00	Hour Part 00
2	01	Hour Part 01
3	02	Hour Part 02
4	03	Hour Part 03
5	04	Hour Part 04
6	05	Hour Part 05
7	06	Hour Part 06
8	07	Hour Part 07
9	08	Hour Part 08
10	09	Hour Part 09
11	10	Hour Part 10
12	11	Hour Part 11
13	12	Hour Part 12
14	13	Hour Part 13
15	14	Hour Part 14
16	15	Hour Part 15
17	16	Hour Part 16
18	17	Hour Part 17
19	18	Hour Part 18
20	19	Hour Part 19
21	20	Hour Part 20
22	21	Hour Part 21
23	22	Hour Part 22
24	23	Hour Part 23
\.


--
-- Data for Name: inventory; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.inventory (rowid, category, artist, song, album, songlength, rndorder, expireson, lastplayed, dateadded, spinstoday, spinsweek, spinstotal, sourcelink) FROM stdin;
\.


--
-- Data for Name: schedule; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schedule (rowid, days, hours, "position", categories, spinstoplay) FROM stdin;
1	MON	00	01	STATIONID	1
2	MON	00	02	ADDSTOH	2
3	MON	00	03	ADDS	1
4	MON	00	04	TOP40	3
5	MON	00	05	MUSIC	1
6	MON	00	06	IMAGINGID	1
7	MON	00	07	ROOTS	1
8	MON	00	08	TOP40	1
9	MON	00	09	MUSIC	1
10	MON	00	10	IMAGINGID	1
11	MON	00	11	TOP40	1
12	MON	00	12	ROOTS	1
13	MON	00	13	ADDS	1
14	MON	00	14	MUSIC	3
15	MON	00	15	FILLTOTOH	1
16	MON	01	01	STATIONID	1
17	MON	01	02	ADDSTOH	2
18	MON	01	03	ADDS	1
19	MON	04	04	TOP40	3
20	MON	01	05	MUSIC	1
21	MON	01	06	IMAGINGID	1
22	MON	01	07	ROOTS	1
23	MON	01	08	TOP40	1
24	MON	01	09	MUSIC	1
25	MON	01	10	IMAGINGID	1
26	MON	01	11	TOP40	1
27	MON	01	12	ROOTS	1
28	MON	01	13	ADDS	1
29	MON	01	14	MUSIC	3
30	MON	01	15	FILLTOTOH	1
31	MON	02	01	STATIONID	1
32	MON	02	02	ADDSTOH	2
33	MON	02	03	ADDS	1
34	MON	02	04	TOP40	3
35	MON	02	05	MUSIC	1
36	MON	02	06	IMAGINGID	1
37	MON	02	07	ROOTS	1
38	MON	02	08	TOP40	1
39	MON	02	09	MUSIC	1
40	MON	02	10	IMAGINGID	1
41	MON	02	11	TOP40	1
42	MON	02	12	ROOTS	1
43	MON	02	13	ADDS	1
44	MON	02	14	MUSIC	3
45	MON	02	15	FILLTOTOH	1
46	MON	03	01	STATIONID	1
47	MON	03	02	ADDSTOH	2
48	MON	03	03	ADDS	1
49	MON	03	04	TOP40	3
50	MON	03	05	MUSIC	1
51	MON	03	06	IMAGINGID	1
52	MON	03	07	ROOTS	1
53	MON	03	08	TOP40	1
54	MON	03	09	MUSIC	1
55	MON	03	10	IMAGINGID	1
56	MON	03	11	TOP40	1
57	MON	03	12	ROOTS	1
58	MON	03	13	ADDS	1
59	MON	03	14	MUSIC	3
60	MON	03	15	FILLTOTOH	1
61	MON	04	01	STATIONID	1
62	MON	04	02	ADDSDRIVETIME	2
63	MON	04	03	ADDS	1
64	MON	04	04	TOP40	3
65	MON	04	05	MUSIC	1
66	MON	04	06	IMAGINGID	1
67	MON	04	07	ROOTS	1
68	MON	04	08	TOP40	1
69	MON	04	09	MUSIC	1
70	MON	04	10	IMAGINGID	1
71	MON	04	11	TOP40	1
72	MON	04	12	ROOTS	1
73	MON	04	13	ADDS	1
74	MON	04	14	MUSIC	3
75	MON	04	15	FILLTOTOH	1
76	MON	05	01	STATIONID	1
77	MON	05	02	ADDSDRIVETIME	2
78	MON	05	03	ADDS	1
79	MON	05	04	TOP40	3
80	MON	05	05	MUSIC	1
81	MON	05	06	IMAGINGID	1
82	MON	05	07	ROOTS	1
83	MON	05	08	TOP40	1
84	MON	05	09	MUSIC	1
85	MON	05	10	IMAGINGID	1
86	MON	05	11	TOP40	1
87	MON	05	12	ROOTS	1
88	MON	05	13	ADDS	1
89	MON	05	14	MUSIC	3
90	MON	05	15	FILLTOTOH	1
91	MON	06	01	STATIONID	1
92	MON	06	02	ADDSDRIVETIME	2
93	MON	06	03	ADDS	1
94	MON	06	04	TOP40	3
95	MON	06	05	MUSIC	1
96	MON	06	06	IMAGINGID	1
97	MON	06	07	ROOTS	1
98	MON	06	08	TOP40	1
99	MON	06	09	MUSIC	1
100	MON	06	10	IMAGINGID	1
101	MON	06	11	TOP40	1
102	MON	06	12	ROOTS	1
103	MON	06	13	ADDS	1
104	MON	06	14	MUSIC	3
105	MON	06	15	FILLTOTOH	1
106	MON	07	01	STATIONID	1
107	MON	07	02	ADDSDRIVETIME	2
108	MON	07	03	ADDS	1
109	MON	07	04	TOP40	3
110	MON	07	05	MUSIC	1
111	MON	07	06	IMAGINGID	1
112	MON	07	07	ROOTS	1
113	MON	07	08	TOP40	1
114	MON	07	09	MUSIC	1
115	MON	07	10	IMAGINGID	1
116	MON	07	11	TOP40	1
117	MON	07	12	ROOTS	1
118	MON	07	13	ADDS	1
119	MON	07	14	MUSIC	3
120	MON	07	15	FILLTOTOH	1
121	MON	08	01	STATIONID	1
122	MON	08	02	ADDSDRIVETIME	2
123	MON	08	03	ADDS	1
124	MON	08	04	TOP40	3
125	MON	08	05	MUSIC	1
126	MON	08	06	IMAGINGID	1
127	MON	08	07	ROOTS	1
128	MON	08	08	TOP40	1
129	MON	08	09	MUSIC	1
130	MON	08	10	IMAGINGID	1
131	MON	08	11	TOP40	1
132	MON	08	12	ROOTS	1
133	MON	08	13	ADDS	1
134	MON	08	14	MUSIC	3
135	MON	08	15	FILLTOTOH	1
136	MON	09	01	STATIONID	1
137	MON	09	02	ADDSDRIVETIME	2
138	MON	09	03	ADDS	1
139	MON	09	04	TOP40	3
140	MON	09	05	MUSIC	1
141	MON	09	06	IMAGINGID	1
142	MON	09	07	ROOTS	1
143	MON	09	08	TOP40	1
144	MON	09	09	MUSIC	1
145	MON	09	10	IMAGINGID	1
146	MON	09	11	TOP40	1
147	MON	09	12	ROOTS	1
148	MON	09	13	ADDS	1
149	MON	09	14	MUSIC	3
150	MON	09	15	FILLTOTOH	1
151	MON	10	01	STATIONID	1
152	MON	10	02	ADDSDRIVETIME	2
153	MON	10	03	ADDS	1
154	MON	10	04	TOP40	3
155	MON	10	05	MUSIC	1
156	MON	10	06	IMAGINGID	1
157	MON	10	07	ROOTS	1
158	MON	10	08	TOP40	1
159	MON	10	09	MUSIC	1
160	MON	10	10	IMAGINGID	1
161	MON	10	11	TOP40	1
162	MON	10	12	ROOTS	1
163	MON	10	13	ADDS	1
164	MON	10	14	MUSIC	3
165	MON	10	15	FILLTOTOH	1
166	MON	11	01	STATIONID	1
167	MON	11	02	ADDSDRIVETIME	2
168	MON	11	03	ADDS	1
169	MON	11	04	TOP40	3
170	MON	11	05	MUSIC	1
171	MON	11	06	IMAGINGID	1
172	MON	11	07	ROOTS	1
173	MON	11	08	TOP40	1
174	MON	11	09	MUSIC	1
175	MON	11	10	IMAGINGID	1
176	MON	11	11	TOP40	1
177	MON	11	12	ROOTS	1
178	MON	11	13	ADDS	1
179	MON	11	14	MUSIC	3
180	MON	11	15	FILLTOTOH	1
181	MON	12	01	STATIONID	1
182	MON	12	02	ADDSDRIVETIME	2
183	MON	12	03	ADDS	1
184	MON	12	04	TOP40	3
185	MON	12	05	MUSIC	1
186	MON	12	06	IMAGINGID	1
187	MON	12	07	ROOTS	1
188	MON	12	08	TOP40	1
189	MON	12	09	MUSIC	1
190	MON	12	10	IMAGINGID	1
191	MON	12	11	TOP40	1
192	MON	12	12	ROOTS	1
193	MON	12	13	ADDS	1
194	MON	12	14	MUSIC	3
195	MON	12	15	FILLTOTOH	1
196	MON	13	01	STATIONID	1
197	MON	13	02	ADDSDRIVETIME	2
198	MON	13	03	ADDS	1
199	MON	13	04	TOP40	3
200	MON	13	05	MUSIC	1
201	MON	13	06	IMAGINGID	1
202	MON	13	07	ROOTS	1
203	MON	13	08	TOP40	1
204	MON	13	09	MUSIC	1
205	MON	13	10	IMAGINGID	1
206	MON	13	11	TOP40	1
207	MON	13	12	ROOTS	1
208	MON	13	13	ADDS	1
209	MON	13	14	MUSIC	3
210	MON	13	15	FILLTOTOH	1
211	MON	14	01	STATIONID	1
212	MON	14	02	ADDSTOH	2
213	MON	14	03	ADDS	1
214	MON	14	04	TOP40	3
215	MON	14	05	MUSIC	1
216	MON	14	06	IMAGINGID	1
217	MON	14	07	ROOTS	1
218	MON	14	08	TOP40	1
219	MON	14	09	MUSIC	1
220	MON	14	10	IMAGINGID	1
221	MON	14	11	TOP40	1
222	MON	14	12	ROOTS	1
223	MON	14	13	ADDS	1
224	MON	14	14	MUSIC	3
225	MON	14	15	FILLTOTOH	1
226	MON	15	01	STATIONID	1
227	MON	15	02	ADDSTOH	2
228	MON	15	03	ADDS	1
229	MON	15	04	TOP40	3
230	MON	15	05	MUSIC	1
231	MON	15	06	IMAGINGID	1
232	MON	15	07	ROOTS	1
233	MON	15	08	TOP40	1
234	MON	15	09	MUSIC	1
235	MON	15	10	IMAGINGID	1
236	MON	15	11	TOP40	1
237	MON	15	12	ROOTS	1
238	MON	15	13	ADDS	1
239	MON	15	14	MUSIC	3
240	MON	15	15	FILLTOTOH	1
241	MON	16	01	STATIONID	1
242	MON	16	02	ADDSDRIVETIME	2
243	MON	16	03	ADDS	1
244	MON	16	04	TOP40	3
245	MON	16	05	MUSIC	1
246	MON	16	06	IMAGINGID	1
247	MON	16	07	ROOTS	1
248	MON	16	08	TOP40	1
249	MON	16	09	MUSIC	1
250	MON	16	10	IMAGINGID	1
251	MON	16	11	TOP40	1
252	MON	16	12	ROOTS	1
253	MON	16	13	ADDS	1
254	MON	16	14	MUSIC	3
255	MON	16	15	FILLTOTOH	1
256	MON	17	01	STATIONID	1
257	MON	17	02	ADDSDRIVETIME	2
258	MON	17	03	ADDS	1
259	MON	17	04	TOP40	3
260	MON	17	05	MUSIC	1
261	MON	17	06	IMAGINGID	1
262	MON	17	07	ROOTS	1
263	MON	17	08	TOP40	1
264	MON	17	09	MUSIC	1
265	MON	17	10	IMAGINGID	1
266	MON	17	11	TOP40	1
267	MON	17	12	ROOTS	1
268	MON	17	13	ADDS	1
269	MON	17	14	MUSIC	3
270	MON	17	15	FILLTOTOH	1
271	MON	18	01	STATIONID	1
272	MON	18	02	ADDSDRIVETIME	2
273	MON	18	03	ADDS	1
274	MON	18	04	TOP40	3
275	MON	18	05	MUSIC	1
276	MON	18	06	IMAGINGID	1
277	MON	18	07	ROOTS	1
278	MON	18	08	TOP40	1
279	MON	18	09	MUSIC	1
280	MON	18	10	IMAGINGID	1
281	MON	18	11	TOP40	1
282	MON	18	12	ROOTS	1
283	MON	18	13	ADDS	1
284	MON	18	14	MUSIC	3
285	MON	18	15	FILLTOTOH	1
286	MON	19	01	STATIONID	1
287	MON	19	02	ADDSTOH	2
288	MON	19	03	ADDS	1
289	MON	19	04	TOP40	3
290	MON	19	05	MUSIC	1
291	MON	19	06	IMAGINGID	1
292	MON	19	07	ROOTS	1
293	MON	19	08	TOP40	1
294	MON	19	09	MUSIC	1
295	MON	19	10	IMAGINGID	1
296	MON	19	11	TOP40	1
297	MON	19	12	ROOTS	1
298	MON	19	13	ADDS	1
299	MON	19	14	MUSIC	3
300	MON	19	15	FILLTOTOH	1
301	MON	20	01	STATIONID	1
302	MON	20	02	ADDSTOH	2
303	MON	20	03	ADDS	1
304	MON	20	04	TOP40	3
305	MON	20	05	MUSIC	1
306	MON	20	06	IMAGINGID	1
307	MON	20	07	ROOTS	1
308	MON	20	08	TOP40	1
309	MON	20	09	MUSIC	1
310	MON	20	10	IMAGINGID	1
311	MON	20	11	TOP40	1
312	MON	20	12	ROOTS	1
313	MON	20	13	ADDS	1
314	MON	20	14	MUSIC	3
315	MON	20	15	FILLTOTOH	1
316	MON	21	01	STATIONID	1
317	MON	21	02	ADDSTOH	2
318	MON	21	03	ADDS	1
319	MON	21	04	TOP40	3
320	MON	21	05	MUSIC	1
321	MON	21	06	IMAGINGID	1
322	MON	21	07	ROOTS	1
323	MON	21	08	TOP40	1
324	MON	21	09	MUSIC	1
325	MON	21	10	IMAGINGID	1
326	MON	21	11	TOP40	1
327	MON	21	12	ROOTS	1
328	MON	21	13	ADDS	1
329	MON	21	14	MUSIC	3
330	MON	21	15	FILLTOTOH	1
331	MON	22	01	STATIONID	1
332	MON	22	02	ADDSTOH	2
333	MON	22	03	ADDS	1
334	MON	22	04	TOP40	3
335	MON	22	05	MUSIC	1
336	MON	22	06	IMAGINGID	1
337	MON	22	07	ROOTS	1
338	MON	22	08	TOP40	1
339	MON	22	09	MUSIC	1
340	MON	22	10	IMAGINGID	1
341	MON	22	11	TOP40	1
342	MON	22	12	ROOTS	1
343	MON	22	13	ADDS	1
344	MON	22	14	MUSIC	3
345	MON	22	15	FILLTOTOH	1
346	MON	23	01	STATIONID	1
347	MON	23	02	ADDSTOH	2
348	MON	23	03	ADDS	1
349	MON	23	04	TOP40	3
350	MON	23	05	MUSIC	1
351	MON	23	06	IMAGINGID	1
352	MON	23	07	ROOTS	1
353	MON	23	08	TOP40	1
354	MON	23	09	MUSIC	1
355	MON	23	10	IMAGINGID	1
356	MON	23	11	TOP40	1
357	MON	23	12	ROOTS	1
358	MON	23	13	ADDS	1
359	MON	23	14	MUSIC	3
360	MON	23	15	FILLTOTOH	1
\.


--
-- Name: categories_rowid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_rowid_seq', 11, true);


--
-- Name: days_rowid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.days_rowid_seq', 7, true);


--
-- Name: hours_rowid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.hours_rowid_seq', 24, true);


--
-- Name: inventory_rowid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.inventory_rowid_seq', 1, false);


--
-- Name: schedule_rowid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.schedule_rowid_seq', 360, true);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (rowid);


--
-- Name: days days_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.days
    ADD CONSTRAINT days_pkey PRIMARY KEY (rowid);


--
-- Name: hours hours_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.hours
    ADD CONSTRAINT hours_pkey PRIMARY KEY (rowid);


--
-- Name: inventory inventory_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inventory
    ADD CONSTRAINT inventory_pkey PRIMARY KEY (rowid);


--
-- Name: schedule schedule_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedule
    ADD CONSTRAINT schedule_pkey PRIMARY KEY (rowid);


--
-- Name: categoriesindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX categoriesindex ON public.categories USING btree (id);


--
-- Name: dayindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX dayindex ON public.days USING btree (dayofweek);


--
-- Name: hoursindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX hoursindex ON public.hours USING btree (id);


--
-- Name: inventorybyartist; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX inventorybyartist ON public.inventory USING btree (artist, song);


--
-- Name: inventorybycategorysong; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX inventorybycategorysong ON public.inventory USING btree (category, song);


--
-- Name: scheduleindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX scheduleindex ON public.schedule USING btree (days, hours, "position");


--
-- PostgreSQL database dump complete
--

