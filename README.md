# ral2jqlog
RDS Audit Log to Json Query Log

## Usage

```
$ ral2jqlog --file testdata/test002.gz --type gzip
{"timeStamp":"20220920 05:38:07","user":"common","host":"1.1.1.1","command":"1.1.1.1","query":"set names \\'ujis\\'"}
{"timeStamp":"20220920 05:53:09","user":"common","host":"1.1.1.1","command":"1.1.1.1","query":"select * from `test` where `id` is null and `type` = \\'35\\' and `flg` = \\'1\\' limit 1"}
```
