#!/bin/bash

filter="java.io.PrintStream javax.management.openmbean.TabularData java.lang.management.MemoryUsage java.util.UUID"
javap="javap -s -classpath /files/data/apache-cassandra-2.0.8/lib/apache-cassandra-2.0.8.jar"
#gen="go run ../jag/cmd/jagen/jagen.go"
gen="jagen"
srcDir=/home/tim/repos/cassandra/src/java
srcName=org.apache.cassandra

#name=tools.NodeProbe
#namePath=$(echo $name | sed 's#\.#/#g')
#outName=$(echo $name | sed 's#\.##g')

#$javap ${srcName}.$name | $gen -src ${srcDir}/${namePath}.java > ${outName}.go

function generate() {
    name=$1
    namePath=$(echo $name | sed 's#\.#/#g')
    outName=$(echo $name | sed 's#\.#_#g')
    $javap $name | $gen -pkg nodeprobe -filter "$filter" -src ${srcDir}/${namePath}.java -trim $srcName > ${outName}.go
}

function deps() {
    name=$1
    $javap $name | $gen -filter "$filter" -d
}

if echo $1 | grep -v $srcName > /dev/null; then
	exit
fi

echo "generating $1"
generate $1

for name in $(deps $1); do
	echo $name >> deps.txt
	./gen.sh $name
done
