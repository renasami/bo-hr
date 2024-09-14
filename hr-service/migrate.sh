#!/bin/bash

# データベースの接続情報
DB_HOST="localhost"
DB_PORT="5432"
DB_USER="user"
DB_NAME="hrdb"
DB_PASSWORD="password"

# SQLファイルのディレクトリ
SQL_DIR="./migrations"

# 環境変数としてPostgreSQLのパスワードを設定
export PGPASSWORD=$DB_PASSWORD

# SQLファイルを順次実行
for sql_file in $SQL_DIR/*.sql; do
  if [ -f "$sql_file" ]; then
    echo "Applying migration: $sql_file"
    
    # psqlコマンドを使用してSQLファイルをデータベースに適用
    psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f "$sql_file"
    
    if [ $? -eq 0 ]; then
      echo "Successfully applied $sql_file"
    else
      echo "Error applying $sql_file"
      exit 1
    fi
  fi
done

echo "All migrations applied successfully."
