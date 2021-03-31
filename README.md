WAFLab 🐾
====

WAFLab is a web-based platform for testing WAFs.

## Live Demo

https://waflab.org/

## Architecture

WAFLab contains 3 parts:

Name | Description | Language | Source code | Release
----|------|----|----|----
Server-frontend | Web frontend UI for WAFLab server-side | Javascript + React + Ant Design | https://github.com/microsoft/waflab/tree/master/web | N/A
Server-backend | RESTful API backend for WAFLab server-side | Golang + Beego + MySQL | https://github.com/microsoft/waflab | N/A

## Installation

### Server-side

- Get the code:

git clone this repo.

- Prepare a [Xorm ORM](https://gitea.com/xorm/xorm) supported database (MySQL is recommended), replace `root:123@tcp(localhost:3306)/` in [conf/app.conf](https://github.com/microsoft/waflab/blob/master/conf/app.conf) with your own connection string. WAFLab will create a database named `waflab` and necessary tables in it if not exist. All Xorm supported databases are listed [here](https://gitea.com/xorm/xorm#user-content-drivers-support).

- Run Server-backend (in port 7070):

```
go run main.go
 ```

- Run Server-frontend (in the same machine's port 7000):

```
cd web
npm install
npm start
```

- Open browser:

http://localhost:7000/

## License

This project is licensed under the [MIT license](LICENSE).

If you have any issues or feature requests, please contact us. PR is welcomed.
- https://github.com/microsoft/waflab/issues
