# LiveATC.net Archiver

![CI](https://github.com/lorenzoaiello/liveatc-archiver/workflows/CI/badge.svg)

A simple application that is designed to be run daily that will download all of the LiveATC.net frequency recordings from the day before and store them on an FTP server.

## Kubernetes Example

```yaml
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: daily-archiver
spec:
  schedule: "0 0 * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
            - name: archiver
              image: laiello/liveatc-archiver:latest
              imagePullPolicy: Always
              env:
              - name: FTP_HOST
                value: ""
              - name: FTP_USER
                value: ""
              - name: FTP_PASS
                value: ""
              - name: FTP_BASE
                value: ""
          restartPolicy: OnFailure

```
