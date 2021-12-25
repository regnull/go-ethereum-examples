# Export Account

This example shows how to export an account as JSON. You can save the
exported JSON to a file and later import it to another keystore.

```
go run main.go --keystore-dir=$HOME/keystore \
    --address=0xADF6b8ff67E5537367156e27c532Be90144ba7D2 \
    --password=supersecretpassword \
    --export-password=anothersecretpassword
```