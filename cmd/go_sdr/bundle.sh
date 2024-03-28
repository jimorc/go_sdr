cd $GOPATH/go_sdr/internal/gosdrgui
fyne bundle -o bundled.go $GOPATH/go_sdr/cmd/go_sdr/images/start.svg
fyne bundle -o bundled.go -append $GOPATH/go_sdr/cmd/go_sdr/images/stop.svg
if [[ "$OSTYPE" == "darwin"* ]]; then
sed -i '' 's/main/gosdrgui/' bundled.go
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
sed 's/main/gosdrgui/' bundled.go
fi
cd $GOPATH/go_sdr/cmd/go_sdr