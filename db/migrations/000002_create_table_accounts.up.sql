CREATE SEQUENCE IF NOT EXISTS account_seq;

CREATE TABLE IF NOT EXISTS accounts (
  id UUID DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY, 
  name TEXT NOT NULL, 
  nik TEXT NOT NULL, 
  phone_number TEXT NOT NULL, 
  account_number TEXT NOT NULL DEFAULT CONCAT('', LPAD(NEXTVAL('account_seq')::TEXT, 15, '0')), 
  balance NUMERIC(14, 2), 
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(), 
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
