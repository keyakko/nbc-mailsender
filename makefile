getdate:
	cat capture | awk '{print $1}' | cut -c 1-8 | uniq -c
getdata:
	sudo tcpdump -i lo port 1025 > capture
run:
	while true; do go run main.go; done
