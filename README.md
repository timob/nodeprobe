nodeprobe
==========

Cassandra DB monitoring library.

####Example
See nodestat example utility in source code.

####Install
Eg. On on Ubuntu with Cassandra installed in /files/data/apache-cassandra-2.0.8
````
go get github.com/timob/nodeprobe
go install github.com/timob/nodeprobe/nodestat

export LD_LIBRARY_PATH=/usr/lib/jvm/default-java/jre/lib/amd64/server
export CLASSPATH=/files/data/apache-cassandra-2.0.8/lib/apache-cassandra-2.0.8.jar:/files/data/apache-cassandra-2.0.8/lib/guava-15.0.jar
nodestat -host dbhost
````

