create table IF NOT EXISTS account_mutations (
  id uuid default uuid_generate_v4() not null primary key,
  account_number text not null,
  transaction_date timestamp with time zone,
  transaction_code text not null,
  amount numeric(14, 2),
  created_at timestamp with time zone,
  updated_at timestamp with time zone
);