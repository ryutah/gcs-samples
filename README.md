# GCS 変更通知サンプル

* GCS Notify登録

```console
$ gsutil notification create -t ${TOPIC_NAME} -f json gs://${BUCKET_NAME}
```

* PubSubへPushSubscriberを登録

```console
$ gcloud beta pubsub subscriptions create ${SUBSCRIPTION_NAME} --push-endpoint ${END_POINT_URL} --topic ${TOPIC_NAME}
```


## メッセージ
### 新規作成
```json
{
  "message": {
    "data": "ewogICJraW5kIjogInN0b3JhZ2Ujb2JqZWN0IiwKICAiaWQiOiAic2FuZGJveC1oYXJhLmFwcHNwb3QuY29tL3NhbXBsZS9vbmVfY2xhc3Nfc3ZtLnBuZy8xNTA3MjYwNTA3NzkyNzk3IiwKICAic2VsZkxpbmsiOiAiaHR0cHM6Ly93d3cuZ29vZ2xlYXBpcy5jb20vc3RvcmFnZS92MS9iL3NhbmRib3gtaGFyYS5hcHBzcG90LmNvbS9vL3NhbXBsZSUyRm9uZV9jbGFzc19zdm0ucG5nIiwKICAibmFtZSI6ICJzYW1wbGUvb25lX2NsYXNzX3N2bS5wbmciLAogICJidWNrZXQiOiAic2FuZGJveC1oYXJhLmFwcHNwb3QuY29tIiwKICAiZ2VuZXJhdGlvbiI6ICIxNTA3MjYwNTA3NzkyNzk3IiwKICAibWV0YWdlbmVyYXRpb24iOiAiMSIsCiAgImNvbnRlbnRUeXBlIjogImltYWdlL3BuZyIsCiAgInRpbWVDcmVhdGVkIjogIjIwMTctMTAtMDZUMDM6Mjg6MjcuNzc3WiIsCiAgInVwZGF0ZWQiOiAiMjAxNy0xMC0wNlQwMzoyODoyNy43NzdaIiwKICAic3RvcmFnZUNsYXNzIjogIlNUQU5EQVJEIiwKICAidGltZVN0b3JhZ2VDbGFzc1VwZGF0ZWQiOiAiMjAxNy0xMC0wNlQwMzoyODoyNy43NzdaIiwKICAic2l6ZSI6ICIxMjI0MiIsCiAgIm1kNUhhc2giOiAiMHVXSnp0aFJxeXlyb0tSSEIyQ3EyQT09IiwKICAibWVkaWFMaW5rIjogImh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL2Rvd25sb2FkL3N0b3JhZ2UvdjEvYi9zYW5kYm94LWhhcmEuYXBwc3BvdC5jb20vby9zYW1wbGUlMkZvbmVfY2xhc3Nfc3ZtLnBuZz9nZW5lcmF0aW9uPTE1MDcyNjA1MDc3OTI3OTcmYWx0PW1lZGlhIiwKICAiY3JjMzJjIjogImxHeUd0UT09IiwKICAiZXRhZyI6ICJDSjNMbXZtRzI5WUNFQUU9Igp9Cg==",
    "attributes": {
      "objectGeneration": "1507260507792797",
      "resource": "projects/_/buckets/sandbox-hara.appspot.com/objects/sample/one_class_svm.png#1507260507792797",
      "bucketId": "sandbox-hara.appspot.com",
      "eventType": "OBJECT_FINALIZE",
      "notificationConfig": "projects/_/buckets/sandbox-hara.appspot.com/notificationConfigs/3",
      "payloadFormat": "JSON_API_V1",
      "objectId": "sample/one_class_svm.png"
    },
    "message_id": "157810539096020",
    "messageId": "157810539096020",
    "publish_time": "2017-10-06T03:28:28.241Z",
    "publishTime": "2017-10-06T03:28:28.241Z"
  },
  "subscription": "projects/sandbox-hara/subscriptions/samplesubsc"
}
```

