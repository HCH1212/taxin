-- 1. 确保安装pgvector扩展
CREATE EXTENSION IF NOT EXISTS vector;

-- 2. 创建验证JSON的函数(需要在创建表之前)
CREATE OR REPLACE FUNCTION is_json(text) RETURNS BOOLEAN AS $$
DECLARE
    v_json json;
BEGIN
    v_json := $1::json;
    RETURN TRUE;
EXCEPTION WHEN others THEN
    RETURN FALSE;
END;
$$ LANGUAGE plpgsql IMMUTABLE;

-- 3. 创建用户表
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    "like" TEXT NOT NULL,
    like_embedding vector (768), -- 使用小写vector类型
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT chk_like_json CHECK (is_json ("like"))
);

-- 4. 创建自动更新timestamp的触发器函数
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 5. 创建触发器
CREATE TRIGGER update_users_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- 6. 创建索引
CREATE INDEX idx_users_user_id ON users (user_id);

-- 7. 创建向量索引(确保lists参数适合您的数据量)
-- 一般建议lists = rows/1000，但不大于2000
CREATE INDEX idx_users_like_embedding ON users USING ivfflat (
    like_embedding vector_cosine_ops
)
WITH (lists = 100);