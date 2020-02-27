1. `conf.yaml` is needed to start the project
```sh
echo "
db_user: your db user
db_password: your db password
db_port: 5432
db_name: sso_dev

server:
  mode: debug

crendentials:
  jwt_secret: YOUR jwt code

aliyun_sms:
  region: 
  ak: 
  aks: 
  sign_name: 
  template_code: 
" >> conf.yaml
```
