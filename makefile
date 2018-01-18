getdate:
	echo "cat capture | awk '{print $1}' | cut -c 1-8 | uniq -c"
	cat capture | awk '{print $1}' | cut -c 1-8 | uniq -c
getdata:
	echo "sudo tcpdump -i lo port 1025 > capture"
	sudo tcpdump -i lo port 1025 > capture
