-- public.tr_lookup definition

-- Drop table

-- DROP TABLE public.tr_lookup;

CREATE TABLE public.tr_lookup (
	code varchar NOT NULL,
	value varchar NOT NULL,
	description varchar NOT NULL,
	notes varchar NULL,
	status_id varchar NOT NULL DEFAULT 1,
	created_by varchar NOT NULL,
	created_date timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_name varchar NOT NULL,
	created_from varchar NOT NULL,
	modified_by varchar NULL,
	modified_date varchar NULL DEFAULT CURRENT_TIMESTAMP,
	modified_name varchar NULL,
	modified_from varchar NULL
);


-- public.tx_weight_record definition

-- Drop table

-- DROP TABLE public.tx_weight_record;

CREATE TABLE public.tx_weight_record (
	weight_record_id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
	"date" date NOT NULL,
	max int2 NOT NULL DEFAULT 0,
	min int2 NOT NULL DEFAULT 0,
	status_id varchar NOT NULL DEFAULT 1,
	created_by varchar NOT NULL,
	created_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_name varchar NOT NULL,
	created_from varchar NOT NULL,
	modified_by varchar NULL,
	modified_date timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	modified_name varchar NULL,
	modified_from varchar NULL
);