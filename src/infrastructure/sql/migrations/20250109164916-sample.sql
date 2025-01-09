
-- +migrate Up
CREATE TABLE Todo (
    id SERIAL PRIMARY KEY,  -- 自動インクリメントの整数型ID
    title VARCHAR(255) NOT NULL,  -- タスクのタイトル
    description TEXT,  -- タスクの説明
    completed BOOLEAN DEFAULT FALSE,  -- タスクが完了したかどうか
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  -- 更新日時
);

-- +migrate Down
DROP TABLE IF EXISTS Todo;
