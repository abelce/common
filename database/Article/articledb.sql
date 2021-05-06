
DROP TABLE IF EXISTS "public"."tbl_article";
CREATE TABLE "public"."tbl_article" (
	"id" uuid NOT NULL,
	"data" jsonb NOT NULL
)
WITH (OIDS=FALSE);
ALTER TABLE "public"."tbl_article" OWNER TO "postgres";

-- ----------------------------
--  Primary key structure for table tbl_article
-- ----------------------------
ALTER TABLE "public"."tbl_article" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;
