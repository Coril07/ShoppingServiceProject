## 背景
大学時代に、javaをバックエンド開発言語にして、ウェブアプリケーションを開発しました。近年、golang言語の高性能や優れる並行処理などのおかげでより多くの企業がgolangを導入し始めています。それで、時代のトレンドにしたがうために、galangの勉強と練習は必要と思っています。ですので、このプロジェクトをしました。  
概要:  
このプロジェクトはvue,golang,Typescriptなどに基づいて自社商品販売のウェブサイトを真似るために、作られたアプリケーションです。このアプリは権限によって、二つの部分に分けられます。管理者権限とユーザー権限です。  
共通権限では主に四つの機能があります。一つ目は会員登録です。会員登録では,ユーザーネーム、パスワードやメールアドレスなどを入力することで、利用者は会員になることができます。二つ目はログインです。ログイン後、利用者にはユーザー権限が与えられ、会員として買い物をすることができます。三つ目は個人情報の更新です。最後は商品の閲覧です。この部分ではページング、カテゴリーに基づく商品の絞り込み、商品詳細情報の閲覧等を実現しました。  
管理者権限を持つユーザーは在庫管理ができます。具体的には、商品とカテゴリーのアップロード、更新ができます。  
最後に、ユーザー権限を持つ利用者が買い物かごの管理と注文管理ができます。  

## 目次
```
app
│  go.mod
│  go.sum
│  package-lock.json
│  package.json
│  shopping_pro.sql　　//sqlファイル
│  
├─client
│  │  env.d.ts
│  │  index.html　　　//入口
│  │  package-lock.json
│  │  package.json
│  │  tsconfig.app.json
│  │  tsconfig.json
│  │  tsconfig.node.json
│  │  vite.config.ts
│  │  
│  ├─.vscode
│  │      
│  ├─node_modules
│  │     
│  ├─public
│  └─src
│      │  App.vue
│      │  main.ts
│      │  router.ts　　　
│      │  store.ts　　　//グローバル変数
│      │  
│      ├─assets　　　//静的リソースのディレクトリ
│      │      cm01.jpg
│      │      cm02.jpg
│      │      cm03.jpg
│      │      cm04.jpg
│      │      iphone15ケース.jpg
│      │      iphone15ケースwafwaf.jpg
│      │      logo.png
│      │      Microsoft QZI-00020 Surface Laptop 5.jpg
│      │      ぱいぱい茶ぱいぱい.jpg
│      │      冷蔵庫.webp
│      │      冷蔵庫test0.webp
│      │      如果如果.txt
│      │      柴犬.jpg
│      │      柴犬test.jpg
│      │      茶.jpg
│      │      辛いもの.jpg
│      │      辛いもの２.jpg
│      │      
│      ├─components
│      └─views   //ビュー
│              admin.vue
│              cart.vue
│              Index.vue
│              login.vue
│              log_out.vue
│              order.vue
│              personalfile.vue
│              product.vue
│              sign_up.vue
│              
│                  
└─server
    │  main.go
    │  
    ├─config
    │      config.go
    │      config.yaml
    │      
    ├─controllers  //コントロール
    │  │  router.go　　
    │  │  
    │  ├─cart
    │  │      controller.go
    │  │      types.go
    │  │      
    │  ├─category
    │  │      controller.go
    │  │      types.go
    │  │      
    │  ├─order
    │  │      controller.go
    │  │      types.go
    │  │      
    │  ├─product
    │  │      controller.go
    │  │      types.go
    │  │      
    │  └─user
    │          controller.go
    │          types.go
    │          
    ├─docs
    ├─domain　　　　//モデルと対応するフック、エラー処理するための対象や方法、sql操作など
    │  ├─cart
    │  │      entity.go
    │  │      errors.go
    │  │      hook.go
    │  │      repository.go
    │  │      service.go
    │  │      
    │  ├─category
    │  │      entity.go
    │  │      errors.go
    │  │      repository.go
    │  │      service.go
    │  │      
    │  ├─order
    │  │      entity.go
    │  │      errors.go
    │  │      hook.go
    │  │      order_item_repository.go
    │  │      order_repository.go
    │  │      service.go
    │  │      
    │  ├─product
    │  │      entity.go
    │  │      errors.go
    │  │      hook.go
    │  │      repository.go
    │  │      service.go
    │  │      
    │  └─user
    │          entity.go
    │          errors.go
    │          hook.go
    │          repository.go
    │          service.go
    │          validation.go
    │          
    └─utils　　//ツール
        ├─api_helper
        │      error_handler.go
        │      query_helper.go
        │      types.go
        │      
        ├─csv_helper
        │      csv_read.go
        │      
        ├─database_handler
        │      db_handler.go
        │      
        ├─hash
        │      hash.go
        │      
        ├─jwt
        │      jwt_helper.go
        │      
        ├─middleware
        │      middleware.go
        │      
        ├─pagination
        │      page.go
        │      
        └─shutdown
                shutdown.go
```
## アプリケーションのフレームワーク：

![フレームワーク](https://github.com/Coril07/ShoppingServiceProject/assets/114814470/d274edf7-05e6-4329-bf9e-da60af18e1f4)

## 説明
html,css,Javascript,vueを使用してuiデザインを作成しました。この部分はMVCモデルのフレームワークにおけるビューにあたります。クライアントは、サーバーサイドからのサービスを獲得するために、サーバーサイドにHTTPリクエストを送信します。そして、コントロールはHTTPリクエストに基づいて要求されたサービスに対応したモデルを探し出し、対応処理を行います。この対応処理にはMYSQLに対する処理が含まれる可能性があります。最後に、コントロールは処理した結果をHTTPレスポンスで返信します。
 
## このプロジェクトのメリット:
構成フレームワークは筋が通る.目次がサービス別に分けられているので、プロジェクトメンバーができたら、正確に理解しすぐに開発業務に携わることができます。
