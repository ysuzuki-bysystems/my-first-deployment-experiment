# 私の初めてのデプロイの実験

縛り

- サーバーはインタネット接続不可
- デプロイ先サーバーは少し古めの RHEL

構成

- リバースプロキシ
- アプリケーション

## 1. 仮のアプリケーションをデプロイ

### ホスト

1. `make` を実行
2. `rpm/dist/my-first-deployment-experiment-*-0.x86_64.rpm` を対象サーバーへコピー

### 対象サーバー

nginx を仮のアプリケーションとする。
NOTE: 80 番ポートで HTTP をサービスするものであれば何でも可
NOTE: ここは縛り解除

1. `sudo dnf install ./my-first-deployment-experiment-*.x86_64.rpm` で RPM をインストール
2. `podman image pull docker.io/library/nginx:latest` を実行
    - nginx イメージを PULL
3. `podman image tag docker.io/library/nginx:latest test-app:latest`
    - イメージのタグ付け
4. `podman image save test-app:latest --format docker-archive --output /usr/local/share/test/app.tar` を実行
    - イメージを tar として出力
5. `xz /usr/local/share/test/app.tar` を実行
    - イメージを xz 圧縮
6. `sudo systemctl start test-app.target` 起動
7. `curl localhost` 動作確認

## 2. 本アプリケーションをデプロイ

### ホスト

1. `make build-image` を実行
2. `app/test-app.tar.xz` を対象サーバへコピー
3. `sudo systemctl restart test-homepage-image.service` を実行
    - イメージを読み込み直す
3. `sudo systemctl restart test-homepage.service` を実行
    - コンテナを再起動

# Hint

- RPM を改変した場合、バージョンを上げること。または、既存 RPM 削除したうえで再インストールすること
