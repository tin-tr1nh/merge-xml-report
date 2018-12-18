# Merge XML Reports Tool #

Old phpunit not support to export cov report file so you can't merge them with phpcov tool
This will help to merge report file in xml file format.

## How to build yourself ##

```bash
docker build . -t merge-xml-report
```

And run it with this

```bash
docker run -v "<path_to_your_result_dir>:/files/result" -v "<path_to_your_reports_dir>:/files/reports" merge-xml-report
```

## How to run for docker hub ##

Log in docker hub of Hamee

```bash
docker login
```

And run it with this

```bash
docker run -v "<path_to_your_result_dir>:/files/result" -v "<path_to_your_reports_dir>:/files/reports" trinhtin/merge-xml-report:latest
```
