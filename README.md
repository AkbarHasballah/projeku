### Update Library GISWisataBandung

```sh
go get -u all
go mod tidy
git tag                                 #cek riwayat versi tag
git tag v1.0.0                          #set versi tag
git push origin v1.0.0     #push tag version ke repo
go list -m github.com/AkbarHasballah/projeku@v1.0.0   #publish ke PKG go Dev

go get github.com/AkbarHasballah/projeku #Jika ingin Menggunakan Package atau library
```

