
DROP TABLE IF EXISTS "public"."tbl_category";
CREATE TABLE "public"."tbl_category" (
	"id" uuid NOT NULL,
	"data" jsonb NOT NULL
)
WITH (OIDS=FALSE);
ALTER TABLE "public"."tbl_category" OWNER TO "postgres";

-- ----------------------------
--  Primary key structure for table tbl_category
-- ----------------------------
ALTER TABLE "public"."tbl_category" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;
