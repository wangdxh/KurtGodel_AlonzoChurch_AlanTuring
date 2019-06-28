#### httptotestcase
get http data from  wireshark or tcpdump pcap files,   
and then generate test case samples from http data.      
     
now support the test case is go4api:  https://github.com/zpsean/go4api    

#### dependencies
 * pip install scapy     
   https://scapy.readthedocs.io/en/latest/installation.html    

#### main.py
  use scapy to get tcp streams from pacp, and then split tcp stream to seperate files   
  in tcpstreamdata   
#### autohttptestcase
  is a golang program that read the tcp data files and then parse them to http, then filter some urls,    
  then generate go4api test case in testcase directory   
  
 







