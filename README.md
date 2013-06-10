GoJsonSqlNullObj
================

Sql NullObj(NullString etc) can friendly output to Json for Golang


About & Usage
================

Once your DB have field allow null, you have to use sql.NullXXX, but it can't directly output to Json as you wish. for example http://play.golang.org/p/BWrNfU1vjP

GoJsonSqlNullObj can help you deal with this, simplely change the import package from "database/sql" to github.com/soarcn/GoJsonSqlNullObj after get this toolkit


Alternative
================

If you don't want to use this toolkit, you can refer this code snippet
http://play.golang.org/p/88lg8EgOw2
