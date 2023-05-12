package docclass

var DocClassEntityTableDDL = `
CREATE TABLE cpx_doc_class (
class_id     varchar(30)   NOT NULL primary key,
name         varchar(80)   NOT NULL,
child_ids    varchar(80)   NULL,
extension    varchar(3)    NOT NULL,
cod_cliente  varchar(10)   NOT NULL,
max_cpx      int           NOT NULL,
max_docs     int           NOT NULL,
max_size     int           NOT NULL,
platform     varchar(10)   NOT NULL,
servizio     varchar(10)   NOT NULL,
procedura    varchar(10)   NOT NULL,
version      varchar(10)   NOT NULL,
pkg_layout   varchar(2)    NULL,
sql_query    varchar(1024) NOT NULL,
distinta_ged bool          not null
);
`
var DocClassEntityTableDropDDL = `DROP TABLE cpx_doc_class`
