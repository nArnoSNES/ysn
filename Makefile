all:
	go build
	cat test.ysn | ./ysn

clean:
	- rm -f ysn.exe ysn
