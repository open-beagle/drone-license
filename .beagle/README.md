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

# create 10year license
go run cmd/license-create/main.go \
--in .beagle/id_beagle \
--iss "Beagle Cloud Team" \
--cus "比格大数据" \
--sub "研发中心" \
--typ "基础版本" \
--exp "87840h0m0s" \
--dat '{"kind":"standard","repos":0,"users":0,"nodes":0,"builds":0}'

# verify
go run cmd/license-verify/main.go \
--pub .beagle/id_beagle.pub \
--file license.txt
```
