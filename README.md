# twitpicker

twitpicのユーザー名から画像をダウンロードしてくる

途中で飽きたので放置

いくつか使いまわせそうな処理をメモ

## HTTPでヘッダーを付けてGET

```go
func getHTTP(url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0")
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
```

## JSONデコード

```go
type Image struct {
	ShortID string `json:"short_id"`
	Type    string `json:"type"`
}

type Photos struct {
	Images []Image
}

func DecodePhotos(phJson []byte) Photos {
	var p Photos
	json.Unmarshal(phJson, &p)
	return p
}
```

## HTTPからダウンロードしたものをファイルに保存

`io.Copy` でReaderからWriterにそのまま出力

```go
	var resp *http.Response
	var err error

	numRetry := 3
	for i := 0; i < numRetry; i++ {
		resp, err = http.Get(img.ToURL())
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			log.Println("bad http status ", resp.StatusCode)
			resp.Body.Close()
			if i == numRetry-1 {
				return errors.New("maxium number of retry reached")
			}
			continue
		}
		defer resp.Body.Close()
		break
	}

	...

	f, err := os.Create(name)

	written, err := io.Copy(f, resp.Body)
```

## 並列数Nで非同期に処理を実行

```go
func downloadImages(photos twitpic.Photos) {
	ch := make(chan twitpic.Image, len(photos.Images))

	var wg sync.WaitGroup
	defer wg.Wait()

	for n := 0; n < numDownloads; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for img := range ch {
				err := img.Download()
				if err != nil {
					fmt.Printf("error: %s\n", err.Error())
					os.Exit(1)
				}
			}
		}()
	}

	for _, img := range photos.Images {
		ch <- img
	}
	close(ch)
}
```
