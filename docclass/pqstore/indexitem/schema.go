package indexitem

var IndexItemEntityTableDDL = `
CREATE TABLE cpx_ndx_item (
class_id      varchar(30) NOT NULL,
ndx_id        int         NOT NULL,
name          varchar(80) NOT NULL,
type          varchar(10) NULL,
data_type     varchar(10) NULL,
format        varchar(20) NULL,
value         varchar(80) NOT NULL,
source_format varchar(30) NULL,
required      bool        ,  
CONSTRAINT PK_class_id_ndx_id PRIMARY KEY (class_id, ndx_id),  
CONSTRAINT FK_cpx_ndx_item_doc_class FOREIGN KEY (class_id)
    REFERENCES cpx_doc_class (class_id)
);
`
var IndexItemEntityTableDropDDL = `DROP TABLE cpx_ndx_item`
