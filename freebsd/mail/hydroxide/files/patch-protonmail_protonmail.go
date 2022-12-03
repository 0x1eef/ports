--- protonmail/protonmail.go.orig	2022-12-03 00:18:27 UTC
+++ protonmail/protonmail.go
@@ -119,7 +119,8 @@ func (c *Client) newJSONRequest(method, path string, b
 }
 
 func (c *Client) do(req *http.Request) (*http.Response, error) {
-	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0")
+	req.Header.Set("User-Agent", "Ubuntu_20.04")
+	req.Header.Set("x-pm-appversion", "Other")
 
 	httpClient := c.HTTPClient
 	if httpClient == nil {
