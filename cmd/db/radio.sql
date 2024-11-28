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
    length integer,
    expireson timestamp without time zone,
    lastplayed timestamp without time zone,
    dateadded timestamp without time zone,
    spinstoday integer,
    spinsweek integer,
    spinstotal integer
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
    "position" integer,
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
1	stationid	Station ID
2	imagingid	Imaging ID
3	next	Play Next
4	live	Live
5	addstop	Advertising Top Of Hour
6	adds	Advertising
7	top40	Top 40 Music
8	roots	Roots Music
9	music	Music Library
10	live	Live
11	fill	Fill Schedule
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

COPY public.inventory (rowid, category, artist, song, album, length, expireson, lastplayed, dateadded, spinstoday, spinsweek, spinstotal) FROM stdin;
\.


--
-- Data for Name: schedule; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schedule (rowid, days, hours, "position", categories, spinstoplay) FROM stdin;
1	MON	00	1	stationid	1
2	MON	00	2	addstop	2
3	MON	00	3	adds	1
4	MON	00	4	top40	3
5	MON	00	5	music	1
6	MON	00	6	imagingid	1
7	MON	00	7	roots	1
8	MON	00	8	top40	1
9	MON	00	9	music	1
10	MON	00	10	imagingid	1
11	MON	00	11	top40	1
12	MON	00	12	roots	1
13	MON	00	13	adds	1
14	MON	00	14	music	3
15	MON	00	15	fill	1
16	MON	01	1	stationid	1
17	MON	01	2	addstop	2
18	MON	01	3	adds	1
19	MON	04	4	top40	3
20	MON	01	5	music	1
21	MON	01	6	imagingid	1
22	MON	01	7	roots	1
23	MON	01	8	top40	1
24	MON	01	9	music	1
25	MON	01	10	imagingid	1
26	MON	01	11	top40	1
27	MON	01	12	roots	1
28	MON	01	13	adds	1
29	MON	01	14	music	3
30	MON	01	15	fill	1
31	MON	02	1	stationid	1
32	MON	02	2	addstop	2
33	MON	02	3	adds	1
34	MON	02	4	top40	3
35	MON	02	5	music	1
36	MON	02	6	imagingid	1
37	MON	02	7	roots	1
38	MON	02	8	top40	1
39	MON	02	9	music	1
40	MON	02	10	imagingid	1
41	MON	02	11	top40	1
42	MON	02	12	roots	1
43	MON	02	13	adds	1
44	MON	02	14	music	3
45	MON	02	15	fill	1
46	MON	03	1	stationid	1
47	MON	03	2	addstop	2
48	MON	03	3	adds	1
49	MON	03	4	top40	3
50	MON	03	5	music	1
51	MON	03	6	imagingid	1
52	MON	03	7	roots	1
53	MON	03	8	top40	1
54	MON	03	9	music	1
55	MON	03	10	imagingid	1
56	MON	03	11	top40	1
57	MON	03	12	roots	1
58	MON	03	13	adds	1
59	MON	03	14	music	3
60	MON	03	15	fill	1
61	MON	04	1	stationid	1
62	MON	04	2	addstop	2
63	MON	04	3	adds	1
64	MON	04	4	top40	3
65	MON	04	5	music	1
66	MON	04	6	imagingid	1
67	MON	04	7	roots	1
68	MON	04	8	top40	1
69	MON	04	9	music	1
70	MON	04	10	imagingid	1
71	MON	04	11	top40	1
72	MON	04	12	roots	1
73	MON	04	13	adds	1
74	MON	04	14	music	3
75	MON	04	15	fill	1
76	MON	05	1	stationid	1
77	MON	05	2	addstop	2
78	MON	05	3	adds	1
79	MON	05	4	top40	3
80	MON	05	5	music	1
81	MON	05	6	imagingid	1
82	MON	05	7	roots	1
83	MON	05	8	top40	1
84	MON	05	9	music	1
85	MON	05	10	imagingid	1
86	MON	05	11	top40	1
87	MON	05	12	roots	1
88	MON	05	13	adds	1
89	MON	05	14	music	3
90	MON	05	15	fill	1
91	MON	06	1	stationid	1
92	MON	06	2	addstop	2
93	MON	06	3	adds	1
94	MON	06	4	top40	3
95	MON	06	5	music	1
96	MON	06	6	imagingid	1
97	MON	06	7	roots	1
98	MON	06	8	top40	1
99	MON	06	9	music	1
100	MON	06	10	imagingid	1
101	MON	06	11	top40	1
102	MON	06	12	roots	1
103	MON	06	13	adds	1
104	MON	06	14	music	3
105	MON	06	15	fill	1
106	MON	07	1	stationid	1
107	MON	07	2	addstop	2
108	MON	07	3	adds	1
109	MON	07	4	top40	3
110	MON	07	5	music	1
111	MON	07	6	imagingid	1
112	MON	07	7	roots	1
113	MON	07	8	top40	1
114	MON	07	9	music	1
115	MON	07	10	imagingid	1
116	MON	07	11	top40	1
117	MON	07	12	roots	1
118	MON	07	13	adds	1
119	MON	07	14	music	3
120	MON	07	15	fill	1
121	MON	08	1	stationid	1
122	MON	08	2	addstop	2
123	MON	08	3	adds	1
124	MON	08	4	top40	3
125	MON	08	5	music	1
126	MON	08	6	imagingid	1
127	MON	08	7	roots	1
128	MON	08	8	top40	1
129	MON	08	9	music	1
130	MON	08	10	imagingid	1
131	MON	08	11	top40	1
132	MON	08	12	roots	1
133	MON	08	13	adds	1
134	MON	08	14	music	3
135	MON	08	15	fill	1
136	MON	09	1	stationid	1
137	MON	09	2	addstop	2
138	MON	09	3	adds	1
139	MON	09	4	top40	3
140	MON	09	5	music	1
141	MON	09	6	imagingid	1
142	MON	09	7	roots	1
143	MON	09	8	top40	1
144	MON	09	9	music	1
145	MON	09	10	imagingid	1
146	MON	09	11	top40	1
147	MON	09	12	roots	1
148	MON	09	13	adds	1
149	MON	09	14	music	3
150	MON	09	15	fill	1
151	MON	10	1	stationid	1
152	MON	10	2	addstop	2
153	MON	10	3	adds	1
154	MON	10	4	top40	3
155	MON	10	5	music	1
156	MON	10	6	imagingid	1
157	MON	10	7	roots	1
158	MON	10	8	top40	1
159	MON	10	9	music	1
160	MON	10	10	imagingid	1
161	MON	10	11	top40	1
162	MON	10	12	roots	1
163	MON	10	13	adds	1
164	MON	10	14	music	3
165	MON	10	15	fill	1
166	MON	11	1	stationid	1
167	MON	11	2	addstop	2
168	MON	11	3	adds	1
169	MON	11	4	top40	3
170	MON	11	5	music	1
171	MON	11	6	imagingid	1
172	MON	11	7	roots	1
173	MON	11	8	top40	1
174	MON	11	9	music	1
175	MON	11	10	imagingid	1
176	MON	11	11	top40	1
177	MON	11	12	roots	1
178	MON	11	13	adds	1
179	MON	11	14	music	3
180	MON	11	15	fill	1
181	MON	12	1	stationid	1
182	MON	12	2	addstop	2
183	MON	12	3	adds	1
184	MON	12	4	top40	3
185	MON	12	5	music	1
186	MON	12	6	imagingid	1
187	MON	12	7	roots	1
188	MON	12	8	top40	1
189	MON	12	9	music	1
190	MON	12	10	imagingid	1
191	MON	12	11	top40	1
192	MON	12	12	roots	1
193	MON	12	13	adds	1
194	MON	12	14	music	3
195	MON	12	15	fill	1
196	MON	13	1	stationid	1
197	MON	13	2	addstop	2
198	MON	13	3	adds	1
199	MON	13	4	top40	3
200	MON	13	5	music	1
201	MON	13	6	imagingid	1
202	MON	13	7	roots	1
203	MON	13	8	top40	1
204	MON	13	9	music	1
205	MON	13	10	imagingid	1
206	MON	13	11	top40	1
207	MON	13	12	roots	1
208	MON	13	13	adds	1
209	MON	13	14	music	3
210	MON	13	15	fill	1
211	MON	14	1	stationid	1
212	MON	14	2	addstop	2
213	MON	14	3	adds	1
214	MON	14	4	top40	3
215	MON	14	5	music	1
216	MON	14	6	imagingid	1
217	MON	14	7	roots	1
218	MON	14	8	top40	1
219	MON	14	9	music	1
220	MON	14	10	imagingid	1
221	MON	14	11	top40	1
222	MON	14	12	roots	1
223	MON	14	13	adds	1
224	MON	14	14	music	3
225	MON	14	15	fill	1
226	MON	15	1	stationid	1
227	MON	15	2	addstop	2
228	MON	15	3	adds	1
229	MON	15	4	top40	3
230	MON	15	5	music	1
231	MON	15	6	imagingid	1
232	MON	15	7	roots	1
233	MON	15	8	top40	1
234	MON	15	9	music	1
235	MON	15	10	imagingid	1
236	MON	15	11	top40	1
237	MON	15	12	roots	1
238	MON	15	13	adds	1
239	MON	15	14	music	3
240	MON	15	15	fill	1
241	MON	16	1	stationid	1
242	MON	16	2	addstop	2
243	MON	16	3	adds	1
244	MON	16	4	top40	3
245	MON	16	5	music	1
246	MON	16	6	imagingid	1
247	MON	16	7	roots	1
248	MON	16	8	top40	1
249	MON	16	9	music	1
250	MON	16	10	imagingid	1
251	MON	16	11	top40	1
252	MON	16	12	roots	1
253	MON	16	13	adds	1
254	MON	16	14	music	3
255	MON	16	15	fill	1
256	MON	17	1	stationid	1
257	MON	17	2	addstop	2
258	MON	17	3	adds	1
259	MON	17	4	top40	3
260	MON	17	5	music	1
261	MON	17	6	imagingid	1
262	MON	17	7	roots	1
263	MON	17	8	top40	1
264	MON	17	9	music	1
265	MON	17	10	imagingid	1
266	MON	17	11	top40	1
267	MON	17	12	roots	1
268	MON	17	13	adds	1
269	MON	17	14	music	3
270	MON	17	15	fill	1
271	MON	18	1	stationid	1
272	MON	18	2	addstop	2
273	MON	18	3	adds	1
274	MON	18	4	top40	3
275	MON	18	5	music	1
276	MON	18	6	imagingid	1
277	MON	18	7	roots	1
278	MON	18	8	top40	1
279	MON	18	9	music	1
280	MON	18	10	imagingid	1
281	MON	18	11	top40	1
282	MON	18	12	roots	1
283	MON	18	13	adds	1
284	MON	18	14	music	3
285	MON	18	15	fill	1
286	MON	19	1	stationid	1
287	MON	19	2	addstop	2
288	MON	19	3	adds	1
289	MON	19	4	top40	3
290	MON	19	5	music	1
291	MON	19	6	imagingid	1
292	MON	19	7	roots	1
293	MON	19	8	top40	1
294	MON	19	9	music	1
295	MON	19	10	imagingid	1
296	MON	19	11	top40	1
297	MON	19	12	roots	1
298	MON	19	13	adds	1
299	MON	19	14	music	3
300	MON	19	15	fill	1
301	MON	20	1	stationid	1
302	MON	20	2	addstop	2
303	MON	20	3	adds	1
304	MON	20	4	top40	3
305	MON	20	5	music	1
306	MON	20	6	imagingid	1
307	MON	20	7	roots	1
308	MON	20	8	top40	1
309	MON	20	9	music	1
310	MON	20	10	imagingid	1
311	MON	20	11	top40	1
312	MON	20	12	roots	1
313	MON	20	13	adds	1
314	MON	20	14	music	3
315	MON	20	15	fill	1
316	MON	21	1	stationid	1
317	MON	21	2	addstop	2
318	MON	21	3	adds	1
319	MON	21	4	top40	3
320	MON	21	5	music	1
321	MON	21	6	imagingid	1
322	MON	21	7	roots	1
323	MON	21	8	top40	1
324	MON	21	9	music	1
325	MON	21	10	imagingid	1
326	MON	21	11	top40	1
327	MON	21	12	roots	1
328	MON	21	13	adds	1
329	MON	21	14	music	3
330	MON	21	15	fill	1
331	MON	22	1	stationid	1
332	MON	22	2	addstop	2
333	MON	22	3	adds	1
334	MON	22	4	top40	3
335	MON	22	5	music	1
336	MON	22	6	imagingid	1
337	MON	22	7	roots	1
338	MON	22	8	top40	1
339	MON	22	9	music	1
340	MON	22	10	imagingid	1
341	MON	22	11	top40	1
342	MON	22	12	roots	1
343	MON	22	13	adds	1
344	MON	22	14	music	3
345	MON	22	15	fill	1
346	MON	23	1	stationid	1
347	MON	23	2	addstop	2
348	MON	23	3	adds	1
349	MON	23	4	top40	3
350	MON	23	5	music	1
351	MON	23	6	imagingid	1
352	MON	23	7	roots	1
353	MON	23	8	top40	1
354	MON	23	9	music	1
355	MON	23	10	imagingid	1
356	MON	23	11	top40	1
357	MON	23	12	roots	1
358	MON	23	13	adds	1
359	MON	23	14	music	3
360	MON	23	15	fill	1
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
-- Name: inventorybycategorydate; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX inventorybycategorydate ON public.inventory USING btree (category, lastplayed);


--
-- Name: scheduleindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX scheduleindex ON public.schedule USING btree (days, hours, "position");


--
-- PostgreSQL database dump complete
--

