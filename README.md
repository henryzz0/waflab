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

### Prerequisites

- [Go](https://golang.org/)
- [Node.js](https://nodejs.org/)
- [Docker](https://www.docker.com/)

### Server-side

Get the source code from Github via Git

```bash
git clone https://github.com/microsoft/waflab.git
```

#### Set up the database

WAFLab use database to store generated testcases and test results.

Prepare a [Xorm ORM](https://gitea.com/xorm/xorm) supported database (MySQL is recommended), replace `root:123@tcp(localhost:3306)/` in [conf/app.conf](https://github.com/microsoft/waflab/blob/master/conf/app.conf) with your own connection string. WAFLab will create a database named `waflab` and necessary tables in it if not exist. All Xorm supported databases are listed [here](https://gitea.com/xorm/xorm#user-content-drivers-support).

#### Setup Server-backend

Git clone the [OWASP ModSecurity Core Rule Set (CRS)](https://github.com/coreruleset/coreruleset) and [WAFBench](https://github.com/microsoft/WAFBench) under a same directory

```bash
git clone https://github.com/microsoft/WAFBench.git
git clone https://github.com/coreruleset/coreruleset.git
```

Pick the CRS version you would like to use. We use CRS v3.2 as an example here.

```bash
cd coreruleset
git checkout --track origin/v3.2/master
```

Set the ```CodeBaseDir``` inside ```waflab/util/const.go``` to the directory of WAFBench and CRS.

```Go
const CodeBaseDir = "DIRECTORY/OF/WAFBENCH/AND/CRS"
```

Run Server-backend (at port 7070 by default):

```bash
cd waflab
go run main.go
 ```

#### Setup Server-frontend

Install all Node.js dependencies with npm.

```bash
cd waflab/web
npm install
```

Run Server-frontend (at port 7000 by default)

```bash
npm start
```

WAFLab web interface is now avaliable at [http://localhost:7000/](http://localhost:7000/).

## License

This project is licensed under the [MIT license](LICENSE).

If you have any issues or feature requests, please contact us. PR is welcomed.
- https://github.com/microsoft/waflab/issues

## Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

## Trademarks

This project may contain trademarks or logos for projects, products, or services. Authorized use of Microsoft
trademarks or logos is subject to and must follow
[Microsoft's Trademark & Brand Guidelines](https://www.microsoft.com/en-us/legal/intellectualproperty/trademarks/usage/general).
Use of Microsoft trademarks or logos in modified versions of this project must not cause confusion or imply Microsoft sponsorship.
Any use of third-party trademarks or logos are subject to those third-party's policies.
