#-------------------------------------------------------
# 製品環境
#-------------------------------------------------------
FROM nginx:1.25.3 as releaser

# 設定ファイルのコピー
COPY conf/* /etc/nginx/conf.d/
# SSL証明書と秘密鍵のコピー
COPY ssl/* /etc/nginx/ssl/
