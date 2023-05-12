package outaction

var OutActionEntityTableDDL = `
CREATE TABLE cpx_out_action (
class_id   varchar(30) NOT NULL,
action_id  int         NOT NULL,
name       varchar(80) NOT NULL,
format     varchar(20) NULL,
value      varchar(80) NOT NULL,
field_name varchar(40) NULL,

CONSTRAINT PK_class_id_action_id PRIMARY KEY (class_id, action_id),

CONSTRAINT FK_cpx_out_Action_doc_class FOREIGN KEY (class_id)
    REFERENCES cpx_doc_class (class_id)
);
`
var OutActionEntityTableDropDDL = `DROP TABLE cpx_out_action`
