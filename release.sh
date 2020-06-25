# $1 directory
function pack() {
  cp -r ../img/ "$1/"
  cp -r ../config.yaml "$1/"
  cd "$1/" || exit
  zip -r "../$1.zip" *
  cd ../
}

GOOS=linux go build -ldflags "-w -s" -o release/linux/teemo
GOOS=darwin go build -ldflags "-w -s" -o release/darwin/teemo
GOOS=windows go build -ldflags "-w -s" -o release/windows/teemo.exe

GOARCH=arm64 GOOS=linux go build -ldflags "-w -s" -o release/linux_arm64/teemo

cd release || exit

pack "linux"
pack "darwin"
pack "windows"
pack "linux_arm64"