### 削除
```json
{
  "message": {
    "data": "ewogICJraW5kIjogInN0b3JhZ2Ujb2JqZWN0IiwKICAiaWQiOiAic2FuZGJveC1oYXJhLmFwcHNwb3QuY29tL29uZV9jbGFzc19zdm0ucG5nLzE1MDcyNjAxOTMxNzY4NTIiLAogICJzZWxmTGluayI6ICJodHRwczovL3d3dy5nb29nbGVhcGlzLmNvbS9zdG9yYWdlL3YxL2Ivc2FuZGJveC1oYXJhLmFwcHNwb3QuY29tL28vb25lX2NsYXNzX3N2bS5wbmciLAogICJuYW1lIjogIm9uZV9jbGFzc19zdm0ucG5nIiwKICAiYnVja2V0IjogInNhbmRib3gtaGFyYS5hcHBzcG90LmNvbSIsCiAgImdlbmVyYXRpb24iOiAiMTUwNzI2MDE5MzE3Njg1MiIsCiAgIm1ldGFnZW5lcmF0aW9uIjogIjEiLAogICJjb250ZW50VHlwZSI6ICJpbWFnZS9wbmciLAogICJ0aW1lQ3JlYXRlZCI6ICIyMDE3LTEwLTA2VDAzOjIzOjEzLjE2MFoiLAogICJ1cGRhdGVkIjogIjIwMTctMTAtMDZUMDM6MjM6MTMuMTYwWiIsCiAgInN0b3JhZ2VDbGFzcyI6ICJTVEFOREFSRCIsCiAgInRpbWVTdG9yYWdlQ2xhc3NVcGRhdGVkIjogIjIwMTctMTAtMDZUMDM6MjM6MTMuMTYwWiIsCiAgInNpemUiOiAiMTIyNDIiLAogICJtZDVIYXNoIjogIjB1V0p6dGhScXl5cm9LUkhCMkNxMkE9PSIsCiAgIm1lZGlhTGluayI6ICJodHRwczovL3d3dy5nb29nbGVhcGlzLmNvbS9kb3dubG9hZC9zdG9yYWdlL3YxL2Ivc2FuZGJveC1oYXJhLmFwcHNwb3QuY29tL28vb25lX2NsYXNzX3N2bS5wbmc/Z2VuZXJhdGlvbj0xNTA3MjYwMTkzMTc2ODUyJmFsdD1tZWRpYSIsCiAgImNyYzMyYyI6ICJsR3lHdFE9PSIsCiAgImV0YWciOiAiQ0pUNmwrT0YyOVlDRUFFPSIKfQo=",
    "attributes": {
      "objectGeneration": "1507260193176852",
      "resource": "projects/_/buckets/sandbox-hara.appspot.com/objects/one_class_svm.png#1507260193176852",
      "bucketId": "sandbox-hara.appspot.com",
      "eventType": "OBJECT_DELETE",
      "notificationConfig": "projects/_/buckets/sandbox-hara.appspot.com/notificationConfigs/3",
      "payloadFormat": "JSON_API_V1",
      "objectId": "one_class_svm.png"
    },
    "message_id": "157818482251757",
    "messageId": "157818482251757",
    "publish_time": "2017-10-06T03:25:38.669Z",
    "publishTime": "2017-10-06T03:25:38.669Z"
  },
  "subscription": "projects/sandbox-hara/subscriptions/samplesubsc"
}
```

