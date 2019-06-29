#coding=utf-8
import os
import sys
import time

try:
    import scapy.all as scapy
except ImportError:
    import scapy

#

def rmdirs(top):
    for root, dirs, files in os.walk(top, topdown=False):
        # 先删除文件夹
        for name in files:
            os.remove(os.path.join(root, name))
        # 再删除空目录
        for name in dirs:
            os.rmdir(os.path.join(root, name))

filedir = './tcpstreamdata'
rmdirs(filedir)
if os.path.isdir(filedir) is False:
    os.makedirs(filedir)

g_filemap = dict()
def newfile(srcipport):
    if srcipport in g_filemap:
        filelast = g_filemap.get(srcipport)
        filelast.close()
    ffile = open('%s/%s_%s.txt' % (filedir, srcipport, time.time()), 'wb+')
    g_filemap[srcipport] = ffile
    return ffile

def writefile(srcipport, data):
    # if not find file
    filewrite = g_filemap.get(srcipport)
    if filewrite is None:
        filewrite = newfile(srcipport)
    filewrite.write(data)
    return 


packets = scapy.rdpcap(sys.argv[1])
#packets = scapy.rdpcap('./13168.pcap')
clientportlist = []
serverportlist = [80, 8080] ### this must must must set

### must set all the time
clientiplist = []
serveriplist = []

def checkin(clientlist, serverlist, src, dst):
        bret = False
        if len(clientlist)>0 and len(serverlist)>0:
            if (src in clientlist and dst in serverlist) or (dst in clientlist and src in serverlist):
                bret = True
        elif len(clientlist)>0:
            if src in clientlist or dst in clientlist:
                bret = True
        elif len(serverlist)>0:
            if src in serverlist or dst in serverlist:
                bret = True
        else:
            bret = True
        return bret

# return true means is ok
def filertokpacket(pack):
    if p["TCP"] == None:
        return False    
    bip = checkin(clientiplist, serveriplist, p["IP"].src, p["IP"].dst)

    btcp = checkin(clientportlist, serverportlist, p["TCP"].sport, p["TCP"].dport)
    if bip and btcp:
        return True
    return False

for p in packets:    
    if filertokpacket(p) is False:
        continue        

    bserver = p["TCP"].sport in serverportlist
    bclient = p["TCP"].dport in serverportlist
    
    srcipport = ""
    if bserver:
        srcipport = "%s_%s" % (p["IP"].dst, p["TCP"].dport)
    if bclient:
        srcipport = "%s_%s" % (p["IP"].src, p["TCP"].sport)        

    if p["TCP"].flags.value == 2:                
        print(p["TCP"].sport, p["TCP"].dport)
        print(p["IP"].src, p["IP"].dst)        
        print("new tcp connect*************************************************************")
        newfile(srcipport)
        continue

    #client send data with sync, and  server send data with sync and ack
    #if 'S' in p["TCP"].flags:
    #    continue
    
    # great
    if 'P' not in p["TCP"].flags:
        continue

    if  isinstance(p['TCP'].payload, scapy.packet.NoPayload) :
        continue
    ##### this is for ethernet padding     5 means 20 0101
    if p['TCP'].urgptr != 0  or p['IP'].ihl != 5:
        print(p["TCP"].sport, p["TCP"].dport)
        print(p["IP"].src, p["IP"].dst)        
        print("error : ----------------------something error  tcp urgptr now not support, or ip header is not 20")
        break    

    rawdatalen =  p['IP'].len - 20 - 20
    print("IP.len", p['IP'].len, rawdatalen)
    if rawdatalen == 0:
        continue    

    load = p['TCP'].payload.load[0:rawdatalen]

    print(type(p['TCP'].payload.load))
        
    print(p["TCP"].sport, p["TCP"].dport)
    print(p["IP"].src, p["IP"].dst)
    print("filename", srcipport)
    #print(p['TCP'].payload.load)
    writefile(srcipport, p['TCP'].payload.load)
for val in g_filemap.values():
    val.close()

def main(argv):
    print(argv[1])
 
if __name__ == "__main__":
    main(sys.argv)