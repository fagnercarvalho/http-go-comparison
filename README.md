### Go HTTP versions comparison

Repository that exposes services in HTTP/1.1, HTTP/2 and HTTP/3 (new QUIC protocol) made in Go.

![](diagram.png)

#### Usage

```docker-compose up -d
docker-compose up -d
```

> :warning: **Important**: Since all services use [self signed certificates](https://en.wikipedia.org/wiki/Self-signed_certificate) you need to bypass the certificate check in Chrome and Firefox.
> 
> For Firefox, access:
> 
> - https://localhost:6123.
> - https://localhost:6124.
> - https://localhost:6125
> 
> And for each site click on 'Advanced' and then 'Accept the Risk and Continue'. 
>  
> 
> For Chrome, access [chrome://flags/#allow-insecure-localhost](), enable the flag and restart your browser.
> 
>
> You should be able to access http://localhost:8080 and test the services right after.


After all containers start go to the client by browsing `http://localhost:8080`.

To test, click in the button for the HTTP version you want to use. You should see the response from the service printed in the screen.

#### Notes

QUIC is still an experimental protocol and to allow requests to services that are using it you need to enable QUIC in your browser. QUIC is supported in most major browsers including Firefox and Chrome. 

For Firefox go to `about:config` and toggle `network.http.http3.enabled` and increase `network.http.http3.recvBufferSize` to `2500000`.

For Chrome go to `chrome://flags` and toggle `Experimental QUIC protocol`.

