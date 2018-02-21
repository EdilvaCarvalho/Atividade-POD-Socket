CREATE TABLE pessoa
    (
        uid serial NOT NULL,
        nome character varying(100) NOT NULL,
        curso character varying(100) NOT NULL,
        cidade character varying(100) NOT NULL,
        CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
    )
    WITH (OIDS=FALSE);

select * from pessoa