# wokehide
A semi-satirical cli tool to loosely "hide" images using base64 encoding/decoding

## Building
```sh
sudo rm "$(which wokehide)"; go build -o wokehide ./src; sudo mv ./wokehide /usr/bin