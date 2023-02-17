package span

const tableCreateSql = "CREATE TABLE `{TABLE_NAME}` (\n" +
	"`id` bigint(11) NOT NULL AUTO_INCREMENT,\n" +
	"`operation_name` varchar(255) NOT NULL,\n" +
	"`start_time` bigint(13) NOT NULL,\n" +
	"`finish_time` bigint(13) NOT NULL,\n" +
	"`parent_id` varchar(10) NOT NULL,\n" +
	"`span_id` varchar(10) NOT NULL,\n" +
	"`trace_id` varchar(30) NOT NULL,\n" +
	"`application` varchar(20) NOT NULL,\n" +
	"`application_group` varchar(20) NOT NULL,\n" +
	"`span_kind` varchar(10) NOT NULL,\n" +
	"`error` tinyint(3) NOT NULL,\n" +
	"`log_datas` text NOT NULL,\n" +
	"`tags` text NOT NULL,\n" +
	"`app_instance` varchar(100) NOT NULL,\n" +
	"`component` varchar(100) NOT NULL DEFAULT '',\n" +
	"PRIMARY KEY (`id`),\n" +
	"KEY `idx_start_time` (`start_time`) USING BTREE,\n" +
	"KEY `span_kind` (`span_kind`) USING HASH,\n" +
	"KEY `idx_error` (`error`) USING HASH,\n" +
	"KEY `trace_id` (`trace_id`) USING HASH,\n" +
	"KEY `idx_group` (`application_group`) USING HASH,\n" +
	"KEY `idx_appName` (`application`) USING HASH\n" +
	") ENGINE=InnoDB DEFAULT CHARSET=utf16;"
