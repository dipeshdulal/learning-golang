### Google Cloud Bucket Storage

From google docs;
> Each project can contain multiple buckets, which are containers to store your objects. For example, you might create a photos bucket for all the image files your app generates and a separate videos bucket.

So this implemenation contains the file upload example from `gin` to `bucket` storage. User uploads file in out `gin` api and then the api uploads file to bucket storage. This repo only contains code example for uploading files in bucket. Deleting, Updating files in bucket are not implemented.

After looking at the example implementation over [here](https://github.com/GoogleCloudPlatform/golang-samples/blob/master/appengine_flexible/storage/storage.go) there are some points that are to be noted;

- At first context has to be created with `Request` object by doing.
    ```go
    ctx := appengine.NewContext(c.Request)
    ```
- Initial setup of the storage client is done using;
    ```go
    storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("serviceAccountKey.json"))   
    ```
    Notice the `option.WithCredentialsFiles` points to the serviceAccountKey of bucket storage.

- The line;
    ```go
    sw := storageClient.Bucket(bucket).Object(uploadedFile.Filename).NewWriter(ctx)
    ```
    Creates a new storage writer for which out file object `f` can be copied to using `io.Copy` function. By doing such we upload the file to cloud storage. 

- `sw.Close()` closes the storage writer and to get information about the uploaded object we can use `sw.Attrs()` to get all the attribute of uploaded files.