### ファイル更新
```json
 {
   "message": {
     "data": "ewogICJraW5kIjogInN0b3JhZ2Ujb2JqZWN0IiwKICAiaWQiOiAic2FuZGJveC1oYXJhLmFwcHNwb3QuY29tL3NhbXBsZS9vbmVfY2xhc3Nfc3ZtLnBuZy8xNTA3MjYwNTA3NzkyNzk3IiwKICAic2VsZkxpbmsiOiAiaHR0cHM6Ly93d3cuZ29vZ2xlYXBpcy5jb20vc3RvcmFnZS92MS9iL3NhbmRib3gtaGFyYS5hcHBzcG90LmNvbS9vL3NhbXBsZSUyRm9uZV9jbGFzc19zdm0ucG5nIiwKICAibmFtZSI6ICJzYW1wbGUvb25lX2NsYXNzX3N2bS5wbmciLAogICJidWNrZXQiOiAic2FuZGJveC1oYXJhLmFwcHNwb3QuY29tIiwKICAiZ2VuZXJhdGlvbiI6ICIxNTA3MjYwNTA3NzkyNzk3IiwKICAibWV0YWdlbmVyYXRpb24iOiAiMSIsCiAgImNvbnRlbnRUeXBlIjogImltYWdlL3BuZyIsCiAgInRpbWVDcmVhdGVkIjogIjIwMTctMTAtMDZUMDM6Mjg6MjcuNzc3WiIsCiAgInVwZGF0ZWQiOiAiMjAxNy0xMC0wNlQwMzoyODoyNy43NzdaIiwKICAic3RvcmFnZUNsYXNzIjogIlNUQU5EQVJEIiwKICAidGltZVN0b3JhZ2VDbGFzc1VwZGF0ZWQiOiAiMjAxNy0xMC0wNlQwMzoyODoyNy43NzdaIiwKICAic2l6ZSI6ICIxMjI0MiIsCiAgIm1kNUhhc2giOiAiMHVXSnp0aFJxeXlyb0tSSEIyQ3EyQT09IiwKICAibWVkaWFMaW5rIjogImh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL2Rvd25sb2FkL3N0b3JhZ2UvdjEvYi9zYW5kYm94LWhhcmEuYXBwc3BvdC5jb20vby9zYW1wbGUlMkZvbmVfY2xhc3Nfc3ZtLnBuZz9nZW5lcmF0aW9uPTE1MDcyNjA1MDc3OTI3OTcmYWx0PW1lZGlhIiwKICAiY3JjMzJjIjogImxHeUd0UT09IiwKICAiZXRhZyI6ICJDSjNMbXZtRzI5WUNFQUU9Igp9Cg==",
     "attributes": {
       "objectGeneration": "1507260507792797",
       "overwrittenByGeneration": "1507260566514012",
       "resource": "projects/_/buckets/sandbox-hara.appspot.com/objects/sample/one_class_svm.png#1507260507792797",
       "bucketId": "sandbox-hara.appspot.com",
       "eventType": "OBJECT_DELETE",
       "notificationConfig": "projects/_/buckets/sandbox-hara.appspot.com/notificationConfigs/3",
       "payloadFormat": "JSON_API_V1",
       "objectId": "sample/one_class_svm.png"
     },
     "message_id": "157825666482053",
     "messageId": "157825666482053",
     "publish_time": "2017-10-06T03:29:27.472Z",
     "publishTime": "2017-10-06T03:29:27.472Z"
   },
   "subscription": "projects/sandbox-hara/subscriptions/samplesubsc"
 }
```

```json
{
  "message": {
    "data": "ewogICJraW5kIjogInN0b3JhZ2Ujb2JqZWN0IiwKICAiaWQiOiAic2FuZGJveC1oYXJhLmFwcHNwb3QuY29tL3NhbXBsZS9vbmVfY2xhc3Nfc3ZtLnBuZy8xNTA3MjYwNTY2NTE0MDEyIiwKICAic2VsZkxpbmsiOiAiaHR0cHM6Ly93d3cuZ29vZ2xlYXBpcy5jb20vc3RvcmFnZS92MS9iL3NhbmRib3gtaGFyYS5hcHBzcG90LmNvbS9vL3NhbXBsZSUyRm9uZV9jbGFzc19zdm0ucG5nIiwKICAibmFtZSI6ICJzYW1wbGUvb25lX2NsYXNzX3N2bS5wbmciLAogICJidWNrZXQiOiAic2FuZGJveC1oYXJhLmFwcHNwb3QuY29tIiwKICAiZ2VuZXJhdGlvbiI6ICIxNTA3MjYwNTY2NTE0MDEyIiwKICAibWV0YWdlbmVyYXRpb24iOiAiMSIsCiAgImNvbnRlbnRUeXBlIjogImltYWdlL3BuZyIsCiAgInRpbWVDcmVhdGVkIjogIjIwMTctMTAtMDZUMDM6Mjk6MjYuNTAyWiIsCiAgInVwZGF0ZWQiOiAiMjAxNy0xMC0wNlQwMzoyOToyNi41MDJaIiwKICAic3RvcmFnZUNsYXNzIjogIlNUQU5EQVJEIiwKICAidGltZVN0b3JhZ2VDbGFzc1VwZGF0ZWQiOiAiMjAxNy0xMC0wNlQwMzoyOToyNi41MDJaIiwKICAic2l6ZSI6ICIxMjI0MiIsCiAgIm1kNUhhc2giOiAiMHVXSnp0aFJxeXlyb0tSSEIyQ3EyQT09IiwKICAibWVkaWFMaW5rIjogImh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL2Rvd25sb2FkL3N0b3JhZ2UvdjEvYi9zYW5kYm94LWhhcmEuYXBwc3BvdC5jb20vby9zYW1wbGUlMkZvbmVfY2xhc3Nfc3ZtLnBuZz9nZW5lcmF0aW9uPTE1MDcyNjA1NjY1MTQwMTImYWx0PW1lZGlhIiwKICAiY3JjMzJjIjogImxHeUd0UT09IiwKICAiZXRhZyI6ICJDTnpTbXBXSDI5WUNFQUU9Igp9Cg==",
    "attributes": {
      "overwroteGeneration": "1507260507792797",
      "objectGeneration": "1507260566514012",
      "resource": "projects/_/buckets/sandbox-hara.appspot.com/objects/sample/one_class_svm.png#1507260566514012",
      "bucketId": "sandbox-hara.appspot.com",
      "eventType": "OBJECT_FINALIZE",
      "notificationConfig": "projects/_/buckets/sandbox-hara.appspot.com/notificationConfigs/3",
      "payloadFormat": "JSON_API_V1",
      "objectId": "sample/one_class_svm.png"
    },
    "message_id": "157795317705531",
    "messageId": "157795317705531",
    "publish_time": "2017-10-06T03:29:27.391Z",
    "publishTime": "2017-10-06T03:29:27.391Z"
  },
  "subscription": "projects/sandbox-hara/subscriptions/samplesubsc"
}
```

