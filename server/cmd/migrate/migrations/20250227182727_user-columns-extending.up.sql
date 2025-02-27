ALTER TABLE user_details
ADD COLUMN IF NOT EXISTS gender varchar(3) NOT NULL DEFAULT 'NS' CONSTRAINT gender_check CHECK (gender in ('M', 'F', 'NS')),
ADD COLUMN IF NOT EXISTS birth_year INT NOT NULL CHECK (birth_year <= EXTRACT(YEAR FROM CURRENT_DATE) - 13);
