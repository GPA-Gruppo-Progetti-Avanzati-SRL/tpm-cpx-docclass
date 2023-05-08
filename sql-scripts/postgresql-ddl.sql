--
--
DROP SEQUENCE If Exists cpx_ndx_item;
DROP TABLE If Exists cpx_out_action;
DROP TABLE If Exists cpx_doc_class;
DROP TABLE If Exists cpx_field_dictionary;
DROP TABLE If Exists cpx_registry;

CREATE TABLE cpx_doc_class
(
    class_id    varchar(15)   NOT NULL primary key,
    name        varchar(40)   NOT NULL,
    childIds    varchar(80)   NULL,
    extension   varchar(3)    NOT NULL,
    cod_cliente varchar(10)   NOT NULL,
    max_cpx     int           NOT NULL,
    max_docs    int           NOT NULL,
    max_size    int           NOT NULL,
    platform    varchar(10)   NOT NULL,
    servizio    varchar(10)   NOT NULL,
    procedura   varchar(10)   NOT NULL,
    version     varchar(10)   NOT NULL,
    pkg_layout  varchar(2)    NULL,
    sql_query   varchar(1024) NOT NULL
);

CREATE SEQUENCE cpx_seq;

CREATE TABLE cpx_ndx_item
(
    class_id      varchar(15) NOT NULL,
    ndx_id        int         NOT NULL,
    name          varchar(40) NOT NULL,
    type          varchar(10) NULL,
    data_type     varchar(10) NULL,
    format        varchar(10) NULL,
    value         varchar(40) NOT NULL,
    source_format varchar(10) NULL,
    required      bit         ,

    CONSTRAINT PK_class_id_ndx_id PRIMARY KEY (class_id, ndx_id),

    CONSTRAINT FK_cpx_ndx_item_doc_class FOREIGN KEY (class_id)
        REFERENCES cpx_doc_class (class_id)
);

CREATE TABLE cpx_out_action
(
    class_id   varchar(15) NOT NULL,
    action_id  int         NOT NULL,
    name       varchar(40) NOT NULL,
    format     varchar(10) NULL,
    value      varchar(40) NOT NULL,
    field_name varchar(40)  NULL,

    CONSTRAINT PK_class_id_action_id PRIMARY KEY (class_id, action_id),

    CONSTRAINT FK_cpx_out_Action_doc_class FOREIGN KEY (class_id)
        REFERENCES cpx_doc_class (class_id)

);

CREATE TABLE cpx_field_dictionary
(
    name             varchar(40) NOT NULL primary key,
    document_mapping varchar(80) NOT NULL
);

CREATE TABLE cpx_registry
(
    name     varchar(40) NOT NULL,
    class_id varchar(15) NOT NULL,
    enabled  bit         not null,

    CONSTRAINT PK_class_id_action_id PRIMARY KEY (name, class_id)
);