### 名前変更
```json
 {
   "message": {
     "data": "ewogICJraW5kIjogInN0b3JhZ2Ujb2JqZWN0IiwKICAiaWQiOiAic2FuZGJveC1oYXJhLmFwcHNwb3QuY29tL3NhbXBsZS9vbmVfY2xhc3Nfc3ZtX2NoYW5nZS5wbmcvMTUwNzI2MDcxMTI0MDQxNiIsCiAgInNlbGZMaW5rIjogImh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL3N0b3JhZ2UvdjEvYi9zYW5kYm94LWhhcmEuYXBwc3BvdC5jb20vby9zYW1wbGUlMkZvbmVfY2xhc3Nfc3ZtX2NoYW5nZS5wbmciLAogICJuYW1lIjogInNhbXBsZS9vbmVfY2xhc3Nfc3ZtX2NoYW5nZS5wbmciLAogICJidWNrZXQiOiAic2FuZGJveC1oYXJhLmFwcHNwb3QuY29tIiwKICAiZ2VuZXJhdGlvbiI6ICIxNTA3MjYwNzExMjQwNDE2IiwKICAibWV0YWdlbmVyYXRpb24iOiAiMSIsCiAgImNvbnRlbnRUeXBlIjogImltYWdlL3BuZyIsCiAgInRpbWVDcmVhdGVkIjogIjIwMTctMTAtMDZUMDM6MzE6NTEuMTQzWiIsCiAgInVwZGF0ZWQiOiAiMjAxNy0xMC0wNlQwMzozMTo1MS4xNDNaIiwKICAic3RvcmFnZUNsYXNzIjogIlNUQU5EQVJEIiwKICAidGltZVN0b3JhZ2VDbGFzc1VwZGF0ZWQiOiAiMjAxNy0xMC0wNlQwMzozMTo1MS4xNDNaIiwKICAic2l6ZSI6ICIxMjI0MiIsCiAgIm1kNUhhc2giOiAiMHVXSnp0aFJxeXlyb0tSSEIyQ3EyQT09IiwKICAibWVkaWFMaW5rIjogImh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL2Rvd25sb2FkL3N0b3JhZ2UvdjEvYi9zYW5kYm94LWhhcmEuYXBwc3BvdC5jb20vby9zYW1wbGUlMkZvbmVfY2xhc3Nfc3ZtX2NoYW5nZS5wbmc/Z2VuZXJhdGlvbj0xNTA3MjYwNzExMjQwNDE2JmFsdD1tZWRpYSIsCiAgImNyYzMyYyI6ICJsR3lHdFE9PSIsCiAgImV0YWciOiAiQ09DRm5OcUgyOVlDRUFFPSIKfQo=",
     "attributes": {
       "objectGeneration": "1507260711240416",
       "resource": "projects/_/buckets/sandbox-hara.appspot.com/objects/sample/one_class_svm_change.png#1507260711240416",
       "bucketId": "sandbox-hara.appspot.com",
       "eventType": "OBJECT_FINALIZE",
       "notificationConfig": "projects/_/buckets/sandbox-hara.appspot.com/notificationConfigs/3",
       "payloadFormat": "JSON_API_V1",
       "objectId": "sample/one_class_svm_change.png"
     },
     "message_id": "157826736163590",
     "messageId": "157826736163590",
     "publish_time": "2017-10-06T03:31:52.288Z",
     "publishTime": "2017-10-06T03:31:52.288Z"
   },
   "subscription": "projects/sandbox-hara/subscriptions/samplesubsc"
 }
```

