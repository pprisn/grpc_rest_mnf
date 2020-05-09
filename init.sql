--
-- PostgreSQL database dump
--

-- Dumped from database version 11.7 (Ubuntu 11.7-2.pgdg18.04+1)
-- Dumped by pg_dump version 11.7 (Ubuntu 11.7-2.pgdg18.04+1)

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

SET default_with_oids = false;

--
-- Name: part; Type: TABLE; Schema: public; Owner: ppri
--

CREATE TABLE public.part (
    id character varying(100) NOT NULL,
    mnf_id character varying(100),
    vendor_code character varying(1000),
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.part OWNER TO ppri;

--
-- Name: part_manufacturer; Type: TABLE; Schema: public; Owner: ppri
--

CREATE TABLE public.part_manufacturer (
    id character varying(100) NOT NULL,
    name character varying(1000) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.part_manufacturer OWNER TO ppri;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: ppri
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO ppri;

--
-- Data for Name: part; Type: TABLE DATA; Schema: public; Owner: ppri
--

COPY public.part (id, mnf_id, vendor_code, created_at, deleted_at) FROM stdin;
eb56e00d-ae91-46ad-87bf-f87a7b7c6498	2fb7f969-2fdd-4f43-a501-4e9e3582ebd7	0358861	2020-05-04 15:48:55.12077+03	\N
70a5cbf2-1a43-45f2-b8e6-abe4b3ac9a48	2fb7f969-2fdd-4f43-a501-4e9e3582ebd7	0358417	2020-05-04 15:48:55.12077+03	\N
b946e027-ec79-4174-ac80-c182f5b38de9	2fb7f969-2fdd-4f43-a501-4e9e3582ebd7	0357230	2020-05-04 15:48:55.12077+03	\N
4daa005b-eb52-4adf-ab49-1f11484d4c66	3b0ee2e5-7629-4256-aba3-526e797ca89a	R2432	2020-05-04 15:48:55.12077+03	\N
15d1ec95-a77f-448f-8416-5db96857b011	3b0ee2e5-7629-4256-aba3-526e797ca89a	R2447	2020-05-04 15:48:55.12077+03	\N
80ec4bea-d95d-4ed7-b891-f008433ccfe5	3b0ee2e5-7629-4256-aba3-526e797ca89a	R0984	2020-05-04 15:48:55.12077+03	\N
41728754-6c58-4b5c-85d8-522b34f3bb59	4bf9a24f-dcb3-41ec-99b9-0e3c4e27bb93	777386	2020-05-04 15:48:55.12077+03	\N
d47ed2fa-d741-4f5f-8054-b2fec40ad831	4bf9a24f-dcb3-41ec-99b9-0e3c4e27bb93	789360	2020-05-04 15:48:55.12077+03	\N
7c5dafad-dd5d-4c97-b6e0-186485d5e60e	4bf9a24f-dcb3-41ec-99b9-0e3c4e27bb93	691957	2020-05-04 15:48:55.12077+03	\N
7ab69a9e-5cb9-4ac0-9294-5736e15e7785	3824a6ee-2020-4df1-8548-22651ac9f714	2329500	2020-05-04 15:48:55.12077+03	\N
c53b2902-5569-462b-a1e8-7fba3e3fca54	3824a6ee-2020-4df1-8548-22651ac9f714	2326700	2020-05-04 15:48:55.12077+03	\N
d87315f9-a1b6-48da-8092-6ff062e1f65c	3824a6ee-2020-4df1-8548-22651ac9f714	2327600	2020-05-04 15:48:55.12077+03	\N
452196fc-0dcb-412d-bbcb-2f14b15d398a	eb9c98ea-4e29-4549-862d-837119082be5	TS01387	2020-05-04 15:48:55.12077+03	\N
7c78f55e-0e98-495a-b2a5-811fbe346a8a	eb9c98ea-4e29-4549-862d-837119082be5	TS01375	2020-05-04 15:48:55.12077+03	\N
76e1e433-a93f-4ce3-b9be-1bd72a81b22b	6926e659-564e-4c6d-adbb-07c9023fdbb1	1021022	2020-05-04 15:48:55.12077+03	\N
4ca8e2a7-fd8e-4ad7-a18d-0a5581fb41fe	6926e659-564e-4c6d-adbb-07c9023fdbb1	1010973	2020-05-04 15:48:55.12077+03	\N
6905bba0-3e32-473b-95db-fce146dc0ba4	6926e659-564e-4c6d-adbb-07c9023fdbb1	1010592	2020-05-04 15:48:55.12077+03	\N
7a5a4ca4-badc-40f7-95a0-cc0d8f4bb606	18a0add4-8c00-4e02-b567-a5fcb1c8c744	538460	2020-05-04 15:48:55.12077+03	\N
a2b782f0-27d3-4267-a3a5-87ac44386891	18a0add4-8c00-4e02-b567-a5fcb1c8c744	538462	2020-05-04 15:48:55.12077+03	\N
8fc4931b-4fc9-44f0-b2c2-dd79b581ea99	18a0add4-8c00-4e02-b567-a5fcb1c8c744	538475	2020-05-04 15:48:55.12077+03	\N
ee82c9a2-d49c-4d8f-be8d-6726919e3df2	63c721af-1c59-4ddd-90a4-9744f02b6ee6	11837	2020-05-04 15:48:55.12077+03	\N
1f487c0f-d868-4615-b23c-248db2d3915d	63c721af-1c59-4ddd-90a4-9744f02b6ee6	9643	2020-05-04 15:48:55.12077+03	\N
23065854-e969-46ec-80de-24e4fc849f10	63c721af-1c59-4ddd-90a4-9744f02b6ee6	10490	2020-05-04 15:48:55.12077+03	\N
bb205e34-5af7-4f34-80c6-20058f779ad5	9b92b546-74be-4628-8d99-b19e171beda1	305175	2020-05-04 15:48:55.12077+03	\N
5c3d4d68-4a0d-4921-a3c9-d3bce1c1a04c	9b92b546-74be-4628-8d99-b19e171beda1	308061	2020-05-04 15:48:55.12077+03	\N
4ca600f6-cc6e-4958-a4f5-80527eb56cb0	9b92b546-74be-4628-8d99-b19e171beda1	308407	2020-05-04 15:48:55.12077+03	\N
71f0d6b4-38f4-4d74-abba-9aee07d306ec	a0d0fe7e-6f44-4a79-9d4c-df951a0f8631	2187133	2020-05-04 15:48:55.12077+03	\N
f179a18b-eb36-40e6-9f3e-3aaed667c99d	a0d0fe7e-6f44-4a79-9d4c-df951a0f8631	2204463	2020-05-04 15:48:55.12077+03	\N
b1a88b27-f4b7-4d61-81d5-db6926e89e4e	a0d0fe7e-6f44-4a79-9d4c-df951a0f8631	2187033	2020-05-04 15:48:55.12077+03	\N
c029f11e-3532-4a0a-b165-dbdaad17ecde	7da4bd51-67b6-4b82-ad48-fa2bfa37a503	EMS	2020-05-06 22:15:05.22501+03	\N
6934e2a3-bc5e-4e3f-9fdd-9b07ebd5e623	eb9c98ea-4e29-4549-862d-837119082be5	TS01345	2020-05-04 15:48:55.12077+03	2020-05-08 18:39:40.820661+03
b1f353cf-7c08-4d2e-af35-f66fd4965c16	390a13d0-ca93-4766-809b-a9994b964f38	GazelNEXT_Update	2020-05-06 22:59:02.714501+03	\N
3b5ce0eb-7544-4660-9b84-f83319b3e3c8	971a3522-47f3-4dd6-b075-5c60f3f60136	VAZ2115_update	2020-05-06 22:59:02.651218+03	\N
b582b36d-d00f-4e0d-bd84-7d3d894618b7	971a3522-47f3-4dd6-b075-5c60f3f60136	VESTA_update	2020-05-06 22:59:02.704834+03	\N
\.


--
-- Data for Name: part_manufacturer; Type: TABLE DATA; Schema: public; Owner: ppri
--

COPY public.part_manufacturer (id, name, created_at) FROM stdin;
2fb7f969-2fdd-4f43-a501-4e9e3582ebd7	Continental	2020-05-04 15:31:14.383516+03
3b0ee2e5-7629-4256-aba3-526e797ca89a	Yokohama	2020-05-04 15:31:14.383516+03
4bf9a24f-dcb3-41ec-99b9-0e3c4e27bb93	Michlen	2020-05-04 15:31:14.383516+03
3824a6ee-2020-4df1-8548-22651ac9f714	Pirelli	2020-05-04 15:31:14.383516+03
eb9c98ea-4e29-4549-862d-837119082be5	TOYO	2020-05-04 15:31:14.383516+03
6926e659-564e-4c6d-adbb-07c9023fdbb1	HANKOOK	2020-05-04 15:31:14.383516+03
18a0add4-8c00-4e02-b567-a5fcb1c8c744	GOODYEAR	2020-05-04 15:31:14.383516+03
63c721af-1c59-4ddd-90a4-9744f02b6ee6	Bridgestone	2020-05-04 15:31:14.383516+03
9b92b546-74be-4628-8d99-b19e171beda1	DUNLOP	2020-05-04 15:31:14.383516+03
a0d0fe7e-6f44-4a79-9d4c-df951a0f8631	KUMHO	2020-05-04 15:31:14.383516+03
96f3f504-bdda-4ac1-a4c1-f0b8e195f8af	Rosinca	2020-05-05 13:02:04.54421+03
390a13d0-ca93-4766-809b-a9994b964f38	GAZ	2020-05-05 16:04:11.295085+03
7da4bd51-67b6-4b82-ad48-fa2bfa37a503	Russian Post	2020-05-05 15:24:17.056683+03
971a3522-47f3-4dd6-b075-5c60f3f60136	АвтоВАЗ update test	2020-05-05 16:04:11.133433+03
35ed13f5-aca4-4b33-b9e1-aad4023d6d77	УАЗ update test	2020-05-05 16:04:11.251617+03
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: ppri
--

COPY public.schema_migrations (version, dirty) FROM stdin;
20200504153203	f
\.


--
-- Name: part_manufacturer part_manufacturer_pkey; Type: CONSTRAINT; Schema: public; Owner: ppri
--

ALTER TABLE ONLY public.part_manufacturer
    ADD CONSTRAINT part_manufacturer_pkey PRIMARY KEY (id);


--
-- Name: part part_pkey; Type: CONSTRAINT; Schema: public; Owner: ppri
--

ALTER TABLE ONLY public.part
    ADD CONSTRAINT part_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: ppri
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: mnfid_index; Type: INDEX; Schema: public; Owner: ppri
--

CREATE INDEX mnfid_index ON public.part USING btree (mnf_id);


--
-- Name: part delete_part; Type: RULE; Schema: public; Owner: ppri
--

CREATE RULE delete_part AS
    ON DELETE TO public.part DO INSTEAD  UPDATE public.part SET deleted_at = now()
  WHERE ((part.id)::text = (old.id)::text);


--
-- Name: part part_mnf_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ppri
--

ALTER TABLE ONLY public.part
    ADD CONSTRAINT part_mnf_id_fkey FOREIGN KEY (mnf_id) REFERENCES public.part_manufacturer(id);


--
-- PostgreSQL database dump complete
--

