# 私の初めてのデプロイの実験

縛り

- サーバーはインタネット接続不可
- デプロイ先サーバーは少し古めの RHEL (v8.9)
- sudo できない。root 権限行使はハードルが高い

構成

- リバースプロキシ (proxy)
- アプリケーション (app)

## 初期構築

### 1. ホスト

1. `make build build-image` を実行
    - 下記ファイルが生成される
        - `rpm/dist/my-first-deployment-experiment-*.x86_64.rpm`
        - `app/test-app.tar.xz`
2. `rpm/dist/my-first-deployment-experiment-*.x86_64.rpm` `app/test-app.tar.xz` を対象サーバーへコピー
    - 例) `scp rpm/dist/my-first-deployment-experiment-*.x86_64.rpm app/test-app.tar.xz 192.168.100.174:`
    - `scp` の利用は好まれないが、簡易に例示するために利用

### 2. ゲスト

1. `su` で root にスイッチする
2. 作業ユーザを `adm` グループに所属させる
    - 例) `usermod -aGadm ysuzuki`
3. `rpm` コマンドでRPMをインストールする
    - `rpm -ivh ./my-first-deployment-experiment-*.x86_64.rpm`
    - NOTE: `yum` や `dnf` はインタネット通信が必要なため、使えない
    - NOTE: RPM インストールによりサービス有効化(起動)される
4. `exit` または Ctrl+d で root ユーザを抜ける
5. `newgrp adm` や再ログインにより、対話セッションが `adm` グループに属するようにする
6. test-app.tar.xz を /usr/local/share/test へ移動(またはコピー)する
    - 例) `mv test-app.tar.xz /usr/local/share/test/`
7. `test-app` サービスを起動
    - `systemctl start test-app`
    - NOTE: RPM インストール直後はイメージファイルが無いため、起動していないことを想定
8. テスト
    - `curl localhost`

# Hint

- RPM を改変した場合、バージョンを上げること。または、既存 RPM 削除したうえで再インストールすること