```json
{
  "message": {
    "data": "ewogICJraW5kIjogInN0b3JhZ2Ujb2JqZWN0IiwKICAiaWQiOiAic2FuZGJveC1oYXJhLmFwcHNwb3QuY29tL3NhbXBsZS9vbmVfY2xhc3Nfc3ZtLnBuZy8xNTA3MjYwNTY2NTE0MDEyIiwKICAic2VsZkxpbmsiOiAiaHR0cHM6Ly93d3cuZ29vZ2xlYXBpcy5jb20vc3RvcmFnZS92MS9iL3NhbmRib3gtaGFyYS5hcHBzcG90LmNvbS9vL3NhbXBsZSUyRm9uZV9jbGFzc19zdm0ucG5nIiwKICAibmFtZSI6ICJzYW1wbGUvb25lX2NsYXNzX3N2bS5wbmciLAogICJidWNrZXQiOiAic2FuZGJveC1oYXJhLmFwcHNwb3QuY29tIiwKICAiZ2VuZXJhdGlvbiI6ICIxNTA3MjYwNTY2NTE0MDEyIiwKICAibWV0YWdlbmVyYXRpb24iOiAiMSIsCiAgImNvbnRlbnRUeXBlIjogImltYWdlL3BuZyIsCiAgInRpbWVDcmVhdGVkIjogIjIwMTctMTAtMDZUMDM6Mjk6MjYuNTAyWiIsCiAgInVwZGF0ZWQiOiAiMjAxNy0xMC0wNlQwMzoyOToyNi41MDJaIiwKICAic3RvcmFnZUNsYXNzIjogIlNUQU5EQVJEIiwKICAidGltZVN0b3JhZ2VDbGFzc1VwZGF0ZWQiOiAiMjAxNy0xMC0wNlQwMzoyOToyNi41MDJaIiwKICAic2l6ZSI6ICIxMjI0MiIsCiAgIm1kNUhhc2giOiAiMHVXSnp0aFJxeXlyb0tSSEIyQ3EyQT09IiwKICAibWVkaWFMaW5rIjogImh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL2Rvd25sb2FkL3N0b3JhZ2UvdjEvYi9zYW5kYm94LWhhcmEuYXBwc3BvdC5jb20vby9zYW1wbGUlMkZvbmVfY2xhc3Nfc3ZtLnBuZz9nZW5lcmF0aW9uPTE1MDcyNjA1NjY1MTQwMTImYWx0PW1lZGlhIiwKICAiY3JjMzJjIjogImxHeUd0UT09IiwKICAiZXRhZyI6ICJDTnpTbXBXSDI5WUNFQUU9Igp9Cg==",
    "attributes": {
      "objectGeneration": "1507260566514012",
      "resource": "projects/_/buckets/sandbox-hara.appspot.com/objects/sample/one_class_svm.png#1507260566514012",
      "bucketId": "sandbox-hara.appspot.com",
      "eventType": "OBJECT_DELETE",
      "notificationConfig": "projects/_/buckets/sandbox-hara.appspot.com/notificationConfigs/3",
      "payloadFormat": "JSON_API_V1",
      "objectId": "sample/one_class_svm.png"
    },
    "message_id": "157811238131329",
    "messageId": "157811238131329",
    "publish_time": "2017-10-06T03:31:52.490Z",
    "publishTime": "2017-10-06T03:31:52.490Z"
  },
  "subscription": "projects/sandbox-hara/subscriptions/samplesubsc"
}
```
