# drone

<https://github.com/drone/go-license>

```bash
git remote add upstream git@github.com:drone/go-license.git

git fetch upstream

git merge v1.0.2
```

## License

```bash
# keygen
go run cmd/license-keygen/main.go --out .beagle/id_beagle

# create
go run cmd/license-create/main.go \
--in .beagle/id_beagle \
--iss "Beagle Cloud Team" \
--cus "比格大数据" \
--sub "研发中心" \
--typ "基础版本" \
--exp "26352h0m0s"

# verify
go run cmd/license-verify/main.go \
--pub .beagle/id_beagle.pub \
--file license.txt
```
