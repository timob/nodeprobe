package main

import (
	"fmt"
	"log"
	"../"
	"github.com/timob/javabind"
	"text/tabwriter"
	"os"
	"strings"
)

type proxy struct {
	keyspace string
	bean *nodeprobe.DbColumnFamilyStoreMBean
}

type stat struct {
	title string
	format string
	f func(p *proxy) interface{}
}

func main() {
	err := javabind.SetupJVM(
		"/files/data/apache-cassandra-2.0.8/lib/apache-cassandra-2.0.8.jar",
		"/files/data/apache-cassandra-2.0.8/lib/guava-15.0.jar",
	)	
	if err != nil {
		log.Fatal(err)
	}
	
	nodeProbe, err := nodeprobe.NewToolsNodeProbe3("localhost")
	if err != nil {
		log.Fatal(err)
	}
		
	writer := tabwriter.NewWriter(os.Stdout, 2, 4, 2, ' ', 0)
	
	stats := make([]stat, 0)
	stats = append(stats, stat{"Keyspace", "%s", func(p *proxy) interface{} { return p.keyspace }})
	stats = append(stats, stat{"CF name", "%s", func(p *proxy) interface{} { return p.bean.GetColumnFamilyName() }})
	stats = append(stats, stat{"DiskSpace", "%d", func(p *proxy) interface{} { return p.bean.GetTotalDiskSpaceUsed() }})
	stats = append(stats, stat{"ReadCount", "%d", func(p *proxy) interface{} { return p.bean.GetReadCount() }})
	stats = append(stats, stat{"WriteCount", "%d", func(p *proxy) interface{} { return p.bean.GetWriteCount() }})
	stats = append(stats, stat{"LiveSSTableCount", "%d", func(p *proxy) interface{} { return p.bean.GetLiveSSTableCount() }})
	
	
//	titles := forall(stats, func(s *stat) string { return s.title}).([]string)
	titles := make([]string, 0)
	formats := make([]string, 0)
	for _, s := range stats {
		titles = append(titles, s.title)
		formats = append(formats, s.format)
	}
	fmt.Fprintf(writer, strings.Join(titles, "\t") + "\n")
	format := strings.Join(formats, "\t") + "\n"
	results := make([]interface{}, len(stats))
	for _, v := range nodeProbe.GetColumnFamilyStoreMBeanProxies() {
		proxy := &proxy{v.Key, v.Value}
		for i, s := range stats {
			results[i] = s.f(proxy)
		}
		fmt.Fprintf(writer, format, results...)
	}
	
	if err = nodeProbe.Close(); err != nil {
		log.Fatal(err)
	}
	writer.Flush()
}

