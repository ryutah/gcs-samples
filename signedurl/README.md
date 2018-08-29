# Signed URL Sample
Create Signed URL at server side and upload file from client side.

## Prepare
### Create Service Account and Download Secret JSON.

### Set Environment variables.
```console
$ export BUCKET=sandbox-hara-signedurl-sample
$ export GOOGLE_ACCESS_ID=[SERVICE_ACCOUNT_NAME]@[PROJECT_ID].iam.gserviceaccount.com
$ export PRIVATE_KEY=[SERVICE_ACCOUNT_PRIVATE_KEY]
```

### Configure GCS CORS.
```console
$ gsutil cors set ./cors.json gs://$BUCKET/ 
```

## Start
```console
$ go run main.go
```

---

## See
* [Signed URL](https://cloud.google.com/storage/docs/access-control/signed-urls)
* [Set Bucket CORS](https://cloud.google.com/storage/docs/xml-api/put-bucket-cors)
