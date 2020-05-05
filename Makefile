all:
	cat test.ysn | python main.py -o test.s

clean:
	- rm -f *.pyc
