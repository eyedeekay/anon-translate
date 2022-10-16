# anon-translate

Translate text anonymously using a hosted service.
Reads text from Standard In, outputs text to Standard Out.

## installation

```sh
go install github.com/eyedeekay/anon-translate@latest
# It is helpful to add this to "$HOME/.bashrc"
PATH="$HOME/go/bin:$PATH"
```

## example

```sh
echo "test" | anon-translate
```
