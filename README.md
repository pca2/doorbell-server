# doorbell-server
Build for Raspberry Pi B+: 
`GOOS=linux GOARCH=arm GOARM=6 go build -o rasp-doorbell-server`


Routes:
/status
/play #play default file path
/play?file=/path/to/sound.mp3

