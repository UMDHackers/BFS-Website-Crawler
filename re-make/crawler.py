import sys
from bs4 import BeautifulSoup
import urllib2

#Get the website
website = sys.argv[1]
#print website
web_queue = []
web_queue.append(website)
temp_queue = []
levels = 0
#print web_queue

def get_links(curr_website):
	html = urllib2.urlopen(curr_website).read()
	soup = BeautifulSoup(html, 'html.parser')
	arr = []
	for i in soup.findAll("a"):
		arr.append(i)
	return arr

while(levels != 3):
	while(web_queue):
		curr_website = web_queue.pop()
		array_links = get_links(curr_website)
		temp_queue.extend(array_links)
	print web_queue
	web_queue.extend(temp_queue)
	levels = levels + 1
	temp_queue = []
